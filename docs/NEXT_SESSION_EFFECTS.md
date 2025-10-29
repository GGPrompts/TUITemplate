# Next Session: Complete TUI Effects Examples

**Date**: 2025-01-26
**Status**: Core library complete ✅, 3 examples remaining
**Location**: `/home/matt/projects/TUITemplate/examples/effects/`

## 🎯 Goal

Complete the remaining 3 effect examples to showcase the full capabilities of the TUI Effects library.

## ✅ What's Already Done

- ✅ **Core Library** (`lib/effects/`) - All 4 packages complete
  - `metaballs/` - Blob physics and metaball engine
  - `waves/` - Grid and distortion effects
  - `rainbow/` - Color cycling
  - `compositor/` - Layer management
- ✅ **Documentation** - Comprehensive READMEs
- ✅ **Example 1** - `metaball-spinner/main.go` (working!)

## 🚀 What to Create

### 1. Wavy Menu Example
**File**: `examples/effects/wavy-menu/main.go`

Create an interactive menu with a flowing wave grid background.

**Features**:
- Wavy grid background (purple theme)
- 5 menu items (like Balatro: NEW GAME, CONTINUE, COLLECTION, OPTIONS, QUIT)
- Arrow key navigation
- Gold highlight on selected item
- Enter to select
- Smooth animation

**Implementation Guide**:

```go
package main

import (
    "github.com/GGPrompts/TUITemplate/lib/effects/waves"
    "github.com/GGPrompts/TUITemplate/lib/effects/compositor"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Model struct {
    grid         *waves.Grid
    selectedItem int
    menuItems    []string
    width        int
    height       int
}

// Menu items
menuItems := []string{
    "NEW GAME",
    "CONTINUE",
    "COLLECTION",
    "OPTIONS",
    "QUIT",
}

// In Update():
// - Handle arrow keys to change selectedItem
// - Handle Enter to take action
// - Update grid on tickMsg

// In View():
// 1. Render the wavy grid background
// 2. Create menu with styled items (gold for selected)
// 3. Add border around menu
// 4. Use compositor or lipgloss.Place to center menu over grid
// 5. Add controls hint at bottom
```

**Reference**: See how `metaball-spinner` handles tick messages and rendering.

**Expected Output**:
```
[Flowing wavy grid background in purple]

    ╭────────────────╮
    │  ► NEW GAME    │  <- Gold, selected
    │    CONTINUE    │  <- Gray
    │    COLLECTION  │
    │    OPTIONS     │
    │    QUIT        │
    ╰────────────────╯

    ↑↓: Navigate | Enter: Select | Q: Quit
```

---

### 2. Rainbow Text Example
**File**: `examples/effects/rainbow-text/main.go`

Demonstrate rainbow color cycling on ASCII art or large text.

**Features**:
- Big ASCII art title (use "RAINBOW" or "EFFECTS")
- Rainbow colors cycling through the text
- Smooth color wave animation
- Show both single-line and multi-line rainbow effects
- FPS counter (optional)

**Implementation Guide**:

```go
package main

import (
    "github.com/GGPrompts/TUITemplate/lib/effects/rainbow"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    cycler *rainbow.Cycler
    width  int
    height int
}

// Define ASCII art
var asciiArt = []string{
    "██████╗  █████╗ ██╗███╗   ██╗██████╗  ██████╗ ██╗    ██╗",
    "██╔══██╗██╔══██╗██║████╗  ██║██╔══██╗██╔═══██╗██║    ██║",
    "██████╔╝███████║██║██╔██╗ ██║██████╔╝██║   ██║██║ █╗ ██║",
    "██╔══██╗██╔══██║██║██║╚██╗██║██╔══██╗██║   ██║██║███╗██║",
    "██║  ██║██║  ██║██║██║ ╚████║██████╔╝╚██████╔╝╚███╔███╔╝",
    "╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝  ╚══╝╚══╝ ",
}

// In View():
// 1. Use cycler.RenderLines(asciiArt) for the title
// 2. Show a separator
// 3. Use cycler.Render() for single-line example text
// 4. Show color palette used
// 5. Show current speed setting
// 6. Add controls to adjust speed (+ and -)
```

**Bonus Features**:
- Press `+` to speed up color cycling
- Press `-` to slow down color cycling
- Press `c` to cycle through different color palettes (rainbow, pastel, fire)

**Expected Output**:
```
[Large "RAINBOW" text with animated rainbow colors]

────────────────────────────────────

Single line: [rainbow colored text]

Multi-line:
  [Line 1 with rainbow wave]
  [Line 2 with rainbow wave]
  [Line 3 with rainbow wave]

Current Speed: 5 frames/shift
Palette: Rainbow (7 colors)

+/-: Speed | C: Change palette | Q: Quit
```

---

### 3. Landing Page (Full Demo)
**File**: `examples/effects/landing-page/main.go`

The complete Balatro-style landing page combining ALL effects!

**Features**:
- Wavy grid background
- 4 floating metaballs (cyan, magenta, yellow, purple)
- Rainbow animated title
- Menu with gold selection
- All composited together
- Smooth 60fps animation

**Implementation Guide**:

```go
package main

import (
    "github.com/GGPrompts/TUITemplate/lib/effects/metaballs"
    "github.com/GGPrompts/TUITemplate/lib/effects/waves"
    "github.com/GGPrompts/TUITemplate/lib/effects/rainbow"
    "github.com/GGPrompts/TUITemplate/lib/effects/compositor"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Model struct {
    grid         *waves.Grid
    metaballs    *metaballs.Engine
    rainbow      *rainbow.Cycler
    selectedItem int
    menuItems    []string
    width        int
    height       int
}

// Initialize all effects:
// 1. Create wavy grid (purple theme)
// 2. Create metaball engine with 4 blobs:
//    - Blob 1: x=w/4, y=h/3, cyan
//    - Blob 2: x=3w/4, y=h/2, magenta
//    - Blob 3: x=w/2, y=2h/3, yellow
//    - Blob 4: x=w/3, y=h/4, purple
// 3. Create rainbow cycler
// 4. Menu items

// Big ASCII art for title
var titleArt = []string{
    "████████╗██╗   ██╗██╗",
    "╚══██╔══╝██║   ██║██║",
    "   ██║   ██║   ██║██║",
    "   ██║   ██║   ██║██║",
    "   ██║   ╚██████╔╝██║",
    "   ╚═╝    ╚═════╝ ╚═╝",
}

// In Update():
// - Update all effects on tick
// - Handle menu navigation
// - Handle selection

// In View():
// LAYER 1: Render grid background
// LAYER 2: Render metaballs
// LAYER 3: Composite grid + metaballs
// LAYER 4: Create rainbow title in bordered box
// LAYER 5: Create menu in bordered box (gold border)
// LAYER 6: Manually composite title + menu onto background
//          (center them vertically, space them apart)
// LAYER 7: Add controls at bottom
```

**Rendering Strategy**:
```go
// 1. Render background (grid + metaballs)
gridRender := m.grid.Render()
blobsRender := m.metaballs.Render()
// Composite these (blobs over grid where blobs have content)

// 2. Render title box
titleText := m.rainbow.RenderLines(titleArt)
titleBox := lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("201")). // Magenta
    Padding(1, 2).
    Render(titleText)

// 3. Render menu box
menu := renderMenu(m.menuItems, m.selectedItem)
menuBox := lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("226")). // Gold
    Padding(1, 2).
    Render(menu)

// 4. Calculate positions
titleHeight := lipgloss.Height(titleBox)
menuHeight := lipgloss.Height(menuBox)
totalHeight := titleHeight + 2 + menuHeight
startY := (m.height - totalHeight) / 2

// 5. Overlay title and menu onto background
// Split background into lines, overlay at calculated positions
```

**Expected Output**:
```
[Purple wavy grid with floating colored metaballs]

    ╭──────────────────╮
    │  [RAINBOW TITLE] │  <- Animated rainbow colors
    ╰──────────────────╯

    ╭────────────────╮
    │  ► NEW GAME    │  <- Gold border, gold selection
    │    CONTINUE    │
    │    COLLECTION  │
    │    OPTIONS     │
    │    QUIT        │
    ╰────────────────╯

↑↓: Navigate | Enter: Select | Q: Quit
```

---

## 🔧 Implementation Tips

### All Examples Should Include:

1. **Tick Message**:
```go
type tickMsg time.Time

func tick() tea.Cmd {
    return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}
```

2. **Window Size Handling**:
```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
    m.grid.Resize(msg.Width, msg.Height)
    // ... resize other effects
```

3. **Alt Screen**:
```go
p := tea.NewProgram(m, tea.WithAltScreen())
```

4. **Quit Handling**:
```go
case "q", "ctrl+c", "esc":
    return m, tea.Quit
```

### Color Palette to Use (Neon Theme)

```go
var neonColors = struct {
    Gold    lipgloss.Color
    Cyan    lipgloss.Color
    Magenta lipgloss.Color
    Green   lipgloss.Color
    Purple  lipgloss.Color
    Yellow  lipgloss.Color
    Red     lipgloss.Color
}{
    Gold:    lipgloss.Color("226"),
    Cyan:    lipgloss.Color("51"),
    Magenta: lipgloss.Color("201"),
    Green:   lipgloss.Color("46"),
    Purple:  lipgloss.Color("129"),
    Yellow:  lipgloss.Color("226"),
    Red:     lipgloss.Color("196"),
}
```

### Testing Each Example

After creating each example:

```bash
cd examples/effects/[example-name]
go mod init example
go mod edit -replace github.com/GGPrompts/TUITemplate=../../..
go mod tidy
go run main.go
```

## 📚 Reference Files

Look at these for guidance:

1. **`examples/effects/metaball-spinner/main.go`** - Complete working example
2. **`lib/effects/metaballs/engine.go`** - API reference
3. **`lib/effects/waves/grid.go`** - Grid API
4. **`lib/effects/rainbow/cycle.go`** - Rainbow API
5. **Original source**: `/home/matt/projects/TUIClassics/games/balatro/landing_page.go`

## ✅ Completion Checklist

For each example, verify:

- [ ] Runs without errors
- [ ] Animations are smooth (20-30 fps)
- [ ] Window resize works correctly
- [ ] Quit keys work (q, Ctrl+C, Esc)
- [ ] Colors look good
- [ ] Controls are documented
- [ ] Code is clean and commented

## 🎉 When Complete

After finishing all 3 examples:

1. Update `examples/effects/README.md` - Remove "(Coming Soon)" tags
2. Test all 4 examples work
3. Take screenshots if possible (for docs)
4. Update main TUITemplate README to mention effects library
5. Consider publishing as standalone module
6. Share with the community! 🌟

## 📝 File Structure When Done

```
examples/effects/
├── README.md
├── metaball-spinner/
│   └── main.go          ✅ DONE
├── wavy-menu/
│   └── main.go          🎯 TODO
├── rainbow-text/
│   └── main.go          🎯 TODO
└── landing-page/
    └── main.go          🎯 TODO
```

## 🚀 Quick Start Command

```bash
cd /home/matt/projects/TUITemplate/examples/effects/wavy-menu
# Create main.go following the guide above
go mod init example
go mod edit -replace github.com/GGPrompts/TUITemplate=../../..
go mod tidy
go run main.go
```

---

**Happy coding! These effects are going to look amazing! ✨**
