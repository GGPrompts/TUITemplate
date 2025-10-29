# Claude Code Development Notes

Documentation of key fixes and patterns for AI-assisted development of this TUI template system.

## Quick Reference - Critical Rules

**The 4 Golden Rules for TUI Layout:**

1. **Always Account for Borders** - Subtract 2 from height calculations BEFORE rendering panels
   ```
   contentHeight = totalHeight - titleLines - statusLines - 2 (borders)
   ```

2. **Never Auto-Wrap in Bordered Panels** - Always truncate text explicitly
   ```go
   maxTextWidth := panelWidth - 4  // -2 borders, -2 padding
   text = truncateString(text, maxTextWidth)
   ```

3. **Match Mouse Detection to Layout** - Use X coords for horizontal, Y coords for vertical
   ```go
   if m.shouldUseVerticalStack() {
       // Use msg.Y
   } else {
       // Use msg.X
   }
   ```

4. **Use Weights, Not Pixels** - Proportional layouts scale perfectly
   ```go
   width := (totalWidth * weight) / totalWeights
   ```

## Critical Layout Fixes

### Issue 1: Panels Covering Header / Border Overflow

**Symptom**: Panels would overflow and cover the title bar, especially on portrait/vertical monitors.

**Root Cause**: Height calculation didn't account for panel borders.

**The Math**:
```
WRONG:
contentHeight = totalHeight - 3 (title) - 1 (status) = totalHeight - 4
Panel renders with borders = contentHeight + 2 (borders)
Actual height used = totalHeight - 4 + 2 = totalHeight - 2 (TOO TALL!)

CORRECT:
contentHeight = totalHeight - 3 (title) - 1 (status) - 2 (borders) = totalHeight - 6
Panel renders with borders = contentHeight + 2
Actual height used = totalHeight - 6 + 2 = totalHeight - 4 ✓
```

**Fix** (`model.go`):
```go
func (m model) calculateLayout() (int, int) {
    contentWidth := m.width
    contentHeight := m.height

    if m.config.UI.ShowTitle {
        contentHeight -= 3 // title bar (3 lines)
    }
    if m.config.UI.ShowStatus {
        contentHeight -= 1 // status bar
    }

    // CRITICAL: Account for panel borders
    contentHeight -= 2 // top + bottom borders

    return contentWidth, contentHeight
}
```

### Issue 2: Panels Misaligned (One Row Different)

**Symptom**: When left panel focused in accordion mode, right panel appeared 1 row lower than left panel.

**Root Cause**: Text wrapping. When right panel was narrower, the subtitle `"Weight: 2 | Size: 80x25"` would wrap to 2 lines, making the panel taller.

**Fix** (`view.go`):
```go
// Calculate max text width to prevent wrapping
maxTextWidth := width - 4 // -2 for borders, -2 for padding

// Truncate ALL text before rendering
title = truncateString(title, maxTextWidth)
subtitle = truncateString(subtitle, maxTextWidth)

// Truncate content lines too
for i := 0; i < availableContentLines && i < len(content); i++ {
    line := truncateString(content[i], maxTextWidth)
    lines = append(lines, line)
}

// Helper function
func truncateString(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen-1] + "…"
}
```

**Key Insight**: NEVER let Lipgloss auto-wrap text in bordered panels. Always truncate explicitly.

### Issue 3: Mouse Clicks Broken in Vertical Stack Mode

**Symptom**: When terminal < 80 cols wide, panels stack vertically but mouse clicks don't change focus.

**Root Cause**: Click detection still using X coordinates (left/right) instead of Y coordinates (top/bottom).

**Fix** (`update_mouse.go`):
```go
func (m model) handleLeftClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
    // ... boundary checks ...

    if m.shouldUseVerticalStack() {
        // Vertical stack mode: use Y coordinates
        topHeight, _ := m.calculateVerticalStackLayout()
        relY := msg.Y - contentStartY

        if relY < topHeight {
            m.focusedPanel = "left"  // Top panel
        } else if relY > topHeight {
            m.focusedPanel = "right" // Bottom panel
        }
    } else {
        // Side-by-side mode: use X coordinates
        leftWidth, _ := m.calculateDualPaneLayout()

        if msg.X < leftWidth {
            m.focusedPanel = "left"
        } else if msg.X > leftWidth {
            m.focusedPanel = "right"
        }
    }

    return m, nil
}
```

**Key Insight**: Mouse detection logic must match the layout orientation (horizontal vs vertical).

## Weight-Based Dynamic Layout Pattern (LazyGit)

### Core Concept

Instead of calculating pixel widths, assign **weights** to panels:

```go
// Calculate weights based on focus
leftWeight, rightWeight := 1, 1

if m.accordionMode && m.focusedPanel == "left" {
    leftWeight = 2  // Focused panel gets 2x weight
}

// Calculate actual widths from weights
totalWeight := leftWeight + rightWeight
leftWidth := (availableWidth * leftWeight) / totalWeight
rightWidth := availableWidth - leftWidth
```

**Result**:
- Equal weights (1:1) = 50/50 split
- Focused weight (2:1) = 66/33 split
- Clean, proportional resizing with no fixed sizes

### Why This Works

1. **Proportional**: Always maintains exact ratios
2. **Simple**: No complex formulas, just division
3. **Immediate**: No animations needed, instant resize
4. **Flexible**: Change weight = instant layout change

## Debugging Checklist for Layout Issues

When panels don't align or render incorrectly:

1. **Check height accounting**:
   - Count title bar lines (usually 3)
   - Count status bar lines (usually 1)
   - Count panel borders (always 2)
   - Formula: `contentHeight = totalHeight - titleLines - statusLines - borderLines`

   **Visual Layout:**
   ```
   ┌─────────────────────────────────┐  ← Title Bar (3 lines)
   │  App Title                      │
   │  Subtitle/Info                  │
   ├─────────────────────────────────┤  ─┐
   │ ┌─────────────┬───────────────┐ │   │
   │ │             │               │ │   │
   │ │   Left      │     Right     │ │   │ Content Height
   │ │   Panel     │     Panel     │ │   │ (minus borders)
   │ │             │               │ │   │
   │ └─────────────┴───────────────┘ │   │
   ├─────────────────────────────────┤  ─┘
   │ Status Bar: Help text here      │  ← Status Bar (1 line)
   └─────────────────────────────────┘

   Panel borders (┌─┐ └─┘) = 2 lines total (top + bottom)
   ```

   **Height Calculation:**
   ```
   Total Terminal Height: 25
   - Title Bar:           -3
   - Status Bar:          -1
   - Panel Borders:       -2
   ─────────────────────────
   Content Height:        19 ✓
   ```

2. **Check text wrapping**:
   - Calculate max text width: `panelWidth - 4` (2 for borders, 2 for padding)
   - Truncate ALL strings before rendering
   - Never rely on auto-wrapping inside bordered panels

3. **Check mouse detection**:
   - Horizontal layout: use X coordinates
   - Vertical layout: use Y coordinates
   - Account for divider width/height (usually 1)
   - Use relative positions: `relY = msg.Y - contentStartY`

4. **Check border consistency**:
   - Use same border style for all panels
   - Don't mix `Height()` setting with natural height
   - Let content determine height, borders add to it

### Debugging Decision Tree

```
Panel Layout Problem?
│
├─ Panels covering title/status bar?
│  └─> Check height accounting (Rule #1)
│      - Did you subtract 2 for borders?
│      - Formula: totalHeight - titleLines - statusLines - 2
│
├─ Panels misaligned (different heights)?
│  └─> Check text wrapping (Rule #2)
│      - Is text wrapping to multiple lines?
│      - maxWidth = panelWidth - 4
│      - Truncate ALL strings explicitly
│
├─ Mouse clicks not working?
│  └─> Check mouse detection (Rule #3)
│      - Vertical stack? → Use msg.Y
│      - Horizontal? → Use msg.X
│      - Match detection to layout mode
│
└─ Accordion/resize janky?
   └─> Check weight calculations (Rule #4)
       - Using weights instead of pixels?
       - Formula: (totalWidth * weight) / totalWeights
```

## Common Pitfalls

❌ **DON'T**: Set explicit `Height()` on Lipgloss styles with borders
```go
// BAD: Can cause misalignment
panelStyle := lipgloss.NewStyle().
    Border(border).
    Height(height)  // ← Don't do this!
```

✅ **DO**: Fill content to exact height, let borders render naturally
```go
// GOOD: Fill content lines to exact height
for len(lines) < innerHeight {
    lines = append(lines, "")
}
panelStyle := lipgloss.NewStyle().
    Border(border)
    // No Height() - let content determine it
```

❌ **DON'T**: Assume layout orientation in mouse handlers
```go
// BAD: Always using X coordinate
if msg.X < leftWidth {
    // This breaks in vertical stack!
}
```

✅ **DO**: Check layout mode first
```go
// GOOD: Different logic per orientation
if m.shouldUseVerticalStack() {
    // Use Y coordinates
} else {
    // Use X coordinates
}
```

## Key Files Reference

- `model.go:80-98` - Layout calculation with border accounting
- `view.go:340-352` - Text truncation helper
- `view.go:143-180` - Dynamic panel rendering
- `update_mouse.go:79-106` - Mouse click detection with mode awareness

---

**Created**: 2025-01-19
**Last Updated**: 2025-10-23
**Purpose**: Preserve critical bug fixes and patterns for future AI-assisted development
