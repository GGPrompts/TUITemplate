# Termux Mobile Demo

Mobile-optimized TUI for Termux with keyboard open

## Features

- 🚀 Fast and lightweight
- 🎨 Beautiful TUI interface
- ⚙️ Configurable (see ~/.config/termux-mobile-demo/config.yaml.example)
- 🖱️ Mouse and keyboard support

## Installation

```bash
go install github.com/Matt/termux-mobile-demo@latest
```

Or build from source:

```bash
git clone https://github.com/Matt/termux-mobile-demo.git
cd termux-mobile-demo
go build
```

## Usage

```bash
termux-mobile-demo
```

### Keyboard Shortcuts

- `q` - Quit
- `?` - Help
- `↑/↓` or `k/j` - Navigate
- `Enter` - Select
- `Space` - Toggle

## Configuration

Copy the example config:

```bash
cp ~/.config/termux-mobile-demo/config.yaml.example ~/.config/termux-mobile-demo/config.yaml
```

Then edit `~/.config/termux-mobile-demo/config.yaml` to customize.

## Development

This project was built with [TUITemplate](https://github.com/yourname/TUITemplate).

### Project Structure

```
termux-mobile-demo/
├── main.go              # Entry point
├── types.go             # Type definitions
├── model.go             # Model initialization
├── update.go            # Message dispatcher
├── update_keyboard.go   # Keyboard handling
├── update_mouse.go      # Mouse handling
├── view.go              # View rendering
├── styles.go            # Lipgloss styles
├── config.go            # Configuration
├── components/          # UI components
└── lib/                 # Utility libraries
```

### Building

```bash
go build
```

### Testing

```bash
go test ./...
```

## License

MIT

## Author

Matt
