# TUI Effects Library

**Added**: 2025-01-26
**Status**: Initial Release - Core Effects Complete ✅
**Location**: `lib/effects/`

## 📖 Overview

A reusable animation effects library extracted from the **Balatro TUI** landing page. Provides physics-based animations that make terminal UIs as visually appealing as modern GUIs.

## 🎨 What's Included

### 4 Effect Packages

| Package | Purpose | Lines | Status |
|---------|---------|-------|--------|
| **metaballs** | Lava lamp floating blobs | ~250 | ✅ Complete |
| **waves** | Sine wave distortions | ~200 | ✅ Complete |
| **rainbow** | Color cycling effects | ~150 | ✅ Complete |
| **compositor** | Multi-layer rendering | ~200 | ✅ Complete |

### File Structure

```
lib/effects/
├── metaballs/
│   ├── blob.go       # Individual blob physics
│   └── engine.go     # Metaball rendering engine
├── waves/
│   ├── grid.go       # Animated wavy grid
│   └── distortion.go # Wave distortion functions
├── rainbow/
│   └── cycle.go      # Rainbow color cycling
├── compositor/
│   └── layer.go      # ANSI-aware layer compositing
└── README.md         # Complete API documentation
```

## 🚀 Examples

Created working examples in `examples/effects/`:

| Example | Status | Description |
|---------|--------|-------------|
| **metaball-spinner** | ✅ Complete | Loading screen with 3 floating blobs |
| **wavy-menu** | ✅ Complete | Menu with animated wave background |
| **rainbow-text** | ✅ Complete | Color-cycling text demo |
| **landing-page** | ✅ Complete | Full TUI-style landing page |

## 🔧 Technical Highlights

### Metaballs Engine
- **Real physics**: Field strength calculated as `radius² / distance²`
- **Organic motion**: Sine/cosine wobble for natural movement
- **Gradient rendering**: Unicode block characters (`░▒▓█`)
- **Color blending**: Blobs take color from strongest contributor
- **Performance**: Optimized for 60fps with up to 10 blobs

### Wave Distortion
- **Mathematical**: Based on `sin(y/5 + frame/20) * amplitude`
- **Customizable**: Amplitude, frequency, and speed parameters
- **Flexible**: Apply to grids, text, or any coordinates
- **Smooth**: Natural flowing motion without jitter

### Rainbow Cycling
- **7-color default**: Red→Orange→Yellow→Green→Cyan→Blue→Magenta
- **Per-character**: Each character gets different color
- **Wave patterns**: Multi-line text creates vertical color wave
- **Customizable**: Use any color palette

### Compositor
- **ANSI-aware**: Properly handles styled text overlay
- **Multi-layer**: Unlimited layers, rendered bottom-to-top
- **Auto-centering**: Automatically centers smaller layers
- **Smart extraction**: Extracts visible characters accounting for ANSI codes

## 💡 Use Cases

### Loading Screens
```go
engine := metaballs.NewEngine(width, height)
engine.AddBlob(/* colorful blob */)
// Shows animated loading with floating blobs
```

### Background Effects
```go
grid := waves.NewGrid(width, height)
// Flowing grid background for menus
```

### Title Screens
```go
cycler := rainbow.NewCycler()
title := cycler.RenderLines(asciiArt)
// Rainbow animated title
```

### Complex UIs
```go
comp := compositor.NewCompositor(width, height)
comp.AddLayer(background)
comp.AddLayer(metaballs)
comp.AddLayer(text)
// Professional multi-layer UI
```

## 📚 Documentation

All packages are fully documented with:
- Package-level godoc comments
- Function documentation
- Example usage in README
- Working code examples

### Quick Links
- **API Docs**: `lib/effects/README.md`
- **Examples**: `examples/effects/README.md`
- **Metaball Spinner**: `examples/effects/metaball-spinner/main.go`

## 🎯 Design Principles

1. **Simple API**: Create, Update, Render pattern
2. **Bubbletea Compatible**: Works seamlessly with Elm architecture
3. **Customizable**: Exposed parameters for colors, speeds, sizes
4. **Performant**: Optimized for smooth 60fps animation
5. **Reusable**: Drop into any TUI project

## 🔮 Future Enhancements

Potential additions (community PRs welcome):

- [ ] **Matrix rain effect** - Falling characters
- [ ] **Particle systems** - Explosions, confetti
- [ ] **Fire simulation** - Animated flames
- [ ] **Starfield** - Scrolling stars background
- [ ] **Plasma effect** - Retro demo scene style
- [ ] **Audio visualizer** - Bars reacting to audio
- [ ] **Game of Life** - Conway's cellular automaton background
- [ ] **Smoke/water** - Fluid simulations

## 📊 Impact

This library enables TUI developers to create:
- **Professional polish**: Animated landing pages
- **Better UX**: Visual feedback during loading
- **Brand identity**: Unique visual signatures
- **User delight**: "Wow" factor in terminals

## 🎨 Inspiration

Extracted from **Balatro TUI**'s landing page, which features:
- Wavy purple grid background
- 4 floating colored metaballs
- Rainbow "BALATRO" title
- All composited into a stunning landing screen

This proved that terminals can be as beautiful as GUIs, and we've made it reusable for everyone!

## 🤝 Integration Examples

### Adding to Existing TUI

```go
import "github.com/GGPrompts/TUITemplate/lib/effects/metaballs"

type Model struct {
    // ... your existing fields
    backgroundEffect *metaballs.Engine
}

func (m Model) Init() tea.Cmd {
    m.backgroundEffect = metaballs.NewEngine(m.width, m.height)
    // ... add blobs
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tickMsg:
        m.backgroundEffect.Update()
        return m, tick()
    }
    return m, nil
}

func (m Model) View() string {
    background := m.backgroundEffect.Render()
    // ... compose with your UI
}
```

### Using as Loading State

```go
if m.loading {
    return m.spinner.Render() // Metaball spinner
}
return m.renderMainUI()
```

### Menu Background

```go
background := m.waveGrid.Render()
menu := m.renderMenu()
return compositor.Overlay(background, menu)
```

## 🎓 Learning Resources

1. **Start**: Read `lib/effects/README.md`
2. **Practice**: Run `examples/effects/metaball-spinner`
3. **Customize**: Modify the example
4. **Build**: Create your own effect
5. **Share**: PR back to the community!

## 📈 Metrics

- **Total Lines**: ~800 LOC across all packages
- **API Surface**: ~30 public functions/methods
- **Dependencies**: Only Bubbletea + Lipgloss
- **Test Coverage**: 0% (TODO: Add tests)
- **Examples**: 4 complete examples

## ✨ Credits

**Original Implementation**: Balatro TUI project (TUI Classics)
**Extraction & Library**: Created for TUITemplate
**Inspiration**: LocalThunk's Balatro game visuals

## 🎉 Conclusion

The TUI Effects library proves that terminals can be **beautiful, animated, and professional**.

Use these effects to:
- Elevate your TUI projects
- Delight your users
- Stand out from CLI tools
- Have fun coding!

---

**Ready to make your TUIs stunning? Check out the examples!** 🚀

```bash
cd examples/effects/metaball-spinner
go run main.go
```
