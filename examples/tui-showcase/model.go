package main

// model.go - Model Management
// Purpose: Model initialization and layout calculations
// When to extend: Add new initialization logic or layout calculation functions here

// initialModel creates the initial application state
func initialModel(cfg Config) model {
	return model{
		config:           cfg,
		width:            0,
		height:           0,
		focusedComponent: "main",
		statusMsg:        "Navigate: Tab/Shift+Tab or use menus | Dynamic tab: Press 'a' for accordion | Press q to quit",
		currentLayout:    "tabbed",
		currentTab:       0,
		focusedPanel:     "left",
		accordionMode:    true,
		leftContent: []string{
			"LEFT PANEL",
			"",
			"This panel demonstrates",
			"dynamic resizing based on focus.",
			"",
			"Click or press '1' to focus.",
			"Press 'a' to toggle accordion mode.",
		},
		rightContent: []string{
			"RIGHT PANEL",
			"",
			"This panel also resizes",
			"when focused.",
			"",
			"Click or press '2' to focus.",
			"Watch the weights change!",
		},
		bottomContent: []string{
			"BOTTOM PANEL",
			"",
			"This panel appears below",
			"the top panels.",
			"",
			"Click or press '3' to focus.",
			"Expands to 66% height when focused!",
		},
	}
}

// setSize updates the model dimensions and recalculates layouts
func (m *model) setSize(width, height int) {
	m.width = width
	m.height = height

	// Recalculate any layout-dependent values here
	// Example:
	// m.viewportHeight = height - 4 // account for title and status bars
	// m.maxVisible = m.viewportHeight - 2
}

// calculateLayout computes layout dimensions based on config
func (m model) calculateLayout() (int, int) {
	contentWidth := m.width
	contentHeight := m.height

	// Adjust for UI elements
	if m.config.UI.ShowTitle {
		contentHeight -= 3 // title bar + button bar height
	}
	if m.config.UI.ShowStatus {
		contentHeight -= 1 // status bar height
	}

	// CRITICAL: Account for panel borders (fixes overflow issue)
	contentHeight -= 2 // top + bottom borders

	return contentWidth, contentHeight
}

// calculateDualPaneLayout computes left and right pane widths
func (m model) calculateDualPaneLayout() (int, int) {
	contentWidth, _ := m.calculateLayout()

	dividerWidth := 0
	if m.config.Layout.ShowDivider {
		dividerWidth = 1
	}

	leftWidth := int(float64(contentWidth-dividerWidth) * m.config.Layout.SplitRatio)
	rightWidth := contentWidth - leftWidth - dividerWidth

	return leftWidth, rightWidth
}

// Helper functions for common operations

// getContentArea returns the available content area dimensions
func (m model) getContentArea() (width, height int) {
	return m.calculateLayout()
}

// isValidSize checks if the terminal size is sufficient
func (m model) isValidSize() bool {
	return m.width >= 40 && m.height >= 10
}

// calculateDynamicPaneLayout computes left and right pane widths using weight-based system
// This implements the LazyGit pattern: focused panel gets 2x weight (accordion mode)
func (m model) calculateDynamicPaneLayout() (int, int) {
	contentWidth, _ := m.calculateLayout()

	dividerWidth := 1 // Space for divider
	availableWidth := contentWidth - dividerWidth

	// Enforce minimum panel width
	minPanelWidth := 30
	if availableWidth < minPanelWidth*2 {
		half := availableWidth / 2
		return half, availableWidth - half
	}

	// Calculate weights based on focus and accordion mode
	leftWeight, rightWeight := 1, 1

	if m.accordionMode {
		if m.focusedPanel == "left" {
			leftWeight = 2 // Focused panel gets 2x weight
		} else if m.focusedPanel == "right" {
			rightWeight = 2
		}
	}

	// Calculate widths based on weights
	totalWeight := leftWeight + rightWeight
	leftWidth := (availableWidth * leftWeight) / totalWeight
	rightWidth := availableWidth - leftWidth

	// Ensure minimum widths
	if leftWidth < minPanelWidth {
		leftWidth = minPanelWidth
		rightWidth = availableWidth - leftWidth
	} else if rightWidth < minPanelWidth {
		rightWidth = minPanelWidth
		leftWidth = availableWidth - rightWidth
	}

	return leftWidth, rightWidth
}

// calculateThreePanelLayout computes dynamic panel dimensions using weight-based system
// This implements the LazyGit pattern for 3-panel layout: focused panel gets 2x weight
func (m model) calculateThreePanelLayout(availableWidth, availableHeight int) (leftWidth, rightWidth, topHeight, bottomHeight int) {
	// Calculate horizontal split for top panels
	dividerWidth := 1
	topAvailableWidth := availableWidth - dividerWidth

	leftWeight, rightWeight := 1, 1
	if m.accordionMode {
		if m.focusedPanel == "left" {
			leftWeight = 2 // Focused panel gets 2x weight (66%)
		} else if m.focusedPanel == "right" {
			rightWeight = 2
		}
	}

	totalHorzWeight := leftWeight + rightWeight
	leftWidth = (topAvailableWidth * leftWeight) / totalHorzWeight
	rightWidth = topAvailableWidth - leftWidth

	// Calculate vertical split
	dividerHeight := 0 // No divider between top and bottom rows
	totalAvailableHeight := availableHeight - dividerHeight

	topWeight, bottomWeight := 1, 1
	if m.accordionMode {
		if m.focusedPanel == "bottom" {
			bottomWeight = 2 // Bottom panel focused gets 2x weight
		} else {
			topWeight = 2 // Either top panel focused gets 2x weight
		}
	} else {
		// When accordion is off, give top panels more space by default (2:1 ratio)
		topWeight = 2
	}

	totalVertWeight := topWeight + bottomWeight
	topHeight = (totalAvailableHeight * topWeight) / totalVertWeight
	bottomHeight = totalAvailableHeight - topHeight

	return leftWidth, rightWidth, topHeight, bottomHeight
}
