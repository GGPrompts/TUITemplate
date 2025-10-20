# Multi-Panel TUI Example

A lazygit-style multi-panel TUI application demonstrating focus management, panel navigation, dynamic resizing, and optimal screen space usage.

## Features

### Layout
- **Left Panel**: File browser/list view
- **Top Right Panel**: Details/preview panel
- **Bottom Right Panel**: Logs/status panel
- **Dynamic Sizing**: Panels resize based on focus (accordion mode)

### Accordion Mode 🎯
**NEW: LazyGit-inspired dynamic panel resizing!**

When accordion mode is **ON** (default):
- Focused panel gets **2x weight** (66% of space)
- Unfocused panels get **1x weight** (33% of space)
- Click or use keyboard to focus → panel expands instantly
- Visual indicator: ● appears in focused panel title

When accordion mode is **OFF**:
- Fixed 2:1 layout ratio
- Top panels always get 66% height
- Left/right panels maintain equal width

### Focus Management
- **Visual Indicators**: Focused panels have highlighted borders and titles
- **Cursor Indicators**: Selected items show ▶ when focused, ▸ when unfocused
- **Panel Switching**: Smooth transitions between panels

### Keyboard Navigation

#### Panel Navigation
- `Tab` - Cycle through panels (forward)
- `Shift+Tab` - Cycle through panels (backward)
- `h` - Move focus to left panel
- `l` - Move focus to right panels (cycles between top/bottom)

#### Within Panels
- `↑` or `k` - Move up / scroll up
- `↓` or `j` - Move down / scroll down
- `Enter` - Select/activate item

#### Global
- `q` or `Ctrl+C` - Quit application
- `r` - Refresh view (adds log entry)
- `a` or `A` - **Toggle accordion mode** (dynamic vs fixed panel sizing)

### Mouse Support

#### Panel Focus
- **Click** on any panel to focus it
- Focused panel shows highlighted border and title

#### Scrolling
- **Mouse wheel up** - Scroll up in focused panel
  - Files panel: Move cursor up
  - Logs panel: Scroll logs up
- **Mouse wheel down** - Scroll down in focused panel
  - Files panel: Move cursor down
  - Logs panel: Scroll logs down

## Running the Example

```bash
cd examples/multi-panel
go run .
```

Or build and run:
```bash
go build -o multi-panel .
./multi-panel
```

## Code Structure

```
multi-panel/
├── main.go      # Entry point
├── model.go     # Application state and data structures
├── update.go    # Event handling and business logic
├── view.go      # Multi-panel rendering
└── styles.go    # Lipgloss styles and theming
```

## Key Concepts

### 1. Panel Focus Management

Each panel has a unique `PanelID`:
```go
type PanelID int

const (
    LeftPanel PanelID = iota
    TopRightPanel
    BottomRightPanel
)
```

The model tracks which panel is focused:
```go
focusedPanel PanelID
```

### 2. Visual Feedback

Focused panels use different styles:
```go
func getPanelBorderStyle(focused bool) lipgloss.Style {
    if focused {
        return focusedBorderStyle  // Purple border
    }
    return unfocusedBorderStyle    // Gray border
}
```

### 3. Dynamic Weight-Based Layout (NEW!)

Panel dimensions are calculated using a **weight system** inspired by LazyGit:

```go
// calculateThreePanelLayout uses weights to determine panel sizes
func (m model) calculateThreePanelLayout(availableWidth, availableHeight int) (leftWidth, rightWidth, topHeight, bottomHeight int) {
    // Calculate horizontal weights
    leftWeight, rightWeight := 1, 1
    if m.accordionMode && m.focusedPanel == LeftPanel {
        leftWeight = 2  // Focused panel gets 2x weight!
    }

    // Convert weights to actual widths
    totalWeight := leftWeight + rightWeight
    leftWidth = (availableWidth * leftWeight) / totalWeight
    rightWidth = availableWidth - leftWidth

    // Similar calculation for vertical split...
}
```

**The Math**:
- Equal weights (1:1) = 50/50 split
- Focused weight (2:1) = 66/33 split
- Formula: `width = (totalWidth * weight) / totalWeights`

**Why weights work**:
- Proportional sizing (not fixed pixels)
- Instant recalculation (no animations needed)
- Simple and predictable ratios

### 4. Context-Aware Input Handling

Different actions based on which panel is focused:
```go
func (m model) handleDown() model {
    switch m.focusedPanel {
    case LeftPanel:
        // Navigate down in file list
    case TopRightPanel:
        // Scroll details down
    case BottomRightPanel:
        // Scroll logs down
    }
}
```

## Customization Ideas

### Add More Panels
1. Add new `PanelID` constant
2. Add to `m.panels` slice
3. Create `renderXXXPanel()` function
4. Adjust layout calculations

### Mouse Support Implementation
The application includes full mouse support with click-to-focus and wheel scrolling:

**Panel Boundaries**: Each panel's screen position is tracked in `m.panelBounds`:
```go
type panelBounds struct {
    x, y, width, height int
}
```

**Click Detection**: Mouse clicks are mapped to panels using boundary checks:
```go
case tea.MouseLeft:
    for panelID, bounds := range m.panelBounds {
        if bounds.contains(msg.X, msg.Y) {
            m.focusedPanel = panelID
        }
    }
```

**Wheel Scrolling**: Separate handlers for each panel type provide context-aware scrolling

### Dynamic Panel Sizing
Add user-controlled panel resizing:
- Store panel ratios in model
- Handle `+`/`-` keys to adjust splits
- Recalculate dimensions in view

### Panel-Specific Commands
Add different keybindings per panel:

```go
case "d":
    if m.focusedPanel == LeftPanel {
        return m.deleteFile()
    } else if m.focusedPanel == TopRightPanel {
        return m.toggleDiff()
    }
```

## Accordion Mode Demo 🎨

Try this sequence to see accordion mode in action:

1. **Launch** the app → Left panel focused, accordion ON
2. **Click top-right panel** → It expands to 66% width
3. **Click bottom-right panel** → It expands to 66% height
4. **Click left panel** → It expands to 66% width
5. **Press 'a'** → Accordion OFF (fixed 2:1 layout)
6. **Click panels** → Focus changes but sizes stay fixed
7. **Press 'a'** → Accordion ON again
8. **Watch panels resize** instantly as you click!

**Visual Indicators**:
- Focused panel title shows **●** in accordion mode
- Border color changes (purple = focused, gray = unfocused)
- Cursor changes (▶ = focused, ▸ = unfocused)

## Comparison to lazygit

This example implements lazygit's core UI patterns:

| Feature | lazygit | This Example |
|---------|---------|--------------|
| Multi-panel layout | ✓ | ✓ |
| Focus indicators | ✓ | ✓ |
| Panel navigation (h/l) | ✓ | ✓ |
| **Weight-based dynamic sizing** | ✓ | ✓ **NEW!** |
| **Accordion mode** | ✓ | ✓ **NEW!** |
| Responsive sizing | ✓ | ✓ |
| Vim-style navigation | ✓ | ✓ |
| Visual border highlights | ✓ | ✓ |
| Mouse click to focus | ✓ | ✓ |
| Mouse wheel scrolling | ✓ | ✓ |

## Use as a Template

This example serves as a production-ready template for building lazygit-style TUI applications. To adapt it:

1. **Replace sample data**: Update `files`, `details`, and `logs` with your data sources
2. **Customize panels**: Modify panel content in `view.go` rendering functions
3. **Add business logic**: Implement actual operations in `update.go` handlers
4. **Extend styling**: Customize colors and styles in `styles.go`
5. **Add features**: Implement search, filtering, sorting, etc.

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework (Elm Architecture)
- `github.com/charmbracelet/lipgloss` - Style definitions and layout

## License

Part of TUITemplate - see project root for license information.
