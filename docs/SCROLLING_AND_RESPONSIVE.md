# Scrolling and Responsive Content Guide

Best practices for handling overflow and making content responsive in apps built from TUITemplate.

## Problem: Content Too Tall for Available Height

When you have more content than available vertical space (common in list views, logs, long text).

---

## Solution 1: Simple Viewport with Offset (Recommended)

**Use Case:** Lists, logs, text viewers - anything that scrolls line-by-line

### Implementation

```go
// types.go - Add to your model
type model struct {
    items      []string // Your content
    cursor     int      // Currently selected item
    offset     int      // Scroll offset (top visible line)

    // ... other fields
}

// view.go - Render visible window
func (m model) renderScrollableList(width, height int) string {
    var lines []string

    // Calculate visible range
    visibleCount := height - 2 // -2 for borders/chrome
    start := m.offset
    end := min(start + visibleCount, len(m.items))

    // Render only visible items
    for i := start; i < end; i++ {
        line := m.items[i]

        // Highlight cursor
        if i == m.cursor {
            line = selectedStyle.Render("▶ " + line)
        } else {
            line = normalStyle.Render("  " + line)
        }

        lines = append(lines, line)
    }

    // Add scroll indicators
    if m.offset > 0 {
        lines = append([]string{dimStyle.Render("↑ more above")}, lines...)
    }
    if end < len(m.items) {
        lines = append(lines, dimStyle.Render("↓ more below"))
    }

    return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// update_keyboard.go - Handle scrolling
func (m model) moveDown() (tea.Model, tea.Cmd) {
    if m.cursor < len(m.items)-1 {
        m.cursor++

        // Auto-scroll: if cursor goes below visible area
        visibleCount := m.height - 4 // Account for title/status
        if m.cursor >= m.offset+visibleCount {
            m.offset = m.cursor - visibleCount + 1
        }
    }
    return m, nil
}

func (m model) moveUp() (tea.Model, tea.Cmd) {
    if m.cursor > 0 {
        m.cursor--

        // Auto-scroll: if cursor goes above visible area
        if m.cursor < m.offset {
            m.offset = m.cursor
        }
    }
    return m, nil
}

// Helper
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

**Key Points:**
- ✅ Only render visible lines (efficient for large lists)
- ✅ `offset` tracks which line is at the top
- ✅ `cursor` tracks selected item
- ✅ Auto-scrolls when cursor moves off-screen
- ✅ Shows scroll indicators (↑↓) when there's more content

---

## Solution 2: Using Bubbles Viewport Component

**Use Case:** Scrollable text blocks, markdown viewers, rich content

### Implementation

```go
// types.go
import "github.com/charmbracelet/bubbles/viewport"

type model struct {
    viewport viewport.Model
    content  string // Full content (can be very long)

    // ... other fields
}

// model.go - Initialize viewport
func initialModel(cfg Config) model {
    vp := viewport.New(80, 20) // width, height
    vp.SetContent("Your very long content here...\n" +
                  "Can have many lines...\n" +
                  "Will scroll automatically!")

    return model{
        viewport: vp,
        // ...
    }
}

// update.go - Handle viewport updates
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.viewport.Width = msg.Width
        m.viewport.Height = msg.Height - 4 // Account for chrome

    case tea.KeyMsg:
        // Viewport handles arrow keys, pgup/pgdown automatically
        m.viewport, cmd = m.viewport.Update(msg)
        return m, cmd
    }

    return m, nil
}

// view.go - Render viewport
func (m model) View() string {
    var sections []string

    if m.config.UI.ShowTitle {
        sections = append(sections, m.renderTitleBar())
    }

    // Viewport renders scrollable content
    sections = append(sections, m.viewport.View())

    if m.config.UI.ShowStatus {
        scrollInfo := fmt.Sprintf("Line %d/%d",
            m.viewport.YOffset+1,
            strings.Count(m.viewport.Content, "\n"))
        sections = append(sections, statusStyle.Render(scrollInfo))
    }

    return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

// Update content dynamically
func (m model) setContent(newContent string) {
    m.viewport.SetContent(newContent)
    m.viewport.GotoTop() // Or GotoBottom()
}
```

**Key Points:**
- ✅ Handles all scroll logic for you
- ✅ Built-in pgup/pgdown, mouse wheel support
- ✅ Tracks scroll position automatically
- ✅ Good for text-heavy content
- ❌ Renders ALL content (can be slow for huge lists)

---

## Solution 3: Page-Based Navigation

**Use Case:** Settings screens, forms, wizards - discrete pages of content

```go
// types.go
type model struct {
    currentPage int
    totalPages  int
    pageContent []string // Each entry is one page's content
}

// view.go
func (m model) renderPaginated(width, height int) string {
    content := m.pageContent[m.currentPage]

    pageInfo := fmt.Sprintf("Page %d/%d", m.currentPage+1, m.totalPages)

    return lipgloss.JoinVertical(lipgloss.Left,
        content,
        dimStyle.Render(pageInfo),
        dimStyle.Render("[PgUp/PgDn to navigate]"),
    )
}

// update_keyboard.go
case "pgdown":
    if m.currentPage < m.totalPages-1 {
        m.currentPage++
    }
    return m, nil

case "pgup":
    if m.currentPage > 0 {
        m.currentPage--
    }
    return m, nil
```

---

## Responsive Width: Making Boxes Use Available Space

### Problem: Fixed-width boxes waste space on wide terminals

**Bad:**
```go
// Always 50 columns, even on 200-column terminal
box := lipgloss.NewStyle().
    Width(50).
    Render(content)
```

**Good:**
```go
// Use percentage of available width
func (m model) renderResponsiveBox(availableWidth int) string {
    // Use 80% of available width, max 100 columns
    boxWidth := min((availableWidth * 80) / 100, 100)

    box := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        Width(boxWidth).
        Render(content)

    return box
}
```

### Adaptive Box Sizing by Terminal Width

```go
func (m model) getBoxWidth(availableWidth int) int {
    switch {
    case availableWidth < 60:
        // Narrow (mobile) - use full width
        return availableWidth - 4

    case availableWidth < 100:
        // Medium - use 90%
        return (availableWidth * 90) / 100

    case availableWidth < 150:
        // Wide - use 80%, max 120
        return min((availableWidth * 80) / 100, 120)

    default:
        // Very wide - cap at 140 for readability
        return 140
    }
}

// Usage
func (m model) renderDialog(width, height int) string {
    boxWidth := m.getBoxWidth(width)

    dialog := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(colorPrimary).
        Width(boxWidth).
        Padding(1, 2).
        Render("Dialog content here")

    // Center the dialog
    return lipgloss.Place(width, height,
        lipgloss.Center, lipgloss.Center,
        dialog)
}
```

### Multi-Column Layout on Wide Terminals

```go
func (m model) renderAdaptiveGrid(width, height int) string {
    items := m.getItems() // Your data

    // Determine columns based on width
    var columns int
    switch {
    case width < 80:
        columns = 1 // Single column on narrow
    case width < 120:
        columns = 2 // Two columns
    case width < 180:
        columns = 3 // Three columns
    default:
        columns = 4 // Four columns on very wide
    }

    // Calculate column width
    colWidth := (width - (columns + 1)) / columns // -1 for gaps

    // Build grid
    var rows []string
    var currentRow []string

    for i, item := range items {
        box := lipgloss.NewStyle().
            Width(colWidth).
            Border(lipgloss.RoundedBorder()).
            Render(item)

        currentRow = append(currentRow, box)

        // Complete row?
        if (i+1) % columns == 0 || i == len(items)-1 {
            rows = append(rows,
                lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
            currentRow = nil
        }
    }

    return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
```

---

## Complete Example: Scrollable File Browser

```go
// types.go
type model struct {
    files       []string // All files
    cursor      int      // Selected file
    offset      int      // Scroll position
    width       int
    height      int
    detailsPane string   // Info about selected file
}

// view.go
func (m model) View() string {
    if m.width < 80 {
        // Narrow: vertical stack
        return m.renderVerticalStack()
    }

    // Wide: side-by-side
    return m.renderSideBySide()
}

func (m model) renderSideBySide() string {
    contentWidth, contentHeight := m.calculateLayout()

    // Left pane: 40% width for file list
    leftWidth := (contentWidth * 40) / 100
    leftPane := m.renderScrollableFileList(leftWidth, contentHeight)

    // Right pane: 60% width for details
    rightWidth := contentWidth - leftWidth - 1
    rightPane := m.renderFileDetails(rightWidth, contentHeight)

    divider := m.renderDivider(contentHeight)

    return lipgloss.JoinHorizontal(lipgloss.Top,
        leftPane, divider, rightPane)
}

func (m model) renderScrollableFileList(width, height int) string {
    var lines []string

    visibleCount := height - 2
    start := m.offset
    end := min(start + visibleCount, len(m.files))

    // Scroll indicators
    if start > 0 {
        lines = append(lines, dimStyle.Render("↑ more above"))
    }

    // Visible files
    for i := start; i < end; i++ {
        file := m.files[i]

        // Truncate if too long
        maxLen := width - 4
        if len(file) > maxLen {
            file = file[:maxLen-1] + "…"
        }

        if i == m.cursor {
            lines = append(lines, selectedStyle.Render("▶ " + file))
        } else {
            lines = append(lines, normalStyle.Render("  " + file))
        }
    }

    if end < len(m.files) {
        lines = append(lines, dimStyle.Render("↓ more below"))
    }

    // Pad to full height
    for len(lines) < height {
        lines = append(lines, "")
    }

    content := strings.Join(lines, "\n")

    return lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(colorBorder).
        Width(width).
        Render(content)
}

func (m model) renderFileDetails(width, height int) string {
    // Details about selected file
    details := m.detailsPane

    // Wrap text to fit width
    maxWidth := width - 4
    wrappedLines := wrapText(details, maxWidth)

    // Truncate to height
    if len(wrappedLines) > height-2 {
        wrappedLines = wrappedLines[:height-2]
        wrappedLines = append(wrappedLines, dimStyle.Render("... (truncated)"))
    }

    content := strings.Join(wrappedLines, "\n")

    return lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(colorBorder).
        Width(width).
        Render(content)
}

// Helper: wrap text to max width
func wrapText(text string, maxWidth int) []string {
    words := strings.Fields(text)
    var lines []string
    var currentLine string

    for _, word := range words {
        testLine := currentLine
        if testLine != "" {
            testLine += " "
        }
        testLine += word

        if len(testLine) > maxWidth {
            if currentLine != "" {
                lines = append(lines, currentLine)
            }
            currentLine = word
        } else {
            currentLine = testLine
        }
    }

    if currentLine != "" {
        lines = append(lines, currentLine)
    }

    return lines
}

// update_keyboard.go
func (m model) moveDown() (tea.Model, tea.Cmd) {
    if m.cursor < len(m.files)-1 {
        m.cursor++

        // Auto-scroll
        visibleCount := m.height - 4
        if m.cursor >= m.offset+visibleCount {
            m.offset = m.cursor - visibleCount + 1
        }

        // Update details pane
        m.detailsPane = m.getFileInfo(m.files[m.cursor])
    }
    return m, nil
}
```

---

## Best Practices Summary

### For Scrolling:
1. ✅ **Use offset + visible range** for lists (most efficient)
2. ✅ **Use bubbles/viewport** for text-heavy content
3. ✅ **Always show scroll indicators** (↑↓) when there's more content
4. ✅ **Auto-scroll** when cursor/selection moves off-screen
5. ✅ **Support both keyboard and mouse wheel** scrolling

### For Responsive Width:
1. ✅ **Use percentages**, not fixed widths: `(width * 80) / 100`
2. ✅ **Set max widths** for readability: `min(calculated, 120)`
3. ✅ **Adapt layout** based on width: single column < 80, multi-column >= 80
4. ✅ **Truncate text** that's too long: `"text…"`
5. ✅ **Test at different sizes**: 40 cols (micro), 80 cols (normal), 120+ cols (wide)

### Don't Do This:
- ❌ Set `.Height()` on styles with borders → causes overflow
- ❌ Render all 10,000 items → use offset/viewport instead
- ❌ Hard-code widths like `.Width(50)` → use dynamic calculation
- ❌ Forget scroll indicators → users won't know there's more content
- ❌ Mix character widths → use monospace-safe truncation

---

## Quick Reference: Which Approach?

| Content Type | Best Solution | Why |
|-------------|---------------|-----|
| **File lists** | Offset + visible range | Efficient, auto-scroll |
| **Log viewer** | bubbles/viewport | Built-in scroll, tail support |
| **Long text** | bubbles/viewport | Handles wrapping, scrolling |
| **Forms/Settings** | Page-based navigation | Discrete sections |
| **Tables/Grids** | Offset + visible range | Row-by-row scrolling |
| **Dialogs** | Responsive width | Adapt to terminal size |
| **Multi-panel** | Weight-based layout | Proportional sizing |

---

**See Also:**
- `examples/tui-showcase/` - Working examples of all patterns
- `CLAUDE.md` - Critical layout rules and debugging
- `TERMUX_MOBILE_GUIDE.md` - Compact mode specifics
