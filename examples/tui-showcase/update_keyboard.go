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
	// Effect mode - allow escape to return to showcase
	if m.activeEffect != "" {
		switch msg.String() {
		case "esc", "q":
			m.activeEffect = ""
			m.statusMsg = "Returned to showcase"
			return m, nil
		}
		// Consume all other keys when in effect mode
		return m, nil
	}

	// Menu navigation (highest priority when menu is open)
	if m.menuOpen {
		switch msg.String() {
		case "esc":
			// Close menu
			m.menuOpen = false
			m.activeMenu = ""
			m.selectedMenuItem = -1
			return m, nil

		case "up", "k":
			// Move selection up
			if m.selectedMenuItem > 0 {
				m.selectedMenuItem--
				// Skip separators
				menus := getMenus()
				menu := menus[m.activeMenu]
				for m.selectedMenuItem >= 0 && menu.Items[m.selectedMenuItem].IsSeparator {
					m.selectedMenuItem--
				}
			}
			return m, nil

		case "down", "j":
			// Move selection down
			menus := getMenus()
			menu := menus[m.activeMenu]
			if m.selectedMenuItem < len(menu.Items)-1 {
				m.selectedMenuItem++
				// Skip separators
				for m.selectedMenuItem < len(menu.Items) && menu.Items[m.selectedMenuItem].IsSeparator {
					m.selectedMenuItem++
				}
			}
			return m, nil

		case "enter":
			// Execute selected item
			if m.selectedMenuItem >= 0 {
				menus := getMenus()
				menu := menus[m.activeMenu]
				if m.selectedMenuItem < len(menu.Items) {
					item := menu.Items[m.selectedMenuItem]
					if !item.IsSeparator && !item.Disabled {
						return m.executeMenuAction(item.Action)
					}
				}
			}
			return m, nil

		case "left", "h":
			// Switch to previous menu
			menuOrder := getMenuOrder()
			for i, key := range menuOrder {
				if key == m.activeMenu {
					prevIndex := (i - 1 + len(menuOrder)) % len(menuOrder)
					m.activeMenu = menuOrder[prevIndex]
					m.selectedMenuItem = -1
					return m, nil
				}
			}
			return m, nil

		case "right", "l":
			// Switch to next menu
			menuOrder := getMenuOrder()
			for i, key := range menuOrder {
				if key == m.activeMenu {
					nextIndex := (i + 1) % len(menuOrder)
					m.activeMenu = menuOrder[nextIndex]
					m.selectedMenuItem = -1
					return m, nil
				}
			}
			return m, nil
		}
	}

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
	// Special handling for Dynamic Panels tab (tab 5) - 3 panels
	if m.currentTab == 5 {
		switch msg.String() {
		case "1":
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
			return m, nil
		case "2":
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
			return m, nil
		case "3":
			m.focusedPanel = "bottom"
			m.statusMsg = "Focused BOTTOM panel"
			return m, nil
		case "a", "A":
			m.accordionMode = !m.accordionMode
			if m.accordionMode {
				m.statusMsg = "Accordion mode: ON (focused panel gets 2x space)"
			} else {
				m.statusMsg = "Accordion mode: OFF (panels equal size)"
			}
			return m, nil
		}
	}

	// Special handling for 4-Panel tab (tab 12) - 4 panels
	if m.currentTab == 12 {
		switch msg.String() {
		case "1":
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel"
			return m, nil
		case "2":
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel"
			return m, nil
		case "3":
			m.focusedPanel = "footer"
			m.statusMsg = "Focused FOOTER panel (bottom)"
			return m, nil
		case "4":
			m.focusedPanel = "header"
			m.statusMsg = "Focused HEADER panel (top)"
			return m, nil
		case "a", "A":
			m.accordionMode = !m.accordionMode
			if m.accordionMode {
				m.statusMsg = "Accordion mode: ON (focused panel gets up to 66% space)"
			} else {
				m.statusMsg = "Accordion mode: OFF (panels at default size)"
			}
			return m, nil
		}
	}

	switch msg.String() {

	// Tab navigation
	case "tab":
		m.currentTab = (m.currentTab + 1) % 13
		tabNames := []string{"Single", "Dual", "Multi", "Borders", "Colors", "Dynamic", "Forms", "Tables", "Dialogs", "Progress", "Tree", "Mobile", "4-Panel"}
		m.statusMsg = "Tab: " + tabNames[m.currentTab]
		return m, nil

	case "shift+tab":
		m.currentTab = (m.currentTab - 1 + 13) % 13
		tabNames := []string{"Single", "Dual", "Multi", "Borders", "Colors", "Dynamic", "Forms", "Tables", "Dialogs", "Progress", "Tree", "Mobile", "4-Panel"}
		m.statusMsg = "Tab: " + tabNames[m.currentTab]
		return m, nil

	// Navigation
	case "up", "k":
		return m.moveUp()

	case "down", "j":
		return m.moveDown()

	case "left", "h":
		return m.moveLeft()

	case "right", "l":
		return m.moveRight()

	case "pgup":
		return m.pageUp()

	case "pgdown":
		return m.pageDown()

	case "home", "g":
		return m.moveToTop()

	case "end", "G":
		return m.moveToBottom()

	// Actions
	case "enter":
		return m.selectItem()

	case " ": // space
		return m.toggleSelection()
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
