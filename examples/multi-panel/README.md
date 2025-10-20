# Multi-Panel TUI Example

A lazygit-style multi-panel TUI application demonstrating focus management, panel navigation, and optimal screen space usage.

## Features

### Layout
- **Left Panel**: File browser/list view (1/3 width)
- **Top Right Panel**: Details/preview panel (2/3 width, top half)
- **Bottom Right Panel**: Logs/status panel (2/3 width, bottom half)

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

### 3. Responsive Layout

Panel dimensions are calculated based on terminal size:
```go
// Left panel takes 1/3 of width
leftWidth := m.width / 3
rightWidth := m.width - leftWidth

// Right panels split vertically
rightTopHeight := availableHeight / 2
rightBottomHeight := availableHeight - rightTopHeight
```

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

## Comparison to lazygit

This example implements lazygit's core UI patterns:

| Feature | lazygit | This Example |
|---------|---------|--------------|
| Multi-panel layout | ✓ | ✓ |
| Focus indicators | ✓ | ✓ |
| Panel navigation (h/l) | ✓ | ✓ |
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
