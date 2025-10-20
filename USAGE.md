# TUITemplate Usage Guide

Complete guide to using TUITemplate for building TUI applications.

## Table of Contents

1. [Creating a New Project](#creating-a-new-project)
2. [Understanding the Template Structure](#understanding-the-template-structure)
3. [Implementing Your Application](#implementing-your-application)
4. [Working with Components](#working-with-components)
5. [Customization](#customization)
6. [Common Patterns](#common-patterns)
7. [Testing](#testing)
8. [Deployment](#deployment)

---

## Creating a New Project

### Interactive Method (Recommended)

```bash
cd TUITemplate
./scripts/new_project.sh
```

You'll be prompted for:

| Prompt | Example | Description |
|--------|---------|-------------|
| App name | `json-viewer` | Lowercase, hyphenated |
| Title | `JSON Viewer` | Display name |
| Description | `Beautiful JSON viewer` | Short description |
| Author | `Matt` | Your name |
| License | `MIT` | License type |
| Layout | `dual_pane` | single/dual_pane/multi_panel/tabbed |
| Components | `panel,list,preview` | Comma-separated list |

The script will:
1. Create `../json-viewer/` directory
2. Copy and process template files
3. Replace `{{.AppName}}`, `{{.AppTitle}}`, etc.
4. Copy selected components
5. Initialize go.mod
6. Create example config
7. Generate README.md

### Manual Method

```bash
mkdir my-app
cd my-app

# Copy template files manually
cp ../TUITemplate/template/*.tmpl .

# Rename .tmpl files
for f in *.tmpl; do mv "$f" "${f%.tmpl}"; done

# Manually replace {{.AppName}} etc. in all files

# Initialize Go module
go mod init github.com/yourname/my-app
go mod tidy
```

---

## Understanding the Template Structure

### File Responsibilities

**main.go** (21 lines)
- Entry point ONLY
- Creates Bubbletea program
- Configures options (mouse, alt screen)
- Runs program loop

```go
func main() {
    p := tea.NewProgram(
        initialModel(loadConfig()),
        tea.WithAltScreen(),
        tea.WithMouseCellMotion(),
    )
    p.Run()
}
```

**types.go** (~150 lines)
- All type definitions
- Application state (`model` struct)
- Configuration structs
- Custom message types

```go
type model struct {
    config Config
    width  int
    height int
    // Your app state here
}

type itemSelectedMsg struct {
    item Item
}
```

**model.go** (~100 lines)
- `initialModel()` - creates initial state
- Layout calculation functions
- Helper functions for model operations

```go
func initialModel(cfg Config) model {
    return model{
        config: cfg,
        // Initialize your state
    }
}
```

**update.go** (~80 lines)
- `Init()` - initialization commands
- `Update()` - main message dispatcher
- Delegates to keyboard/mouse handlers
- Handles window resize, custom messages

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        return m.handleKeyPress(msg)
    case tea.MouseMsg:
        return m.handleMouseEvent(msg)
    // Your custom messages
    }
}
```

**update_keyboard.go** (~200 lines)
- All keyboard input handling
- Global keybindings (quit, help)
- Mode-specific keybindings
- Navigation helpers

**update_mouse.go** (~150 lines)
- All mouse input handling
- Click detection (left, right, double)
- Scroll handling
- Region detection

**view.go** (~200 lines)
- Main `View()` function
- Layout rendering (single/dual/multi)
- Component rendering
- Helper functions

**styles.go** (~150 lines)
- All Lipgloss style definitions
- Color palette
- Theme application
- Style helpers

**config.go** (~200 lines)
- Configuration loading
- YAML parsing
- Default values
- Config validation

---

## Implementing Your Application

### Step 1: Define Your Data

**In types.go:**

```go
// Add to model struct
type model struct {
    // ... existing fields

    // Your data
    items    []Item
    cursor   int
    selected map[string]bool
    filter   string
}

// Define your types
type Item struct {
    ID      string
    Name    string
    Value   string
    Created time.Time
}

// Custom messages
type itemsLoadedMsg struct {
    items []Item
}

type itemSelectedMsg struct {
    item Item
}
```

### Step 2: Initialize State

**In model.go:**

```go
func initialModel(cfg Config) model {
    return model{
        config:   cfg,
        items:    loadItems(),      // Load your data
        cursor:   0,
        selected: make(map[string]bool),
        filter:   "",
    }
}

// Helper function
func loadItems() []Item {
    // Load from file, database, API, etc.
    return []Item{}
}
```

### Step 3: Handle Events

**In update.go:**

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    // ... existing cases

    case itemsLoadedMsg:
        m.items = msg.items
        return m, nil

    case itemSelectedMsg:
        m.selected[msg.item.ID] = true
        return m, nil
    }
}
```

**In update_keyboard.go:**

```go
func (m model) handleMainKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "enter":
        if m.cursor < len(m.items) {
            return m, selectItem(m.items[m.cursor])
        }

    case "d":
        return m, deleteItem(m.items[m.cursor])

    case "/":
        m.filterMode = true
        return m, nil
    }
}
```

### Step 4: Render UI

**In view.go:**

```go
func (m model) renderMainContent() string {
    var content strings.Builder

    // Title
    content.WriteString("Items\n\n")

    // List
    for i, item := range m.items {
        cursor := " "
        if i == m.cursor {
            cursor = ">"
        }

        selected := " "
        if m.selected[item.ID] {
            selected = "✓"
        }

        line := fmt.Sprintf("%s %s %s\n", cursor, selected, item.Name)

        if i == m.cursor {
            content.WriteString(selectedStyle.Render(line))
        } else {
            content.WriteString(line)
        }
    }

    return content.String()
}
```

---

## Working with Components

### Using Built-in Components

**List Component:**

```go
import "github.com/yourname/TUITemplate/components/list"

// In model
type model struct {
    list list.SimpleList
}

// In initialModel
list := list.NewSimpleList(items, width, height)

// In Update
case tea.KeyMsg:
    m.list, cmd = m.list.Update(msg)

// In View
return m.list.View()
```

### Creating Custom Components

Components should implement:

```go
type Component interface {
    Init() tea.Cmd
    Update(tea.Msg) (Component, tea.Cmd)
    View() string
    SetSize(width, height int)
}
```

---

## Customization

### Themes

**Built-in themes:**
- `dark` (default)
- `light`
- `solarized`
- `dracula`
- `nord`

**Custom theme:**

```yaml
theme: "custom"
custom_theme:
  primary: "#FF6B6B"
  secondary: "#4ECDC4"
  background: "#1A1A2E"
  foreground: "#E8E8E8"
  accent: "#95E1D3"
  error: "#FF6B6B"
```

### Keybindings

**Presets:**
- `default` - Standard keys
- `vim` - Vim-style (hjkl)
- `emacs` - Emacs-style

**Custom:**

```yaml
keybindings: "custom"
custom_keybindings:
  quit: "q"
  help: "F1"
  search: "/"
  delete: "d"
  edit: "e"
```

---

## Common Patterns

### Async Data Loading

```go
// In types.go
type dataLoadedMsg struct {
    data []Item
}

// Create command
func loadDataCmd() tea.Cmd {
    return func() tea.Msg {
        data := fetchFromAPI()
        return dataLoadedMsg{data}
    }
}

// In Init()
func (m model) Init() tea.Cmd {
    return loadDataCmd()
}

// In Update()
case dataLoadedMsg:
    m.items = msg.data
    m.loading = false
```

### Spinner/Progress

```go
import "github.com/charmbracelet/bubbles/spinner"

// In model
type model struct {
    spinner spinner.Model
    loading bool
}

// Initialize
s := spinner.New()
s.Spinner = spinner.Dot

// Update
if m.loading {
    m.spinner, cmd = m.spinner.Update(msg)
}

// View
if m.loading {
    return m.spinner.View() + " Loading..."
}
```

### Confirmation Dialog

```go
// In types.go
type confirmDialogMsg struct {
    message string
    confirmed bool
}

// Show dialog
m.showingDialog = true
m.dialogMessage = "Delete this item?"

// In update_keyboard.go
if m.showingDialog {
    switch msg.String() {
    case "y", "Y":
        return m, confirmDialog(true)
    case "n", "N":
        return m, confirmDialog(false)
    }
}
```

---

## Testing

### Unit Tests

```go
// model_test.go
func TestInitialModel(t *testing.T) {
    cfg := getDefaultConfig()
    m := initialModel(cfg)

    if m.cursor != 0 {
        t.Errorf("expected cursor 0, got %d", m.cursor)
    }
}
```

### Integration Tests

```go
import "github.com/charmbracelet/x/exp/teatest"

func TestNavigation(t *testing.T) {
    m := initialModel(getDefaultConfig())
    tm := teatest.NewTestModel(t, m)

    // Send down arrow
    tm.Send(tea.KeyMsg{Type: tea.KeyDown})

    // Check state
    if tm.Get().(model).cursor != 1 {
        t.Error("cursor should move down")
    }
}
```

---

## Deployment

### Building

```bash
# Development build
go build -o myapp

# Production build (smaller binary)
go build -ldflags="-s -w" -o myapp

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=darwin GOARCH=amd64 go build -o myapp-mac
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

### Installation

```bash
# Local install
sudo cp myapp /usr/local/bin/

# Or user install
mkdir -p ~/.local/bin
cp myapp ~/.local/bin/
```

### Distribution

**GitHub Releases:**

```bash
# Tag version
git tag v1.0.0
git push --tags

# Use goreleaser
goreleaser release
```

**Package managers:**
- Homebrew: Create tap with formula
- AUR: Create PKGBUILD
- Nix: Create derivation

---

## Tips & Tricks

### Debugging

Enable logging in config:

```yaml
logging:
  enabled: true
  level: "debug"
  file: "~/.local/share/myapp/debug.log"
```

Then tail the log:

```bash
tail -f ~/.local/share/myapp/debug.log
```

### Performance

For large lists (1000+ items):

```go
// Virtual scrolling
visibleStart := max(0, m.cursor - m.height/2)
visibleEnd := min(len(m.items), visibleStart + m.height)
visibleItems := m.items[visibleStart:visibleEnd]

// Render only visible items
for i, item := range visibleItems {
    // render item
}
```

### Terminal Compatibility

Check terminal capabilities:

```go
import "github.com/charmbracelet/lipgloss/term"

// Check colors
if term.ColorProfile() == term.TrueColor {
    // Use 24-bit colors
} else {
    // Fall back to 256 colors
}

// Check size
width, height := term.Size()
```

---

## Next Steps

1. ✅ Create your first app with `new_project.sh`
2. 📖 Read the [Architecture Guide](ARCHITECTURE.md)
3. 🔍 Explore [Example Apps](examples/)
4. 📚 Review [Research Documentation](docs/research/)
5. 🎨 Customize theme and keybindings
6. 🚀 Build and share your app!

---

For more help:
- See [Examples](examples/)
- Check [API Documentation](docs/api/)
- Browse [Component Guide](docs/guides/components.md)
- Join [Discussions](https://github.com/yourname/TUITemplate/discussions)
