# Terminal Size Detector for Termux

Real-time terminal dimension detector - perfect for finding exact Termux sizes with/without keyboard.

## Purpose

Helps you determine:
- Exact rows and columns in Termux
- Size with keyboard open vs closed
- Optimal layout dimensions for mobile TUIs

## Quick Start

```bash
# Build
go build

# Run in Termux
./termux-size-detector

# Or use go run
go run .
```

## Usage

1. Run the app
2. Open/close the keyboard to see size changes
3. Note the dimensions shown
4. Press 'q' to quit

## Features

- **Real-time Size Display** - Shows current width × height
- **Max/Min Tracking** - Records largest and smallest sizes seen
- **Visual Bars** - Graphical representation of dimensions
- **Size Categories** - Automatically categorizes terminal size
- **Layout Recommendations** - Suggests optimal layouts for current size
- **Change History** - Shows last 10 size changes

## Typical Termux Sizes

### With Keyboard Open
- **Portrait:** ~70-80 cols × 8-12 rows
- **Landscape:** ~130-150 cols × 8-10 rows

### Without Keyboard
- **Portrait:** ~70-80 cols × 35-45 rows
- **Landscape:** ~130-150 cols × 20-25 rows

*Note: Exact dimensions vary by device, keyboard, and font size*

## Use This To

1. **Find your exact Termux dimensions**
2. **Design mobile-optimized TUI layouts**
3. **Test responsive terminal apps**
4. **Determine breakpoints for adaptive layouts**

## Recommendations by Size

### MICRO (≤10 rows) - Termux with keyboard
- Single line title
- 3-5 visible items
- No multi-panel layouts
- Compact everything

### SMALL (11-20 rows) - Termux full screen
- Minimal title bar
- 10-15 items visible
- Single pane layouts
- Tight margins

### MEDIUM+ (21+ rows) - Desktop terminals
- All features available
- Multi-panel layouts work
- Full title/status bars

## Building for Termux

```bash
# On your development machine
GOOS=linux GOARCH=arm64 go build -o termux-size-detector

# Or if your device is arm (32-bit)
GOOS=linux GOARCH=arm go build -o termux-size-detector

# Transfer to Termux and run
./termux-size-detector
```

## Next Steps

Once you know your Termux dimensions, use them to:
- Create a mobile-optimized template
- Set up responsive breakpoints
- Design compact layouts
- Test your TUI apps in constrained spaces

---

**Made for TUITemplate** - Build better mobile TUIs! 📱
