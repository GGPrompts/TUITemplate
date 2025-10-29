package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/GGPrompts/TUITemplate/lib/effects/metaballs"
)

// Model represents the application state
type Model struct {
	engine *metaballs.Engine
	width  int
	height int
}

// tickMsg is sent on every animation frame
type tickMsg time.Time

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
		if m.engine != nil {
			m.engine.Resize(msg.Width, msg.Height)
		}
		return m, nil

	case tickMsg:
		if m.engine != nil {
			m.engine.Update()
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if m.engine == nil {
		return "Loading..."
	}

	// Render the metaballs
	metaballsView := m.engine.Render()

	// Add loading text overlay
	loadingText := "Loading..."
	textStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("0")).
		Padding(0, 2)

	styledText := textStyle.Render(loadingText)

	// Center the text
	lines := lipgloss.Height(metaballsView)
	textY := lines / 2

	metaballLines := lipgloss.SplitLines(metaballsView)
	if textY >= 0 && textY < len(metaballLines) {
		textWidth := lipgloss.Width(styledText)
		textX := (m.width - textWidth) / 2
		if textX < 0 {
			textX = 0
		}

		// Insert the text into the middle of the metaballs
		metaballLines[textY] = overlayText(metaballLines[textY], styledText, textX)
	}

	result := lipgloss.JoinVertical(lipgloss.Left, metaballLines...)

	// Add controls hint at bottom
	controls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("\n\nPress 'q' to quit")

	return result + controls
}

// overlayText overlays text onto a line at the specified x position
func overlayText(line, text string, x int) string {
	lineRunes := []rune(line)
	textRunes := []rune(text)

	if x < 0 || x >= len(lineRunes) {
		return line
	}

	// Simple overlay - replace characters
	for i, r := range textRunes {
		if x+i < len(lineRunes) {
			lineRunes[x+i] = r
		}
	}

	return string(lineRunes)
}

func main() {
	// Create initial model
	m := Model{
		width:  80,
		height: 24,
	}

	// Create metaball engine with 3 colorful blobs
	m.engine = metaballs.NewEngine(m.width, m.height)

	// Add 3 blobs in triangle formation
	m.engine.AddBlob(metaballs.NewBlob(
		float64(m.width)/3,
		float64(m.height)/2,
		0.3, 0.2,
		6,
		lipgloss.Color("51"), // Cyan
	))

	m.engine.AddBlob(metaballs.NewBlob(
		float64(m.width)*2/3,
		float64(m.height)/2,
		-0.25, 0.15,
		7,
		lipgloss.Color("201"), // Magenta
	))

	m.engine.AddBlob(metaballs.NewBlob(
		float64(m.width)/2,
		float64(m.height)*2/3,
		0.2, -0.3,
		5,
		lipgloss.Color("226"), // Yellow
	))

	// Run the program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
