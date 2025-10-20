# TUITemplate Starters

Simple, focused tutorial examples for learning TUI development. Start here if you're new to Bubbletea and TUITemplate!

## What Are Starters?

Starters are **small, single-concept examples** designed to teach specific TUI patterns. Each starter:
- Focuses on ONE main concept
- Is easy to understand and modify
- Builds progressively in complexity
- Includes inline comments explaining key decisions

**Looking for complete applications?** Check out [../examples/](../examples/) for full-featured demos.

---

## Available Starters

### 1. Hello World (`hello/`)

**Complexity:** 🟢 Beginner
**Lines of Code:** ~80
**Status:** ✅ Ready to use

**What You'll Learn:**
- Minimal viable TUI application structure
- Basic architecture (main.go + model.go)
- Simple state management
- Keyboard input handling (press any key)
- Lipgloss styling basics

**Run:**
```bash
cd hello
go mod tidy
go run .
```

**Perfect for:** Your very first TUI app. Start here!

---

### 2. File Browser (`file_browser/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~300
**Status:** 🚧 Coming soon

**What You'll Learn:**
- File system navigation
- List rendering with cursor selection
- Keyboard navigation (up/down/enter/back)
- Dual-pane layout (list + preview)
- Handling user input for navigation

**Use cases:** File managers, directory pickers, navigation menus

---

### 3. Todo App (`todo_app/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~400
**Status:** 🚧 Coming soon

**What You'll Learn:**
- CRUD operations (Create, Read, Update, Delete)
- Multiple views (list/add/edit modes)
- Data persistence (saving to JSON file)
- Forms and input dialogs
- State management patterns

**Use cases:** Task managers, note apps, any CRUD application

---

### 4. Log Viewer (`log_viewer/`)

**Complexity:** 🟡 Intermediate
**Lines of Code:** ~350
**Status:** 🚧 Coming soon

**What You'll Learn:**
- Real-time file tailing
- Filtering and search functionality
- Syntax highlighting
- Scrolling viewport implementation
- Performance optimization for large files

**Use cases:** Log viewers, monitoring tools, text readers

---

### 5. JSON Viewer (`json_viewer/`)

**Complexity:** 🔴 Advanced
**Lines of Code:** ~500
**Status:** 🚧 Coming soon

**What You'll Learn:**
- Dual-pane layout (tree + detail)
- Tree navigation (expandable/collapsible nodes)
- Syntax highlighting
- Search functionality
- Multiple file format support (JSON, YAML, TOML)

**Use cases:** Config viewers, API explorers, data inspectors

---

### 6. Dashboard (`dashboard/`)

**Complexity:** 🔴 Advanced
**Lines of Code:** ~600
**Status:** 🚧 Coming soon

**What You'll Learn:**
- Multi-panel layout composition
- Real-time data updates
- Charts and graphs (ASCII art)
- Tabbed interface
- Component composition patterns

**Use cases:** System monitors, analytics dashboards, admin panels

---

## Learning Path

### 🟢 Beginners (New to TUI development)

1. **Start with `hello/`**
   - Build and run it
   - Read all the code comments
   - Modify the text and colors
   - Try adding your own key bindings

2. **Read the docs**
   - [TUITemplate Architecture](../ARCHITECTURE.md)
   - [Usage Guide](../USAGE.md)
   - [Bubbletea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)

3. **Experiment**
   - Add a counter that increments on key press
   - Add multiple "pages" that switch on different keys
   - Style it with different Lipgloss themes

### 🟡 Intermediate (Understand basics)

1. **Study `file_browser/`** (when ready)
   - Learn list rendering patterns
   - Understand cursor navigation
   - See how to structure dual-pane layouts

2. **Explore `todo_app/`** (when ready)
   - Learn state management for CRUD
   - Understand view switching
   - See data persistence patterns

3. **Try `log_viewer/`** (when ready)
   - Learn viewport scrolling
   - Understand performance techniques
   - See real-time update patterns

### 🔴 Advanced (Ready for complex patterns)

1. **Analyze `json_viewer/`** (when ready)
   - Complex navigation logic
   - Tree data structures
   - Advanced layout composition

2. **Study `dashboard/`** (when ready)
   - Multi-component orchestration
   - Real-time data flow
   - Advanced state management

3. **Build your own**
   - Use `./scripts/new_project.sh`
   - Check out [../examples/](../examples/) for reference

---

## Running a Starter

Each starter is self-contained:

```bash
cd <starter-name>
go mod tidy  # Download dependencies
go run .     # Run the app
```

Example:
```bash
cd hello
go mod tidy
go run .
```

---

## Code Patterns You'll Learn

### State Management
- **Immutable updates** - Never mutate state directly
- **Message-based communication** - Use Bubbletea messages
- **Command pattern** - Async operations return Cmds

### UI Patterns
- **List with cursor** - Up/down navigation with selection
- **Tree navigation** - Expand/collapse hierarchical data
- **Dual-pane layout** - Split screen with focus switching
- **Modal dialogs** - Overlays for user input
- **Status bars** - Fixed headers/footers

### Performance
- **Virtual scrolling** - Only render visible items
- **Lazy loading** - Load data on demand
- **Caching** - Store computed values
- **Debouncing** - Throttle rapid updates

### Architecture
- **Modular files** - Separate concerns (update, view, model)
- **Component composition** - Build complex UIs from simple parts
- **Reusable utilities** - DRY (Don't Repeat Yourself)

---

## After Learning

Once you've completed the starters:

### Check Out Full Examples
See [../examples/](../examples/) for complete, production-ready applications:
- `layout-demo` - Multiple layouts with dropdown menus
- `dynamic-panels-demo` - LazyGit-style resizable panels
- `termux-mobile-demo` - Mobile-optimized TUI

### Build Your Own App
```bash
# Return to TUITemplate root
cd ..

# Create your project
./scripts/new_project.sh

# Choose components and start coding!
```

### Dive Deeper
- Read [Research Documentation](../docs/research/) for advanced patterns
- Study [LAZYGIT_ANALYSIS.md](../docs/research/LAZYGIT_ANALYSIS.md) for weight-based layouts
- Explore [MOUSE_SUPPORT_GUIDE.md](../docs/research/MOUSE_SUPPORT_GUIDE.md) for mouse handling

---

## Contributing Starters

Have an idea for a new starter example?

**Guidelines:**
- Keep it **focused** on ONE concept
- Make it **simple** and easy to understand
- Add **inline comments** explaining decisions
- Include a **README** with description
- Follow **TUITemplate architecture** patterns
- Keep it **under 500 lines** total

**Submit:**
1. Fork the repo
2. Add your starter to `starters/`
3. Update this README
4. Submit a pull request

---

## Need Help?

- **TUITemplate Docs:** [README.md](../README.md)
- **Usage Guide:** [USAGE.md](../USAGE.md)
- **Architecture:** [ARCHITECTURE.md](../ARCHITECTURE.md)
- **Bubbletea Docs:** https://github.com/charmbracelet/bubbletea
- **Lipgloss Docs:** https://github.com/charmbracelet/lipgloss

---

**Start small, build big!** 🚀

Begin with `hello/` and work your way through. Each starter builds on concepts from the previous ones.
