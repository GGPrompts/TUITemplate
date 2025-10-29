# TUI Effects Examples

**Visual demonstrations of the TUI Effects library.**

These examples show you how to use each effect in your own applications.

## 🎯 Examples

### 1. Metaball Spinner
**File**: `metaball-spinner/main.go`

A loading screen with 3 colorful metaballs floating around "Loading..." text.

**Demonstrates**:
- Creating a metaball engine
- Adding multiple blobs with different colors
- Physics-based animation
- Text overlay

```bash
cd metaball-spinner
go run main.go
```

**What you'll see**:
- Cyan, magenta, and yellow blobs
- Organic floating motion
- Blobs merging and separating
- Smooth gradient rendering

---

### 2. Wavy Menu
**File**: `wavy-menu/main.go`

An animated menu with a flowing grid background.

**Demonstrates**:
- Wave distortion effects
- Grid rendering
- Menu integration
- Background layers

```bash
cd wavy-menu
go run main.go
```

**What you'll see**:
- Purple wavy grid background
- Interactive menu with 5 options
- Gold highlight on selected item
- Arrow key navigation

---

### 3. Rainbow Text
**File**: `rainbow-text/main.go`

Text with animated rainbow colors.

**Demonstrates**:
- Color cycling
- Text effects
- Custom color palettes
- Multi-line rainbow waves

```bash
cd rainbow-text
go run main.go
```

**What you'll see**:
- Large "RAINBOW" ASCII art
- Animated color cycling
- Multiple color palettes (Rainbow, Pastel, Fire, Neon)
- Speed controls (+/- keys)

---

### 4. Landing Page (Full Demo)
**File**: `landing-page/main.go`

The complete TUI landing page with all effects combined.

**Demonstrates**:
- Multi-layer compositing
- Metaballs + grid background
- Rainbow title text
- Professional UI layout
- All effects working together

```bash
cd landing-page
go run main.go
```

**What you'll see**:
- Purple wavy grid background
- 4 floating colored metaballs
- Rainbow animated title
- Interactive menu with gold border
- All effects seamlessly composited

---

## 🚀 Quick Start

### Prerequisites

```bash
go mod tidy
```

### Run Any Example

```bash
cd [example-name]
go run main.go
```

### Controls

All examples support:
- **q** or **Ctrl+C** - Quit
- **Esc** - Also quits

## 📚 Learning Path

Recommended order for learning:

1. **Start with metaball-spinner** - Simplest example
2. **Then wavy-menu** - Add wave effects
3. **Then rainbow-text** - Color cycling
4. **Finally landing-page** - See it all together

## 🎨 Customization Ideas

Try modifying the examples:

### Metaball Spinner
- Change blob colors
- Add more blobs
- Adjust velocities
- Change gradient characters

### Wavy Menu
- Modify wave amplitude
- Change grid size
- Custom colors
- Different wave speeds

### Rainbow Text
- Use brand colors instead of rainbow
- Adjust cycling speed
- Vertical vs horizontal waves
- Different text sizes

### Landing Page
- Combine different color schemes
- Add your own logo/title
- Custom menu items
- Different layer arrangements

## 🔧 Code Snippets

### Basic Animation Loop

```go
type Model struct {
    effect EffectType
    frame  int
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        m.effect.Update() // Update the effect
        return m, tick()
    }
    return m, nil
}

func (m Model) View() string {
    return m.effect.Render() // Render the effect
}
```

### Combining Multiple Effects

```go
// Create effects
grid := waves.NewGrid(width, height)
metaball := metaballs.NewEngine(width, height)
rainbow := rainbow.NewCycler()

// Create compositor
comp := compositor.NewCompositor(width, height)
comp.AddLayer(grid)      // Background
comp.AddLayer(metaball)  // Middle
comp.AddLayer(textLayer) // Foreground

// Update all
grid.Update()
metaball.Update()
rainbow.Update()

// Render
return comp.Composite()
```

## 💡 Tips

1. **Frame Rate**: 20-30 fps is plenty smooth (50ms tick)
2. **Terminal Size**: Handle window resize messages
3. **Alt Screen**: Use `tea.WithAltScreen()` for clean exit
4. **Colors**: Test on different terminals (some support fewer colors)
5. **Performance**: Profile if needed - these effects are fast!

## 🎓 Advanced Topics

### Custom Effects

Want to create your own effect? Follow this pattern:

```go
type MyEffect struct {
    Frame  int
    Width  int
    Height int
    // ... custom fields
}

func NewMyEffect(width, height int) *MyEffect {
    return &MyEffect{
        Width:  width,
        Height: height,
    }
}

func (e *MyEffect) Update() {
    e.Frame++
    // ... update logic
}

func (e *MyEffect) Render() string {
    // ... render logic
    return result
}

func (e *MyEffect) Resize(width, height int) {
    e.Width = width
    e.Height = height
}
```

### Integrating with Bubbletea

Effects work great with the Elm architecture:

- **Init**: Create your effects
- **Update**: Call effect.Update() on tick messages
- **View**: Call effect.Render() and compose with your UI

### Performance Optimization

If you notice lag:
- Reduce blob count
- Increase tick duration (slower frame rate)
- Use simpler gradient characters
- Profile with `go tool pprof`

## 📖 Reference

See the main library README for:
- API documentation
- All available methods
- Customization options
- Best practices

**Location**: `../../lib/effects/README.md`

## 🤝 Contributing

Have a cool effect idea? Create it and PR!

Ideas we'd love to see:
- Matrix rain
- Particle explosions
- Fireworks
- Audio visualizer
- Conway's Game of Life background
- Starfield
- Plasma effect

## 🎉 Credits

These effects were extracted from the **Balatro TUI** project and made reusable for the community.

---

**Happy coding! Make your TUIs beautiful! ✨**
