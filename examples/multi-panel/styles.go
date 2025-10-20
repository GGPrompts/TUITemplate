package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Color scheme
	primaryColor   = lipgloss.Color("63")  // Purple
	secondaryColor = lipgloss.Color("42")  // Green
	accentColor    = lipgloss.Color("214") // Orange
	dimColor       = lipgloss.Color("240") // Gray
	bgColor        = lipgloss.Color("235") // Dark gray
	fgColor        = lipgloss.Color("255") // White

	// Title styles
	focusedTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("0")).
				Background(primaryColor).
				Padding(0, 1)

	unfocusedTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(dimColor).
				Background(bgColor).
				Padding(0, 1)

	// Panel border styles
	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(primaryColor)

	unfocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(dimColor)

	// Content styles
	contentStyle = lipgloss.NewStyle().
			Padding(1)

	// List item styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("0")).
				Background(primaryColor).
				Bold(true)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(fgColor)

	dimItemStyle = lipgloss.NewStyle().
			Foreground(dimColor)

	// Status bar style
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(secondaryColor).
			Padding(0, 1)

	// Log styles
	logLineStyle = lipgloss.NewStyle().
			Foreground(fgColor)

	logTimestampStyle = lipgloss.NewStyle().
				Foreground(accentColor)
)

// getPanelTitleStyle returns the title style based on focus
func getPanelTitleStyle(focused bool) lipgloss.Style {
	if focused {
		return focusedTitleStyle
	}
	return unfocusedTitleStyle
}

// getPanelBorderStyle returns the border style based on focus
func getPanelBorderStyle(focused bool) lipgloss.Style {
	if focused {
		return focusedBorderStyle
	}
	return unfocusedBorderStyle
}
