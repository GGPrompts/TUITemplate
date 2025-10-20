# TUITemplate - Bubbletea/Lipgloss Application Template

## Project Vision

**TUITemplate** is a production-ready, modular template for building terminal user interface (TUI) applications using the modern Go stack: Bubbletea + Lipgloss + Bubbles. It extracts best practices from the TFE project and provides reusable components, patterns, and architecture for rapid TUI development.

## Goals

1. **Speed up TUI development** - Go from idea to working app in hours, not days
2. **Consistent architecture** - All apps follow the same modular pattern
3. **Production-ready code** - Error handling, logging, configuration built-in
4. **Extensible components** - Pick and choose what you need
5. **Well-documented** - Clear examples and integration guides
6. **Easy customization** - Themes, keybindings, layouts configurable
7. **Integration-ready** - Apps can be launched from TFE or standalone

## Template Architecture

### Core Principle: Modular & Composable

Based on TFE's successful refactoring from 1668 lines to 21-line main.go, TUITemplate enforces modular architecture from day one.

### Project Structure

```
TUITemplate/
├── README.md                    # Quick start guide
├── ARCHITECTURE.md              # Detailed architecture guide
├── USAGE.md                     # How to use the template
├── go.mod                       # Dependencies
├── go.sum
│
├── template/                    # The actual template files
│   ├── main.go.tmpl            # Entry point template
│   ├── types.go.tmpl           # Type definitions
│   ├── styles.go.tmpl          # Lipgloss styles
│   ├── model.go.tmpl           # Model initialization
│   ├── update.go.tmpl          # Update dispatcher
│   ├── update_keyboard.go.tmpl # Keyboard handling
│   ├── update_mouse.go.tmpl    # Mouse handling
│   ├── view.go.tmpl            # View rendering
│   └── config.go.tmpl          # Configuration system
│
├── components/                  # Reusable Bubbletea components
│   ├── panel/                  # Panel system
│   │   ├── panel.go           # Base panel interface
│   │   ├── dual_pane.go       # Dual-pane layout
│   │   ├── multi_panel.go     # Multi-panel layout
│   │   └── tabbed.go          # Tabbed interface
│   ├── list/                   # List components
│   │   ├── simple_list.go     # Basic list view
│   │   ├── filtered_list.go   # List with filtering
│   │   └── tree.go            # Tree view
│   ├── input/                  # Input components
│   │   ├── text_input.go      # Single-line input
│   │   ├── multiline.go       # Multi-line editor
│   │   ├── form.go            # Form builder (Huh wrapper)
│   │   └── autocomplete.go    # Autocomplete input
│   ├── dialog/                 # Dialog components
│   │   ├── confirm.go         # Confirmation dialog
│   │   ├── input.go           # Input dialog
│   │   ├── progress.go        # Progress dialog
│   │   └── modal.go           # Generic modal
│   ├── menu/                   # Menu components
│   │   ├── context_menu.go    # Right-click menu
│   │   ├── command_palette.go # Fuzzy command finder
│   │   └── menubar.go         # Top menu bar
│   ├── status/                 # Status components
│   │   ├── statusbar.go       # Bottom status bar
│   │   ├── titlebar.go        # Top title bar
│   │   └── breadcrumb.go      # Breadcrumb navigation
│   ├── preview/                # Preview components
│   │   ├── text.go            # Text preview
│   │   ├── markdown.go        # Markdown (Glamour)
│   │   ├── syntax.go          # Code with highlighting
│   │   ├── image.go           # Image preview (terminal protocols)
│   │   └── hex.go             # Hex viewer
│   └── table/                  # Table components
│       ├── simple.go          # Basic table
│       └── interactive.go     # Interactive table (bubble-table)
│
├── lib/                        # Utility libraries
│   ├── config/                # Configuration management
│   │   ├── loader.go         # Load YAML/JSON config
│   │   ├── defaults.go       # Default values
│   │   └── validator.go      # Config validation
│   ├── theme/                 # Theme system
│   │   ├── theme.go          # Theme interface
│   │   ├── presets.go        # Built-in themes
│   │   └── loader.go         # Custom theme loader
│   ├── keybindings/           # Keybinding system
│   │   ├── bindings.go       # Keybinding definitions
│   │   ├── presets.go        # Vim/Emacs/Default presets
│   │   └── parser.go         # Parse user keybindings
│   ├── logger/                # Logging utilities
│   │   ├── logger.go         # Structured logger
│   │   └── debug.go          # Debug mode helpers
│   ├── clipboard/             # Clipboard integration
│   │   ├── clipboard.go      # Cross-platform clipboard
│   │   └── detect.go         # Tool detection
│   └── terminal/              # Terminal utilities
│       ├── detect.go         # Feature detection
│       ├── size.go           # Terminal size
│       └── protocols.go      # Image protocol detection
│
├── examples/                   # Example applications
│   ├── hello/                 # Minimal "Hello World"
│   ├── file_browser/          # Simple file browser
│   ├── todo_app/              # Todo list application
│   ├── log_viewer/            # Log file viewer
│   ├── json_viewer/           # JSON/YAML viewer
│   └── dashboard/             # Multi-panel dashboard
│
├── docs/                       # Documentation & Research
│   ├── research/              # Research from TFE project
│   │   ├── ECOSYSTEM_RESEARCH_2025.md
│   │   ├── ECOSYSTEM_QUICK_REFERENCE.md
│   │   ├── TUI_APPLICATIONS_RESEARCH.md
│   │   └── TUI_TOOLS_QUICK_REFERENCE.md
│   ├── guides/                # How-to guides
│   │   ├── getting_started.md
│   │   ├── components.md      # Component usage guide
│   │   ├── customization.md   # Customization guide
│   │   ├── patterns.md        # Design patterns
│   │   └── integration.md     # TFE integration guide
│   ├── api/                   # API documentation
│   │   ├── components.md      # Component APIs
│   │   ├── lib.md            # Library APIs
│   │   └── template.md       # Template variables
│   └── examples/              # Example walkthroughs
│       ├── simple_app.md
│       └── complex_app.md
│
├── scripts/                    # Helper scripts
│   ├── new_project.sh         # Create new project from template
│   ├── add_component.sh       # Add component to existing project
│   └── update_template.sh     # Update template in existing project
│
└── tests/                     # Tests
    ├── components/            # Component tests
    └── lib/                   # Library tests
```

## Template Features

### 1. Panel System

**Purpose**: Flexible layout management for different UI arrangements

**Layouts Provided:**
- **Single Panel**: Full-screen single view
- **Dual Pane**: Side-by-side panels (TFE-style)
- **Multi-Panel**: 3+ panels with configurable sizes
- **Tabbed**: Multiple views with tab switching
- **Dashboard**: Grid layout with widgets

**Usage Example:**
```go
import "github.com/yourname/tuitemplat/components/panel"

// Dual-pane layout
layout := panel.NewDualPane(
    leftPanel,     // Panel interface
    rightPanel,    // Panel interface
    0.5,          // Split ratio
    true,         // Show divider
)
```

### 2. Component Library

**Pre-built Bubbletea components ready to use:**

#### Lists & Trees
- Simple list (keyboard navigation, selection)
- Filtered list (with search)
- Tree view (expandable/collapsible)

#### Input Widgets
- Text input (single-line)
- Multiline editor (textarea)
- Form builder (Huh wrapper)
- Autocomplete input

#### Dialogs
- Confirm dialog (yes/no)
- Input dialog (prompt user)
- Progress dialog (long operations)
- Generic modal system

#### Menus
- Context menu (right-click)
- Command palette (fuzzy finder)
- Menu bar (top navigation)

#### Status Displays
- Status bar (bottom info)
- Title bar (top title/path)
- Breadcrumb navigation

#### Preview Panes
- Text preview
- Markdown rendering (Glamour)
- Syntax highlighting (Chroma)
- Image preview (Kitty/iTerm2/Sixel)
- Hex viewer

#### Tables
- Simple table
- Interactive table (bubble-table integration)

### 3. Utility Libraries

**lib/config**: Configuration management
- Load from YAML/JSON
- Environment variable support
- Validation with defaults
- Hot-reload support

**lib/theme**: Theme system
- Built-in themes (dark, light, solarized, dracula, nord)
- Custom theme loading
- Per-component style overrides
- Color scheme adapters

**lib/keybindings**: Keybinding system
- Configurable key bindings
- Presets: vim, emacs, default
- Conflict detection
- Help text generation

**lib/logger**: Logging utilities
- Structured logging (log/slog)
- Debug mode file logging
- Performance profiling
- Error tracking

**lib/clipboard**: Clipboard integration
- Cross-platform (Linux, macOS, Windows/WSL)
- Auto-detect available tools
- Fallback chain

**lib/terminal**: Terminal utilities
- Feature detection (mouse, colors, protocols)
- Size calculation
- Image protocol detection

### 4. Configuration System

**Config File Format** (~/.config/yourapp/config.yaml):
```yaml
# Theme
theme: "dark"  # dark, light, solarized, dracula, nord, custom

# Custom theme (if theme: custom)
custom_theme:
  primary: "#61AFEF"
  secondary: "#C678DD"
  background: "#282C34"
  foreground: "#ABB2BF"
  accent: "#98C379"
  error: "#E06C75"

# Keybindings
keybindings: "default"  # default, vim, emacs, custom

# Custom keybindings (if keybindings: custom)
custom_keys:
  quit: "q"
  help: "?"
  search: "/"
  confirm: "enter"
  cancel: "esc"

# Layout
layout:
  type: "dual_pane"  # single, dual_pane, multi_panel, tabbed
  split_ratio: 0.5   # For dual_pane
  show_divider: true

# UI Elements
ui:
  show_title: true
  show_status: true
  show_line_numbers: true
  mouse_enabled: true
  show_icons: true
  icon_set: "nerd_font"  # nerd_font, ascii, unicode

# Performance
performance:
  lazy_loading: true
  cache_size: 100
  async_operations: true

# Logging
logging:
  enabled: false
  level: "info"  # debug, info, warn, error
  file: "~/.local/share/yourapp/debug.log"
```

### 5. Template Variables

When creating a new project, the template system replaces variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `{{.AppName}}` | Application name | `json-viewer` |
| `{{.AppTitle}}` | Display title | `JSON Viewer` |
| `{{.PackageName}}` | Go package name | `jsonviewer` |
| `{{.Description}}` | App description | `A TUI JSON viewer` |
| `{{.Author}}` | Your name | `Matt` |
| `{{.License}}` | License type | `MIT` |
| `{{.Components}}` | Enabled components | `panel, list, preview` |
| `{{.Layout}}` | Default layout | `dual_pane` |

## Usage Workflow

### Creating a New App

```bash
# Clone TUITemplate
git clone https://github.com/yourname/TUITemplate.git
cd TUITemplate

# Run new project script
./scripts/new_project.sh

# Interactive prompts:
# App name: json-viewer
# Title: JSON Viewer
# Description: A beautiful JSON/YAML viewer for the terminal
# Author: Matt
# License: MIT
# Layout (single/dual_pane/multi_panel/tabbed): dual_pane
# Components (comma-separated): panel,list,preview,menu

# This creates:
cd ../json-viewer
# ├── main.go (from template)
# ├── types.go
# ├── model.go
# ├── update.go
# ├── view.go
# ├── config.yaml.example
# ├── components/ (selected components copied)
# ├── lib/ (all libs copied)
# └── README.md

# Run your new app
go mod tidy
go run .
```

### Customizing an Existing App

```bash
# Add a new component
cd json-viewer
../TUITemplate/scripts/add_component.sh table

# Component added:
# - components/table/ copied
# - Import added to go.mod
# - Usage example printed

# Update template files (get latest improvements)
../TUITemplate/scripts/update_template.sh
# This updates core files but preserves your customizations
```

## Component Design Patterns

### 1. Panel Interface

All panels implement a common interface:

```go
type Panel interface {
    // Bubbletea methods
    Init() tea.Cmd
    Update(tea.Msg) (Panel, tea.Cmd)
    View() string

    // Panel-specific
    SetSize(width, height int)
    GetSize() (width, height int)
    Focus()
    Blur()
    IsFocused() bool
}
```

### 2. Component Composition

Components are designed to be composable:

```go
// File browser component
fileBrowser := list.NewFilteredList(
    files,                    // Items
    preview.NewText(),        // Preview component
    menu.NewContextMenu(),    // Context menu
)

// Wrap in dual-pane
layout := panel.NewDualPane(
    fileBrowser,
    preview.NewSyntax(),     // Syntax highlighting preview
    0.6,                     // 60% left, 40% right
    true,                    // Show divider
)
```

### 3. Message-Based Communication

Components communicate via Bubbletea messages:

```go
// Custom messages
type fileSelectedMsg struct {
    path string
}

type previewRequestMsg struct {
    content string
    language string
}

// Components emit and handle these
```

## Integration with TFE

### Launching Apps from TFE

Apps built with TUITemplate can be launched from TFE's context menu or F-key shortcuts:

**In TFE's `editor.go`:**
```go
func openTUIApp(app string, path string) tea.Cmd {
    return tea.ExecProcess(
        exec.Command(app, path),
        func(err error) tea.Msg {
            return appFinishedMsg{err}
        },
    )
}
```

**TFE Context Menu:**
```
Open With >
  ├─ micro (text editor)
  ├─ lazygit (git client)
  ├─ json-viewer (TUITemplate app)  <-- Your app
  └─ hex-editor (TUITemplate app)   <-- Your app
```

### Returning to TFE

Apps can signal completion and return control to TFE:

```go
// In your TUITemplate app
if key.Matches(msg, keys.Quit) {
    // Clean exit, TFE resumes
    return m, tea.Quit
}
```

### Sharing Configuration

Apps can share TFE's theme/config:

```go
// Load TFE theme if available
tfeConfig := config.LoadTFEConfig()
if tfeConfig != nil {
    m.theme = tfeConfig.Theme
}
```

## Best Practices from TFE

### 1. Modular Architecture
- Keep `main.go` minimal (entry point only)
- One file, one responsibility
- Maximum file size: 800 lines (preferably <500)

### 2. Type Organization
- All types in `types.go`
- Enums and constants in `types.go`
- Custom messages in `types.go`

### 3. Event Handling
- Separate keyboard and mouse handling
- Use dispatcher pattern in `update.go`
- Handlers in `update_keyboard.go` and `update_mouse.go`

### 4. View Rendering
- Separate view logic from business logic
- Composable render functions
- Cache expensive rendering operations

### 5. Error Handling
- Never crash on errors
- Show user-friendly error messages
- Log errors to debug file if logging enabled
- Graceful degradation (missing dependencies)

### 6. Performance
- Lazy loading for large datasets
- Virtual scrolling (render only visible)
- Async operations for I/O
- Caching for expensive computations

## Development Roadmap

### Phase 1: Core Template (Week 1-2)
- [ ] Set up project structure
- [ ] Create template files (main, types, model, update, view)
- [ ] Implement config system (lib/config)
- [ ] Implement theme system (lib/theme)
- [ ] Create new_project.sh script
- [ ] Write basic documentation

### Phase 2: Essential Components (Week 3-4)
- [ ] Panel system (single, dual_pane, multi_panel)
- [ ] List components (simple, filtered)
- [ ] Input components (text_input, form)
- [ ] Dialog components (confirm, input)
- [ ] Status components (statusbar, titlebar)

### Phase 3: Advanced Components (Week 5-6)
- [ ] Tree view component
- [ ] Context menu component
- [ ] Command palette (fuzzy finder)
- [ ] Preview components (text, markdown, syntax)
- [ ] Table components

### Phase 4: Utilities & Polish (Week 7-8)
- [ ] Keybinding system
- [ ] Logger system
- [ ] Clipboard integration
- [ ] Terminal detection
- [ ] add_component.sh script
- [ ] update_template.sh script

### Phase 5: Examples & Documentation (Week 9-10)
- [ ] Example: Hello World
- [ ] Example: File Browser
- [ ] Example: Todo App
- [ ] Example: JSON Viewer
- [ ] Example: Dashboard
- [ ] Complete API documentation
- [ ] Write comprehensive guides

### Phase 6: Testing & Release (Week 11-12)
- [ ] Unit tests for components
- [ ] Integration tests
- [ ] Performance benchmarks
- [ ] GitHub Actions CI/CD
- [ ] Release v1.0.0

## Example Apps to Build with TUITemplate

### High Priority (Integrate with TFE)
1. **JSON/YAML Viewer** - Formatted viewing with syntax highlighting
2. **Hex Editor** - Binary file editing
3. **CSV Viewer** - Table view with sorting/filtering
4. **Log Viewer** - Real-time log tailing with filtering
5. **Markdown Editor** - Live preview editing
6. **Image Viewer** - Terminal image display (Kitty/iTerm2)

### Medium Priority
7. **Todo Manager** - Task list with priorities
8. **Note Taker** - Quick notes with search
9. **Password Manager** - Encrypted password storage
10. **Git Diff Viewer** - Beautiful diff display
11. **API Tester** - HTTP request builder
12. **SQL Browser** - Database query tool

### Lower Priority
13. **System Monitor** - CPU/memory/disk dashboard
14. **Network Monitor** - Connection viewer
15. **Process Manager** - htop alternative
16. **RSS Reader** - Feed aggregator

## Dependencies

**Core:**
```go
require (
    github.com/charmbracelet/bubbletea v1.3.10
    github.com/charmbracelet/lipgloss v1.1.1
    github.com/charmbracelet/bubbles v0.21.0
)
```

**Components:**
```go
require (
    github.com/charmbracelet/glamour v0.10.0      // Markdown
    github.com/charmbracelet/huh v0.7.0           // Forms
    github.com/alecthomas/chroma/v2 v2.14.0       // Syntax highlighting
    github.com/evertras/bubble-table v0.17.0      // Tables
    github.com/koki-develop/go-fzf v1.2.0         // Fuzzy finder
    github.com/lrstanley/bubblezone v0.0.0        // Mouse zones
)
```

**Utilities:**
```go
require (
    github.com/fsnotify/fsnotify v1.7.0           // File watching
    github.com/spf13/viper v1.19.0                // Config management
    gopkg.in/yaml.v3 v3.0.1                       // YAML parsing
)
```

**Optional (for advanced features):**
```go
require (
    github.com/blacktop/go-termimg v1.0.0         // Image preview
    github.com/BourgeoisBear/rasterm v0.0.0       // Advanced image protocols
    github.com/mholt/archiver/v4 v4.0.0           // Archive handling
    github.com/go-git/go-git/v5 v5.12.0           // Git integration
)
```

## License

MIT License - Use freely for any project

## Contributing

Contributions welcome! Areas of focus:
- New components
- Additional themes
- Example applications
- Documentation improvements
- Performance optimizations
- Bug fixes

## Support

- GitHub Issues: Report bugs, request features
- Discussions: Ask questions, share apps
- Wiki: Community recipes and patterns

---

**TUITemplate** - Build beautiful terminal applications fast! 🚀
