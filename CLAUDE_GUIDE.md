# Guide for Claude Code - Building TUI Apps with TUITemplate

This guide helps Claude Code efficiently build TUI applications using TUITemplate.

## Quick Context

When working with a TUITemplate-based project, Claude should know:

1. **Tech Stack:**
   - Go with Bubbletea (TUI framework)
   - Lipgloss (styling)
   - Bubbles (components)

2. **Architecture Pattern:**
   - `main.go` - Entry point ONLY (keep minimal, ~20 lines)
   - `types.go` - All type definitions and structs
   - `model.go` - Model initialization and layout
   - `update.go` - Message dispatcher
   - `update_keyboard.go` - Keyboard event handling
   - `update_mouse.go` - Mouse event handling
   - `view.go` - View rendering
   - `styles.go` - Lipgloss style definitions
   - `config.go` - Configuration management

3. **Key Principles:**
   - One file, one responsibility
   - Never add business logic to `main.go`
   - Keep files under 800 lines (ideally <500)
   - Separate keyboard and mouse handling
   - Use message-based communication

## Available Research

Location: `docs/research/`

**Quick references:**
- `ECOSYSTEM_QUICK_REFERENCE.md` - 96 lines, Bubbletea libraries
- `TUI_TOOLS_QUICK_REFERENCE.md` - 273 lines, 80+ TUI tools

**Detailed research:**
- `ECOSYSTEM_RESEARCH_2025.md` - 989 lines, comprehensive Charm ecosystem guide
- `TUI_APPLICATIONS_RESEARCH.md` - 1,249 lines, 80+ TUI tool descriptions

## Common Tasks

### Creating a New App

```bash
cd TUITemplate
./scripts/new_project.sh
```

Prompts for: name, title, description, author, layout, components

### Adding Application State

**In types.go:**
```go
type model struct {
    config Config
    width  int
    height int

    // Add your state here
    items  []Item
    cursor int
}

type Item struct {
    // Your data structure
}
```

### Handling User Input

**In update_keyboard.go:**
```go
func (m model) handleMainKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "enter":
        return m.selectItem()
    case "d":
        return m.deleteItem()
    }
}
```

### Rendering UI

**In view.go:**
```go
func (m model) renderMainContent() string {
    var content strings.Builder

    for i, item := range m.items {
        if i == m.cursor {
            content.WriteString(selectedStyle.Render(item.Name))
        } else {
            content.WriteString(item.Name)
        }
        content.WriteString("\n")
    }

    return content.String()
}
```

### Loading Data Asynchronously

**In types.go:**
```go
type dataLoadedMsg struct {
    data []Item
}
```

**In model.go:**
```go
func loadDataCmd() tea.Cmd {
    return func() tea.Msg {
        data := fetchData() // Your data loading logic
        return dataLoadedMsg{data}
    }
}
```

**In update.go:**
```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case dataLoadedMsg:
        m.items = msg.data
        return m, nil
    }
}
```

**In model.go Init():**
```go
func (m model) Init() tea.Cmd {
    return loadDataCmd()
}
```

## Component Library

Available components (in `components/`):

- **panel/** - Layout management (single, dual, multi, tabbed)
- **list/** - List views (simple, filtered, tree)
- **input/** - Input widgets (text, multiline, form, autocomplete)
- **dialog/** - Dialogs (confirm, input, progress, modal)
- **menu/** - Menus (context, command palette, menubar)
- **status/** - Status displays (statusbar, titlebar, breadcrumb)
- **preview/** - Preview panes (text, markdown, syntax, image, hex)
- **table/** - Tables (simple, interactive)

## Common Libraries to Suggest

When user needs specific functionality:

**Syntax Highlighting:**
```go
import "github.com/alecthomas/chroma/v2"
```

**Markdown Rendering:**
```go
import "github.com/charmbracelet/glamour"
```

**Forms:**
```go
import "github.com/charmbracelet/huh"
```

**Tables:**
```go
import "github.com/evertras/bubble-table"
```

**Fuzzy Finding:**
```go
import "github.com/koki-develop/go-fzf"
```

See `docs/research/ECOSYSTEM_QUICK_REFERENCE.md` for complete list.

## Integration with Existing Tools

For file-specific operations, suggest TUI tools:

| File Type | Tool | Package |
|-----------|------|---------|
| JSON/YAML | fx | `go install github.com/antonmedv/fx@latest` |
| CSV/Excel | visidata | `pip install visidata` |
| Hex/Binary | HexPatch | `go install github.com/Zenithar/hexapatch@latest` |
| Images | timg | Via package manager |
| SQL | harlequin | `pip install harlequin` |

See `docs/research/TUI_TOOLS_QUICK_REFERENCE.md` for 80+ tools.

## Performance Tips

For large datasets (1000+ items):

**Virtual Scrolling:**
```go
visibleStart := max(0, m.cursor - m.height/2)
visibleEnd := min(len(m.items), visibleStart + m.height)
visibleItems := m.items[visibleStart:visibleEnd]

for _, item := range visibleItems {
    // Render only visible items
}
```

**Async Operations:**
```go
func expensiveOperationCmd() tea.Cmd {
    return func() tea.Msg {
        result := doExpensiveWork()
        return resultMsg{result}
    }
}
```

**Caching:**
```go
if cached, ok := m.cache[key]; ok {
    return cached
}
result := computeExpensiveValue()
m.cache[key] = result
```

## Testing Pattern

```go
import "github.com/charmbracelet/x/exp/teatest"

func TestNavigation(t *testing.T) {
    m := initialModel(getDefaultConfig())
    tm := teatest.NewTestModel(t, m)

    tm.Send(tea.KeyMsg{Type: tea.KeyDown})

    if tm.Get().(model).cursor != 1 {
        t.Error("cursor should move down")
    }
}
```

## Common Pitfalls to Avoid

1. ❌ Adding logic to `main.go` - Keep it minimal
2. ❌ Files over 800 lines - Split into modules
3. ❌ Mixing keyboard/mouse in one file - Separate concerns
4. ❌ Rendering all items - Use virtual scrolling for large lists
5. ❌ Blocking operations in Update() - Use commands for async
6. ❌ Not handling window resize - Always implement `tea.WindowSizeMsg`

## Debugging

Enable logging in config:
```yaml
logging:
  enabled: true
  level: "debug"
  file: "~/.local/share/app/debug.log"
```

Then tail the log:
```bash
tail -f ~/.local/share/app/debug.log
```

## File Operations Reference

**Reading files:**
```go
content, err := os.ReadFile(path)
```

**Writing files:**
```go
err := os.WriteFile(path, []byte(content), 0644)
```

**Directory listing:**
```go
entries, err := os.ReadDir(path)
for _, entry := range entries {
    if entry.IsDir() {
        // Directory
    } else {
        // File
    }
}
```

**File watching:**
```go
import "github.com/fsnotify/fsnotify"

watcher, _ := fsnotify.NewWatcher()
watcher.Add(path)

for {
    select {
    case event := <-watcher.Events:
        // Handle file change
    }
}
```

## When to Use What Layout

**Single Pane:**
- Simple lists
- Forms
- Menus
- Focus on one task

**Dual Pane:**
- File browsers with preview
- List + detail view
- Source + preview
- Navigation + content

**Multi-Panel:**
- Dashboards
- Multiple data sources
- Comparison views
- Complex layouts

**Tabbed:**
- Multiple independent views
- Settings screens
- Multi-document interfaces

## Resources

**Within TUITemplate:**
- `README.md` - Quick start
- `USAGE.md` - Comprehensive usage guide
- `ARCHITECTURE.md` - Full architecture documentation
- `examples/` - Working examples
- `docs/research/` - Research on tools and libraries

**External:**
- [Bubbletea Docs](https://github.com/charmbracelet/bubbletea)
- [Lipgloss Docs](https://github.com/charmbracelet/lipgloss)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)

## Quick Checklist for New Apps

When creating a new TUI app:

- [ ] Run `./scripts/new_project.sh`
- [ ] Define data structures in `types.go`
- [ ] Initialize state in `model.go`
- [ ] Implement rendering in `view.go`
- [ ] Add keyboard shortcuts in `update_keyboard.go`
- [ ] Handle messages in `update.go`
- [ ] Style with Lipgloss in `styles.go`
- [ ] Add configuration options in `config.go`
- [ ] Test with various terminal sizes
- [ ] Add error handling
- [ ] Create README with usage instructions

---

**This guide provides Claude with all context needed to efficiently build TUI applications following TUITemplate's architecture and best practices.**
