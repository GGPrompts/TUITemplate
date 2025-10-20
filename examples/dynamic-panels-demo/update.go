package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// update.go - Main Update Dispatcher
// Purpose: Message dispatching and non-input event handling
// When to extend: Add new message types or top-level event handlers here

// Init is called when the program starts
func (m model) Init() tea.Cmd {
	// Return any initialization commands here
	// Example:
	// return tea.Batch(
	//     loadDataCmd(),
	//     tea.EnterAltScreen,
	// )
	return nil
}

// Update handles all messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window resize
	case tea.WindowSizeMsg:
		m.setSize(msg.Width, msg.Height)
		return m, nil

	// Keyboard input
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	// Mouse input
	case tea.MouseMsg:
		return m.handleMouseEvent(msg)

	// Custom messages
	case errMsg:
		m.err = msg.err
		m.statusMsg = "Error: " + msg.err.Error()
		return m, nil

	case statusMsg:
		m.statusMsg = msg.message
		return m, nil

	// Add handlers for your custom messages here
	// Example:
	// case itemSelectedMsg:
	//     return m.handleItemSelected(msg)
	//
	// case dataLoadedMsg:
	//     return m.handleDataLoaded(msg)
	}

	return m, nil
}

// Helper functions for message handling

// sendStatus creates a status message command
func sendStatus(message string) tea.Cmd {
	return func() tea.Msg {
		return statusMsg{message: message}
	}
}

// sendError creates an error message command
func sendError(err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg{err: err}
	}
}

// isSpecialKey checks if a key is a special key (not printable)
func isSpecialKey(key tea.KeyMsg) bool {
	return key.Type != tea.KeyRunes
}
