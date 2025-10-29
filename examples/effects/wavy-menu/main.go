package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/GGPrompts/TUITemplate/lib/effects/waves"
)

// Model represents the application state
type Model struct {
	grid         *waves.Grid
	selectedItem int
	menuItems    []string
	width        int
	height       int
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
		if m.grid != nil {
			m.grid.Resize(msg.Width, msg.Height)
		}
		return m, nil

	case tickMsg:
		if m.grid != nil {
			m.grid.Update()
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
	if m.grid == nil {
		return "Loading..."
	}

	// Render the wavy grid background
	background := m.grid.Render()

	// Create menu with styled items
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

	// Create menu box
	menu := strings.Join(menuLines, "\n")
	menuBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("129")). // Purple
		Padding(1, 2).
		Render(menu)

	// Add controls hint
	controls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("↑↓: Navigate | Enter: Select | Q: Quit")

	// Combine menu and controls
	content := lipgloss.JoinVertical(lipgloss.Center, menuBox, "", controls)

	// Calculate centering positions
	menuHeight := lipgloss.Height(content)
	menuWidth := lipgloss.Width(content)
	topPadding := (m.height - menuHeight) / 2
	leftPadding := (m.width - menuWidth) / 2

	if topPadding < 0 {
		topPadding = 0
	}
	if leftPadding < 0 {
		leftPadding = 0
	}

	// Overlay menu onto background
	backgroundLines := strings.Split(background, "\n")
	contentLines := strings.Split(content, "\n")

	// Ensure we have enough background lines
	for len(backgroundLines) < m.height {
		backgroundLines = append(backgroundLines, strings.Repeat(" ", m.width))
	}

	// Overlay content onto background
	for i, line := range contentLines {
		targetY := topPadding + i
		if targetY >= 0 && targetY < len(backgroundLines) {
			backgroundLines[targetY] = overlayLine(backgroundLines[targetY], line, leftPadding, m.width)
		}
	}

	return strings.Join(backgroundLines, "\n")
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

	// Overlay content
	for i, r := range contentRunes {
		pos := x + i
		if pos >= 0 && pos < len(bgRunes) {
			bgRunes[pos] = r
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

	// Create wavy grid with purple theme
	m.grid = waves.NewGrid(m.width, m.height)
	m.grid.SetGridSize(10)
	m.grid.SetColors(waves.GridColors{
		Intersection: lipgloss.Color("129"), // Purple
		Vertical:     lipgloss.Color("61"),  // Dark purple
		Horizontal:   lipgloss.Color("61"),  // Dark purple
	})

	// Run the program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
