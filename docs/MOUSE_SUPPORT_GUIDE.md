# Mouse Support Guide for TUITemplate

Complete guide to implementing mouse support in your TUI applications.

## Table of Contents

1. [Overview](#overview)
2. [Quick Start](#quick-start)
3. [Mouse Events](#mouse-events)
4. [Click Detection](#click-detection)
5. [Hover Effects](#hover-effects)
6. [Best Practices](#best-practices)
7. [Common Patterns](#common-patterns)

---

## Overview

TUITemplate has built-in mouse support that's enabled by default. The template handles mouse events in `update_mouse.go`, keeping your code organized and maintainable.

### What's Already Set Up

✅ Mouse motion tracking enabled in `main.go`
✅ Mouse event dispatcher in `update.go`
✅ Dedicated mouse handler in `update_mouse.go`
✅ Click region detection helpers
✅ Hover effect support

## Quick Start

Mouse support is enabled by default in the config:

```yaml
# config.yaml
ui:
  mouse_enabled: true
```

The template automatically configures Bubbletea with mouse support in `main.go`:

```go
if cfg.UI.MouseEnabled {
    opts = append(opts, tea.WithMouseCellMotion())
}
```

## Mouse Events

Bubbletea provides these mouse event types:

### Click Events

- `tea.MouseLeft` - Left button click
- `tea.MouseRight` - Right button click
- `tea.MouseMiddle` - Middle button click

### Scroll Events

- `tea.MouseWheelUp` - Scroll wheel up
- `tea.MouseWheelDown` - Scroll wheel down

### Motion Events

- `tea.MouseMotion` - Mouse movement (for hover effects)

## Click Detection

### Basic Click Handler

```go
// In update_mouse.go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    x, y := msg.X, msg.Y

    // Check regions from top to bottom
    if m.isInTitleBar(x, y) {
        return m.handleTitleBarClick(x, y)
    }

    if m.isInContentArea(x, y) {
        return m.handleContentClick(x, y)
    }

    if m.isInStatusBar(x, y) {
        return m.handleStatusBarClick(x, y)
    }

    return m, nil
}
```

### Region Detection Helper

```go
func (m model) isInRegion(x, y, startX, startY, width, height int) bool {
    return x >= startX && x < startX+width &&
           y >= startY && y < startY+height
}
```

### Button Click Detection

```go
// Example: Detect clicks on buttons at specific positions
func (m model) handleButtonClick(x, y int) (tea.Model, tea.Cmd) {
    // Button 1 at position (2, 1), width 10
    if x >= 2 && x < 12 && y == 1 {
        return m.handleButton1()
    }

    // Button 2 at position (14, 1), width 10
    if x >= 14 && x < 24 && y == 1 {
        return m.handleButton2()
    }

    return m, nil
}
```

## Hover Effects

### Track Mouse Position

Add to your model in `types.go`:

```go
type model struct {
    // ... other fields

    // Mouse tracking
    mouseX      int
    mouseY      int
    hoveredItem string
}
```

### Handle Mouse Motion

```go
func (m model) handleMouseMotion(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    m.mouseX = msg.X
    m.mouseY = msg.Y

    // Detect what's being hovered
    if m.isOverButton1(msg.X, msg.Y) {
        m.hoveredItem = "button-1"
    } else if m.isOverButton2(msg.X, msg.Y) {
        m.hoveredItem = "button-2"
    } else {
        m.hoveredItem = ""
    }

    return m, nil
}
```

### Render Hover State

In `view.go`:

```go
func (m model) renderButton(label string, id string) string {
    var style lipgloss.Style

    if m.hoveredItem == id {
        style = buttonHoverStyle  // Highlighted
    } else {
        style = buttonStyle       // Normal
    }

    return style.Render(label)
}
```

Define hover style in `styles.go`:

```go
var buttonHoverStyle = lipgloss.NewStyle().
    Foreground(colorPrimary).
    Background(lipgloss.Color("#3E4451")).
    Bold(true).
    Padding(0, 2)
```

## Best Practices

### 1. **Always Check Boundaries**

```go
// Good
if x >= startX && x < endX && y >= startY && y < endY {
    // Handle click
}

// Bad - may miss edge cases
if x > startX && x <= endX {
    // Handle click
}
```

### 2. **Layer Click Regions (Top to Bottom)**

```go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    // Check smaller/overlapping regions first
    if m.isInButton(msg.X, msg.Y) {
        return m.handleButtonClick(msg.X, msg.Y)
    }

    // Then check larger regions
    if m.isInPanel(msg.X, msg.Y) {
        return m.handlePanelClick(msg.X, msg.Y)
    }

    return m, nil
}
```

### 3. **Provide Visual Feedback**

Always show when something is:
- Hovered (change color/brightness)
- Clicked (show active state)
- Selected (highlight)

### 4. **Make Click Targets Large Enough**

Minimum recommended size: 3x1 characters for buttons

```go
// Good - easy to click
button := "[ Click Me ]"  // 12 chars wide

// Bad - hard to click
button := "[X]"  // Only 3 chars wide
```

### 5. **Update Status Bar**

Show mouse interaction feedback:

```go
func (m model) renderStatusBar() string {
    status := "Ready"

    if m.hoveredItem != "" {
        status = "Hovering: " + m.hoveredItem
    }

    if m.config.UI.MouseEnabled {
        status += " | Mouse: " + fmt.Sprintf("%d,%d", m.mouseX, m.mouseY)
    }

    return statusStyle.Render(status)
}
```

## Common Patterns

### Pattern 1: Clickable List Items

```go
type model struct {
    items       []string
    cursor      int
    listStartY  int
}

func (m model) handleContentClick(x, y int) (tea.Model, tea.Cmd) {
    // Calculate which item was clicked
    itemIndex := (y - m.listStartY)

    if itemIndex >= 0 && itemIndex < len(m.items) {
        m.cursor = itemIndex
        return m.selectItem()
    }

    return m, nil
}
```

### Pattern 2: Dual Pane Click Detection

```go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    leftWidth, _ := m.calculateDualPaneLayout()

    if msg.X < leftWidth {
        // Click in left pane
        m.focusedPane = "left"
        return m.handleLeftPaneClick(msg)
    } else {
        // Click in right pane
        m.focusedPane = "right"
        return m.handleRightPaneClick(msg)
    }
}
```

### Pattern 3: Tab Bar Clicks

```go
func (m model) handleTabBarClick(x, y int) (tea.Model, tea.Cmd) {
    tabNames := []string{"Tab 1", "Tab 2", "Tab 3"}
    xPos := 0

    for i, name := range tabNames {
        tabWidth := len(name) + 4  // "[ " + name + " ]"

        if x >= xPos && x < xPos+tabWidth {
            m.currentTab = i
            return m, nil
        }

        xPos += tabWidth + 1  // +1 for space between tabs
    }

    return m, nil
}
```

### Pattern 4: Context Menu on Right Click

```go
type model struct {
    showContextMenu bool
    contextMenuX    int
    contextMenuY    int
}

func (m model) handleRightClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    m.showContextMenu = true
    m.contextMenuX = msg.X
    m.contextMenuY = msg.Y
    return m, nil
}

func (m model) renderContextMenu() string {
    if !m.showContextMenu {
        return ""
    }

    menu := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(colorPrimary).
        Render("Copy\nPaste\nDelete")

    // Position menu at click location
    return lipgloss.PlaceHorizontal(
        m.width,
        lipgloss.Left,
        menu,
        lipgloss.WithWhitespaceChars(" "),
        lipgloss.WithWhitespaceForeground(lipgloss.Color("0")),
    )
}
```

### Pattern 5: Scroll Handling

```go
func (m model) handleWheelUp(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    // Scroll up by 3 lines
    m.scrollOffset = max(0, m.scrollOffset-3)
    return m, nil
}

func (m model) handleWheelDown(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    // Scroll down by 3 lines
    maxScroll := max(0, len(m.items)-m.viewportHeight)
    m.scrollOffset = min(maxScroll, m.scrollOffset+3)
    return m, nil
}
```

## Testing Mouse Support

### Manual Testing Checklist

- [ ] Click all buttons and verify they respond
- [ ] Hover over interactive elements and check visual feedback
- [ ] Test scrolling with mouse wheel
- [ ] Click in different panels/regions
- [ ] Test with different terminal sizes
- [ ] Try right-click (if implemented)
- [ ] Verify mouse position tracking in status bar

### Debugging Tips

1. **Show mouse position in status bar** for debugging click regions
2. **Log click coordinates** to verify detection ranges
3. **Use distinct colors** for different hover states
4. **Test in different terminals** (some terminals handle mouse differently)

## Example: Complete Mouse-Enabled Button

```go
// types.go
type model struct {
    hoveredButton string
    buttons       []Button
}

type Button struct {
    id    string
    label string
    x     int
    y     int
    width int
}

// update_mouse.go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    for _, btn := range m.buttons {
        if m.isClickOnButton(msg.X, msg.Y, btn) {
            return m.handleButtonClick(btn.id)
        }
    }
    return m, nil
}

func (m model) handleMouseMotion(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    m.hoveredButton = ""

    for _, btn := range m.buttons {
        if m.isClickOnButton(msg.X, msg.Y, btn) {
            m.hoveredButton = btn.id
            break
        }
    }

    return m, nil
}

func (m model) isClickOnButton(x, y int, btn Button) bool {
    return x >= btn.x && x < btn.x+btn.width && y == btn.y
}

// view.go
func (m model) renderButton(btn Button) string {
    var style lipgloss.Style

    if m.hoveredButton == btn.id {
        style = buttonHoverStyle
    } else {
        style = buttonStyle
    }

    return lipgloss.PlaceHorizontal(
        btn.width,
        lipgloss.Center,
        style.Render(btn.label),
    )
}
```

## Reference: layout-demo Implementation

See the `layout-demo` example in TUITemplate for a complete working example:

- Clickable layout switcher buttons
- Clickable tabs with hover effects
- Panel click detection
- Mouse position tracking
- Real-time hover feedback

Location: `examples/layout-demo/`

Files to review:
- `update_mouse.go` - Mouse event handling
- `view.go` - Rendering buttons with hover states
- `types.go` - Mouse tracking state

---

**Happy mouse clicking!** 🖱️✨

For more information, see:
- [Bubbletea Mouse Documentation](https://github.com/charmbracelet/bubbletea#mouse-support)
- [TUITemplate Architecture Guide](../ARCHITECTURE.md)
- [Layout Demo Example](../../examples/layout-demo/)
