# Layout Demo

Comprehensive demonstration of all TUITemplate layout types with labeled panels for easy reference.

## Purpose

This app demonstrates all four layout types available in TUITemplate:
1. **Single Pane** - Full-screen single view
2. **Dual Pane** - Side-by-side panels
3. **Multi-Panel** - Three panels (2 top, 1 bottom)
4. **Tabbed** - Tabbed interface with 3 tabs

Each panel displays its **Panel ID** and purpose, making it easy to give instructions to Claude about specific panels.

## Quick Start

```bash
# Run the app
go run .

# Or build and run
go build
./layout-demo
```

## Keyboard Shortcuts

### Layout Switching
- `1` - Switch to Single Pane layout
- `2` - Switch to Dual Pane layout
- `3` - Switch to Multi-Panel layout
- `4` - Switch to Tabbed layout

### Tab Navigation (in Tabbed layout)
- `Tab` - Next tab
- `Shift+Tab` - Previous tab

### General
- `q` or `Ctrl+C` - Quit
- `?` - Help
- `Ctrl+R` - Refresh

### Mouse Support 🖱️
- **Click** layout buttons (1-4) to switch layouts
- **Click** tabs in tabbed layout
- **Click** panels to interact
- **Hover** over buttons/tabs for visual feedback
- **Scroll** with mouse wheel to navigate
- Mouse position shown in status bar

## Panel Reference

Use these Panel IDs when giving instructions to Claude:

### Single Pane Layout
- **main-content** - The main content area

### Dual Pane Layout
- **left-pane** - Left panel (navigation, lists, menus)
- **right-pane** - Right panel (preview, details, content)

### Multi-Panel Layout
- **top-left** - Top left panel (primary navigation, controls)
- **top-right** - Top right panel (details, preview)
- **bottom-panel** - Bottom panel (logs, console, output)

### Tabbed Layout
- **tab1-content** - Overview tab (First tab content)
- **tab2-content** - Content tab (Second tab content)
- **tab3-content** - Settings tab (Third tab content)
- **border-showcase** - Border Styles tab (Shows all Lipgloss border styles)
- **color-palette** - Colors tab (Shows available color palette)

## Example Instructions for Claude

"Update the **left-pane** to show a file list"
"Add syntax highlighting to the **right-pane**"
"Display logs in the **bottom-panel**"
"Show user settings in **tab2-content**"

## Features Demonstrated

- ✅ All 4 layout types
- ✅ Layout switching at runtime
- ✅ Tab navigation (5 tabs total)
- ✅ Labeled panels with clear IDs
- ✅ Keyboard shortcuts
- ✅ Status bar with context
- ✅ Title bar showing current layout
- ✅ **Border Showcase** - All Lipgloss border styles (Rounded, Double, Thick, Normal, Hidden)
- ✅ **Color Palette** - Complete color reference with hex codes
- ✅ **Full Mouse Support** - Clickable buttons, tabs, and panels with hover effects
- ✅ **Real-time Mouse Tracking** - See mouse position and hovered elements in status bar

## Built With

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - Components
- [TUITemplate](https://github.com/yourname/TUITemplate) - Project template

## License

MIT
