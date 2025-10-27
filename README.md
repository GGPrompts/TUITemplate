# TUITemplate

**A production-ready template for building beautiful terminal user interfaces with Go, Bubbletea, and Lipgloss.**

## Overview

TUITemplate provides a modular, well-architected starting point for creating TUI applications. It extracts best practices from the [TFE (Terminal File Explorer)](https://github.com/GGPrompts/tfe) project and packages them into reusable components, patterns, and templates.

### Features

- 🏗️ **Modular Architecture** - Clean separation of concerns (main, types, model, update, view)
- 🎨 **Beautiful Styling** - Pre-configured Lipgloss styles with theme support
- ⚙️ **Configuration System** - YAML config with hot-reload and validation
- 🖱️ **Mouse & Keyboard** - Full input handling with customizable keybindings
- 📐 **Flexible Layouts** - Single-pane, dual-pane, multi-panel, and tabbed interfaces
- 🔧 **Utility Libraries** - Config, theme, keybindings, logging, clipboard, terminal detection
- ✨ **Effects Library** - Physics-based animations: metaballs, wave distortions, rainbow cycling
- 📚 **Comprehensive Docs** - Research on 80+ TUI tools and 20+ Bubbletea libraries
- 🚀 **Quick Start** - Generate new projects in seconds with `new_project.sh`

## Quick Start

### 1. Clone the Template

```bash
git clone https://github.com/yourname/TUITemplate.git
cd TUITemplate
```

### 2. Create a New Project

```bash
./scripts/new_project.sh

# Follow the interactive prompts:
# App name: my-tui-app
# Title: My TUI App
# Description: A beautiful TUI application
# Author: Your Name
# Layout: dual_pane
# Components: panel,list,preview
```

### 3. Build and Run

```bash
cd ../my-tui-app
go mod tidy
go run .
```

That's it! You now have a fully functional TUI application with:
- Clean modular architecture
- Keyboard and mouse support
- Configuration system
- Beautiful styling
- Example layout

## Project Structure

```
TUITemplate/
├── template/           # Core template files (.tmpl)
├── components/         # Reusable Bubbletea components
├── lib/               # Utility libraries
├── examples/          # Example applications
├── docs/              # Documentation and research
└── scripts/           # Helper scripts
```

## Template Files

The `template/` directory contains production-ready Go files:

| File | Purpose | Lines |
|------|---------|-------|
| `main.go.tmpl` | Entry point (minimal, 21 lines) | 21 |
| `types.go.tmpl` | Type definitions, structs, enums | ~150 |
| `model.go.tmpl` | Model initialization & layout | ~100 |
| `update.go.tmpl` | Message dispatcher | ~80 |
| `update_keyboard.go.tmpl` | Keyboard handling | ~200 |
| `update_mouse.go.tmpl` | Mouse handling | ~150 |
| `view.go.tmpl` | View rendering & layouts | ~200 |
| `styles.go.tmpl` | Lipgloss style definitions | ~150 |
| `config.go.tmpl` | Configuration management | ~200 |

## Available Components

### Panel System

Pre-built panel layouts for different UI arrangements:

- **Single Panel** - Full-screen single view
- **Dual Pane** - Side-by-side panels (like TFE)
- **Multi-Panel** - 3+ panels with configurable sizes
- **Tabbed** - Multiple views with tab switching

### UI Components

- **Lists** - Simple list, filtered list, tree view
- **Input** - Text input, multiline, forms, autocomplete
- **Dialogs** - Confirm, input, progress, modal
- **Menus** - Context menu, command palette, menu bar
- **Status** - Status bar, title bar, breadcrumbs
- **Preview** - Text, markdown, syntax highlighting, images, hex
- **Tables** - Simple and interactive tables

## TUI Effects Library

Beautiful, physics-based animations for terminal UIs:

### Available Effects

- 🔮 **Metaballs** - Lava lamp-style floating blobs with physics simulation
- 🌊 **Wave Effects** - Sine wave distortions for grids and content
- 🌈 **Rainbow Cycling** - Animated color gradients for text
- 🎭 **Layer Compositor** - ANSI-aware multi-layer rendering

### Quick Example

```go
import "github.com/GGPrompts/TUITemplate/lib/effects/metaballs"

// Create engine with floating blobs
engine := metaballs.NewEngine(width, height)
engine.AddBlob(metaballs.NewBlob(x, y, vx, vy, radius, color))

// Update and render
engine.Update()
return engine.Render()
```

### Complete Examples

```bash
cd examples/effects/metaball-spinner  # Loading screen
cd examples/effects/wavy-menu         # Animated menu
cd examples/effects/rainbow-text      # Color cycling
cd examples/effects/landing-page      # Full demo
```

**Documentation:** `lib/effects/README.md` and `docs/EFFECTS_LIBRARY.md`

## Configuration

Apps built with TUITemplate use YAML configuration:

```yaml
# ~/.config/your-app/config.yaml

theme: "dark"
keybindings: "default"

layout:
  type: "dual_pane"
  split_ratio: 0.5

ui:
  show_title: true
  show_status: true
  mouse_enabled: true
  show_icons: true
```

## Research Documentation

TUITemplate includes comprehensive research on:

### TUI Ecosystem (989 lines)
- 20+ Bubbletea/Charm libraries
- Integration patterns and code examples
- Priority rankings and recommendations

### TUI Applications (1,249 lines)
- 80+ actively maintained TUI tools
- Image editors, hex editors, CSV viewers, databases
- File type detection and tool integration

**Location:** `docs/research/`

## Examples

### Hello World

```bash
cd examples/hello
go run .
```

Minimal example showing basic structure.

### File Browser

```bash
cd examples/file_browser
go run .
```

Simple file browser with navigation and preview.

### Todo App

```bash
cd examples/todo_app
go run .
```

Todo list with add/delete/toggle functionality.

## Customization

### Adding Components

```bash
cd your-app
../TUITemplate/scripts/add_component.sh table
```

This copies the table component and updates your imports.

### Updating Template

```bash
cd your-app
../TUITemplate/scripts/update_template.sh
```

Gets latest template improvements while preserving your customizations.

### Custom Themes

Create a custom theme in your config:

```yaml
theme: "custom"
custom_theme:
  primary: "#61AFEF"
  secondary: "#C678DD"
  background: "#282C34"
  foreground: "#ABB2BF"
  accent: "#98C379"
  error: "#E06C75"
```

## Integration with TFE

Apps built with TUITemplate can be launched from TFE:

1. Place your binary in PATH or a known location
2. Add to TFE's context menu in `~/.config/tfe/tools.yaml`:

```yaml
tools:
  - name: "JSON Viewer"
    command: "json-viewer {{file}}"
    icon: "📄"
    showFor: "files"
```

3. Right-click a file in TFE → Open With → JSON Viewer

## Best Practices

### Architecture

- Keep `main.go` minimal (entry point only)
- All types in `types.go`
- Separate keyboard and mouse handling
- One file, one responsibility
- Maximum file size: 800 lines (ideally <500)

### Error Handling

- Never crash on errors
- Show user-friendly error messages
- Log errors to debug file if logging enabled
- Graceful degradation for missing dependencies

### Performance

- Lazy loading for large datasets
- Virtual scrolling (render only visible items)
- Async operations for I/O
- Cache expensive computations

### Testing

- Write tests as features stabilize
- Use table-driven tests
- Test public APIs and edge cases

## Development Workflow

### 1. Design Phase

- Sketch the UI layout
- Identify required components
- Plan data structures

### 2. Implementation Phase

```bash
# Create project from template
./scripts/new_project.sh

# Add needed components
cd ../your-app
../TUITemplate/scripts/add_component.sh list
../TUITemplate/scripts/add_component.sh preview

# Implement features
# - Update types.go with your data structures
# - Implement rendering in view.go
# - Add keyboard shortcuts in update_keyboard.go
# - Add mouse interactions in update_mouse.go
```

### 3. Testing Phase

```bash
# Run your app
go run .

# Test with various terminal sizes
# Test keyboard shortcuts
# Test mouse interactions
```

### 4. Release Phase

```bash
# Build binary
go build -o your-app

# Install
sudo mv your-app /usr/local/bin/

# Or package for distribution
```

## Dependencies

**Core:**
```
github.com/charmbracelet/bubbletea
github.com/charmbracelet/lipgloss
github.com/charmbracelet/bubbles
gopkg.in/yaml.v3
```

**Optional** (uncomment in go.mod as needed):
```
github.com/charmbracelet/glamour       # Markdown
github.com/charmbracelet/huh           # Forms
github.com/alecthomas/chroma/v2        # Syntax highlighting
github.com/evertras/bubble-table       # Tables
github.com/koki-develop/go-fzf         # Fuzzy finder
```

See `docs/research/ECOSYSTEM_QUICK_REFERENCE.md` for complete list.

## Documentation

### Guides

- **[Emoji Width Alignment Fix](docs/EMOJI_WIDTH_FIX.md)** - Fix emoji alignment issues in WezTerm/Termux
- **[Mouse Support Guide](docs/MOUSE_SUPPORT_GUIDE.md)** - Implementing mouse interactions
- **[Termux Mobile Guide](docs/TERMUX_MOBILE_GUIDE.md)** - Building TUIs for Android/mobile
- **[Scrolling & Responsive Design](docs/SCROLLING_AND_RESPONSIVE.md)** - Viewport and layout handling
- **[Effects Library](docs/EFFECTS_LIBRARY.md)** - Physics-based animations and visual effects

### Research & Analysis

- **[Lazygit Analysis](docs/LAZYGIT_ANALYSIS.md)** - Architecture study of popular TUI
- **Ecosystem Research** - 80+ TUI tools and 20+ Bubbletea libraries analyzed

## Contributing

Contributions welcome! Areas of focus:

- New components
- Additional themes
- Example applications
- Documentation improvements
- Bug fixes

## License

MIT License - Use freely for any project

## Resources

- [Bubbletea Documentation](https://github.com/charmbracelet/bubbletea)
- [Lipgloss Documentation](https://github.com/charmbracelet/lipgloss)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)
- [TFE Source Code](https://github.com/GGPrompts/tfe)

## Support

- GitHub Issues: Report bugs, request features
- Discussions: Ask questions, share apps
- Wiki: Community recipes and patterns

---

**Build beautiful terminal applications fast!** 🚀

Made with ❤️ using [Charm](https://charm.sh/)
