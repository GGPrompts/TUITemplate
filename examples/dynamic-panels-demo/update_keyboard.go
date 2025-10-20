package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// update_keyboard.go - Keyboard Event Handling
// Purpose: All keyboard input processing
// When to extend: Add new keyboard shortcuts or key bindings here

// handleKeyPress handles keyboard input
func (m model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Global keybindings (work in all modes)
	switch {
	case key.Matches(msg, keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, keys.Help):
		return m.showHelp()

	case key.Matches(msg, keys.Refresh):
		return m.refresh()
	}

	// Mode-specific keybindings
	switch m.focusedComponent {
	case "main":
		return m.handleMainKeys(msg)

	// Add handlers for other components/modes
	// case "dialog":
	//     return m.handleDialogKeys(msg)
	//
	// case "menu":
	//     return m.handleMenuKeys(msg)
	}

	return m, nil
}

// handleMainKeys handles keys in main view
func (m model) handleMainKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {

	// Panel focus switching
	case "1":
		m.focusedPanel = "left"
		m.statusMsg = "Focused LEFT panel"
		return m, nil

	case "2":
		m.focusedPanel = "right"
		m.statusMsg = "Focused RIGHT panel"
		return m, nil

	case "3":
		if m.showThreePanel {
			m.focusedPanel = "bottom"
			m.statusMsg = "Focused BOTTOM panel"
		} else {
			m.statusMsg = "Bottom panel not visible (press 'm' to enable 3-panel mode)"
		}
		return m, nil

	// Toggle accordion mode (focused panel gets 2x weight)
	case "a", "A":
		m.accordionMode = !m.accordionMode
		if m.accordionMode {
			m.statusMsg = "Accordion mode: ON (focused panel gets 2x space)"
		} else {
			m.statusMsg = "Accordion mode: OFF (all panels equal size)"
		}
		return m, nil

	// Toggle 3-panel mode
	case "m", "M":
		m.showThreePanel = !m.showThreePanel
		if m.showThreePanel {
			m.statusMsg = "3-Panel mode: ON (added bottom panel)"
			// Default to left panel when enabling 3-panel
			if m.focusedPanel == "bottom" {
				m.focusedPanel = "left"
			}
		} else {
			m.statusMsg = "3-Panel mode: OFF (only left/right panels)"
			// If bottom was focused, switch to left
			if m.focusedPanel == "bottom" {
				m.focusedPanel = "left"
			}
		}
		return m, nil

	// Cycle through panels with Tab
	case "tab":
		return m.cycleFocusForward()

	case "shift+tab":
		return m.cycleFocusBackward()

	// Navigation arrows also cycle panels
	case "left", "h":
		if m.focusedPanel == "right" {
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
		return m, nil

	case "right", "l":
		if m.focusedPanel == "left" {
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
		}
		return m, nil

	case "down", "j":
		if m.showThreePanel && (m.focusedPanel == "left" || m.focusedPanel == "right") {
			m.focusedPanel = "bottom"
			m.statusMsg = "Focused BOTTOM panel"
		}
		return m, nil

	case "up", "k":
		if m.focusedPanel == "bottom" {
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
		return m, nil
	}

	return m, nil
}

// Focus cycling functions

func (m model) cycleFocusForward() (tea.Model, tea.Cmd) {
	if m.showThreePanel {
		// Cycle: left -> right -> bottom -> left
		switch m.focusedPanel {
		case "left":
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
		case "right":
			m.focusedPanel = "bottom"
			m.statusMsg = "Focused BOTTOM panel"
		case "bottom":
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
	} else {
		// Cycle: left -> right -> left
		if m.focusedPanel == "left" {
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
		} else {
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
	}
	return m, nil
}

func (m model) cycleFocusBackward() (tea.Model, tea.Cmd) {
	if m.showThreePanel {
		// Cycle: left -> bottom -> right -> left
		switch m.focusedPanel {
		case "left":
			m.focusedPanel = "bottom"
			m.statusMsg = "Focused BOTTOM panel"
		case "bottom":
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
		case "right":
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
	} else {
		// Cycle: left -> right -> left
		if m.focusedPanel == "left" {
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
		} else {
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
		}
	}
	return m, nil
}

// Navigation helper functions

func (m model) moveUp() (tea.Model, tea.Cmd) {
	// Implement up navigation
	// Example: m.cursor = max(0, m.cursor-1)
	return m, nil
}

func (m model) moveDown() (tea.Model, tea.Cmd) {
	// Implement down navigation
	// Example: m.cursor = min(len(m.items)-1, m.cursor+1)
	return m, nil
}

func (m model) moveLeft() (tea.Model, tea.Cmd) {
	// Implement left navigation
	return m, nil
}

func (m model) moveRight() (tea.Model, tea.Cmd) {
	// Implement right navigation
	return m, nil
}

func (m model) pageUp() (tea.Model, tea.Cmd) {
	// Implement page up
	// Example: m.cursor = max(0, m.cursor-m.viewportHeight)
	return m, nil
}

func (m model) pageDown() (tea.Model, tea.Cmd) {
	// Implement page down
	// Example: m.cursor = min(len(m.items)-1, m.cursor+m.viewportHeight)
	return m, nil
}

func (m model) moveToTop() (tea.Model, tea.Cmd) {
	// Example: m.cursor = 0
	return m, nil
}

func (m model) moveToBottom() (tea.Model, tea.Cmd) {
	// Example: m.cursor = len(m.items) - 1
	return m, nil
}

// Action helper functions

func (m model) selectItem() (tea.Model, tea.Cmd) {
	// Implement item selection
	return m, nil
}

func (m model) toggleSelection() (tea.Model, tea.Cmd) {
	// Implement toggle selection
	return m, nil
}

func (m model) switchFocus() (tea.Model, tea.Cmd) {
	// Implement focus switching between components
	return m, nil
}

func (m model) showHelp() (tea.Model, tea.Cmd) {
	// Show help dialog
	m.statusMsg = "Help: q=quit, ?=help, ↑↓=navigate, enter=select"
	return m, nil
}

func (m model) refresh() (tea.Model, tea.Cmd) {
	// Refresh the current view
	m.statusMsg = "Refreshed"
	return m, nil
}

// Key bindings definition
type keyMap struct {
	Quit    key.Binding
	Help    key.Binding
	Refresh key.Binding
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Select  key.Binding
	Toggle  key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "refresh"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "right"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Toggle: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle"),
	),
}
