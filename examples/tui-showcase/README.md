# TUI Showcase

**The Comprehensive TUI Pattern Library**

A single, unified application demonstrating **ALL** TUI patterns from basic layouts to advanced components.

## Quick Start

```bash
go run .
```

Press **q** to quit.

## What's Inside?

This application consolidates ALL TUI patterns into one comprehensive tab-based showcase:
- **12 Interactive Tabs** - From basic layouts to advanced components
- **Dropdown Menu System** - Professional menu bar with mouse support
- **Full Mouse & Keyboard Support** - Click or type your way around
- **Unified Navigation** - Everything accessible via tabs (no mode switching!)

## Navigation

### Using Menus
Click the menu bar to open dropdowns:
- **File** - New, Open, Save, Quit
- **View** - Jump to layout demos (Single, Dual, Multi-Panel, Borders, Colors, Dynamic)
- **Components** - Jump to component showcases (Forms, Tables, Dialogs, Progress, Tree, Mobile)
- **Help** - Keyboard shortcuts and about

### Using Keyboard
```
Tabs:
  Tab       - Next tab
  Shift+Tab - Previous tab

Dynamic Tab (Tab 5):
  1 - Focus left panel
  2 - Focus right panel
  3 - Focus bottom panel
  a - Toggle accordion mode

Other:
  ? - Help
  q - Quit
```

## The 12 Showcases

### Layout Demo Tabs (0-2)

**Tab 0: Single Pane**
- Full-screen single content area demo
- Best for: simple apps, focus mode, readers, editors
- Shows the single-pane layout pattern

**Tab 1: Dual Pane**
- Side-by-side panels with divider
- Best for: file browsers, previews, compare views
- Live demo using actual dual-pane rendering

**Tab 2: Multi-Panel**
- 3-panel layout (top-left, top-right, bottom)
- Best for: IDEs, dashboards, complex UIs
- Live demo with clickable panels

### Core Showcase Tabs (3-5)

**Tab 3: Borders**
- Lipgloss border styles showcase
- Rounded, Double, Thick, Normal borders
- Border color demonstrations

**Tab 4: Colors**
- Theme color palette showcase
- Primary, Secondary, Accent colors
- Error, Warning, Info, Success colors
- Visual examples with hex codes

**Tab 5: Dynamic Panels**
- LazyGit-style accordion panels (3-panel layout)
- Weight-based layout system (1:1 = 50/50, 2:1 = 66/33)
- Interactive: Click panels to focus OR press 1/2/3
- Press 'a' to toggle accordion mode
- Watch panels resize dynamically!

### Component Showcase Tabs (6-11)

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
