package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tea.MouseMsg:
		return m.handleMouseEvent(msg)
	}

	return m, nil
}

// handleKeyPress handles keyboard input
func (m model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "tab", "shift+tab":
		return m.cycleFocus(msg.String() == "shift+tab"), nil

	case "h":
		// Move focus left
		return m.moveFocusLeft(), nil

	case "l":
		// Move focus right
		return m.moveFocusRight(), nil

	case "up", "k":
		return m.handleUp(), nil

	case "down", "j":
		return m.handleDown(), nil

	case "enter":
		return m.handleEnter(), nil

	case "r":
		m.addLog("Refreshed view")
		return m, nil

	case "a", "A":
		// Toggle accordion mode
		m.accordionMode = !m.accordionMode
		if m.accordionMode {
			m.addLog("Accordion mode: ON (focused panel gets 2x space)")
			m.statusMsg = "Accordion mode: ON • Focused panel expands to 66%"
		} else {
			m.addLog("Accordion mode: OFF (fixed 2:1 layout)")
			m.statusMsg = "Accordion mode: OFF • Fixed 2:1 layout"
		}
		return m, nil
	}

	return m, nil
}

// handleMouseEvent handles mouse events
func (m model) handleMouseEvent(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.MouseLeft:
		// Click to focus panel
		for panelID, bounds := range m.panelBounds {
			if bounds.contains(msg.X, msg.Y) {
				if m.focusedPanel != panelID {
					m.focusedPanel = panelID
					m.addLog(fmt.Sprintf("Clicked on %s panel", panelID.String()))
				}
				break
			}
		}
		return m, nil

	case tea.MouseWheelUp:
		// Scroll up in focused panel
		return m.handleMouseScrollUp(), nil

	case tea.MouseWheelDown:
		// Scroll down in focused panel
		return m.handleMouseScrollDown(), nil
	}

	return m, nil
}

// handleMouseScrollUp handles mouse wheel scroll up
func (m model) handleMouseScrollUp() model {
	switch m.focusedPanel {
	case LeftPanel:
		// Scroll up in file list
		if m.cursor > 0 {
			m.cursor--
			m.updateDetails()
		}

	case TopRightPanel:
		// Could scroll details up if implemented
		m.addLog("Scrolling details up")

	case BottomRightPanel:
		// Scroll logs up
		if m.logScroll > 0 {
			m.logScroll--
		}
	}
	return m
}

// handleMouseScrollDown handles mouse wheel scroll down
func (m model) handleMouseScrollDown() model {
	switch m.focusedPanel {
	case LeftPanel:
		// Scroll down in file list
		if m.cursor < len(m.files)-1 {
			m.cursor++
			m.updateDetails()
		}

	case TopRightPanel:
		// Could scroll details down if implemented
		m.addLog("Scrolling details down")

	case BottomRightPanel:
		// Scroll logs down
		if m.logScroll < len(m.logs)-10 {
			m.logScroll++
		}
	}
	return m
}

// cycleFocus cycles through panels (Tab)
func (m model) cycleFocus(reverse bool) model {
	currentIdx := 0
	for i, p := range m.panels {
		if p == m.focusedPanel {
			currentIdx = i
			break
		}
	}

	if reverse {
		currentIdx--
		if currentIdx < 0 {
			currentIdx = len(m.panels) - 1
		}
	} else {
		currentIdx++
		if currentIdx >= len(m.panels) {
			currentIdx = 0
		}
	}

	m.focusedPanel = m.panels[currentIdx]
	m.addLog(fmt.Sprintf("Switched to %s panel", m.focusedPanel.String()))
	return m
}

// moveFocusLeft moves focus to the left panel
func (m model) moveFocusLeft() model {
	switch m.focusedPanel {
	case TopRightPanel, BottomRightPanel:
		m.focusedPanel = LeftPanel
		m.addLog("Switched to Files panel")
	}
	return m
}

// moveFocusRight moves focus to the right panels
func (m model) moveFocusRight() model {
	switch m.focusedPanel {
	case LeftPanel:
		m.focusedPanel = TopRightPanel
		m.addLog("Switched to Details panel")
	case TopRightPanel:
		m.focusedPanel = BottomRightPanel
		m.addLog("Switched to Logs panel")
	case BottomRightPanel:
		m.focusedPanel = TopRightPanel
		m.addLog("Switched to Details panel")
	}
	return m
}

// handleUp handles up arrow / k key
func (m model) handleUp() model {
	switch m.focusedPanel {
	case LeftPanel:
		// Navigate up in file list
		if m.cursor > 0 {
			m.cursor--
			m.updateDetails()
		}

	case TopRightPanel:
		// Could scroll details up
		m.addLog("Scrolling details up")

	case BottomRightPanel:
		// Scroll logs up
		if m.logScroll > 0 {
			m.logScroll--
		}
	}
	return m
}

// handleDown handles down arrow / j key
func (m model) handleDown() model {
	switch m.focusedPanel {
	case LeftPanel:
		// Navigate down in file list
		if m.cursor < len(m.files)-1 {
			m.cursor++
			m.updateDetails()
		}

	case TopRightPanel:
		// Could scroll details down
		m.addLog("Scrolling details down")

	case BottomRightPanel:
		// Scroll logs down
		if m.logScroll < len(m.logs)-10 {
			m.logScroll++
		}
	}
	return m
}

// handleEnter handles enter key
func (m model) handleEnter() model {
	switch m.focusedPanel {
	case LeftPanel:
		file := m.getSelectedFile()
		m.addLog(fmt.Sprintf("Selected: %s", file))
		m.updateDetails()

	case TopRightPanel:
		m.addLog("Action in Details panel")

	case BottomRightPanel:
		m.addLog("Action in Logs panel")
	}
	return m
}

// updateDetails updates the details panel based on selected file
func (m *model) updateDetails() {
	file := m.getSelectedFile()
	if file == "" {
		m.details = "No file selected"
		return
	}

	// Generate some example details
	m.details = fmt.Sprintf(`File: %s

Type: Go source file
Size: 2.4 KB
Modified: 2025-01-15 14:30:22

This is a sample details panel that would
show file contents, metadata, or other
relevant information based on the selected
item in the left panel.

In a real application, this could show:
• File contents
• Git diff
• File history
• Metadata
• Preview`, file)
}
