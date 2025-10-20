# TUITemplate Examples

Example applications demonstrating TUITemplate usage.

## Available Examples

### 1. Hello World (`hello/`)

**Complexity:** Beginner
**Lines of Code:** ~80
**Demonstrates:**
- Minimal viable TUI application
- Basic architecture (main.go + model.go)
- Simple state management
- Keyboard input handling
- Lipgloss styling

**Run:**
```bash
cd hello && go run .
```

### 2. File Browser (`file_browser/`)

**Complexity:** Intermediate
**Lines of Code:** ~300
**Demonstrates:**
- File system navigation
- List rendering with selection
- Keyboard navigation (up/down/enter)
- Dual-pane layout
- File preview

**Status:** 🚧 Coming soon

### 3. Todo App (`todo_app/`)

**Complexity:** Intermediate
**Lines of Code:** ~400
**Demonstrates:**
- CRUD operations
- Multiple views (list/add/edit)
- Data persistence (JSON file)
- Forms and input dialogs
- State management patterns

**Status:** 🚧 Coming soon

### 4. Log Viewer (`log_viewer/`)

**Complexity:** Intermediate
**Lines of Code:** ~350
**Demonstrates:**
- Real-time file tailing
- Filtering and search
- Syntax highlighting
- Scrolling viewport
- Performance optimization

**Status:** 🚧 Coming soon

### 5. JSON Viewer (`json_viewer/`)

**Complexity:** Advanced
**Lines of Code:** ~500
**Demonstrates:**
- Dual-pane layout
- Tree navigation (expandable/collapsible)
- Syntax highlighting
- Search functionality
- Multiple file format support (JSON, YAML, TOML)

**Status:** 🚧 Coming soon

### 6. Dashboard (`dashboard/`)

**Complexity:** Advanced
**Lines of Code:** ~600
**Demonstrates:**
- Multi-panel layout
- Real-time data updates
- Charts and graphs (ASCII)
- Tabbed interface
- Component composition

**Status:** 🚧 Coming soon

## Running Examples

Each example is self-contained:

```bash
cd <example-name>
go mod tidy
go run .
```

## Learning Path

**Beginners:**
1. Start with `hello/` - understand the basic structure
2. Read the code comments
3. Modify and experiment

**Intermediate:**
1. Study `file_browser/` - learn navigation patterns
2. Explore `todo_app/` - understand state management
3. Try `log_viewer/` - see performance techniques

**Advanced:**
1. Analyze `json_viewer/` - complex layout and navigation
2. Study `dashboard/` - multi-component composition
3. Create your own app with `./scripts/new_project.sh`

## Code Patterns Demonstrated

### State Management
- Immutable updates
- Message-based communication
- Async command pattern

### UI Patterns
- List with cursor selection
- Tree with expand/collapse
- Dual-pane with focus switching
- Modal dialogs
- Status bars and breadcrumbs

### Performance
- Virtual scrolling
- Lazy loading
- Caching strategies
- Debouncing input

### Architecture
- Modular file organization
- Component composition
- Separation of concerns
- Reusable utilities

## Building Your Own

After exploring the examples:

```bash
# Return to TUITemplate root
cd ../..

# Create your project
./scripts/new_project.sh

# Choose components and layout
# Start coding!
```

## Contributing Examples

Have a cool TUI app built with TUITemplate?

1. Fork the repo
2. Add your example to `examples/`
3. Include a README with screenshots
4. Submit a pull request

Guidelines:
- Keep examples focused on one concept
- Add inline comments explaining key decisions
- Include a README with description and screenshots
- Keep code simple and readable
- Follow TUITemplate's architecture patterns

## Resources

- [TUITemplate Documentation](../../README.md)
- [Usage Guide](../../USAGE.md)
- [Research Documentation](../../docs/research/)
- [Bubbletea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)

---

**Start with `hello/` and work your way up!** 🚀
