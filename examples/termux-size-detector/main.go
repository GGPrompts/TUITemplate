package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Terminal Size Detector for Termux
// Shows real-time terminal dimensions - perfect for finding Termux sizes

type model struct {
	width       int
	height      int
	maxWidth    int
	maxHeight   int
	minWidth    int
	minHeight   int
	history     []string
	updateCount int
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func initialModel() model {
	return model{
		width:     0,
		height:    0,
		maxWidth:  0,
		maxHeight: 0,
		minWidth:  9999,
		minHeight: 9999,
		history:   []string{},
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateCount++

		// Track max/min
		if m.width > m.maxWidth {
			m.maxWidth = m.width
		}
		if m.height > m.maxHeight {
			m.maxHeight = m.height
		}
		if m.width < m.minWidth {
			m.minWidth = m.width
		}
		if m.height < m.minHeight {
			m.minHeight = m.height
		}

		// Add to history
		record := fmt.Sprintf("[%d] %dx%d", m.updateCount, m.width, m.height)
		m.history = append(m.history, record)
		if len(m.history) > 10 {
			m.history = m.history[1:]
		}

		return m, nil

	case tickMsg:
		return m, tickCmd()
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Detecting terminal size..."
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6")).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("6")).
		Padding(0, 1)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("4")).
		Padding(1, 2)

	highlightStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("2")).
		Bold(true)

	dimStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8"))

	// Title
	title := titleStyle.Render("📏 TERMINAL SIZE DETECTOR")

	// Current size (large display)
	current := fmt.Sprintf(
		"Current Size: %s × %s\n"+
			"             (Width × Height)",
		highlightStyle.Render(fmt.Sprintf("%d", m.width)),
		highlightStyle.Render(fmt.Sprintf("%d", m.height)),
	)

	// Stats
	stats := fmt.Sprintf(
		"Max: %dx%d  |  Min: %dx%d  |  Updates: %d",
		m.maxWidth, m.maxHeight,
		m.minWidth, m.minHeight,
		m.updateCount,
	)

	// Size categories
	category := m.getSizeCategory()
	categoryText := fmt.Sprintf("Category: %s", highlightStyle.Render(category))

	// Recommendations
	recommendations := m.getRecommendations()

	// History
	historyText := "Recent Changes:\n"
	if len(m.history) == 0 {
		historyText += dimStyle.Render("  (no changes yet)")
	} else {
		for i := len(m.history) - 1; i >= 0; i-- {
			historyText += dimStyle.Render("  "+m.history[i]) + "\n"
		}
	}

	// Visual representation
	visual := m.renderVisualSize()

	// Instructions
	instructions := dimStyle.Render(
		"\nInstructions:\n" +
			"  • Open/close keyboard to see size changes\n" +
			"  • Resize terminal to see updates\n" +
			"  • Press 'q' to quit\n",
	)

	// Combine everything
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		current,
		stats,
		categoryText,
		"",
		visual,
		"",
		boxStyle.Render(recommendations),
		"",
		historyText,
		instructions,
	)

	return content
}

func (m model) getSizeCategory() string {
	// Categorize the terminal size
	if m.height <= 10 {
		return "MICRO (Termux with keyboard)"
	} else if m.height <= 20 {
		return "SMALL (Termux without keyboard)"
	} else if m.height <= 30 {
		return "MEDIUM"
	} else if m.height <= 45 {
		return "LARGE"
	}
	return "EXTRA LARGE"
}

func (m model) getRecommendations() string {
	if m.height <= 10 {
		return "💡 MICRO LAYOUT TIPS (Termux + Keyboard):\n" +
			"   • Use single line for title\n" +
			"   • Show only 3-5 items at a time\n" +
			"   • Use compact status bar\n" +
			"   • Avoid multi-panel layouts\n" +
			"   • Consider horizontal scrolling"
	} else if m.height <= 20 {
		return "💡 SMALL LAYOUT TIPS (Termux):\n" +
			"   • Use minimal title bar\n" +
			"   • Show 10-15 items\n" +
			"   • Single pane works best\n" +
			"   • Keep margins tight"
	} else if m.height <= 30 {
		return "💡 MEDIUM LAYOUT TIPS:\n" +
			"   • All standard layouts work\n" +
			"   • Dual pane is comfortable\n" +
			"   • Title + content + status bar fits well"
	}
	return "💡 LARGE+ LAYOUT TIPS:\n" +
		"   • Full features available\n" +
		"   • Multi-panel layouts work great\n" +
		"   • Plenty of space for content"
}

func (m model) renderVisualSize() string {
	barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	fillStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

	// Width bar (out of 120)
	maxW := 120
	widthBars := m.width * 40 / maxW
	if widthBars > 40 {
		widthBars = 40
	}
	widthBar := fillStyle.Render(repeat("█", widthBars)) + barStyle.Render(repeat("░", 40-widthBars))

	// Height bar (out of 60)
	maxH := 60
	heightBars := m.height * 20 / maxH
	if heightBars > 20 {
		heightBars = 20
	}
	heightBar := fillStyle.Render(repeat("█", heightBars)) + barStyle.Render(repeat("░", 20-heightBars))

	return fmt.Sprintf(
		"Width:  [%s] %d\nHeight: [%s] %d",
		widthBar, m.width,
		heightBar, m.height,
	)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
