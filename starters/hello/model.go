package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width   int
	height  int
	counter int
}

func initialModel() model {
	return model{
		counter: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case " ", "enter":
			m.counter++
			return m, nil

		case "r":
			m.counter = 0
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return ""
	}

	// Styles
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#61AFEF")).
		Bold(true).
		Padding(1, 2)

	contentStyle := lipgloss.NewStyle().
		Padding(1, 2)

	counterStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#98C379")).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5C6370")).
		Padding(1, 2)

	// Content
	title := titleStyle.Render("Hello, TUITemplate! 👋")

	content := contentStyle.Render(
		fmt.Sprintf("You pressed the button %s times!\n\n"+
			"This is a minimal example showing TUITemplate's architecture:\n"+
			"  • main.go - Entry point (10 lines)\n"+
			"  • model.go - Model, Update, View (70 lines)\n\n"+
			"Press Space/Enter to increment counter\n"+
			"Press R to reset\n"+
			"Press Q to quit",
			counterStyle.Render(fmt.Sprintf("%d", m.counter)),
		),
	)

	help := helpStyle.Render("Space/Enter: Count | R: Reset | Q: Quit")

	// Layout
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		content,
		help,
	)
}
