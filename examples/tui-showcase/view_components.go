package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// view_components.go - Component Showcase Renderers
// Purpose: Rendering functions for component demonstration tabs
// When to extend: Add new component showcases here

// renderFormsTab demonstrates form components
func (m model) renderFormsTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║      FORM COMPONENTS SHOWCASE         ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Text Input
	textInputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(50).
		Render("┌─ Text Input ────────────────┐\n│ Username: [john_doe______] │\n│ Email:    [john@example.com]│\n└─────────────────────────────┘")
	content.WriteString(textInputBox)
	content.WriteString("\n\n")

	// Checkbox Group
	checkboxBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(50).
		Render("┌─ Checkboxes ────────────────┐\n│ [✓] Enable notifications    │\n│ [ ] Dark mode              │\n│ [✓] Auto-save              │\n└─────────────────────────────┘")
	content.WriteString(checkboxBox)
	content.WriteString("\n\n")

	// Radio Buttons
	radioBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(50).
		Render("┌─ Radio Buttons ─────────────┐\n│ ( ) Small                  │\n│ (•) Medium                 │\n│ ( ) Large                  │\n└─────────────────────────────┘")
	content.WriteString(radioBox)
	content.WriteString("\n\n")

	// Validation Example
	validationBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorError).
		Padding(1, 2).
		Width(50).
		Render("┌─ Validation ────────────────┐\n│ Password: [********]       │\n│ ⚠ Must be 8+ characters    │\n└─────────────────────────────┘")
	content.WriteString(validationBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderTablesTab demonstrates table/grid components
func (m model) renderTablesTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║      TABLE COMPONENTS SHOWCASE        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Data Table
	tableBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(70).
		Render(
			"┌─ User Data Table ──────────────────────────────────┐\n" +
			"│ ID │ Name        │ Email              │ Status   │\n" +
			"├────┼─────────────┼────────────────────┼──────────┤\n" +
			"│ 1  │ Alice Smith │ alice@example.com  │ ✓ Active │\n" +
			"│ 2  │ Bob Jones   │ bob@example.com    │ ✓ Active │\n" +
			"│ 3  │ Carol White │ carol@example.com  │ ⚠ Pending│\n" +
			"└────┴─────────────┴────────────────────┴──────────┘")
	content.WriteString(tableBox)
	content.WriteString("\n\n")

	// Sorting & Filtering
	controlsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(70).
		Render(
			"┌─ Table Controls ──────────────────────────────────┐\n" +
			"│ Sort by: Name ▼  │  Filter: [____]  │  3 rows  │\n" +
			"└───────────────────────────────────────────────────┘")
	content.WriteString(controlsBox)
	content.WriteString("\n\n")

	// Row Selection
	selectionBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(70).
		Render(
			"┌─ Selection ───────────────────────────────────────┐\n" +
			"│ [✓] Row 1 - Alice Smith                          │\n" +
			"│ [ ] Row 2 - Bob Jones                            │\n" +
			"│ [✓] Row 3 - Carol White                          │\n" +
			"│                                      2 selected   │\n" +
			"└───────────────────────────────────────────────────┘")
	content.WriteString(selectionBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderDialogsTab demonstrates dialog/modal components
func (m model) renderDialogsTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║     DIALOG COMPONENTS SHOWCASE        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Confirmation Dialog
	confirmBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(50).
		Align(lipgloss.Center).
		Render(
			"⚠  Confirm Action\n\n" +
			"Are you sure you want to delete\n" +
			"this item? This cannot be undone.\n\n" +
			"[  Cancel  ]  [  Delete  ]")
	content.WriteString(confirmBox)
	content.WriteString("\n\n")

	// Alert Dialog
	alertBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(50).
		Align(lipgloss.Center).
		Render(
			"✓  Success\n\n" +
			"Your changes have been saved\n" +
			"successfully!\n\n" +
			"[     OK     ]")
	content.WriteString(alertBox)
	content.WriteString("\n\n")

	// Error Dialog
	errorBox := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(colorError).
		Padding(1, 2).
		Width(50).
		Align(lipgloss.Center).
		Render(
			"✗  Error\n\n" +
			"Failed to connect to server.\n" +
			"Please check your connection.\n\n" +
			"[   Retry   ]  [  Cancel  ]")
	content.WriteString(errorBox)
	content.WriteString("\n\n")

	// Input Prompt
	promptBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorInfo).
		Padding(1, 2).
		Width(50).
		Render(
			"  Enter Your Name\n\n" +
			"Name: [________________]\n\n" +
			"[  Cancel  ]  [   OK   ]")
	content.WriteString(promptBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderProgressTab demonstrates progress indicators
func (m model) renderProgressTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║    PROGRESS COMPONENTS SHOWCASE       ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Determinate Progress Bar
	progressBox1 := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(60).
		Render(
			"Downloading file.zip\n\n" +
			"[████████████░░░░░░░░] 65%\n\n" +
			"325 MB / 500 MB")
	content.WriteString(progressBox1)
	content.WriteString("\n\n")

	// Multiple Progress Bars
	progressBox2 := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(60).
		Render(
			"Build Progress\n\n" +
			"Compiling:  [████████████████████] 100%\n" +
			"Testing:    [██████████░░░░░░░░░░]  50%\n" +
			"Packaging:  [░░░░░░░░░░░░░░░░░░░░]   0%")
	content.WriteString(progressBox2)
	content.WriteString("\n\n")

	// Spinners
	spinnerBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(60).
		Render(
			"Loading States\n\n" +
			"◐ Loading...          │ ⠋ Processing...\n" +
			"⣾ Fetching data...    │ ⠙ Connecting...\n" +
			"◓ Please wait...      │ ⠹ Syncing...")
	content.WriteString(spinnerBox)
	content.WriteString("\n\n")

	// Indeterminate Progress
	indeterminateBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorInfo).
		Padding(1, 2).
		Width(60).
		Render(
			"Indeterminate Progress\n\n" +
			"[░░░░████░░░░░░░░░░░░] Calculating...")
	content.WriteString(indeterminateBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderTreeViewTab demonstrates tree/hierarchical components
func (m model) renderTreeViewTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║     TREE VIEW COMPONENTS SHOWCASE     ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// File Browser Tree
	treeBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(60).
		Render(
			"📁 File Browser\n\n" +
			"▼ 📁 project/\n" +
			"  ▼ 📁 src/\n" +
			"    ▸ 📁 components/\n" +
			"    ▸ 📁 utils/\n" +
			"      📄 main.go\n" +
			"      📄 config.go\n" +
			"  ▼ 📁 tests/\n" +
			"      📄 main_test.go\n" +
			"    📄 README.md\n" +
			"    📄 go.mod")
	content.WriteString(treeBox)
	content.WriteString("\n\n")

	// Hierarchical Data
	hierarchyBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(60).
		Render(
			"Organization Structure\n\n" +
			"▼ Engineering\n" +
			"  ▸ Frontend Team\n" +
			"  ▼ Backend Team\n" +
			"    • Alice (Lead)\n" +
			"    • Bob (Senior)\n" +
			"    • Carol (Junior)\n" +
			"  ▸ DevOps Team")
	content.WriteString(hierarchyBox)
	content.WriteString("\n\n")

	// Tree Controls
	controlsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(60).
		Render(
			"Tree Controls\n\n" +
			"▸ Click to expand/collapse\n" +
			"→/← Arrow keys to navigate\n" +
			"Space to select items\n" +
			"/ to search in tree")
	content.WriteString(controlsBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderMobileTab demonstrates mobile/touch-friendly patterns
func (m model) renderMobileTab(width, height int) string {
	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║   MOBILE-OPTIMIZED UI PATTERNS        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Vertical Stack Layout (Termux pattern)
	stackBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(min(width-4, 50)).
		Render(
			"Vertical Stack Layout\n\n" +
			"Automatically activates when\n" +
			"terminal width < 80 columns\n\n" +
			"Perfect for:\n" +
			"• Mobile terminals (Termux)\n" +
			"• Portrait mode displays\n" +
			"• Narrow terminal windows")
	content.WriteString(stackBox)
	content.WriteString("\n\n")

	// Touch-Friendly Buttons
	buttonsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorAccent).
		Padding(1, 2).
		Width(min(width-4, 50)).
		Render(
			"Touch-Friendly Buttons\n\n" +
			"┌────────────────────┐\n" +
			"│   Large Button     │ (3+ lines tall)\n" +
			"└────────────────────┘\n\n" +
			"Easy to tap on small screens")
	content.WriteString(buttonsBox)
	content.WriteString("\n\n")

	// Adaptive Layout
	adaptiveBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorSecondary).
		Padding(1, 2).
		Width(min(width-4, 50)).
		Render(
			"Adaptive Design\n\n" +
			fmt.Sprintf("Current width: %d cols\n", width) +
			"Layout mode: " + getLayoutMode(width) + "\n\n" +
			"Resize terminal to see\n" +
			"automatic layout changes")
	content.WriteString(adaptiveBox)
	content.WriteString("\n\n")

	// Mobile Tips
	tipsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorInfo).
		Padding(1, 2).
		Width(min(width-4, 50)).
		Render(
			"Mobile Best Practices\n\n" +
			"✓ Vertical stacking < 80 cols\n" +
			"✓ Large tap targets (3+ lines)\n" +
			"✓ Clear focus indicators\n" +
			"✓ Swipe-like key navigation\n" +
			"✓ Minimal text wrapping")
	content.WriteString(tipsBox)

	// Fit content to available height to prevent overflow
	// Leave small buffer to account for any rendering overhead
	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1 // -1 for safety buffer
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderMetaballsTab demonstrates metaballs effect
func (m model) renderMetaballsTab(width, height int) string {
	if m.metaballEngine == nil {
		return contentStyle.Width(width).Render("Metaballs engine not initialized")
	}

	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║      METABALLS EFFECT SHOWCASE        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	content.WriteString(lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Physics-based floating blobs with organic motion"))
	content.WriteString("\n\n")

	// Render the metaballs effect
	metaballsRender := m.metaballEngine.Render()
	content.WriteString(metaballsRender)

	content.WriteString("\n\n")
	content.WriteString(lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Field strength = radius² / distance² | Gradient rendering with Unicode blocks"))

	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderWavyMenuTab demonstrates wave grid effect
func (m model) renderWavyMenuTab(width, height int) string {
	if m.waveGrid == nil {
		return contentStyle.Width(width).Render("Wave grid not initialized")
	}

	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║      WAVY GRID EFFECT SHOWCASE        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	content.WriteString(lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Sine wave distortion for animated grid backgrounds"))
	content.WriteString("\n\n")

	// Render the wavy grid effect
	gridRender := m.waveGrid.Render()
	content.WriteString(gridRender)

	content.WriteString("\n\n")
	content.WriteString(lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("wave = sin(y/5 + frame/20) × amplitude | Perfect for menu backgrounds"))

	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderRainbowTab demonstrates rainbow text effect
func (m model) renderRainbowTab(width, height int) string {
	if m.rainbowCycler == nil {
		return contentStyle.Width(width).Render("Rainbow cycler not initialized")
	}

	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║      RAINBOW TEXT EFFECT SHOWCASE     ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	content.WriteString(lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("Animated rainbow colors cycling through text"))
	content.WriteString("\n\n")

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
	content.WriteString(rainbowArt)

	content.WriteString("\n\n")

	// Single line example
	exampleText := m.rainbowCycler.Render("The quick brown fox jumps over the lazy dog")
	content.WriteString(exampleText)

	content.WriteString("\n\n")
	content.WriteString(lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("Per-character coloring with frame-based shifting | Vertical wave patterns"))

	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// renderLandingPageTab demonstrates all effects combined
func (m model) renderLandingPageTab(width, height int) string {
	if m.waveGrid == nil || m.metaballEngine == nil || m.rainbowCycler == nil {
		return contentStyle.Width(width).Render("Effects not initialized")
	}

	// This would ideally use the compositor to combine grid + metaballs + rainbow title
	// For simplicity in the showcase, let's just show the metaballs over the grid

	var content strings.Builder

	content.WriteString(titleStyle.Render("╔═══════════════════════════════════════╗"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("║     LANDING PAGE - ALL EFFECTS        ║"))
	content.WriteString("\n")
	content.WriteString(titleStyle.Render("╚═══════════════════════════════════════╝"))
	content.WriteString("\n\n")

	// Rainbow title
	titleArt := []string{
		"████████╗██╗   ██╗██╗",
		"╚══██╔══╝██║   ██║██║",
		"   ██║   ██║   ██║██║",
	}
	rainbowTitle := m.rainbowCycler.RenderLines(titleArt)
	content.WriteString(rainbowTitle)

	content.WriteString("\n\n")

	// Show metaballs effect
	metaballsRender := m.metaballEngine.Render()
	metaballLines := strings.Split(metaballsRender, "\n")
	if len(metaballLines) > 10 {
		metaballLines = metaballLines[:10]
	}
	content.WriteString(strings.Join(metaballLines, "\n"))

	content.WriteString("\n\n")
	content.WriteString(lipgloss.NewStyle().
		Foreground(colorInfo).
		Render("✨ Wavy Grid + Metaballs + Rainbow = Beautiful TUIs ✨"))

	content.WriteString("\n\n")
	content.WriteString(lipgloss.NewStyle().
		Foreground(colorDimmed).
		Render("See examples/effects/ for standalone demos | Use in your own apps!"))

	contentStr := content.String()
	lines := strings.Split(contentStr, "\n")
	maxLines := height - 1
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return contentStyle.Width(width).Render(strings.Join(lines, "\n"))
}

// Helper functions

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getLayoutMode(width int) string {
	if width < 80 {
		return "Vertical Stack"
	}
	return "Side-by-Side"
}
