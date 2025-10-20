# TUI Showcase

**The Comprehensive TUI Pattern Library**

A single, unified application demonstrating **ALL** TUI patterns from basic layouts to advanced components.

## Quick Start

```bash
go run .
```

Press **q** to quit.

## What's Inside?

This application consolidates patterns from multiple demos into one comprehensive showcase with:
- **4 Layout Modes** - Single, dual pane, multi-panel, tabbed
- **12 Interactive Tabs** - From basic layouts to advanced components
- **Dropdown Menu System** - Professional menu bar with mouse support
- **Full Mouse & Keyboard Support** - Click or type your way around

## Navigation

### Using Menus
Click the menu bar to open dropdowns:
- **File** - New, Open, Save, Quit
- **View** - Switch layouts and tabs
- **Components** - Jump to component showcases
- **Help** - Keyboard shortcuts and about

### Using Keyboard
```
Layouts:
  1 - Single Pane
  2 - Dual Pane
  3 - Multi-Panel
  4 - Tabbed

Tabs (in Tabbed mode):
  Tab       - Next tab
  Shift+Tab - Previous tab

Other:
  ? - Help
  q - Quit
```

## The 12 Showcases

### Core Layout Tabs (0-5)

**Tab 0: Overview**
- Introduction to the showcase
- Navigation instructions
- Quick reference

**Tab 1: Content**
- Basic content area demonstration
- Single-pane layout example

**Tab 2: Settings**
- Settings panel example
- Configuration UI patterns

**Tab 3: Borders**
- Lipgloss border styles showcase
- Rounded, Double, Thick, Normal borders
- Border color demonstrations

**Tab 4: Colors**
- Theme color palette
- Primary, Secondary, Accent colors
- Error, Warning, Info colors
- Visual color examples

**Tab 5: Dynamic Panels**
- LazyGit-style accordion panels
- Weight-based layout system (1:1 = 50/50, 2:1 = 66/33)
- Click panels to focus them
- Press 'a' to toggle accordion mode
- Press 1/2 to switch focus

### Component Tabs (6-11)

**Tab 6: Forms** (Components → Forms)
- Text input fields
- Checkboxes
- Radio buttons
- Form validation examples

**Tab 7: Tables** (Components → Tables)
- Data grid with headers
- Row selection
- Sorting and filtering
- Table controls

**Tab 8: Dialogs** (Components → Dialogs)
- Confirmation dialogs
- Alert/success messages
- Error dialogs
- Input prompts

**Tab 9: Progress** (Components → Progress Bars)
- Determinate progress bars
- Multiple progress tracking
- Spinner animations
- Indeterminate progress

**Tab 10: Tree View** (Components → Tree View)
- File browser tree
- Expandable/collapsible nodes
- Hierarchical data display
- Tree navigation controls

**Tab 11: Mobile** (Components → Mobile Patterns)
- Termux-optimized UI
- Vertical stacking for narrow terminals
- Touch-friendly buttons
- Adaptive layout examples
- Mobile best practices

## Code Organization

```
tui-showcase/
├── main.go              - Entry point
├── types.go             - Type definitions
├── model.go             - Model init, layout calculations
├── update.go            - Main update dispatcher
├── update_keyboard.go   - Keyboard event handling
├── update_mouse.go      - Mouse event handling
├── view.go              - View rendering (layouts & core tabs)
├── view_components.go   - Component tab renderers (NEW!)
├── menu.go              - Dropdown menu system
├── styles.go            - Lipgloss styles
└── config.go            - Configuration
```

## Key Patterns Demonstrated

### 1. Dropdown Menus
- Clean menu bar rendering
- Dropdown overlay without corruption
- Mouse hover and click detection
- Keyboard navigation (arrows, enter, esc)

**See:** `menu.go`

### 2. Weight-Based Layouts
```go
// LazyGit pattern: focused panel gets 2x weight
leftWeight, rightWeight := 1, 1
if accordionMode && focusedPanel == "left" {
    leftWeight = 2  // Focused gets 66% (2:1 ratio)
}

leftWidth = (availableWidth * leftWeight) / (leftWeight + rightWeight)
```

**See:** `model.go:106-190` (Dynamic Panels tab)

### 3. Component Showcases

All component tabs demonstrate visual patterns without interactive logic:
- **Forms** - Input layouts and validation UI
- **Tables** - Grid structures and controls
- **Dialogs** - Modal window designs
- **Progress** - Progress bar styles
- **Tree** - Hierarchical data display
- **Mobile** - Responsive layout patterns

**See:** `view_components.go`

### 4. Border Height Accounting

```go
// CRITICAL: Account for panel borders
contentHeight -= 2 // top + bottom borders
```

This prevents panels from overflowing and covering the title bar.

**See:** `model.go:74`, `../../CLAUDE.md` for full explanation

### 5. Mobile Optimization

- Vertical stacking when width < 80 columns
- Touch-friendly UI elements (large tap targets)
- Adaptive layout switching
- Clear focus indicators

**See:** Tab 11 (Mobile Patterns)

## Building Your Own App

After exploring this showcase:

1. **Study the code** - Each file is documented with purpose and usage
2. **Copy patterns you need** - All code is reusable
3. **Reference this showcase** - Use as your TUI pattern library
4. **Read CLAUDE.md** - Critical bug fixes and layout patterns

## Related Documentation

- [../README.md](../README.md) - Examples overview
- [../../CLAUDE.md](../../CLAUDE.md) - Critical layout fixes and patterns
- [../../USAGE.md](../../USAGE.md) - TUITemplate usage guide
- [../../ARCHITECTURE.md](../../ARCHITECTURE.md) - Architecture patterns

## Tips

### Responsive Design
The showcase automatically adapts to terminal size:
- **< 80 cols** - Switches to vertical stacking
- **≥ 80 cols** - Uses side-by-side layouts

Resize your terminal to see adaptive behavior!

### Exploring Components
Navigate to component tabs via:
1. **Components menu** - Click "Components" → Select component
2. **Keyboard** - Press 4 (tabbed mode) then Tab to component tabs

### Understanding Code Structure
Each pattern is isolated:
- **Layout logic** - `model.go`, `view.go`
- **Components** - `view_components.go`
- **Menus** - `menu.go`
- **Styles** - `styles.go`

Copy the files you need for your project!

## Troubleshooting

**Panels covering title bar?**
- Check border accounting: `contentHeight -= 2`
- See `../../CLAUDE.md` for detailed explanation

**Mouse clicks not working?**
- Check terminal mouse support
- Verify coordinate calculations in `update_mouse.go`

**Text wrapping in panels?**
- Always truncate text: `maxTextWidth := width - 4`
- Never rely on Lipgloss auto-wrapping

**Layout issues on mobile?**
- Test with `termux-size-detector` to verify terminal dimensions
- Check vertical stacking logic in `view.go`

## Contributing

Found a pattern that should be added?

1. Add rendering function to `view_components.go`
2. Add tab to `renderTabBar()` in `view.go`
3. Add case to `renderTabbed()` switch statement
4. Wire up Components menu in `menu.go`
5. Update this README

## Built With

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - Components
- [TUITemplate](https://github.com/GGPrompts/TUITemplate) - Project template

## License

MIT

---

**Explore all 12 showcases and build amazing TUIs!** 🚀
