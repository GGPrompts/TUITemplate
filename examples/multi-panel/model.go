package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// PanelID identifies different panels
type PanelID int

const (
	LeftPanel PanelID = iota
	TopRightPanel
	BottomRightPanel
)

// String returns the name of the panel
func (p PanelID) String() string {
	switch p {
	case LeftPanel:
		return "Files"
	case TopRightPanel:
		return "Details"
	case BottomRightPanel:
		return "Logs"
	default:
		return "Unknown"
	}
}

// panelBounds tracks the screen boundaries of a panel
type panelBounds struct {
	x      int
	y      int
	width  int
	height int
}

// contains checks if a point is within the panel bounds
func (b panelBounds) contains(x, y int) bool {
	return x >= b.x && x < b.x+b.width && y >= b.y && y < b.y+b.height
}

// model represents the application state
type model struct {
	width  int
	height int

	// Focus management
	focusedPanel PanelID
	panels       []PanelID

	// Panel boundaries for mouse detection
	panelBounds map[PanelID]panelBounds

	// Panel data
	files       []string
	cursor      int
	details     string
	logs        []string
	logScroll   int

	// Status
	statusMsg string
}

// initialModel creates the initial model
func initialModel() model {
	// Sample data
	files := []string{
		"main.go",
		"model.go",
		"update.go",
		"view.go",
		"styles.go",
		"README.md",
		"go.mod",
		"go.sum",
		".gitignore",
		"Makefile",
	}

	logs := []string{
		fmt.Sprintf("[%s] Application started", time.Now().Format("15:04:05")),
		fmt.Sprintf("[%s] Loaded %d files", time.Now().Format("15:04:05"), len(files)),
		fmt.Sprintf("[%s] Multi-panel layout initialized", time.Now().Format("15:04:05")),
		fmt.Sprintf("[%s] Ready for input", time.Now().Format("15:04:05")),
	}

	return model{
		focusedPanel: LeftPanel,
		panels:       []PanelID{LeftPanel, TopRightPanel, BottomRightPanel},
		panelBounds:  make(map[PanelID]panelBounds),
		files:        files,
		cursor:       0,
		details:      "Select a file to view details",
		logs:         logs,
		logScroll:    0,
		statusMsg:    "Tab/h/l: switch • ↑↓: navigate • Mouse: click/wheel",
	}
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// getSelectedFile returns the currently selected file
func (m model) getSelectedFile() string {
	if m.cursor >= 0 && m.cursor < len(m.files) {
		return m.files[m.cursor]
	}
	return ""
}

// addLog adds a new log entry
func (m *model) addLog(msg string) {
	timestamp := time.Now().Format("15:04:05")
	m.logs = append(m.logs, fmt.Sprintf("[%s] %s", timestamp, msg))
	// Auto-scroll to bottom
	m.logScroll = max(0, len(m.logs)-10)
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
