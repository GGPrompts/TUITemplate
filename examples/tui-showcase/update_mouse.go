package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// update_mouse.go - Mouse Event Handling
// Purpose: All mouse input processing
// When to extend: Add new mouse interactions or clickable elements here

// handleMouseEvent handles mouse input
func (m model) handleMouseEvent(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	if !m.config.UI.MouseEnabled {
		return m, nil
	}

	switch msg.Type {
	case tea.MouseLeft:
		return m.handleLeftClick(msg)

	case tea.MouseRight:
		return m.handleRightClick(msg)

	case tea.MouseWheelUp:
		return m.handleWheelUp(msg)

	case tea.MouseWheelDown:
		return m.handleWheelDown(msg)

	case tea.MouseMotion:
		// Handle mouse motion if needed (for hover effects)
		return m.handleMouseMotion(msg)
	}

	return m, nil
}

// handleLeftClick handles left mouse button clicks
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	x, y := msg.X, msg.Y

	// Menu bar clicks (first priority)
	if m.isInMenuBar(x, y) {
		menuKey := m.getMenuAtPosition(x)
		if menuKey != "" {
			if m.menuOpen && m.activeMenu == menuKey {
				// Clicking same menu closes it
				m.menuOpen = false
				m.activeMenu = ""
				m.selectedMenuItem = -1
			} else {
				// Open menu
				m.menuOpen = true
				m.activeMenu = menuKey
				m.selectedMenuItem = -1
			}
			return m, nil
		}
	}

	// Dropdown menu clicks (second priority)
	if m.menuOpen && m.isInDropdown(x, y) {
		itemIndex := m.getMenuItemAtPosition(y)
		if itemIndex >= 0 {
			menus := getMenus()
			menu := menus[m.activeMenu]
			if itemIndex < len(menu.Items) {
				item := menu.Items[itemIndex]
				// Execute action if not separator or disabled
				if !item.IsSeparator && !item.Disabled {
					return m.executeMenuAction(item.Action)
				}
			}
		}
		return m, nil
	}

	// Click outside menu closes it
	if m.menuOpen {
		m.menuOpen = false
		m.activeMenu = ""
		m.selectedMenuItem = -1
		return m, nil
	}

	// Tab clicks
	if m.isInTabBar(x, y) {
		return m.handleTabBarClick(x, y)
	}

	// Dynamic panels click (when on tab 5)
	if m.currentTab == 5 {
		return m.handleDynamicPanelClick(msg)
	}

	// Panel clicks - detect which panel was clicked
	if m.currentLayout == "dual_pane" {
		contentWidth, _ := m.calculateLayout()
		leftWidth, _ := m.calculateDualPaneLayout()

		// Rough detection of which pane was clicked
		if x < leftWidth {
			m.lastClicked = "left-pane"
			m.statusMsg = "Clicked on left-pane"
		} else if x >= leftWidth && x < contentWidth {
			m.lastClicked = "right-pane"
			m.statusMsg = "Clicked on right-pane"
		}
		return m, nil
	}

	// Check if clicked on UI elements
	if m.isInTitleBar(x, y) {
		return m.handleTitleBarClick(x, y)
	}

	if m.isInStatusBar(x, y) {
		return m.handleStatusBarClick(x, y)
	}

	return m, nil
}

// handleRightClick handles right mouse button clicks
func (m model) handleRightClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	x, y := msg.X, msg.Y

	// Example: show context menu
	// return m.showContextMenu(x, y)

	_ = x
	_ = y
	return m, nil
}

// handleWheelUp handles mouse wheel scroll up
func (m model) handleWheelUp(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Scroll up in the focused component
	return m.moveUp()
}

// handleWheelDown handles mouse wheel scroll down
func (m model) handleWheelDown(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Scroll down in the focused component
	return m.moveDown()
}

// handleMouseMotion handles mouse movement (for hover effects)
func (m model) handleMouseMotion(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	m.mouseX = msg.X
	m.mouseY = msg.Y

	// Detect hovering over dropdown menu items (when menu is open)
	if m.menuOpen && m.isInDropdown(msg.X, msg.Y) {
		itemIndex := m.getMenuItemAtPosition(msg.Y)
		if itemIndex >= 0 {
			menus := getMenus()
			menu := menus[m.activeMenu]
			if itemIndex < len(menu.Items) && !menu.Items[itemIndex].IsSeparator {
				m.selectedMenuItem = itemIndex
				return m, nil
			}
		}
	}

	// Detect hovering over menu bar (auto-switch menu when one is open)
	if m.isInMenuBar(msg.X, msg.Y) {
		menuKey := m.getMenuAtPosition(msg.X)
		if menuKey != "" && m.menuOpen && menuKey != m.activeMenu {
			// Auto-switch to hovered menu if a menu is already open
			m.activeMenu = menuKey
			m.selectedMenuItem = -1
			return m, nil
		}
	}

	// Detect hovering over layout buttons
	if msg.Y == 1 && m.config.UI.ShowTitle {
		if msg.X >= 2 && msg.X <= 4 {
			m.hoveredItem = "button-1"
		} else if msg.X >= 6 && msg.X <= 8 {
			m.hoveredItem = "button-2"
		} else if msg.X >= 10 && msg.X <= 12 {
			m.hoveredItem = "button-3"
		} else if msg.X >= 14 && msg.X <= 16 {
			m.hoveredItem = "button-4"
		} else {
			m.hoveredItem = ""
		}
		return m, nil
	}

	// Detect hovering over tabs
	if m.currentLayout == "tabbed" && m.isInTabBar(msg.X, msg.Y) {
		// Tab detection logic
		tabNames := []string{"Overview", "Content", "Settings", "Borders", "Colors", "Dynamic"}
		xPos := 0
		for i, name := range tabNames {
			tabWidth := len(name) + 4 // "[ " + name + " ]"
			if msg.X >= xPos && msg.X < xPos+tabWidth {
				m.hoveredItem = "tab-" + string(rune(i+'0'))
				return m, nil
			}
			xPos += tabWidth + 1 // +1 for space between tabs
		}
	}

	m.hoveredItem = ""
	return m, nil
}

// Helper functions for click region detection

func (m model) isInTitleBar(x, y int) bool {
	if !m.config.UI.ShowTitle {
		return false
	}
	return y < 2
}

func (m model) isInStatusBar(x, y int) bool {
	if !m.config.UI.ShowStatus {
		return false
	}
	return y >= m.height-1
}

func (m model) handleTitleBarClick(x, y int) (tea.Model, tea.Cmd) {
	// Example: click on breadcrumb navigation
	// or click on window control buttons
	_ = x
	_ = y
	return m, nil
}

func (m model) handleStatusBarClick(x, y int) (tea.Model, tea.Cmd) {
	// Example: click on status bar items
	_ = x
	_ = y
	return m, nil
}

// Double-click detection (if needed)
type clickTracker struct {
	lastClickX    int
	lastClickY    int
	lastClickTime int64
}

var tracker clickTracker

func (m model) isDoubleClick(msg tea.MouseMsg) bool {
	// Implement double-click detection
	// Compare with tracker.lastClickTime
	// Reset tracker.lastClickX, tracker.lastClickY, tracker.lastClickTime
	return false
}

// isInTabBar checks if position is in the tab bar
func (m model) isInTabBar(x, y int) bool {
	if m.currentLayout != "tabbed" {
		return false
	}
	titleHeight := 0
	if m.config.UI.ShowTitle {
		titleHeight = 2
	}
	return y == titleHeight
}

// handleTabBarClick handles clicks on tabs
func (m model) handleTabBarClick(x, y int) (tea.Model, tea.Cmd) {
	tabNames := []string{"Single", "Dual", "Multi", "Borders", "Colors", "Dynamic", "Forms", "Tables", "Dialogs", "Progress", "Tree", "Mobile"}
	xPos := 0

	for i, name := range tabNames {
		// Tab format: "[ name ]" (4 chars) + Padding(0,1) (2 chars) = 6 + len(name)
		tabWidth := len(name) + 6 // "[ " + name + " ]" + padding(0,1)
		if x >= xPos && x < xPos+tabWidth {
			m.currentTab = i
			m.statusMsg = "Tab: " + name
			m.lastClicked = "tab-" + string(rune(i+'0'))
			return m, nil
		}
		xPos += tabWidth + 1 // +1 for space between tabs
	}

	return m, nil
}

// handleDynamicPanelClick handles clicks on dynamic panels (tab 5) - 3-panel layout
func (m model) handleDynamicPanelClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Calculate content area boundaries
	titleHeight := 3 // title + instructions + separator + tab bar
	if m.config.UI.ShowTitle {
		titleHeight += 1 // add tab bar
	}
	statusHeight := 1
	contentStartY := titleHeight
	contentEndY := m.height - statusHeight

	// Only process clicks within content area
	if msg.Y < contentStartY || msg.Y >= contentEndY {
		return m, nil
	}

	// Calculate content dimensions
	contentHeight := contentEndY - contentStartY
	contentWidth := m.width

	// Calculate 3-panel layout dimensions
	leftWidth, _, topHeight, _ := m.calculateThreePanelLayout(contentWidth, contentHeight)

	// Calculate relative Y position
	relY := msg.Y - contentStartY

	// Determine which panel was clicked
	if relY < topHeight {
		// Click in top panels area (left or right)
		if msg.X < leftWidth {
			m.focusedPanel = "left"
			m.statusMsg = "Focused LEFT panel (clicked)"
		} else if msg.X > leftWidth { // Skip divider at X == leftWidth
			m.focusedPanel = "right"
			m.statusMsg = "Focused RIGHT panel (clicked)"
		}
	} else {
		// Click in bottom panel area
		m.focusedPanel = "bottom"
		m.statusMsg = "Focused BOTTOM panel (clicked)"
	}

	return m, nil
}
