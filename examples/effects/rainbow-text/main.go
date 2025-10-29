package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/GGPrompts/TUITemplate/lib/effects/rainbow"
)

// Model represents the application state
type Model struct {
	cycler         *rainbow.Cycler
	width          int
	height         int
	speed          int
	currentPalette int
	palettes       [][]lipgloss.Color
	paletteNames   []string
}

// tickMsg is sent on every animation frame
type tickMsg time.Time

// ASCII art for "RAINBOW"
var asciiArt = []string{
	"██████╗  █████╗ ██╗███╗   ██╗██████╗  ██████╗ ██╗    ██╗",
	"██╔══██╗██╔══██╗██║████╗  ██║██╔══██╗██╔═══██╗██║    ██║",
	"██████╔╝███████║██║██╔██╗ ██║██████╔╝██║   ██║██║ █╗ ██║",
	"██╔══██╗██╔══██║██║██║╚██╗██║██╔══██╗██║   ██║██║███╗██║",
	"██║  ██║██║  ██║██║██║ ╚████║██████╔╝╚██████╔╝╚███╔███╔╝",
	"╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝  ╚══╝╚══╝ ",
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
		return m, nil

	case tickMsg:
		if m.cycler != nil {
			m.cycler.Update()
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "+", "=":
			// Speed up
			if m.speed > 1 {
				m.speed--
				m.cycler.SetSpeed(m.speed)
			}

		case "-", "_":
			// Slow down
			if m.speed < 10 {
				m.speed++
				m.cycler.SetSpeed(m.speed)
			}

		case "c", "C":
			// Cycle through palettes
			m.currentPalette = (m.currentPalette + 1) % len(m.palettes)
			m.cycler.SetColors(m.palettes[m.currentPalette])
		}
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if m.cycler == nil {
		return "Loading..."
	}

	var sections []string

	// Title: Rainbow animated ASCII art
	titleLines := m.cycler.RenderLines(asciiArt)
	sections = append(sections, titleLines)

	// Separator
	separator := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(strings.Repeat("─", 60))
	sections = append(sections, "", separator, "")

	// Single line example
	singleLineStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Bold(true)
	singleLineText := m.cycler.Render("The quick brown fox jumps over the lazy dog")
	sections = append(sections, singleLineStyle.Render("Single line: ")+singleLineText, "")

	// Multi-line example
	multiLineStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Bold(true)
	sections = append(sections, multiLineStyle.Render("Multi-line:"))

	multiLineText := []string{
		"Terminal UIs can be beautiful too!",
		"Rainbow colors cycling through text.",
		"Created with the TUI Effects library.",
	}
	multiLineRainbow := m.cycler.RenderLines(multiLineText)
	sections = append(sections, multiLineRainbow, "")

	// Current settings
	settingsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	settings := fmt.Sprintf("Current Speed: %d frames/shift", m.speed)
	paletteInfo := fmt.Sprintf("Palette: %s (%d colors)",
		m.paletteNames[m.currentPalette],
		len(m.palettes[m.currentPalette]))

	sections = append(sections, settingsStyle.Render(settings))
	sections = append(sections, settingsStyle.Render(paletteInfo), "")

	// Controls
	controlsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))
	controls := controlsStyle.Render("+/-: Speed | C: Change palette | Q: Quit")
	sections = append(sections, controls)

	// Join all sections
	content := strings.Join(sections, "\n")

	// Center the content
	contentLines := strings.Split(content, "\n")
	totalHeight := len(contentLines)
	topPadding := (m.height - totalHeight) / 2
	if topPadding < 0 {
		topPadding = 0
	}

	// Add top padding
	paddedLines := make([]string, topPadding)
	for i := 0; i < topPadding; i++ {
		paddedLines[i] = ""
	}
	paddedLines = append(paddedLines, contentLines...)

	// Center each line horizontally
	var centeredLines []string
	for _, line := range paddedLines {
		lineWidth := lipgloss.Width(line)
		leftPadding := (m.width - lineWidth) / 2
		if leftPadding < 0 {
			leftPadding = 0
		}
		centeredLines = append(centeredLines, strings.Repeat(" ", leftPadding)+line)
	}

	return strings.Join(centeredLines, "\n")
}

func main() {
	// Define color palettes
	rainbowPalette := []lipgloss.Color{
		lipgloss.Color("196"), // Red
		lipgloss.Color("208"), // Orange
		lipgloss.Color("226"), // Yellow
		lipgloss.Color("46"),  // Green
		lipgloss.Color("51"),  // Cyan
		lipgloss.Color("21"),  // Blue
		lipgloss.Color("201"), // Magenta
	}

	pastelPalette := []lipgloss.Color{
		lipgloss.Color("213"), // Pink
		lipgloss.Color("228"), // Pale Yellow
		lipgloss.Color("156"), // Pale Green
		lipgloss.Color("117"), // Pale Blue
		lipgloss.Color("183"), // Lavender
	}

	firePalette := []lipgloss.Color{
		lipgloss.Color("196"), // Red
		lipgloss.Color("202"), // Orange-Red
		lipgloss.Color("208"), // Orange
		lipgloss.Color("214"), // Orange-Yellow
		lipgloss.Color("226"), // Yellow
		lipgloss.Color("255"), // White
	}

	neonPalette := []lipgloss.Color{
		lipgloss.Color("51"),  // Cyan
		lipgloss.Color("201"), // Magenta
		lipgloss.Color("46"),  // Green
		lipgloss.Color("226"), // Yellow
		lipgloss.Color("129"), // Purple
	}

	// Create initial model
	m := Model{
		width:          80,
		height:         24,
		speed:          5,
		currentPalette: 0,
		palettes: [][]lipgloss.Color{
			rainbowPalette,
			pastelPalette,
			firePalette,
			neonPalette,
		},
		paletteNames: []string{
			"Rainbow",
			"Pastel",
			"Fire",
			"Neon",
		},
	}

	// Create rainbow cycler with default rainbow palette
	m.cycler = rainbow.NewCycler()
	m.cycler.SetColors(rainbowPalette)
	m.cycler.SetSpeed(m.speed)

	// Run the program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
