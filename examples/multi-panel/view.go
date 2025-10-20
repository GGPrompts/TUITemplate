package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the entire application
func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	// Check minimum terminal size
	minWidth := 60
	minHeight := 15
	if m.width < minWidth || m.height < minHeight {
		msg := lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")).
			Bold(true).
			Render("Terminal too small!\n\n") +
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Render(fmt.Sprintf("Minimum size: %dx%d\n", minWidth, minHeight)) +
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Render("Current size: ") +
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Render(fmt.Sprintf("%dx%d", m.width, m.height)) +
			"\n\n" +
			lipgloss.NewStyle().
				Foreground(dimColor).
				Render("Press q to quit")
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, msg)
	}

	// Calculate dimensions
	// Status bar takes 1 line
	statusHeight := 1
	availableHeight := m.height - statusHeight

	// Use dynamic weight-based layout calculation
	leftWidth, rightWidth, rightTopHeight, rightBottomHeight := m.calculateThreePanelLayout(m.width, availableHeight)

	// Store panel boundaries for mouse detection
	m.panelBounds[LeftPanel] = panelBounds{
		x:      0,
		y:      0,
		width:  leftWidth,
		height: availableHeight,
	}
	m.panelBounds[TopRightPanel] = panelBounds{
		x:      leftWidth,
		y:      0,
		width:  rightWidth,
		height: rightTopHeight,
	}
	m.panelBounds[BottomRightPanel] = panelBounds{
		x:      leftWidth,
		y:      rightTopHeight,
		width:  rightWidth,
		height: rightBottomHeight,
	}

	// Render panels
	leftPanel := m.renderLeftPanel(leftWidth, availableHeight)
	rightTopPanel := m.renderTopRightPanel(rightWidth, rightTopHeight)
	rightBottomPanel := m.renderBottomRightPanel(rightWidth, rightBottomHeight)

	// Stack right panels vertically
	rightPanels := lipgloss.JoinVertical(
		lipgloss.Left,
		rightTopPanel,
		rightBottomPanel,
	)

	// Join left and right horizontally
	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		rightPanels,
	)

	// Add status bar
	statusBar := m.renderStatusBar()

	// Final layout
	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		statusBar,
	)
}

// renderLeftPanel renders the left panel (file list)
func (m model) renderLeftPanel(width, height int) string {
	focused := m.focusedPanel == LeftPanel

	// Calculate content dimensions accounting for all decorations
	borderWidth := 2   // left + right border
	borderHeight := 2  // top + bottom border
	titleHeight := 1   // title bar
	paddingHeight := 2 // contentStyle Padding(1) = top + bottom

	contentWidth := width - borderWidth - 2  // -2 for padding left+right
	contentHeight := height - borderHeight - titleHeight - paddingHeight

	// Title with weight indicator
	titleText := " " + LeftPanel.String() + " "
	if m.accordionMode && focused {
		titleText += "●" // Indicator for focused + accordion mode
	}
	title := getPanelTitleStyle(focused).Render(titleText)

	// Build file list
	var items []string
	visibleStart := max(0, m.cursor-contentHeight+3)
	visibleEnd := min(len(m.files), visibleStart+contentHeight)

	for i := visibleStart; i < visibleEnd; i++ {
		file := m.files[i]
		prefix := "  "
		isSelected := i == m.cursor

		if isSelected {
			if focused {
				prefix = "▶ "
			} else {
				prefix = "▸ "
			}
		}

		line := prefix + file
		// Truncate if too long
		if len(line) > contentWidth {
			line = line[:contentWidth-3] + "..."
		}

		// Apply styling
		if isSelected && focused {
			line = lipgloss.NewStyle().
				Width(contentWidth).
				Render(selectedItemStyle.Render(line))
		} else if isSelected {
			line = dimItemStyle.Render(line)
		} else {
			line = normalItemStyle.Render(line)
		}

		items = append(items, line)
	}

	// Fill remaining space
	for len(items) < contentHeight {
		items = append(items, "")
	}

	// Don't set Height on contentStyle since we already have exact line count
	content := contentStyle.
		Width(contentWidth).
		Render(strings.Join(items, "\n"))

	// Combine title and content
	panel := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		content,
	)

	// Apply border (no size constraints - let it wrap naturally around content)
	return getPanelBorderStyle(focused).Render(panel)
}

// renderTopRightPanel renders the top right panel (details)
func (m model) renderTopRightPanel(width, height int) string {
	focused := m.focusedPanel == TopRightPanel

	// Calculate content dimensions accounting for all decorations
	borderWidth := 2   // left + right border
	borderHeight := 2  // top + bottom border
	titleHeight := 1   // title bar
	paddingHeight := 2 // contentStyle Padding(1) = top + bottom

	contentWidth := width - borderWidth - 2  // -2 for padding left+right
	contentHeight := height - borderHeight - titleHeight - paddingHeight

	// Title with weight indicator
	titleText := " " + TopRightPanel.String() + " "
	if m.accordionMode && focused {
		titleText += "●" // Indicator for focused + accordion mode
	}
	title := getPanelTitleStyle(focused).Render(titleText)

	// Build details content
	lines := strings.Split(m.details, "\n")
	var visibleLines []string

	for i := 0; i < contentHeight && i < len(lines); i++ {
		line := lines[i]
		// Truncate if too long
		if len(line) > contentWidth {
			line = line[:contentWidth-3] + "..."
		}
		visibleLines = append(visibleLines, line)
	}

	// Fill remaining space
	for len(visibleLines) < contentHeight {
		visibleLines = append(visibleLines, "")
	}

	// Don't set Height on contentStyle since we already have exact line count
	content := contentStyle.
		Width(contentWidth).
		Render(strings.Join(visibleLines, "\n"))

	// Combine title and content
	panel := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		content,
	)

	// Apply border (no size constraints - let it wrap naturally around content)
	return getPanelBorderStyle(focused).Render(panel)
}

// renderBottomRightPanel renders the bottom right panel (logs)
func (m model) renderBottomRightPanel(width, height int) string {
	focused := m.focusedPanel == BottomRightPanel

	// Calculate content dimensions accounting for all decorations
	borderWidth := 2   // left + right border
	borderHeight := 2  // top + bottom border
	titleHeight := 1   // title bar
	paddingHeight := 2 // contentStyle Padding(1) = top + bottom

	contentWidth := width - borderWidth - 2  // -2 for padding left+right
	contentHeight := height - borderHeight - titleHeight - paddingHeight

	// Title with weight indicator
	titleText := " " + BottomRightPanel.String() + " "
	if m.accordionMode && focused {
		titleText += "●" // Indicator for focused + accordion mode
	}
	title := getPanelTitleStyle(focused).Render(titleText)

	// Build logs content
	var visibleLines []string
	start := max(0, m.logScroll)
	end := min(len(m.logs), start+contentHeight)

	for i := start; i < end; i++ {
		line := m.logs[i]
		// Truncate if too long
		if len(line) > contentWidth {
			line = line[:contentWidth-3] + "..."
		}
		visibleLines = append(visibleLines, logLineStyle.Render(line))
	}

	// Fill remaining space
	for len(visibleLines) < contentHeight {
		visibleLines = append(visibleLines, "")
	}

	// Add scroll indicator if needed
	if len(m.logs) > contentHeight {
		scrollInfo := lipgloss.NewStyle().
			Foreground(dimColor).
			Render(lipgloss.PlaceHorizontal(
				contentWidth,
				lipgloss.Right,
				"⌃⌄",
			))
		if len(visibleLines) > 0 {
			visibleLines[len(visibleLines)-1] = scrollInfo
		}
	}

	// Don't set Height on contentStyle since we already have exact line count
	content := contentStyle.
		Width(contentWidth).
		Render(strings.Join(visibleLines, "\n"))

	// Combine title and content
	panel := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		content,
	)

	// Apply border (no size constraints - let it wrap naturally around content)
	return getPanelBorderStyle(focused).Render(panel)
}

// renderStatusBar renders the status bar
func (m model) renderStatusBar() string {
	// Left side - focused panel indicator
	leftContent := "Panel: " + m.focusedPanel.String()

	// Right side - help text
	rightContent := m.statusMsg

	// Available width (accounting for padding in statusBarStyle which is 2)
	availableWidth := m.width - 2

	// Minimum spacing between left and right content
	minSpacing := 2

	// Calculate how much space we have for the right content
	leftWidth := lipgloss.Width(leftContent)
	rightMaxWidth := availableWidth - leftWidth - minSpacing

	// Truncate right content if needed
	if rightMaxWidth < 0 {
		// Terminal too narrow, truncate left content too
		if availableWidth > 3 {
			leftContent = leftContent[:availableWidth-3] + "..."
			rightContent = ""
			minSpacing = 0
		} else {
			return statusBarStyle.Width(availableWidth).Render("")
		}
	} else if lipgloss.Width(rightContent) > rightMaxWidth {
		// Truncate right content to fit
		if rightMaxWidth > 3 {
			// Count visual width, not byte length
			truncated := ""
			for _, r := range rightContent {
				if lipgloss.Width(truncated+string(r)) > rightMaxWidth-3 {
					break
				}
				truncated += string(r)
			}
			rightContent = truncated + "..."
		} else {
			rightContent = ""
		}
	}

	// Calculate final spacing
	spacing := availableWidth - lipgloss.Width(leftContent) - lipgloss.Width(rightContent)
	if spacing < 0 {
		spacing = 0
	}

	status := leftContent + strings.Repeat(" ", spacing) + rightContent

	// Ensure the status is exactly terminal width
	statusWidth := lipgloss.Width(status)
	if statusWidth < m.width {
		status += strings.Repeat(" ", m.width-statusWidth)
	} else if statusWidth > m.width {
		// Truncate carefully by visual width
		truncated := ""
		for _, r := range status {
			if lipgloss.Width(truncated+string(r)) >= m.width {
				break
			}
			truncated += string(r)
		}
		status = truncated
	}

	// Render with color but no padding to avoid height issues
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(secondaryColor).
		Render(status)
}
