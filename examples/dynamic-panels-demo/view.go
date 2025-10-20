package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// view.go - View Rendering
// Purpose: Dynamic panel rendering with focus-based resizing
// Pattern: LazyGit-inspired weight-based layout system

// View renders the entire application
func (m model) View() string {
	// Check if terminal size is sufficient
	if !m.isValidSize() {
		return m.renderMinimalView()
	}

	// Handle errors
	if m.err != nil {
		return m.renderErrorView()
	}

	// Render based on panel mode
	if m.showThreePanel {
		return m.renderThreePanelLayout()
	}
	return m.renderTwoPanelLayout()
}

// renderTwoPanelLayout renders left and right panels with dynamic widths
func (m model) renderTwoPanelLayout() string {
	var sections []string

	// Title bar
	sections = append(sections, m.renderTitleBar())

	// Check if we should stack vertically (portrait/narrow mode)
	if m.shouldUseVerticalStack() {
		return m.renderVerticalStackLayout()
	}

	// Calculate pane dimensions using weight-based system
	leftWidth, rightWidth := m.calculateDualPaneLayout()

	// Get content height
	_, contentHeight := m.calculateLayout()

	// Render left panel
	leftPanel := m.renderPanel("left", leftWidth, contentHeight, m.leftContent)

	// Render divider
	divider := m.renderVerticalDivider(contentHeight)

	// Render right panel
	rightPanel := m.renderPanel("right", rightWidth, contentHeight, m.rightContent)

	// Join panes horizontally
	panes := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, divider, rightPanel)
	sections = append(sections, panes)

	// Status bar
	sections = append(sections, m.renderStatusBar())

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderVerticalStackLayout renders panels stacked vertically (for narrow terminals)
func (m model) renderVerticalStackLayout() string {
	var sections []string

	// Title bar
	sections = append(sections, m.renderTitleBar())

	// Calculate heights using weight-based system
	topHeight, bottomHeight := m.calculateVerticalStackLayout()

	// Get content width (full width for each panel)
	contentWidth, _ := m.calculateLayout()

	// Render top panel (left content, but full width)
	topPanel := m.renderPanel("left", contentWidth, topHeight, m.leftContent)

	// Render horizontal divider
	hDivider := m.renderHorizontalDivider()

	// Render bottom panel (right content, but full width)
	bottomPanel := m.renderPanel("right", contentWidth, bottomHeight, m.rightContent)

	// Join sections vertically
	sections = append(sections, topPanel)
	sections = append(sections, hDivider)
	sections = append(sections, bottomPanel)

	// Status bar
	sections = append(sections, m.renderStatusBar())

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderThreePanelLayout renders top (left/right) and bottom panels
func (m model) renderThreePanelLayout() string {
	var sections []string

	// Title bar
	sections = append(sections, m.renderTitleBar())

	// Calculate panel dimensions
	leftWidth, rightWidth, topHeight, bottomHeight := m.calculateThreePanelLayout()

	// Render top left panel
	leftPanel := m.renderPanel("left", leftWidth, topHeight, m.leftContent)

	// Render vertical divider
	vDivider := m.renderVerticalDivider(topHeight)

	// Render top right panel
	rightPanel := m.renderPanel("right", rightWidth, topHeight, m.rightContent)

	// Join top panels horizontally
	topPanes := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, vDivider, rightPanel)

	// Render horizontal divider
	hDivider := m.renderHorizontalDivider()

	// Render bottom panel
	contentWidth, _ := m.calculateLayout()
	bottomPanel := m.renderPanel("bottom", contentWidth, bottomHeight, m.bottomContent)

	// Join all sections
	sections = append(sections, topPanes)
	sections = append(sections, hDivider)
	sections = append(sections, bottomPanel)

	// Status bar
	sections = append(sections, m.renderStatusBar())

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderPanel renders a single panel with border and content
func (m model) renderPanel(panelName string, width, height int, content []string) string {
	isFocused := m.focusedPanel == panelName

	// Create border style based on focus
	borderStyle := lipgloss.RoundedBorder()
	borderColor := lipgloss.Color("240") // Dim gray
	titleColor := lipgloss.Color("252")  // Light gray

	if isFocused {
		borderColor = colorPrimary       // Bright blue
		titleColor = colorPrimary        // Bright blue
	}

	// Calculate max text width (account for borders and padding)
	maxTextWidth := width - 4 // -2 for borders, -2 for padding
	if maxTextWidth < 1 {
		maxTextWidth = 1
	}

	// Create panel title
	title := strings.ToUpper(panelName) + " PANEL"
	if isFocused {
		title += " ●" // Indicator for focused panel
	}
	// Truncate title if too long
	title = truncateString(title, maxTextWidth)

	// Calculate weight for display
	weight := 1
	if m.accordionMode && isFocused {
		weight = 2
	}

	subtitle := fmt.Sprintf("Weight: %d | Size: %dx%d", weight, width, height)
	// Truncate subtitle if too long
	subtitle = truncateString(subtitle, maxTextWidth)

	// Create header
	header := lipgloss.NewStyle().
		Foreground(titleColor).
		Bold(true).
		Render(title)

	subheader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Render(subtitle)

	// Build content lines
	var lines []string
	lines = append(lines, header)
	lines = append(lines, subheader)
	lines = append(lines, "")

	// Calculate exact content area height
	// height = total height available for panel
	// - 2 for top/bottom borders
	// - 2 for top/bottom padding (0 padding set, but border takes space)
	// - 3 for header lines (title + subtitle + blank line)
	innerHeight := height - 2 // Remove borders
	availableContentLines := innerHeight - 3 // Remove header lines

	if availableContentLines < 1 {
		availableContentLines = 1
	}

	// Add content lines (truncate if too long to prevent wrapping)
	for i := 0; i < availableContentLines && i < len(content); i++ {
		line := truncateString(content[i], maxTextWidth)
		lines = append(lines, line)
	}

	// Fill remaining space to ensure consistent height
	for len(lines) < innerHeight {
		lines = append(lines, "")
	}

	contentStr := strings.Join(lines, "\n")

	// Create styled panel - let height be determined by content
	panelStyle := lipgloss.NewStyle().
		Border(borderStyle).
		BorderForeground(borderColor).
		Width(width - 2) // Account for left/right borders

	return panelStyle.Render(contentStr)
}

// renderTitleBar renders the title and controls
func (m model) renderTitleBar() string {
	title := titleStyle.Render("🎨 Dynamic Panels Demo")

	// Show mode indicators
	mode := ""
	if m.accordionMode {
		mode += "[Accordion: ON] "
	} else {
		mode += "[Accordion: OFF] "
	}

	if m.showThreePanel {
		mode += "[3-Panel Mode]"
	} else if m.shouldUseVerticalStack() {
		mode += "[Vertical Stack]"
	} else {
		mode += "[2-Panel Mode]"
	}

	modeStyle := lipgloss.NewStyle().
		Foreground(colorSecondary).
		Render(mode)

	titleLine := lipgloss.JoinHorizontal(lipgloss.Left, title, "  ", modeStyle)

	// Instructions
	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Render("Focus: 1/2/3 or click | Toggle: a=accordion, m=3-panel | q=quit")

	return lipgloss.JoinVertical(lipgloss.Left,
		titleLine,
		instructions,
		strings.Repeat("─", m.width),
	)
}

// renderStatusBar renders the status bar
func (m model) renderStatusBar() string {
	// Show focused panel and mouse position
	leftInfo := fmt.Sprintf("Focused: %s", strings.ToUpper(m.focusedPanel))

	rightInfo := fmt.Sprintf("Mouse: (%d,%d)", m.mouseX, m.mouseY)

	// Status message (if any)
	statusText := m.statusMsg
	if m.hoveredItem != "" {
		statusText = fmt.Sprintf("Hovering: %s", m.hoveredItem)
	}

	leftStyle := lipgloss.NewStyle().Foreground(colorPrimary)
	rightStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

	// Calculate spacing
	leftPart := leftStyle.Render(leftInfo)
	rightPart := rightStyle.Render(rightInfo)
	statusPart := lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render(statusText)

	spacingLeft := m.width - lipgloss.Width(leftPart) - lipgloss.Width(statusPart) - lipgloss.Width(rightPart) - 4
	if spacingLeft < 2 {
		spacingLeft = 2
	}

	return leftPart + " | " + statusPart + strings.Repeat(" ", spacingLeft) + rightPart
}

// renderVerticalDivider renders a vertical divider between panels
func (m model) renderVerticalDivider(height int) string {
	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	var lines []string
	for i := 0; i < height; i++ {
		lines = append(lines, "│")
	}

	return dividerStyle.Render(strings.Join(lines, "\n"))
}

// renderHorizontalDivider renders a horizontal divider
func (m model) renderHorizontalDivider() string {
	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	return dividerStyle.Render(strings.Repeat("─", m.width))
}

// renderMinimalView renders a minimal view for small terminals
func (m model) renderMinimalView() string {
	msg := "Terminal too small!\nResize to at least 40x10"
	return lipgloss.NewStyle().
		Foreground(colorError).
		Bold(true).
		Render(msg)
}

// renderErrorView renders error messages
func (m model) renderErrorView() string {
	errorStyle := lipgloss.NewStyle().
		Foreground(colorError).
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorError)

	return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
}

// truncateString truncates a string to maxLen, adding "…" if truncated
func truncateString(s string, maxLen int) string {
	if maxLen < 1 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen == 1 {
		return "…"
	}
	return s[:maxLen-1] + "…"
}
