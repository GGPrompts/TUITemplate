package main

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

	// Layout demo specific state
	currentLayout string // Current layout being displayed
	currentTab    int    // For tabbed layout (0-5: Overview, Content, Settings, Borders, Colors, Dynamic)

	// Mouse tracking
	mouseX      int    // Current mouse X position
	mouseY      int    // Current mouse Y position
	hoveredItem string // Currently hovered item (for visual feedback)
	lastClicked string // Last clicked panel/button

	// Dynamic panels (Tab 5) state
	focusedPanel  string   // Which panel is focused for dynamic resizing
	accordionMode bool     // Whether accordion mode is enabled
	leftContent   []string // Content for left panel
	rightContent  []string // Content for right panel
	bottomContent []string // Content for bottom panel (3-panel mode)

	// Menu system state
	menuOpen        bool   // Whether any menu is currently open
	activeMenu      string // Which menu is active ("file", "view", "components", "help", "")
	selectedMenuItem int    // Index of selected item in active menu (-1 = none)
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

// Menu types

// MenuItem represents a single menu item
type MenuItem struct {
	Label    string   // Display text
	Action   string   // Action identifier (e.g., "quit", "switch-layout-single")
	Shortcut string   // Keyboard shortcut display (e.g., "Ctrl+Q")
	Disabled bool     // Whether item is disabled
	IsSeparator bool  // Whether this is a separator line
}

// Menu represents a dropdown menu
type Menu struct {
	Label string     // Menu label in menu bar
	Items []MenuItem // Menu items
}
