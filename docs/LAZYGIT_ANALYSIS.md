# LazyGit TUI Architecture Analysis

Deep dive into how LazyGit achieves its excellent panel management and dynamic resizing.

## 🔑 Key Insights

### 1. Box Layout System

LazyGit uses a **hierarchical box layout** system (from `lazycore/pkg/boxlayout`):

```go
type Box struct {
    Direction Direction  // ROW or COLUMN
    Children  []*Box
    Window    string     // Window name (if leaf node)
    Size      int        // Static size (fixed height/width)
    Weight    int        // Dynamic size (proportional)
}
```

**Key Concept**: You MUST choose either `Size` (static) OR `Weight` (dynamic), not both!

- **Size**: "I want exactly 3 rows" or "I want exactly 20 columns"
- **Weight**: "Give me a proportion of remaining space"

**Example:**
```go
// Two boxes with weights 1 and 2
// First gets 33% of space, second gets 66%
[]*Box{
    {Window: "panel1", Weight: 1},
    {Window: "panel2", Weight: 2},
}
```

### 2. Focus-Based Dynamic Resizing

This is the MAGIC! 🪄

LazyGit changes panel **weights** based on which panel is focused:

```go
// From window_arrangement_helper.go:238-266
func getMidSectionWeights(args WindowArrangementArgs) (int, int) {
    sideSectionWeight := 40  // Default ratio
    mainSectionWeight := 80

    // KEY: When main panel is focused in HALF/FULL mode
    if args.CurrentWindow == "main" {
        if args.ScreenMode == types.SCREEN_HALF ||
           args.ScreenMode == types.SCREEN_FULL {
            sideSectionWeight = 0  // Hide side panel!
        }
    } else {
        // When side panel is focused
        if args.ScreenMode == types.SCREEN_HALF {
            mainSectionWeight = sideSectionWeight  // Equal split
        } else if args.ScreenMode == types.SCREEN_FULL {
            mainSectionWeight = 0  // Hide main panel!
        }
    }

    return sideSectionWeight, mainSectionWeight
}
```

**Result**: Panels expand/shrink based on focus!

### 3. Accordion Mode for Side Panels

The "snap to larger" effect in side panels:

```go
// From window_arrangement_helper.go:447-457
accordionMode := args.UserConfig.Gui.ExpandFocusedSidePanel

accordionBox := func(defaultBox *Box) *Box {
    if accordionMode && defaultBox.Window == args.CurrentSideWindow {
        return &Box{
            Window: defaultBox.Window,
            Weight: 2,  // 2x weight when focused!
        }
    }
    return defaultBox  // Weight: 1 when not focused
}

// Applied to each panel
panels := []*Box{
    {Window: "status", Size: 3},  // Always 3 rows
    accordionBox(&Box{Window: "files", Weight: 1}),
    accordionBox(&Box{Window: "branches", Weight: 1}),
    accordionBox(&Box{Window: "commits", Weight: 1}),
}
```

**Result**:
- Unfocused panels: Weight 1 (smaller)
- Focused panel: Weight 2 (2x larger!)

### 4. Screen Modes

Three modes that dramatically change layout:

1. **SCREEN_NORMAL**: All panels visible, normal proportions
2. **SCREEN_HALF**: Focused section gets more space
3. **SCREEN_FULL**: Only focused section visible (other weight = 0)

### 5. Conditional Sizing Based on Terminal Size

```go
// From window_arrangement_helper.go:422-498
func sidePanelChildren(args) func(width, height int) []*Box {
    return func(width, height int) []*Box {
        if args.ScreenMode == types.SCREEN_FULL {
            // Full screen: Only show focused panel
            return fullHeightBoxes()
        } else if height >= 28 {
            // Normal mode: Accordion if enabled
            return normalPanels()
        } else {
            // Compact mode (<28 rows): Squashed panels
            return squashedPanels()
        }
    }
}
```

**Result**: Layout adapts to terminal size!

### 6. The Layout Flow

```
User focuses panel
    ↓
Focus changes CurrentWindow/CurrentSideWindow
    ↓
layout() is called (on every render)
    ↓
getWindowDimensions() calculates new dimensions
    ↓
getMidSectionWeights() adjusts weights based on focus
    ↓
sidePanelChildren() adjusts side panel weights
    ↓
boxlayout.ArrangeWindows() calculates pixel positions
    ↓
Panels resize smoothly!
```

## 🎯 Implementation Patterns

### Pattern 1: Weight-Based Focus Expansion

```go
func calculatePanelWeights(focusedPanel string) (leftWeight, rightWeight int) {
    if focusedPanel == "left" {
        return 2, 1  // Left gets 66%, right gets 33%
    } else if focusedPanel == "right" {
        return 1, 2  // Left gets 33%, right gets 66%
    }
    return 1, 1  // Equal split (50/50)
}
```

### Pattern 2: Size-Based Focus Expansion

```go
func calculatePanelSizes(totalWidth int, focusedPanel string) (leftWidth, rightWidth int) {
    if focusedPanel == "left" {
        return int(totalWidth * 0.7), int(totalWidth * 0.3)
    } else if focusedPanel == "right" {
        return int(totalWidth * 0.3), int(totalWidth * 0.7)
    }
    return totalWidth / 2, totalWidth / 2
}
```

### Pattern 3: Multi-Panel Accordion

```go
type Panel struct {
    Name     string
    IsFocused bool
}

func calculateMultiPanelHeights(panels []Panel, totalHeight int, numPanels int) []int {
    baseHeight := 3  // Minimum for unfocused

    // Count focused panels
    focusedCount := 0
    for _, p := range panels {
        if p.IsFocused {
            focusedCount++
        }
    }

    if focusedCount == 0 {
        // Equal distribution
        perPanel := totalHeight / numPanels
        return []int{perPanel, perPanel, perPanel, ...}
    }

    // Give focused panel most space
    unfocusedTotal := (numPanels - focusedCount) * baseHeight
    focusedHeight := totalHeight - unfocusedTotal

    heights := []int{}
    for _, p := range panels {
        if p.IsFocused {
            heights = append(heights, focusedHeight)
        } else {
            heights = append(heights, baseHeight)
        }
    }

    return heights
}
```

## 📊 LazyGit's Layout Strategy

### Width Allocation (Normal Mode)

```
Terminal Width: 120 characters
├── Side Section: 40 chars (33%)
│   ├── Status: 3 rows (fixed)
│   ├── Files: Weight 1
│   ├── Branches: Weight 1 (or 2 if focused with accordion)
│   ├── Commits: Weight 1
│   └── Stash: 3 rows (or Weight 1 if focused)
└── Main Section: 80 chars (67%)
    ├── Main: Weight 1
    └── Secondary: Weight 1 (if split)
```

### Height Allocation (Side Panels, Normal Mode)

```
Total Height: 30 rows
├── Status: 3 rows (fixed Size)
├── Files: 6.75 rows (Weight 1)
├── Branches: 13.5 rows (Weight 2 - FOCUSED!)
├── Commits: 6.75 rows (Weight 1)
└── Stash: 3 rows (fixed Size)
```

### Height Allocation (Side Panels, Accordion Disabled)

```
Total Height: 30 rows
├── Status: 3 rows
├── Files: 9 rows (Weight 1, equal split)
├── Branches: 9 rows (Weight 1, equal split)
├── Commits: 9 rows (Weight 1, equal split)
└── Stash: 3 rows (focused) or 1 row (unfocused)
```

## 💡 Key Takeaways for TUITemplate

### 1. **Dynamic Weight System**
- Store current focused panel in model
- Recalculate weights on every layout pass
- Use weight ratios (1:2, 1:3, etc.) for smooth resizing

### 2. **Accordion Mode**
- Make it a config option: `ExpandFocusedPanel bool`
- When enabled, focused panel gets 2x weight
- When disabled, all panels equal weight

### 3. **Screen Modes** (Optional)
- Normal: All visible
- Half: Focused section larger
- Full: Only focused section

### 4. **Minimum Sizes**
- Unfocused panels: 3 rows minimum
- Prevents panels from disappearing completely
- User can still see what's there

### 5. **Smooth Transitions**
- No animations needed (Bubbletea handles redraws)
- Weights create natural proportions
- Feels smooth because it's immediate

## 🔧 Implementation Checklist for TUITemplate

- [ ] Create box layout utility (or use existing layout calc)
- [ ] Add `focusedPanel string` to model
- [ ] Add `ExpandFocusedPanel bool` to config
- [ ] Implement weight calculator based on focus
- [ ] Update layout function to use dynamic weights
- [ ] Add keyboard shortcuts to switch focus
- [ ] Add mouse click to change focus
- [ ] Test with 2-4 panels
- [ ] Document the pattern

## 📝 Minimal Example

```go
// types.go
type model struct {
    focusedPanel     string  // "left", "right", "bottom"
    expandOnFocus    bool    // Config: accordion mode
}

// view.go
func (m model) calculateDualPaneLayout() (int, int) {
    totalWidth := m.width

    if !m.expandOnFocus {
        // Equal split
        return totalWidth / 2, totalWidth / 2
    }

    // Accordion mode
    if m.focusedPanel == "left" {
        // Left gets 70%, right gets 30%
        leftWidth := int(float64(totalWidth) * 0.7)
        rightWidth := totalWidth - leftWidth
        return leftWidth, rightWidth
    } else {
        // Left gets 30%, right gets 70%
        leftWidth := int(float64(totalWidth) * 0.3)
        rightWidth := totalWidth - leftWidth
        return leftWidth, rightWidth
    }
}

// update_keyboard.go
func (m model) handleMainKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "1":
        m.focusedPanel = "left"
        return m, nil
    case "2":
        m.focusedPanel = "right"
        return m, nil
    }
    return m, nil
}

// update_mouse.go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    leftWidth, _ := m.calculateDualPaneLayout()

    if msg.X < leftWidth {
        m.focusedPanel = "left"
    } else {
        m.focusedPanel = "right"
    }

    return m, nil
}
```

## 🎨 Visual Example

```
BEFORE (both panels unfocused, 50/50):
╔═══════════════╦═══════════════╗
║   LEFT        ║   RIGHT       ║
║   PANEL       ║   PANEL       ║
║               ║               ║
╚═══════════════╩═══════════════╝

AFTER (left panel focused with accordion, 70/30):
╔════════════════════╦══════════╗
║   LEFT             ║  RIGHT   ║
║   PANEL            ║  PANEL   ║
║   (FOCUSED)        ║          ║
╚════════════════════╩══════════╝
```

## 🚀 Why This Works So Well

1. **Immediate**: No animations, instant resize
2. **Proportional**: Weight system maintains ratios naturally
3. **Predictable**: Same behavior every time
4. **Configurable**: Accordion mode can be toggled
5. **Space-Efficient**: Focused panel gets attention, unfocused stay visible but small

---

**Source Code**: `github.com/jesseduffield/lazygit`
**Key Files**:
- `pkg/gui/controllers/helpers/window_arrangement_helper.go`
- `vendor/github.com/jesseduffield/lazycore/pkg/boxlayout/boxlayout.go`

**License**: MIT (LazyGit is open source!)

This analysis is for educational purposes to learn layout patterns. Implement your own version for TUITemplate! 🎉
