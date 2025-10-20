# Hello World Example

Minimal TUITemplate example demonstrating the basic architecture.

## Features

- Clean separation: `main.go` (entry point) + `model.go` (logic & rendering)
- Simple counter that increments on button press
- Keyboard shortcuts
- Lipgloss styling

## Run

```bash
cd examples/hello
go run .
```

## Code Structure

**main.go** (10 lines)
- Entry point only
- Creates Bubbletea program
- Runs the app

**model.go** (70 lines)
- `model` struct - application state
- `initialModel()` - initialize state
- `Init()` - startup commands
- `Update()` - handle messages
- `View()` - render UI

## What it demonstrates

1. **Minimal viable TUI** - All essentials in ~80 lines
2. **Clean architecture** - Separation of concerns
3. **Event handling** - Keyboard input
4. **Styling** - Lipgloss styling
5. **Layout** - Vertical composition

## Keyboard Shortcuts

- `Space` / `Enter` - Increment counter
- `R` - Reset counter
- `Q` - Quit

## Next Steps

See more complex examples:
- `examples/file_browser` - Navigation and lists
- `examples/todo_app` - CRUD operations
- `examples/json_viewer` - Dual-pane layout with preview

Or create your own:

```bash
cd ../..
./scripts/new_project.sh
```
