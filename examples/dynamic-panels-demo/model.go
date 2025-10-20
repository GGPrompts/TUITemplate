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
		focusedPanel:     "left",
		accordionMode:    true, // Start with accordion mode enabled
		showThreePanel:   false, // Start with 2-panel layout
		statusMsg:        "Click panels or press 1/2/3 to focus | Press 'a' for accordion | 'm' for 3-panel mode | 'q' to quit",

		// Initialize panel content
		leftContent: []string{
			"LEFT PANEL",
			"",
			"This is the left panel.",
			"It contains scrollable content.",
			"",
			"When focused, it expands!",
			"",
			"Try clicking or pressing '1'",
			"",
			"Line 10",
			"Line 11",
			"Line 12",
			"Line 13",
			"Line 14",
			"Line 15",
		},
		rightContent: []string{
			"RIGHT PANEL",
			"",
			"This is the right panel.",
			"It also contains scrollable content.",
			"",
			"Press '2' to focus this panel",
			"",
			"When focused, it gets more space!",
			"",
			"Line 10",
			"Line 11",
			"Line 12",
			"Line 13",
			"Line 14",
			"Line 15",
		},
		bottomContent: []string{
			"BOTTOM PANEL",
			"",
			"This panel appears in 3-panel mode.",
			"Press 'm' to toggle 3-panel mode",
			"Press '3' to focus this panel",
			"",
			"Line 7",
			"Line 8",
			"Line 9",
			"Line 10",
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
		contentHeight -= 3 // title bar (3 lines: title, instructions, separator)
	}
	if m.config.UI.ShowStatus {
		contentHeight -= 1 // status bar height
	}

	// Account for panel borders (top + bottom = 2 lines)
	// Since panels are rendered WITH borders, we need to reserve space for them
	contentHeight -= 2

	return contentWidth, contentHeight
}

// shouldUseVerticalStack checks if terminal is too narrow for side-by-side panels
func (m model) shouldUseVerticalStack() bool {
	contentWidth, _ := m.calculateLayout()
	// Minimum comfortable width for side-by-side panels is ~80 chars
	// Below that, stack vertically
	return contentWidth < 80
}

// calculateDualPaneLayout computes left and right pane widths using weight-based system
// This implements the LazyGit pattern: focused panel gets 2x weight (accordion mode)
func (m model) calculateDualPaneLayout() (int, int) {
	contentWidth, _ := m.calculateLayout()

	dividerWidth := 1 // Space for divider
	availableWidth := contentWidth - dividerWidth

	// Enforce minimum panel width (at least 30 chars per panel)
	minPanelWidth := 30
	if availableWidth < minPanelWidth*2 {
		// Too narrow, but still provide reasonable values
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

// calculateVerticalStackLayout computes heights for vertically stacked panels
func (m model) calculateVerticalStackLayout() (int, int) {
	_, contentHeight := m.calculateLayout()

	dividerHeight := 1 // Space for horizontal divider
	availableHeight := contentHeight - dividerHeight

	// Calculate weights based on focus and accordion mode
	topWeight, bottomWeight := 1, 1

	if m.accordionMode {
		if m.focusedPanel == "left" {
			topWeight = 2 // Top (left) panel focused
		} else if m.focusedPanel == "right" {
			bottomWeight = 2 // Bottom (right) panel focused
		}
	}

	// Calculate heights based on weights
	totalWeight := topWeight + bottomWeight
	topHeight := (availableHeight * topWeight) / totalWeight
	bottomHeight := availableHeight - topHeight

	return topHeight, bottomHeight
}

// calculateThreePanelLayout computes top (left/right) and bottom heights
// Bottom panel gets 1/3 of space when unfocused, 2/3 when focused (accordion)
func (m model) calculateThreePanelLayout() (int, int, int, int) {
	contentWidth, contentHeight := m.calculateLayout()

	// Calculate horizontal split for top panels (same as dual pane)
	dividerWidth := 1
	availableWidth := contentWidth - dividerWidth

	leftWeight, rightWeight := 1, 1
	if m.accordionMode && m.focusedPanel == "left" {
		leftWeight = 2
	} else if m.accordionMode && m.focusedPanel == "right" {
		rightWeight = 2
	}

	totalHorzWeight := leftWeight + rightWeight
	leftWidth := (availableWidth * leftWeight) / totalHorzWeight
	rightWidth := availableWidth - leftWidth

	// Calculate vertical split
	dividerHeight := 1
	availableHeight := contentHeight - dividerHeight

	topWeight, bottomWeight := 2, 1 // Default: top gets 2/3, bottom gets 1/3

	if m.accordionMode && m.focusedPanel == "bottom" {
		bottomWeight = 2 // When bottom focused, it gets 2/3
		topWeight = 1
	}

	totalVertWeight := topWeight + bottomWeight
	topHeight := (availableHeight * topWeight) / totalVertWeight
	bottomHeight := availableHeight - topHeight

	return leftWidth, rightWidth, topHeight, bottomHeight
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
