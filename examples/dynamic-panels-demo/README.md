# Dynamic Panels Demo

A demonstration of **LazyGit-inspired dynamic panel resizing** with focus-based weight allocation.

## Features

🎯 **Weight-Based Layout System**
- Panels use weight ratios (1:1, 1:2, 2:1) instead of fixed sizes
- Smooth, proportional resizing based on focus

🎨 **Accordion Mode**
- Focused panel automatically gets 2x the space
- Unfocused panels shrink to minimum size
- Toggle on/off to compare behaviors

📐 **Two Layout Modes**
- **2-Panel Mode**: Left and right panels side-by-side
- **3-Panel Mode**: Left/right on top, bottom panel below

🖱️ **Full Mouse & Keyboard Support**
- Click any panel to focus it
- Keyboard shortcuts for quick navigation
- Visual feedback for focused and hovered panels

## How It Works

Based on LazyGit's box layout system:

### Weight-Based Sizing

Instead of calculating pixel widths manually, panels have **weights**:

```
Two panels with equal weight (1:1):
╔══════════════════╦══════════════════╗
║   LEFT (w=1)     ║   RIGHT (w=1)    ║
║   50% of space   ║   50% of space   ║
╚══════════════════╩══════════════════╝
```

```
Accordion mode - left focused (2:1):
╔═══════════════════════════╦═════════╗
║   LEFT (w=2)              ║ RIGHT   ║
║   66% of space            ║ (w=1)   ║
║                           ║ 33%     ║
╚═══════════════════════════╩═════════╝
```

```
Accordion mode - right focused (1:2):
╔═════════╦═══════════════════════════╗
║ LEFT    ║   RIGHT (w=2)             ║
║ (w=1)   ║   66% of space            ║
║ 33%     ║                           ║
╚═════════╩═══════════════════════════╝
```

### Dynamic Weight Calculation

```go
// From model.go:96-120
func (m model) calculateDualPaneLayout() (int, int) {
    // Calculate weights based on focus and accordion mode
    leftWeight, rightWeight := 1, 1

    if m.accordionMode {
        if m.focusedPanel == "left" {
            leftWeight = 2  // Focused panel gets 2x weight
        } else if m.focusedPanel == "right" {
            rightWeight = 2
        }
    }

    // Calculate actual widths from weights
    totalWeight := leftWeight + rightWeight
    leftWidth := (availableWidth * leftWeight) / totalWeight
    rightWidth := availableWidth - leftWidth

    return leftWidth, rightWidth
}
```

This creates the **snap-to-larger/smaller** effect seen in LazyGit!

## Controls

### Focus Navigation

| Key | Action |
|-----|--------|
| `1` | Focus left panel |
| `2` | Focus right panel |
| `3` | Focus bottom panel (3-panel mode only) |
| `Tab` | Cycle focus forward |
| `Shift+Tab` | Cycle focus backward |
| `←/→` | Navigate between left/right panels |
| `↑/↓` | Navigate to/from bottom panel |

### Mode Toggles

| Key | Action |
|-----|--------|
| `a` | Toggle accordion mode (ON/OFF) |
| `m` | Toggle 3-panel mode (2/3 panels) |

### General

| Key | Action |
|-----|--------|
| `q` | Quit |
| `Ctrl+C` | Quit |

### Mouse

- **Click any panel** to focus it
- **Hover** over panels to see which one you're targeting

## Visual Feedback

- **Focused panel**: Bright blue border + `●` indicator
- **Unfocused panel**: Dim gray border
- **Panel header**: Shows weight and size (e.g., "Weight: 2 | Size: 80x25")
- **Status bar**: Shows currently focused panel and mouse position
- **Title bar**: Shows current modes (Accordion ON/OFF, 2/3-Panel)

## Installation & Usage

```bash
# Build the demo
go build -o dynamic-panels-demo

# Run it
./dynamic-panels-demo
```

**Recommended**: Run in a terminal at least 80×25 characters.

## Try This!

1. **Start with accordion mode ON** (default)
2. **Click the left panel** → Watch it expand to 66% width
3. **Click the right panel** → Watch it expand while left shrinks
4. **Press `a`** → Disable accordion mode
5. **Click panels again** → Notice they stay 50/50 now
6. **Press `m`** → Enable 3-panel mode
7. **Click the bottom panel** → Watch it expand vertically
8. **Press `3` then `1`** → See weights change in real-time

## Implementation Highlights

### Key Files

- **`model.go`**: Weight-based layout calculations
  - `calculateDualPaneLayout()` - 2-panel weight system
  - `calculateThreePanelLayout()` - 3-panel weight system

- **`view.go`**: Dynamic rendering
  - `renderPanel()` - Shows weight & size in each panel
  - Focus-based border styling

- **`update_keyboard.go`**: Focus switching
  - Number keys (1/2/3) for direct focus
  - Arrow keys for navigation
  - Mode toggles (a/m)

- **`update_mouse.go`**: Click-to-focus
  - Panel boundary detection
  - Focus on click

## Pattern Summary

This demo implements the LazyGit pattern:

1. **Panels have weights, not fixed sizes**
2. **Focused panel gets 2x weight** (accordion mode)
3. **Weights are recalculated on every focus change**
4. **Layout is recomputed immediately** (no animations needed)
5. **Result: Smooth, proportional resize behavior**

## Resources

- **LazyGit Analysis**: See `/home/matt/projects/TUITemplate/docs/LAZYGIT_ANALYSIS.md`
- **LazyGit Source**: https://github.com/jesseduffield/lazygit
- **Box Layout Package**: `lazycore/pkg/boxlayout`

---

**Built with**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) • [Lip Gloss](https://github.com/charmbracelet/lipgloss)

**License**: MIT

**Pattern inspired by**: [LazyGit](https://github.com/jesseduffield/lazygit) by Jesse Duffield
