# TUITemplate Examples

Complete, production-ready demo applications showcasing advanced TUI patterns. These are full-featured reference implementations.

## What Are Examples?

Examples are **comprehensive applications** that demonstrate multiple TUI patterns working together. Each example:
- Is a complete, working application
- Combines multiple concepts and patterns
- Shows production-ready code organization
- Serves as a reference for real-world projects

**New to TUI development?** Start with [../starters/](../starters/) for focused tutorials.

---

## Available Examples

### 1. TUI Showcase (`tui-showcase/`)

**Complexity:** 🔴 Advanced
**Lines of Code:** ~2500
**Status:** ✅ Complete

**The Comprehensive TUI Pattern Library**

A single, unified application demonstrating **ALL** TUI patterns from basic layouts to advanced components. Navigate using dropdown menus or keyboard shortcuts to explore 12 different showcases.

**What It Demonstrates:**

**Layout Modes (Press 1-4):**
- **Single Pane** - Basic single-panel layout
- **Dual Pane** - Side-by-side panels with divider
- **Multi-Panel** - 3-panel layout (top-left, top-right, bottom)
- **Tabbed** - Tab-based navigation with 12 different tabs

**Tabbed Showcases (Tab/Shift+Tab or Components menu):**

**Core Tabs (0-5):**
- **Overview** - Introduction and navigation guide
- **Content** - Content area demonstration
- **Settings** - Settings panel example
- **Borders** - Lipgloss border style showcase (Rounded, Double, Thick, Normal)
- **Colors** - Theme color palette demonstration
- **Dynamic Panels** - LazyGit-style accordion with weight-based sizing

**Component Tabs (6-11):**
- **Forms** - Text inputs, checkboxes, radio buttons, validation
- **Tables** - Data grids with sorting, filtering, row selection
- **Dialogs** - Modal dialogs, confirmations, alerts, prompts
- **Progress** - Progress bars (determinate/indeterminate), spinners
- **Tree View** - Expandable/collapsible hierarchical data with file browser
- **Mobile** - Termux-optimized patterns, vertical stacking, touch-friendly UI

**Key Features:**
- **Dropdown Menu System** - File/View/Components/Help menus with mouse support
- **Mouse Support** - Click menus, tabs, panels, all interactive elements
- **Keyboard Navigation** - Full keyboard shortcuts for all actions
- **Responsive Design** - Adapts to terminal size (< 80 cols = vertical stack)
- **Professional UI** - Clean rendering, no visual corruption, proper borders
- **Weight-Based Layouts** - Dynamic panel resizing (1:1 = 50/50, 2:1 = 66/33)

**Navigation:**
```
Menus:   Click menu bar or use mouse
Layouts: Press 1 (single), 2 (dual), 3 (multi), 4 (tabbed)
Tabs:    Tab/Shift+Tab or Components menu
Help:    Press ? for keyboard shortcuts
Quit:    Press q
```

**Run:**
```bash
cd tui-showcase
go run .
```

**Perfect for:** Understanding ALL TUI patterns in one comprehensive application. This is your complete reference implementation.

---

### 2. Termux Mobile Demo (`termux-mobile-demo/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~700
**Status:** ✅ Complete

**What It Demonstrates:**
- **Mobile-optimized UI** - Designed for portrait phone screens
- **Touch-friendly controls** - Larger touch targets
- **Portrait/landscape adaptation** - Responsive layout switching
- **Vertical stacking** - Panels stack when width < 80
- **Termux compatibility** - Tested on Android Termux

**Key Optimizations:**
- Larger buttons for touch accuracy
- Vertical panel stacking for portrait mode
- Simplified navigation for small screens
- Termux-specific workarounds

**Run:**
```bash
cd termux-mobile-demo
go run .
```

**Perfect for:** Building TUIs for mobile terminals (Termux, SSH on phone).

---

### 3. Termux Size Detector (`termux-size-detector/`)

**Complexity:** 🟢 Simple
**Lines of Code:** ~100
**Status:** ✅ Complete (Utility)

**What It Demonstrates:**
- **Terminal size detection** - Real-time width/height display
- **Resize handling** - Updates on terminal resize
- **Termux debugging** - Verify terminal dimensions
- **Window size messages** - Bubbletea resize event handling

**Use Case:**
Debug terminal size issues on Termux or any platform. Shows exactly what dimensions your TUI sees.

**Run:**
```bash
cd termux-size-detector
go run .
```

**Perfect for:** Debugging layout issues and understanding terminal dimensions.

---

## Running an Example

Each example is self-contained:

```bash
cd <example-name>
go mod tidy  # Download dependencies
go run .     # Run the application
```

Example:
```bash
cd layout-demo
go run .
```

---

## What You'll Learn from Examples

### Advanced Patterns

**Dropdown Menus** (tui-showcase)
- Menu bar rendering and positioning
- Dropdown overlay without corruption
- Mouse hover auto-switching
- Keyboard navigation (arrows, enter, esc)

**Weight-Based Layouts** (tui-showcase - Dynamic Panels tab)
- Proportional panel sizing
- Focus-based resizing (accordion mode)
- Instant layout updates
- Clean, maintainable layout code

**Component Showcases** (tui-showcase - Component tabs)
- Forms with inputs and validation
- Tables with sorting and filtering
- Dialogs and modal windows
- Progress indicators and spinners
- Tree views with hierarchical data

**Mobile Optimization** (tui-showcase - Mobile tab, termux-mobile-demo)
- Portrait vs landscape detection
- Vertical stacking for narrow screens
- Touch-friendly UI elements
- Termux-specific considerations

### Production-Ready Code

All examples demonstrate:
- **Modular file organization** - Separate files for concerns
- **Proper state management** - Immutable updates
- **Mouse + keyboard support** - Complete input handling
- **Error handling** - Graceful degradation
- **Performance** - Efficient rendering
- **Documentation** - Code comments and READMEs

---

## Architecture Patterns

### File Organization
```
example/
├── main.go              # Entry point, program setup
├── types.go             # Type definitions, structs
├── model.go             # Model init, layout calculations
├── update.go            # Main update dispatcher
├── update_keyboard.go   # Keyboard event handling
├── update_mouse.go      # Mouse event handling
├── view.go              # View rendering
├── styles.go            # Lipgloss styles
├── config.go            # Configuration loading
└── menu.go              # (Optional) Menu system
```

### Key Concepts

**Separation of Concerns:**
- **types.go** - Data structures
- **model.go** - Business logic
- **update.go** - Event handling
- **view.go** - Rendering
- **styles.go** - Visual styling

**State Management:**
- Immutable model updates
- Message-based communication
- Command pattern for async ops

**Layout Calculation:**
- Account for borders: `height -= 2`
- Use Lipgloss.Width() for accurate sizing
- Truncate text to prevent wrapping

---

## Comparison: Starters vs Examples

| Aspect | Starters | Examples |
|--------|----------|----------|
| **Purpose** | Learn ONE concept | Reference implementation |
| **Complexity** | Simple, focused | Complete, real-world |
| **Size** | 80-600 lines | 600-2000+ lines |
| **Scope** | Single pattern | Multiple patterns combined |
| **Use Case** | Educational | Production reference |
| **Status** | Mostly TODO | All complete |

**When to use:**
- **Learning?** → Use [../starters/](../starters/)
- **Building?** → Use examples/ as reference

---

## Building Your Own App

After studying the examples:

### 1. Create Your Project
```bash
# Return to TUITemplate root
cd ..

# Generate new project
./scripts/new_project.sh

# Choose components and layout
```

### 2. Reference Examples
- **Need menus?** → Study `tui-showcase/menu.go`
- **Need panels?** → Study `tui-showcase/model.go` (Dynamic Panels)
- **Need components?** → Study `tui-showcase/view_components.go`
- **Need mobile support?** → Study `tui-showcase/` (Mobile tab) or `termux-mobile-demo/`
- **Debugging layout?** → Use `termux-size-detector/`

### 3. Read Documentation
- [CLAUDE.md](../CLAUDE.md) - Critical bug fixes and patterns
- [LAZYGIT_ANALYSIS.md](../docs/research/LAZYGIT_ANALYSIS.md) - Weight-based layouts
- [MOUSE_SUPPORT_GUIDE.md](../docs/research/MOUSE_SUPPORT_GUIDE.md) - Mouse handling

---

## Contributing Examples

Have a cool TUI app to share?

**Guidelines:**
- Should be a **complete application**
- Demonstrates **multiple patterns** working together
- Includes **comprehensive README** with screenshots
- Has **inline comments** explaining key decisions
- Follows **TUITemplate architecture**
- Is **well-tested** and production-ready

**Submit:**
1. Fork the repository
2. Add your example to `examples/`
3. Include README with description and usage
4. Update this README
5. Submit a pull request

---

## Resources

- **TUITemplate:** [README.md](../README.md)
- **Usage Guide:** [USAGE.md](../USAGE.md)
- **Architecture:** [ARCHITECTURE.md](../ARCHITECTURE.md)
- **Starters:** [../starters/](../starters/) - Simple tutorials
- **Research:** [../docs/research/](../docs/research/) - Advanced patterns
- **Bubbletea:** https://github.com/charmbracelet/bubbletea
- **Lipgloss:** https://github.com/charmbracelet/lipgloss

---

## Quick Reference

### Dropdown Menus
See: `tui-showcase/menu.go`

### Weight-Based Panels
See: `tui-showcase/model.go:106-190` (Dynamic Panels tab)

### Component Showcases
See: `tui-showcase/view_components.go` (Forms, Tables, Dialogs, Progress, Tree, Mobile)

### Border Height Fix
```go
// CRITICAL: Account for panel borders
contentHeight -= 2 // top + bottom borders
```
See: `CLAUDE.md` for full explanation

### Mouse Click Detection
See: `tui-showcase/update_mouse.go`

### Mobile Optimization
See: `tui-showcase/view_components.go:330+` (Mobile tab) or `termux-mobile-demo/view.go`

---

**Study these examples, then build something amazing!** 🚀
