package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// types.go - Type Definitions
// Purpose: All type definitions, structs, enums, and constants
// When to extend: Add new types here when introducing new data structures

// Model represents the application state
type model struct {
	// Configuration
	config Config

	// UI State
	width  int
	height int

	// Focus management
	focusedComponent string

	// Error handling
	err       error
	statusMsg string

	// Add your application-specific state here
	// Example:
	// items []Item
	// cursor int
	// selected map[string]bool
}

// Config holds application configuration
type Config struct {
	// Theme
	Theme       string
	CustomTheme ThemeColors

	// Keybindings
	Keybindings       string
	CustomKeybindings map[string]string

	// Layout
	Layout LayoutConfig

	// UI Elements
	UI UIConfig

	// Performance
	Performance PerformanceConfig

	// Logging
	Logging LogConfig
}

// ThemeColors defines a color theme
type ThemeColors struct {
	Primary    string
	Secondary  string
	Background string
	Foreground string
	Accent     string
	Error      string
}

// LayoutConfig defines layout settings
type LayoutConfig struct {
	Type        string  // single, dual_pane, multi_panel, tabbed
	SplitRatio  float64 // For dual_pane
	ShowDivider bool
}

// UIConfig defines UI element settings
type UIConfig struct {
	ShowTitle       bool
	ShowStatus      bool
	ShowLineNumbers bool
	MouseEnabled    bool
	ShowIcons       bool
	IconSet         string
}

// PerformanceConfig defines performance settings
type PerformanceConfig struct {
	LazyLoading     bool
	CacheSize       int
	AsyncOperations bool
}

// LogConfig defines logging settings
type LogConfig struct {
	Enabled bool
	Level   string
	File    string
}

// Custom message types
// Add your application-specific messages here

type errMsg struct {
	err error
}

type statusMsg struct {
	message string
}

type resizeMsg struct {
	width  int
	height int
}

// Example custom messages:
// type itemSelectedMsg struct { item Item }
// type dataLoadedMsg struct { data []Item }
// type operationCompleteMsg struct { success bool }
