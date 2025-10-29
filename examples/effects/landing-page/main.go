package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/GGPrompts/TUITemplate/lib/effects/compositor"
	"github.com/GGPrompts/TUITemplate/lib/effects/metaballs"
	"github.com/GGPrompts/TUITemplate/lib/effects/rainbow"
	"github.com/GGPrompts/TUITemplate/lib/effects/waves"
)

// Model represents the application state
type Model struct {
	grid         *waves.Grid
	metaballs    *metaballs.Engine
	rainbow      *rainbow.Cycler
	selectedItem int
	menuItems    []string
	width        int
	height       int
}

// tickMsg is sent on every animation frame
type tickMsg time.Time

// ASCII art for title
var titleArt = []string{
	"████████╗██╗   ██╗██╗",
	"╚══██╔══╝██║   ██║██║",
	"   ██║   ██║   ██║██║",
	"   ██║   ██║   ██║██║",
	"   ██║   ╚██████╔╝██║",
	"   ╚═╝    ╚═════╝ ╚═╝",
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tick(),
		tea.WindowSize(),
	)
}

// tick returns a command that sends a tick message every 50ms (20fps)
func tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.grid != nil {
			m.grid.Resize(msg.Width, msg.Height)
		}
		if m.metaballs != nil {
			m.metaballs.Resize(msg.Width, msg.Height)
		}
		return m, nil

	case tickMsg:
		// Update all effects
		if m.grid != nil {
			m.grid.Update()
		}
		if m.metaballs != nil {
			m.metaballs.Update()
		}
		if m.rainbow != nil {
			m.rainbow.Update()
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.selectedItem > 0 {
				m.selectedItem--
			}

		case "down", "j":
			if m.selectedItem < len(m.menuItems)-1 {
				m.selectedItem++
			}

		case "enter", " ":
			// Handle menu selection
			switch m.selectedItem {
			case 0:
				// NEW GAME
				return m, nil
			case 1:
				// CONTINUE
				return m, nil
			case 2:
				// COLLECTION
				return m, nil
			case 3:
				// OPTIONS
				return m, nil
			case 4:
				// QUIT
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if m.grid == nil || m.metaballs == nil || m.rainbow == nil {
		return "Loading..."
	}

	// LAYER 1: Render wavy grid background
	gridRender := m.grid.Render()

	// LAYER 2: Render metaballs
	blobsRender := m.metaballs.Render()

	// LAYER 3: Composite grid + metaballs
	comp := compositor.NewCompositor(m.width, m.height)
	comp.AddLayer(gridRender)
	comp.AddLayer(blobsRender)
	background := comp.Composite()

	// LAYER 4: Create rainbow title in bordered box
	titleText := m.rainbow.RenderLines(titleArt)
	titleBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("201")). // Magenta
		Padding(1, 2).
		Render(titleText)

	// LAYER 5: Create menu in bordered box
	var menuLines []string
	for i, item := range m.menuItems {
		var line string
		if i == m.selectedItem {
			// Gold highlight for selected item
			style := lipgloss.NewStyle().
				Foreground(lipgloss.Color("226")).
				Bold(true)
			line = style.Render("  ► " + item)
		} else {
			// Gray for unselected items
			style := lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))
			line = style.Render("    " + item)
		}
		menuLines = append(menuLines, line)
	}

	menu := strings.Join(menuLines, "\n")
	menuBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("226")). // Gold
		Padding(1, 2).
		Render(menu)

	// LAYER 6: Calculate positions and overlay title + menu onto background
	titleHeight := lipgloss.Height(titleBox)
	menuHeight := lipgloss.Height(menuBox)
	totalContentHeight := titleHeight + 2 + menuHeight // 2 lines spacing between title and menu

	startY := (m.height - totalContentHeight) / 2
	if startY < 0 {
		startY = 0
	}

	titleY := startY
	menuY := titleY + titleHeight + 2

	// Overlay title and menu onto background
	backgroundLines := strings.Split(background, "\n")

	// Ensure we have enough background lines
	for len(backgroundLines) < m.height {
		backgroundLines = append(backgroundLines, strings.Repeat(" ", m.width))
	}

	// Overlay title
	titleLines := strings.Split(titleBox, "\n")
	titleWidth := lipgloss.Width(titleBox)
	titleX := (m.width - titleWidth) / 2
	if titleX < 0 {
		titleX = 0
	}

	for i, line := range titleLines {
		targetY := titleY + i
		if targetY >= 0 && targetY < len(backgroundLines) {
			backgroundLines[targetY] = overlayLine(backgroundLines[targetY], line, titleX, m.width)
		}
	}

	// Overlay menu
	menuLines2 := strings.Split(menuBox, "\n")
	menuWidth := lipgloss.Width(menuBox)
	menuX := (m.width - menuWidth) / 2
	if menuX < 0 {
		menuX = 0
	}

	for i, line := range menuLines2 {
		targetY := menuY + i
		if targetY >= 0 && targetY < len(backgroundLines) {
			backgroundLines[targetY] = overlayLine(backgroundLines[targetY], line, menuX, m.width)
		}
	}

	result := strings.Join(backgroundLines, "\n")

	// LAYER 7: Add controls at bottom
	controls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("\n\n↑↓: Navigate | Enter: Select | Q: Quit")

	return result + controls
}

// overlayLine overlays content onto a background line at the specified x position
func overlayLine(background, content string, x, maxWidth int) string {
	if x < 0 {
		x = 0
	}

	bgRunes := []rune(background)
	contentRunes := []rune(content)

	// Ensure background is wide enough
	for len(bgRunes) < maxWidth {
		bgRunes = append(bgRunes, ' ')
	}

	// Overlay content (character by character, preserving ANSI codes)
	// For simplicity, we'll do a basic overlay
	for i, r := range contentRunes {
		pos := x + i
		if pos >= 0 && pos < len(bgRunes) {
			// Only overlay non-space characters from content
			if r != ' ' && r != '\x1b' { // Don't overlay spaces or ANSI escape sequences
				bgRunes[pos] = r
			}
		}
	}

	return string(bgRunes)
}

func main() {
	// Menu items
	menuItems := []string{
		"NEW GAME",
		"CONTINUE",
		"COLLECTION",
		"OPTIONS",
		"QUIT",
	}

	// Create initial model
	m := Model{
		menuItems:    menuItems,
		selectedItem: 0,
		width:        80,
		height:       24,
	}

	// Create wavy grid (purple theme)
	m.grid = waves.NewGrid(m.width, m.height)
	m.grid.SetGridSize(10)
	m.grid.SetColors(waves.GridColors{
		Intersection: lipgloss.Color("129"), // Purple
		Vertical:     lipgloss.Color("61"),  // Dark purple
		Horizontal:   lipgloss.Color("61"),  // Dark purple
	})

	// Create metaball engine with 4 blobs
	m.metaballs = metaballs.NewEngine(m.width, m.height)

	// Blob 1: Cyan (top-left quadrant)
	m.metaballs.AddBlob(metaballs.NewBlob(
		float64(m.width)/4,
		float64(m.height)/3,
		0.3, 0.2,
		6,
		lipgloss.Color("51"), // Cyan
	))

	// Blob 2: Magenta (top-right quadrant)
	m.metaballs.AddBlob(metaballs.NewBlob(
		float64(m.width)*3/4,
		float64(m.height)/2,
		-0.25, 0.15,
		7,
		lipgloss.Color("201"), // Magenta
	))

	// Blob 3: Yellow (bottom-center)
	m.metaballs.AddBlob(metaballs.NewBlob(
		float64(m.width)/2,
		float64(m.height)*2/3,
		0.2, -0.3,
		5,
		lipgloss.Color("226"), // Yellow
	))

	// Blob 4: Purple (left-center)
	m.metaballs.AddBlob(metaballs.NewBlob(
		float64(m.width)/3,
		float64(m.height)/4,
		-0.15, 0.25,
		6,
		lipgloss.Color("129"), // Purple
	))

	// Create rainbow cycler
	m.rainbow = rainbow.NewCycler()

	// Run the program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
