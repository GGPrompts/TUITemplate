package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// main.go - Application Entry Point
// Purpose: ONLY contains the main() function
// Rule: Never add business logic to this file. Keep it minimal.

func main() {
	// Load configuration
	cfg := loadConfig()

	// Create program with options based on config
	opts := []tea.ProgramOption{
		tea.WithAltScreen(),
	}

	if cfg.UI.MouseEnabled {
		opts = append(opts, tea.WithMouseCellMotion())
	}

	p := tea.NewProgram(
		initialModel(cfg),
		opts...,
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
