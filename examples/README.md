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

### 1. Layout Demo (`layout-demo/`)

**Complexity:** 🔴 Advanced
**Lines of Code:** ~2000
**Status:** ✅ Complete

**What It Demonstrates:**
- **Dropdown menu system** - File/View/Components/Help menus
- **Multiple layouts** - Single pane, dual pane, multi-panel, tabbed
- **Mouse support** - Click menus, buttons, tabs, panels
- **Keyboard navigation** - Full menu and layout keyboard shortcuts
- **Border showcase** - Different Lipgloss border styles
- **Color palette** - Theme color demonstrations
- **Dynamic panels** - Accordion-style resizing (LazyGit pattern)

**Key Features:**
- Professional menu bar with hover effects
- Clean dropdown rendering without visual corruption
- Accurate mouse click detection
- Tab switching with visual feedback
- 3-panel dynamic layout with weight-based sizing

**Run:**
```bash
cd layout-demo
go run .
```

**Perfect for:** Understanding complete layout systems, dropdown menus, and professional UI organization.

---

### 2. Dynamic Panels Demo (`dynamic-panels-demo/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~800
**Status:** ✅ Complete

**What It Demonstrates:**
- **LazyGit-style accordion panels** - Focused panel gets 2x space
- **Weight-based layout system** - Proportional sizing (1:1 = 50/50, 2:1 = 66/33)
- **Mouse panel focus** - Click panels to focus them
- **Keyboard focus switching** - Press 1/2 to switch focus
- **Accordion toggle** - Press 'a' to enable/disable accordion mode
- **Clean resizing** - Instant, proportional panel resizing

**Key Patterns:**
```go
// Weight-based sizing
leftWeight, rightWeight := 1, 1
if accordionMode && focusedPanel == "left" {
    leftWeight = 2  // Focused gets 66%
}
leftWidth = (totalWidth * leftWeight) / (leftWeight + rightWeight)
```

**Run:**
```bash
cd dynamic-panels-demo
go run .
```

**Perfect for:** Understanding LazyGit-style layouts and weight-based panel systems.

---

### 3. Multi-Panel Example (`multi-panel/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~600
**Status:** ✅ Complete

**What It Demonstrates:**
- **3-panel layout** - Top-left, top-right, bottom (full-width)
- **Panel borders** - Rounded borders with proper height accounting
- **Content truncation** - Preventing text wrapping issues
- **Mouse click detection** - Clicking panels with coordinate detection
- **Vertical stacking** - Responsive layout for narrow terminals

**Critical Patterns:**
- Border height accounting: `contentHeight -= 2`
- Text truncation to prevent wrapping
- Relative coordinate calculations for clicks

**Run:**
```bash
cd multi-panel
go run .
```

**Perfect for:** Multi-panel layouts and proper border handling.

---

### 4. Termux Mobile Demo (`termux-mobile-demo/`)

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

### 5. Termux Size Detector (`termux-size-detector/`)

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

**Dropdown Menus** (layout-demo)
- Menu bar rendering and positioning
- Dropdown overlay without corruption
- Mouse hover auto-switching
- Keyboard navigation (arrows, enter, esc)

**Weight-Based Layouts** (dynamic-panels-demo)
- Proportional panel sizing
- Focus-based resizing (accordion mode)
- Instant layout updates
- Clean, maintainable layout code

**Border Height Accounting** (multi-panel)
- Critical: `contentHeight -= 2` for borders
- Preventing panel overflow
- Text truncation to avoid wrapping
- Proper panel alignment

**Mobile Optimization** (termux-mobile-demo)
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
- **Need menus?** → Study `layout-demo/menu.go`
- **Need panels?** → Study `dynamic-panels-demo/model.go`
- **Need mobile support?** → Study `termux-mobile-demo/`
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
See: `layout-demo/menu.go`

### Weight-Based Panels
See: `dynamic-panels-demo/model.go:103-144`

### Border Height Fix
```go
// CRITICAL: Account for panel borders
contentHeight -= 2 // top + bottom borders
```
See: `CLAUDE.md` for full explanation

### Mouse Click Detection
See: `layout-demo/update_mouse.go`

### Mobile Optimization
See: `termux-mobile-demo/view.go`

---

**Study these examples, then build something amazing!** 🚀
