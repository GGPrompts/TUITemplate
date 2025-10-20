package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// menu.go - Dropdown Menu System
// Purpose: Menu bar rendering and interaction logic
// When to extend: Add new menus or menu items here

// getMenus returns all available menus
func getMenus() map[string]Menu {
	return map[string]Menu{
		"file": {
			Label: "File",
			Items: []MenuItem{
				{Label: "New", Action: "file-new", Shortcut: "Ctrl+N"},
				{Label: "Open", Action: "file-open", Shortcut: "Ctrl+O"},
				{Label: "Save", Action: "file-save", Shortcut: "Ctrl+S"},
				{Label: "Save As...", Action: "file-save-as", Shortcut: "Ctrl+Shift+S"},
				{IsSeparator: true},
				{Label: "Quit", Action: "quit", Shortcut: "Q"},
			},
		},
		"view": {
			Label: "View",
			Items: []MenuItem{
				{Label: "Single Pane", Action: "switch-layout-single", Shortcut: "1"},
				{Label: "Dual Pane", Action: "switch-layout-dual", Shortcut: "2"},
				{Label: "Multi-Panel", Action: "switch-layout-multi", Shortcut: "3"},
				{Label: "Tabbed", Action: "switch-layout-tabbed", Shortcut: "4"},
				{IsSeparator: true},
				{Label: "Overview Tab", Action: "switch-tab-0"},
				{Label: "Content Tab", Action: "switch-tab-1"},
				{Label: "Settings Tab", Action: "switch-tab-2"},
				{Label: "Borders Tab", Action: "switch-tab-3"},
				{Label: "Colors Tab", Action: "switch-tab-4"},
				{Label: "Dynamic Panels Tab", Action: "switch-tab-5"},
			},
		},
		"components": {
			Label: "Components",
			Items: []MenuItem{
				{Label: "Forms", Action: "show-forms", Disabled: true},
				{Label: "Tables", Action: "show-tables", Disabled: true},
				{Label: "Dialogs", Action: "show-dialogs", Disabled: true},
				{Label: "Progress Bars", Action: "show-progress", Disabled: true},
				{Label: "Tree View", Action: "show-tree", Disabled: true},
				{IsSeparator: true},
				{Label: "(Coming soon...)", Disabled: true},
			},
		},
		"help": {
			Label: "Help",
			Items: []MenuItem{
				{Label: "Keyboard Shortcuts", Action: "show-help-keys", Shortcut: "?"},
				{Label: "About Layout Demo", Action: "show-about"},
				{IsSeparator: true},
				{Label: "GitHub Repository", Action: "open-github"},
			},
		},
	}
}

// getMenuOrder returns the order of menus in the menu bar
func getMenuOrder() []string {
	return []string{"file", "view", "components", "help"}
}

// renderMenuBar renders the menu bar
func (m model) renderMenuBar() string {
	menus := getMenus()
	menuOrder := getMenuOrder()

	var renderedMenus []string

	for _, menuKey := range menuOrder {
		menu := menus[menuKey]

		// Style based on active state
		var style lipgloss.Style
		if m.activeMenu == menuKey && m.menuOpen {
			style = menuActiveStyle
		} else {
			style = menuInactiveStyle
		}

		// Render menu label (style already has Padding(0,1))
		renderedMenu := style.Render(menu.Label)
		renderedMenus = append(renderedMenus, renderedMenu)
	}

	// Join with single space (matching tab behavior)
	menuBarContent := strings.Join(renderedMenus, " ")
	padding := m.width - lipgloss.Width(menuBarContent)
	if padding < 0 {
		padding = 0
	}

	return menuBarContent + strings.Repeat(" ", padding)
}

// renderActiveDropdown renders the currently active dropdown menu
func (m model) renderActiveDropdown() string {
	if !m.menuOpen || m.activeMenu == "" {
		return ""
	}

	menus := getMenus()
	menu, exists := menus[m.activeMenu]
	if !exists {
		return ""
	}

	// Calculate dropdown position (X coordinate)
	menuOrder := getMenuOrder()
	xPos := 2
	for _, menuKey := range menuOrder {
		if menuKey == m.activeMenu {
			break
		}
		menuLabel := menus[menuKey].Label
		xPos += len(menuLabel) + 3 // +3 for " " + label + " "
	}

	// Build dropdown panel
	var lines []string
	maxWidth := 0

	// First pass: calculate max width
	for _, item := range menu.Items {
		if item.IsSeparator {
			continue
		}
		width := len(item.Label)
		if item.Shortcut != "" {
			width += len(item.Shortcut) + 3 // spacing
		}
		if width > maxWidth {
			maxWidth = width
		}
	}

	// Add padding
	maxWidth += 4 // 2 chars padding on each side
	if maxWidth < 20 {
		maxWidth = 20
	}

	// Second pass: render items
	for i, item := range menu.Items {
		if item.IsSeparator {
			lines = append(lines, strings.Repeat("─", maxWidth-2))
			continue
		}

		// Determine style
		var itemStyle lipgloss.Style
		if item.Disabled {
			itemStyle = menuItemDisabledStyle
		} else if i == m.selectedMenuItem {
			itemStyle = menuItemSelectedStyle
		} else {
			itemStyle = menuItemStyle
		}

		// Build item line
		label := item.Label
		shortcut := item.Shortcut

		// Pad label
		labelWidth := maxWidth - 4
		if shortcut != "" {
			labelWidth -= len(shortcut) + 1
		}

		line := " " + padRight(label, labelWidth)
		if shortcut != "" {
			line += " " + shortcut
		}
		line += " "

		lines = append(lines, itemStyle.Render(line))
	}

	// Create dropdown panel
	dropdown := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Width(maxWidth).
		Render(strings.Join(lines, "\n"))

	// Position dropdown (this returns just the panel, positioning done in view)
	return dropdown
}

// getMenuXPosition calculates the X position for a menu
func (m model) getMenuXPosition(menuKey string) int {
	menus := getMenus()
	menuOrder := getMenuOrder()

	xPos := 0 // Start at 0 (matching tab behavior)
	for _, key := range menuOrder {
		if key == menuKey {
			return xPos
		}
		menu := menus[key]
		// Use actual rendered width to match tab calculation
		var style lipgloss.Style
		if m.activeMenu == key && m.menuOpen {
			style = menuActiveStyle
		} else {
			style = menuInactiveStyle
		}
		renderedMenu := style.Render(menu.Label)
		xPos += lipgloss.Width(renderedMenu) + 1 // +1 for space separator
	}
	return xPos
}

// isInMenuBar checks if position is in the menu bar
func (m model) isInMenuBar(x, y int) bool {
	if !m.config.UI.ShowTitle {
		return false
	}
	// Menu bar is on line 2 (after title line)
	return y == 1
}

// getMenuAtPosition returns which menu is at the given X position
func (m model) getMenuAtPosition(x int) string {
	menus := getMenus()
	menuOrder := getMenuOrder()

	xPos := 0 // Start at 0 (matching tab and render behavior)
	for _, menuKey := range menuOrder {
		menu := menus[menuKey]

		// Calculate actual rendered width
		var style lipgloss.Style
		if m.activeMenu == menuKey && m.menuOpen {
			style = menuActiveStyle
		} else {
			style = menuInactiveStyle
		}
		renderedMenu := style.Render(menu.Label)
		menuWidth := lipgloss.Width(renderedMenu)

		if x >= xPos && x < xPos+menuWidth {
			return menuKey
		}

		xPos += menuWidth + 1 // +1 for space separator
	}

	return ""
}

// isInDropdown checks if position is within the active dropdown
func (m model) isInDropdown(x, y int) bool {
	if !m.menuOpen || m.activeMenu == "" {
		return false
	}

	// Dropdown starts at y=2 (after menu bar)
	if y < 2 {
		return false
	}

	menus := getMenus()
	menu, exists := menus[m.activeMenu]
	if !exists {
		return false
	}

	// Calculate dropdown bounds
	menuX := m.getMenuXPosition(m.activeMenu)

	// Count non-separator items for height
	height := 0
	for range menu.Items {
		height++
	}
	height += 2 // borders

	// Estimate width (will be at least 20)
	maxWidth := 20
	for _, item := range menu.Items {
		if item.IsSeparator {
			continue
		}
		width := len(item.Label)
		if item.Shortcut != "" {
			width += len(item.Shortcut) + 3
		}
		width += 4 // padding
		if width > maxWidth {
			maxWidth = width
		}
	}

	return x >= menuX && x < menuX+maxWidth && y >= 2 && y < 2+height
}

// getMenuItemAtPosition returns the menu item index at the given Y position in dropdown
func (m model) getMenuItemAtPosition(y int) int {
	if !m.menuOpen || m.activeMenu == "" {
		return -1
	}

	// Dropdown content starts at y=3 (after border)
	itemY := y - 3
	if itemY < 0 {
		return -1
	}

	menus := getMenus()
	menu, exists := menus[m.activeMenu]
	if !exists {
		return -1
	}

	if itemY >= len(menu.Items) {
		return -1
	}

	return itemY
}

// executeMenuAction executes a menu item action
func (m model) executeMenuAction(action string) (tea.Model, tea.Cmd) {
	switch action {
	case "quit":
		return m, tea.Quit

	// Layout switching
	case "switch-layout-single":
		m.currentLayout = "single"
		m.statusMsg = "Switched to Single Pane layout"
	case "switch-layout-dual":
		m.currentLayout = "dual_pane"
		m.statusMsg = "Switched to Dual Pane layout"
	case "switch-layout-multi":
		m.currentLayout = "multi_panel"
		m.statusMsg = "Switched to Multi-Panel layout"
	case "switch-layout-tabbed":
		m.currentLayout = "tabbed"
		m.statusMsg = "Switched to Tabbed layout"

	// Tab switching (only works in tabbed layout)
	case "switch-tab-0":
		m.currentLayout = "tabbed"
		m.currentTab = 0
		m.statusMsg = "Switched to Overview tab"
	case "switch-tab-1":
		m.currentLayout = "tabbed"
		m.currentTab = 1
		m.statusMsg = "Switched to Content tab"
	case "switch-tab-2":
		m.currentLayout = "tabbed"
		m.currentTab = 2
		m.statusMsg = "Switched to Settings tab"
	case "switch-tab-3":
		m.currentLayout = "tabbed"
		m.currentTab = 3
		m.statusMsg = "Switched to Borders tab"
	case "switch-tab-4":
		m.currentLayout = "tabbed"
		m.currentTab = 4
		m.statusMsg = "Switched to Colors tab"
	case "switch-tab-5":
		m.currentLayout = "tabbed"
		m.currentTab = 5
		m.statusMsg = "Switched to Dynamic Panels tab"

	// File operations (placeholder)
	case "file-new":
		m.statusMsg = "File → New (not implemented)"
	case "file-open":
		m.statusMsg = "File → Open (not implemented)"
	case "file-save":
		m.statusMsg = "File → Save (not implemented)"
	case "file-save-as":
		m.statusMsg = "File → Save As (not implemented)"

	// Help
	case "show-help-keys":
		m.statusMsg = "Help: q=quit, 1-4=layouts, Tab/Shift+Tab=switch tabs, ?=help"
	case "show-about":
		m.statusMsg = "Layout Demo - TUI Template System | GitHub: GGPrompts/TUITemplate"
	case "open-github":
		m.statusMsg = "GitHub: https://github.com/GGPrompts/TUITemplate"

	default:
		m.statusMsg = "Action: " + action + " (not implemented)"
	}

	// Close menu after action
	m.menuOpen = false
	m.activeMenu = ""
	m.selectedMenuItem = -1

	return m, nil
}
