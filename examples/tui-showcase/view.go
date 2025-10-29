package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// view.go - View Rendering
// Purpose: Top-level view rendering and layout
// When to extend: Add new view modes or modify layout logic

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

	// If an effect is active, render it full-screen
	if m.activeEffect != "" {
		return m.renderFullScreenEffect()
	}

	// Always render tabbed layout
	baseView := m.renderTabbed()

	// Overlay dropdown menu if open
	if m.menuOpen && m.activeMenu != "" {
		dropdown := m.renderActiveDropdown()
		if dropdown != "" {
			// Position dropdown below menu bar
			menuX := m.getMenuXPosition(m.activeMenu)
			menuY := 2 // Below title and menu bar

			// Overlay dropdown on base view
			baseView = m.overlayDropdown(baseView, dropdown, menuX, menuY)
		}
	}

	return baseView
}

// renderSinglePane renders a single-pane layout
func (m model) renderSinglePane() string {
	var sections []string

	// Title bar
	if m.config.UI.ShowTitle {
		sections = append(sections, m.renderTitleBar())
	}

	// Main content
	sections = append(sections, m.renderMainContent())

	// Status bar
	if m.config.UI.ShowStatus {
		sections = append(sections, m.renderStatusBar())
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderDualPane renders a dual-pane layout (side-by-side)
func (m model) renderDualPane() string {
	var sections []string

	// Title bar
	if m.config.UI.ShowTitle {
		sections = append(sections, m.renderTitleBar())
	}

	// Calculate pane dimensions
	leftWidth, rightWidth := m.calculateDualPaneLayout()

	// Left pane
	leftPane := m.renderLeftPane(leftWidth)

	// Divider
	divider := ""
	if m.config.Layout.ShowDivider {
		divider = m.renderDivider()
	}

	// Right pane
	rightPane := m.renderRightPane(rightWidth)

	// Join panes horizontally
	panes := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, divider, rightPane)
	sections = append(sections, panes)

	// Status bar
	if m.config.UI.ShowStatus {
		sections = append(sections, m.renderStatusBar())
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderMultiPanel renders a multi-panel layout
func (m model) renderMultiPanel() string {
	var sections []string

	// Title bar
	if m.config.UI.ShowTitle {
		sections = append(sections, m.renderTitleBar())
	}

	contentWidth, contentHeight := m.calculateLayout()

	// Top row (2 panels side by side)
	topHeight := contentHeight / 2
	panelWidth := contentWidth / 2

	topLeft := m.renderTopLeftPanel(panelWidth, topHeight)
	topRight := m.renderTopRightPanel(panelWidth, topHeight)
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, topLeft, topRight)

	// Bottom panel (full width)
	bottomHeight := contentHeight - topHeight
	bottom := m.renderBottomPanel(contentWidth, bottomHeight)

	sections = append(sections, topRow, bottom)

	// Status bar
	if m.config.UI.ShowStatus {
		sections = append(sections, m.renderStatusBar())
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderTabbed renders a tabbed interface
func (m model) renderTabbed() string {
	var sections []string

	// Title bar
	if m.config.UI.ShowTitle {
		sections = append(sections, m.renderTitleBar())
	}

	// Tab bar
	sections = append(sections, m.renderTabBar())

	// Active tab content
	contentWidth, contentHeight := m.calculateLayout()
	contentHeight -= 2 // Account for tab bar

	switch m.currentTab {
	case 0:
		sections = append(sections, m.renderTab1Content(contentWidth, contentHeight))
	case 1:
		sections = append(sections, m.renderTab2Content(contentWidth, contentHeight))
	case 2:
		sections = append(sections, m.renderTab3Content(contentWidth, contentHeight))
	case 3:
		sections = append(sections, m.renderBorderShowcase(contentWidth, contentHeight))
	case 4:
		sections = append(sections, m.renderColorPalette(contentWidth, contentHeight))
	case 5:
		sections = append(sections, m.renderDynamicPanels(contentWidth, contentHeight))
	case 6:
		sections = append(sections, m.renderFormsTab(contentWidth, contentHeight))
	case 7:
		sections = append(sections, m.renderTablesTab(contentWidth, contentHeight))
	case 8:
		sections = append(sections, m.renderDialogsTab(contentWidth, contentHeight))
	case 9:
		sections = append(sections, m.renderProgressTab(contentWidth, contentHeight))
	case 10:
		sections = append(sections, m.renderTreeViewTab(contentWidth, contentHeight))
	case 11:
		sections = append(sections, m.renderMobileTab(contentWidth, contentHeight))
	case 12:
		sections = append(sections, m.renderFourPanelTab(contentWidth, contentHeight))
	default:
		sections = append(sections, m.renderTab1Content(contentWidth, contentHeight))
	}

	// Status bar
	if m.config.UI.ShowStatus {
		sections = append(sections, m.renderStatusBar())
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// renderTabBar renders the tab navigation bar
func (m model) renderTabBar() string {
	tabNames := []string{"Single", "Dual", "Multi", "Borders", "Colors", "Dynamic", "Forms", "Tables", "Dialogs", "Progress", "Tree", "Mobile", "4-Panel"}
	var renderedTabs []string

	for i, name := range tabNames {
		tabLabel := "[ " + name + " ]"
		if m.currentTab == i {
			renderedTabs = append(renderedTabs, tabActiveStyle.Render(tabLabel))
		} else {
			renderedTabs = append(renderedTabs, tabInactiveStyle.Render(tabLabel))
		}
	}

	tabs := strings.Join(renderedTabs, " ")
	padding := m.width - lipgloss.Width(tabs)
	if padding < 0 {
		padding = 0
	}

	return tabs + strings.Repeat(" ", padding)
}

// Multi-panel rendering functions
func (m model) renderTopLeftPanel(width, height int) string {
	var content strings.Builder
	content.WriteString("╔════════════════╗\n")
	content.WriteString("║  TOP LEFT (1)  ║\n")
	content.WriteString("╚════════════════╝\n\n")
	content.WriteString("Panel ID: top-left\n")
	content.WriteString("Type: multi_panel\n")
	content.WriteString("Position: Top-Left\n\n")
	content.WriteString("Typically used for:\n")
	content.WriteString("• Primary navigation\n")
	content.WriteString("• Main controls\n")

	return contentStyle.Width(width).Height(height).Render(content.String())
}

func (m model) renderTopRightPanel(width, height int) string {
	var content strings.Builder
	content.WriteString("╔═════════════════╗\n")
	content.WriteString("║  TOP RIGHT (2)  ║\n")
	content.WriteString("╚═════════════════╝\n\n")
	content.WriteString("Panel ID: top-right\n")
	content.WriteString("Type: multi_panel\n")
	content.WriteString("Position: Top-Right\n\n")
	content.WriteString("Typically used for:\n")
	content.WriteString("• Details\n")
	content.WriteString("• Preview\n")

	return contentStyle.Width(width).Height(height).Render(content.String())
}

func (m model) renderBottomPanel(width, height int) string {
	var content strings.Builder
	content.WriteString("╔══════════════════════════════════╗\n")
	content.WriteString("║         BOTTOM PANEL (3)         ║\n")
	content.WriteString("╚══════════════════════════════════╝\n\n")
	content.WriteString("Panel ID: bottom-panel\n")
	content.WriteString("Type: multi_panel\n")
	content.WriteString("Position: Bottom (full width)\n\n")
	content.WriteString("Typically used for: Logs, Console, Output, Terminal, Status details\n")

	return contentStyle.Width(width).Height(height).Render(content.String())
}

// Layout demo tabs - these showcase different layout patterns
func (m model) renderTab1Content(width, height int) string {
	// Tab 0: Single Pane Layout Demo
	var content strings.Builder
	content.WriteString("╔═══════════════════════════════════════╗\n")
	content.WriteString("║      SINGLE PANE LAYOUT DEMO          ║\n")
	content.WriteString("╚═══════════════════════════════════════╝\n\n")
	content.WriteString("Purpose: Full-screen single content area\n\n")
	content.WriteString("Best for:\n")
	content.WriteString("• Simple applications with one main view\n")
	content.WriteString("• Focus mode - no distractions\n")
	content.WriteString("• Document viewers, readers\n")
	content.WriteString("• Full-screen editors\n\n")
	content.WriteString("Pattern:\n")
	content.WriteString("┌─────────────────────────────┐\n")
	content.WriteString("│                             │\n")
	content.WriteString("│      MAIN CONTENT           │\n")
	content.WriteString("│                             │\n")
	content.WriteString("└─────────────────────────────┘\n")

	return contentStyle.Width(width).Height(height).Render(content.String())
}

func (m model) renderTab2Content(width, height int) string {
	// Tab 1: Dual Pane Layout Demo
	leftPane := m.renderLeftPane(width/2)
	rightPane := m.renderRightPane(width/2)
	divider := m.renderDivider()

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPane, divider, rightPane)
}

func (m model) renderTab3Content(width, height int) string {
	// Tab 2: Multi-Panel Layout Demo
	topHeight := height / 2
	panelWidth := width / 2

	topLeft := m.renderTopLeftPanel(panelWidth, topHeight)
	topRight := m.renderTopRightPanel(panelWidth, topHeight)
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, topLeft, topRight)

	bottomHeight := height - topHeight
	bottom := m.renderBottomPanel(width, bottomHeight)

	return lipgloss.JoinVertical(lipgloss.Left, topRow, bottom)
}

// renderBorderShowcase displays different border styles
func (m model) renderBorderShowcase(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║   LIPGLOSS BORDER STYLE SHOWCASE      ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Rounded Border
	roundedBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(35).
		Render("Rounded Border\nPerfect for dialogs and cards")
	content.WriteString(roundedBox)
	content.WriteString("\n\n")

	// Double Border
	doubleBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(35).
		Render("Double Border\nElegant and distinctive")
	content.WriteString(doubleBox)
	content.WriteString("\n\n")

	// Thick Border
	thickBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(35).
		Render("Thick Border\nBold and prominent")
	content.WriteString(thickBox)
	content.WriteString("\n\n")

	// Normal Border
	normalBox := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(colorInfo).
		Padding(1, 2).
		Width(35).
		Render("Normal Border\nClean and simple")
	content.WriteString(normalBox)
	content.WriteString("\n\n")

	// Hidden Border (for spacing)
	hiddenBox := lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Padding(1, 2).
		Width(35).
		Background(colorBorder).
		Render("Hidden Border\nInvisible but provides spacing")
	content.WriteString(hiddenBox)
	content.WriteString("\n\n")

	content.WriteString("Panel ID: border-showcase\n")
	content.WriteString("Use these borders in your TUI apps!")

	return contentStyle.Width(width).Height(height).Render(content.String())
}

// renderColorPalette displays available color palette
func (m model) renderColorPalette(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("════════════════════════════════════"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("    COLOR PALETTE SHOWCASE"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("════════════════════════════════════"))
	content.WriteString("\n\n")

	// Theme Colors
	content.WriteString(lipgloss.NewStyle().Bold(true).Render("🎨 Theme Colors"))
	content.WriteString("\n\n")

	colors := []struct {
		name  string
		hex   string
		color lipgloss.Color
	}{
		{"Primary (Blue)", "#61AFEF", colorPrimary},
		{"Secondary (Purple)", "#C678DD", colorSecondary},
		{"Accent (Green)", "#98C379", colorAccent},
		{"Error (Red)", "#E06C75", colorError},
		{"Warning (Yellow)", "#E5C07B", colorWarning},
		{"Info (Cyan)", "#56B6C2", colorInfo},
		{"Foreground", "#ABB2BF", colorForeground},
		{"Background", "#282C34", colorBackground},
	}

	for _, c := range colors {
		colorBox := lipgloss.NewStyle().
			Background(c.color).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			Render("███")

		label := lipgloss.NewStyle().
			Width(25).
			Render(c.name)

		hex := lipgloss.NewStyle().
			Foreground(colorDimmed).
			Render(c.hex)

		content.WriteString(colorBox + "  " + label + " " + hex)
		content.WriteString("\n")
	}

	content.WriteString("\n")
	content.WriteString(lipgloss.NewStyle().Bold(true).Render("🌈 Standard Terminal Colors"))
	content.WriteString("\n\n")

	standardColors := []struct {
		name  string
		color lipgloss.Color
	}{
		{"Black", lipgloss.Color("0")},
		{"Red", lipgloss.Color("1")},
		{"Green", lipgloss.Color("2")},
		{"Yellow", lipgloss.Color("3")},
		{"Blue", lipgloss.Color("4")},
		{"Magenta", lipgloss.Color("5")},
		{"Cyan", lipgloss.Color("6")},
		{"White", lipgloss.Color("7")},
	}

	for _, c := range standardColors {
		colorBox := lipgloss.NewStyle().
			Background(c.color).
			Padding(0, 2).
			Render("   ")

		label := lipgloss.NewStyle().
			Width(15).
			Render(c.name)

		content.WriteString(colorBox + "  " + label)
		content.WriteString("\n")
	}

	content.WriteString("\n")
	content.WriteString(lipgloss.NewStyle().Foreground(colorDimmed).Render("Panel ID: color-palette"))
	content.WriteString("\n")
	content.WriteString(lipgloss.NewStyle().Foreground(colorDimmed).Render("Use lipgloss.Color(\"#HEX\") or lipgloss.Color(\"0-255\")"))

	return contentStyle.Width(width).Height(height).Render(content.String())
}

// Component rendering functions

// renderTitleBar renders the title bar with dropdown menu bar
func (m model) renderTitleBar() string {
	layoutName := map[string]string{
		"single":      "Single Pane",
		"dual_pane":   "Dual Pane",
		"multi_panel": "Multi-Panel",
		"tabbed":      "Tabbed",
	}[m.currentLayout]

	title := titleStyle.Render("Layout Demo - " + layoutName)

	// Add menu bar instead of layout buttons
	menuBar := m.renderMenuBar()

	firstLine := title + strings.Repeat(" ", m.width-lipgloss.Width(title))
	secondLine := menuBar

	// If a dropdown menu is open, render it overlaid on the content
	titleBar := firstLine + "\n" + secondLine

	return titleBar
}

// renderLayoutButtons renders clickable buttons for switching layouts
func (m model) renderLayoutButtons() string {
	var buttons []string

	layouts := []struct {
		num  string
		name string
	}{
		{"1", "Single"},
		{"2", "Dual"},
		{"3", "Multi"},
		{"4", "Tabs"},
	}

	for i, layout := range layouts {
		buttonID := "button-" + string(rune(i+1+'0'))

		// Style based on current layout and hover state
		var style lipgloss.Style
		isActive := (i == 0 && m.currentLayout == "single") ||
			(i == 1 && m.currentLayout == "dual_pane") ||
			(i == 2 && m.currentLayout == "multi_panel") ||
			(i == 3 && m.currentLayout == "tabbed")

		if isActive {
			style = buttonActiveStyle
		} else if m.hoveredItem == buttonID {
			style = buttonHoverStyle
		} else {
			style = buttonStyle
		}

		buttons = append(buttons, style.Render(" "+layout.num+":"+layout.name+" "))
	}

	return "  " + strings.Join(buttons, " ")
}

// renderStatusBar renders the status bar with mouse position
func (m model) renderStatusBar() string {
	status := m.statusMsg

	// Add mouse position info
	mouseInfo := ""
	if m.config.UI.MouseEnabled {
		mouseInfo = lipgloss.NewStyle().
			Foreground(colorDimmed).
			Render(fmt.Sprintf(" | Mouse: %d,%d", m.mouseX, m.mouseY))

		if m.hoveredItem != "" {
			mouseInfo += " [" + m.hoveredItem + "]"
		}
	}

	fullStatus := status + mouseInfo
	width := m.width - lipgloss.Width(fullStatus) - 4

	if width < 0 {
		// Truncate status if too long
		maxLen := m.width - 4
		if maxLen > 0 && len(status) > maxLen {
			status = status[:maxLen-3] + "..."
			fullStatus = status
		}
		width = 0
	}

	return statusStyle.Render(fullStatus + strings.Repeat(" ", width))
}

// renderMainContent renders the main content area
func (m model) renderMainContent() string {
	contentWidth, contentHeight := m.calculateLayout()

	var content strings.Builder
	content.WriteString("╔══════════════════════════════╗\n")
	content.WriteString("║   SINGLE PANE MAIN CONTENT   ║\n")
	content.WriteString("╚══════════════════════════════╝\n\n")
	content.WriteString("This is the main content area for single pane layout.\n\n")
	content.WriteString("Panel ID: main-content\n")
	content.WriteString("Type: single_pane\n\n")
	content.WriteString("Keyboard Shortcuts:\n")
	content.WriteString("  1 - Single Pane Layout\n")
	content.WriteString("  2 - Dual Pane Layout\n")
	content.WriteString("  3 - Multi-Panel Layout\n")
	content.WriteString("  4 - Tabbed Layout\n")
	content.WriteString("  q - Quit\n\n")
	content.WriteString("When giving instructions to Claude, you can reference\n")
	content.WriteString("this panel as 'main-content' or 'the main content area'.")

	return contentStyle.Width(contentWidth).Height(contentHeight).Render(content.String())
}

// renderLeftPane renders the left pane in dual-pane mode
func (m model) renderLeftPane(width int) string {
	_, contentHeight := m.calculateLayout()

	var content strings.Builder
	content.WriteString("╔═══════════════╗\n")
	content.WriteString("║   LEFT PANE   ║\n")
	content.WriteString("╚═══════════════╝\n\n")
	content.WriteString("This is the left pane.\n")
	content.WriteString("Typically used for:\n")
	content.WriteString("• Navigation\n")
	content.WriteString("• File lists\n")
	content.WriteString("• Menu items\n")
	content.WriteString("• Tree views\n\n")
	content.WriteString("Panel ID: left-pane\n")
	content.WriteString("Type: dual_pane\n\n")
	content.WriteString("When giving instructions\n")
	content.WriteString("to Claude, reference this\n")
	content.WriteString("as 'left-pane' or\n")
	content.WriteString("'the left panel'.")

	return leftPaneStyle.Width(width).Height(contentHeight).Render(content.String())
}

// renderRightPane renders the right pane in dual-pane mode
func (m model) renderRightPane(width int) string {
	_, contentHeight := m.calculateLayout()

	var content strings.Builder
	content.WriteString("╔════════════════╗\n")
	content.WriteString("║   RIGHT PANE   ║\n")
	content.WriteString("╚════════════════╝\n\n")
	content.WriteString("This is the right pane.\n")
	content.WriteString("Typically used for:\n")
	content.WriteString("• Preview content\n")
	content.WriteString("• Detail views\n")
	content.WriteString("• File contents\n")
	content.WriteString("• Output display\n\n")
	content.WriteString("Panel ID: right-pane\n")
	content.WriteString("Type: dual_pane\n\n")
	content.WriteString("When giving instructions\n")
	content.WriteString("to Claude, reference this\n")
	content.WriteString("as 'right-pane' or\n")
	content.WriteString("'the right panel'.")

	return rightPaneStyle.Width(width).Height(contentHeight).Render(content.String())
}

// renderDivider renders the vertical divider between panes
func (m model) renderDivider() string {
	_, contentHeight := m.calculateLayout()
	divider := strings.Repeat("│\n", contentHeight)
	return dividerStyle.Render(divider)
}

// Error and minimal views

// renderErrorView renders an error message
func (m model) renderErrorView() string {
	content := "Error: " + m.err.Error() + "\n\n"
	content += "Press q to quit"
	return errorStyle.Render(content)
}

// renderMinimalView renders a minimal view for small terminals
func (m model) renderMinimalView() string {
	content := "Terminal too small\n"
	content += "Minimum: 40x10\n"
	content += "Press q to quit"
	return errorStyle.Render(content)
}

// Helper functions

// truncateString truncates a string to fit within maxWidth
func truncateString(s string, maxWidth int) string {
	if len(s) <= maxWidth {
		return s
	}
	if maxWidth <= 3 {
		return s[:maxWidth]
	}
	return s[:maxWidth-3] + "..."
}

// padRight pads a string with spaces to reach the desired width
func padRight(s string, width int) string {
	currentWidth := lipgloss.Width(s)
	if currentWidth >= width {
		return s
	}
	return s + strings.Repeat(" ", width-currentWidth)
}

// centerString centers a string within the given width
func centerString(s string, width int) string {
	strWidth := lipgloss.Width(s)
	if strWidth >= width {
		return s
	}
	leftPad := (width - strWidth) / 2
	rightPad := width - strWidth - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

// renderDynamicPanels renders the dynamic panels tab (Tab 5) - 3-panel layout
func (m model) renderDynamicPanels(totalWidth, totalHeight int) string {
	// Calculate panel dimensions using weight-based system (3-panel layout)
	leftWidth, rightWidth, topHeight, bottomHeight := m.calculateThreePanelLayout(totalWidth, totalHeight)

	// Render top left panel
	leftPanel := m.renderDynamicPanel("left", leftWidth, topHeight, m.leftContent)

	// Render vertical divider for top panels
	divider := renderVerticalDivider(topHeight)

	// Render top right panel
	rightPanel := m.renderDynamicPanel("right", rightWidth, topHeight, m.rightContent)

	// Join top panels horizontally
	topPanels := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, divider, rightPanel)

	// Render bottom panel (full width)
	bottomPanel := m.renderDynamicPanel("bottom", totalWidth, bottomHeight, m.bottomContent)

	// Stack all panels vertically
	return lipgloss.JoinVertical(lipgloss.Left, topPanels, bottomPanel)
}

// renderDynamicPanel renders a single dynamic panel with border and content
func (m model) renderDynamicPanel(panelName string, width, height int, content []string) string {
	isFocused := m.focusedPanel == panelName

	// Create border style based on focus
	borderStyle := lipgloss.RoundedBorder()
	borderColor := lipgloss.Color("240") // Dim gray
	titleColor := lipgloss.Color("252")  // Light gray

	if isFocused {
		borderColor = colorPrimary // Bright blue
		titleColor = colorPrimary  // Bright blue
	}

	// Calculate max text width
	maxTextWidth := width - 4 // -2 for borders, -2 for padding
	if maxTextWidth < 1 {
		maxTextWidth = 1
	}

	// Create panel title
	title := strings.ToUpper(panelName) + " PANEL"
	if isFocused {
		title += " ●" // Indicator for focused panel
	}
	title = truncateString(title, maxTextWidth)

	// Calculate weight for display
	weight := 1
	if m.accordionMode && isFocused {
		weight = 2
	}

	subtitle := fmt.Sprintf("Weight: %d | Size: %dx%d", weight, width, height)
	subtitle = truncateString(subtitle, maxTextWidth)

	// Accordion mode status
	modeInfo := "Accordion: "
	if m.accordionMode {
		modeInfo += "ON (press 'a' to toggle)"
	} else {
		modeInfo += "OFF (press 'a' to toggle)"
	}
	modeInfo = truncateString(modeInfo, maxTextWidth)

	// Create header
	header := lipgloss.NewStyle().
		Foreground(titleColor).
		Bold(true).
		Render(title)

	subheader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Render(subtitle)

	modeHeader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Render(modeInfo)

	// Build content lines
	var lines []string
	lines = append(lines, header)
	lines = append(lines, subheader)
	lines = append(lines, modeHeader)
	lines = append(lines, "")

	// Calculate exact content area height
	innerHeight := height - 2          // Remove borders
	availableContentLines := innerHeight - 4 // Remove header lines (3 header + 1 blank)

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

	// Create styled panel
	panelStyle := lipgloss.NewStyle().
		Border(borderStyle).
		BorderForeground(borderColor).
		Width(width - 2) // Account for left/right borders

	return panelStyle.Render(contentStr)
}

// renderVerticalDivider renders a vertical divider between panels
func renderVerticalDivider(height int) string {
	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	var lines []string
	for i := 0; i < height; i++ {
		lines = append(lines, "│")
	}

	return dividerStyle.Render(strings.Join(lines, "\n"))
}

// renderFourPanelTab renders the 4-panel dynamic layout (Tab 12)
func (m model) renderFourPanelTab(totalWidth, totalHeight int) string {
	// Calculate panel dimensions using weight-based system (4-panel layout)
	headerHeight, middleHeight, footerHeight, leftWidth, rightWidth := m.calculateFourPanelLayout(totalWidth, totalHeight)

	// Render header panel
	headerPanel := m.renderDynamicPanel("header", totalWidth, headerHeight, m.headerContent)

	// Render middle row panels
	leftPanel := m.renderDynamicPanel("left", leftWidth, middleHeight, m.leftContent)
	divider := renderVerticalDivider(middleHeight)
	rightPanel := m.renderDynamicPanel("right", rightWidth, middleHeight, m.rightContent)

	// Join middle panels horizontally
	middlePanels := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, divider, rightPanel)

	// Render footer panel
	footerPanel := m.renderDynamicPanel("footer", totalWidth, footerHeight, m.bottomContent)

	// Stack all panels vertically
	return lipgloss.JoinVertical(lipgloss.Left, headerPanel, middlePanels, footerPanel)
}

// overlayDropdown overlays a dropdown menu on the base view at the specified position
// Uses proper ANSI-aware overlay to preserve background content (from TFE)
func (m model) overlayDropdown(baseView, dropdown string, x, y int) string {
	// Split base view into lines
	baseLines := strings.Split(baseView, "\n")
	dropdownLines := strings.Split(dropdown, "\n")

	// Ensure we have enough base lines
	for len(baseLines) < m.height {
		baseLines = append(baseLines, "")
	}

	// Overlay each dropdown line onto the base view
	for i, dropdownLine := range dropdownLines {
		targetLine := y + i
		if targetLine < 0 || targetLine >= len(baseLines) {
			continue
		}

		baseLine := baseLines[targetLine]

		// We need to overlay dropdownLine at visual column x
		// Use a string builder to construct the new line
		var newLine strings.Builder

		// Get the part of baseLine before position x
		// We need to handle ANSI codes properly
		visualPos := 0
		bytePos := 0
		inAnsi := false
		baseRunes := []rune(baseLine)

		// Scan through base line until we reach visual position x
		for bytePos < len(baseRunes) && visualPos < x {
			if baseRunes[bytePos] == '\033' {
				inAnsi = true
			} else if inAnsi {
				if (baseRunes[bytePos] >= 'A' && baseRunes[bytePos] <= 'Z') ||
					(baseRunes[bytePos] >= 'a' && baseRunes[bytePos] <= 'z') {
					inAnsi = false
				}
			} else {
				// Count visual width (handles wide characters)
				visualPos++
			}
			bytePos++
		}

		// Add the left part of the base line (up to position x)
		if bytePos > 0 && bytePos <= len(baseRunes) {
			newLine.WriteString(string(baseRunes[:bytePos]))
		}

		// Pad with spaces if needed to reach position x
		for visualPos < x {
			newLine.WriteRune(' ')
			visualPos++
		}

		// Add the dropdown line
		newLine.WriteString(dropdownLine)

		// Now preserve the right side of the base line (after the dropdown)
		dropdownWidth := lipgloss.Width(dropdownLine)
		endVisualPos := x + dropdownWidth

		// Continue from where we left off and skip to the end position
		for bytePos < len(baseRunes) && visualPos < endVisualPos {
			if baseRunes[bytePos] == '\033' {
				inAnsi = true
			} else if inAnsi {
				if (baseRunes[bytePos] >= 'A' && baseRunes[bytePos] <= 'Z') ||
					(baseRunes[bytePos] >= 'a' && baseRunes[bytePos] <= 'z') {
					inAnsi = false
				}
			} else {
				visualPos++
			}
			bytePos++
		}

		// Add the remaining right part of the base line
		if bytePos < len(baseRunes) {
			newLine.WriteString(string(baseRunes[bytePos:]))
		}

		baseLines[targetLine] = newLine.String()
	}

	return strings.Join(baseLines, "\n")
}

// renderFullScreenEffect renders a full-screen effect view
func (m model) renderFullScreenEffect() string {
	switch m.activeEffect {
	case "metaballs":
		return m.renderFullScreenMetaballs()
	case "wavy-menu":
		return m.renderFullScreenWavyMenu()
	case "rainbow":
		return m.renderFullScreenRainbow()
	case "landing":
		return m.renderFullScreenLanding()
	default:
		return "Unknown effect"
	}
}

// renderFullScreenMetaballs renders the metaballs effect full-screen
func (m model) renderFullScreenMetaballs() string {
	if m.metaballEngine == nil {
		return "Metaballs engine not initialized"
	}

	// Render the metaballs effect
	metaballsRender := m.metaballEngine.Render()

	// Add title and instructions at the top
	title := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true).
		Render("METABALLS EFFECT")

	subtitle := lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Physics-based floating blobs with organic motion")

	controls := lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Press Esc or Q to return to showcase")

	// Combine title, subtitle, effect, and controls
	header := lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "")
	footer := lipgloss.JoinVertical(lipgloss.Left, "", controls)

	return lipgloss.JoinVertical(lipgloss.Left, header, metaballsRender, footer)
}

// renderFullScreenWavyMenu renders the wavy grid effect full-screen
func (m model) renderFullScreenWavyMenu() string {
	if m.waveGrid == nil {
		return "Wave grid not initialized"
	}

	// Render the wavy grid effect
	gridRender := m.waveGrid.Render()

	// Add title and instructions at the top
	title := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true).
		Render("WAVY GRID EFFECT")

	subtitle := lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Sine wave distortion for animated grid backgrounds")

	controls := lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Press Esc or Q to return to showcase")

	// Combine title, subtitle, effect, and controls
	header := lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "")
	footer := lipgloss.JoinVertical(lipgloss.Left, "", controls)

	return lipgloss.JoinVertical(lipgloss.Left, header, gridRender, footer)
}

// renderFullScreenRainbow renders the rainbow text effect full-screen
func (m model) renderFullScreenRainbow() string {
	if m.rainbowCycler == nil {
		return "Rainbow cycler not initialized"
	}

	// ASCII art for demo
	asciiArt := []string{
		"████████╗██╗   ██╗██╗",
		"╚══██╔══╝██║   ██║██║",
		"   ██║   ██║   ██║██║",
		"   ██║   ██║   ██║██║",
		"   ██║   ╚██████╔╝██║",
		"   ╚═╝    ╚═════╝ ╚═╝",
	}

	// Render rainbow ASCII art
	rainbowArt := m.rainbowCycler.RenderLines(asciiArt)

	// Single line example
	exampleText := m.rainbowCycler.Render("The quick brown fox jumps over the lazy dog")

	// Title and instructions
	title := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true).
		Render("RAINBOW TEXT EFFECT")

	subtitle := lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Animated rainbow colors cycling through text")

	controls := lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Press Esc or Q to return to showcase")

	// Combine everything
	header := lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "")
	content := lipgloss.JoinVertical(lipgloss.Left, rainbowArt, "", exampleText)
	footer := lipgloss.JoinVertical(lipgloss.Left, "", controls)

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

// renderFullScreenLanding renders all effects combined full-screen
func (m model) renderFullScreenLanding() string {
	if m.waveGrid == nil || m.metaballEngine == nil || m.rainbowCycler == nil {
		return "Effects not initialized"
	}

	// Rainbow title
	titleArt := []string{
		"████████╗██╗   ██╗██╗",
		"╚══██╔══╝██║   ██║██║",
		"   ██║   ██║   ██║██║",
	}
	rainbowTitle := m.rainbowCycler.RenderLines(titleArt)

	// Render metaballs
	metaballsRender := m.metaballEngine.Render()

	// Title and instructions
	title := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true).
		Render("LANDING PAGE - ALL EFFECTS COMBINED")

	subtitle := lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("✨ Wavy Grid + Metaballs + Rainbow = Beautiful TUIs ✨")

	controls := lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Press Esc or Q to return to showcase")

	// Combine everything
	header := lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "")
	content := lipgloss.JoinVertical(lipgloss.Left, rainbowTitle, "", metaballsRender)
	footer := lipgloss.JoinVertical(lipgloss.Left, "", controls)

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}
