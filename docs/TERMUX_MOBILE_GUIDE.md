# Termux Mobile TUI Development Guide

Complete guide for building mobile-optimized TUI applications for Termux.

## Quick Reference: Typical Termux Dimensions

### With Keyboard Open (COMPACT MODE)
- **Portrait:** 70-80 cols × **8-12 rows** ⚠️ Very constrained!
- **Landscape:** 130-150 cols × **8-10 rows**

### Without Keyboard (MOBILE MODE)
- **Portrait:** 70-80 cols × **35-45 rows**
- **Landscape:** 130-150 cols × **20-25 rows**

### Desktop Terminal (FULL MODE)
- **Standard:** 80-120 cols × **24-60 rows**

*Note: Exact sizes vary by device, keyboard app, and font size*

## Step 1: Detect Your Terminal Size

Use the terminal size detector:

```bash
cd /home/matt/projects/termux-size-detector
go build

# Transfer to Termux (or build there directly)
./termux-size-detector
```

The detector will show:
- Real-time dimensions as you resize
- Max/min sizes recorded
- Category (MICRO/SMALL/MEDIUM/LARGE)
- Layout recommendations

## Step 2: Design for Mobile Constraints

### MICRO Layout (≤10 rows) - Termux with Keyboard

**Critical Design Rules:**
1. ❌ **NO status bar** - every line counts!
2. ✅ Ultra-compact title (1 line only)
3. ✅ Show only 5-7 items at a time
4. ✅ Use scroll indicators (↑↓)
5. ✅ Minimal padding/margins
6. ✅ Large touch targets for mouse
7. ✅ Single column layout only

**Example Layout (10 rows):**
```
Row 1:  📱 App Name (7×80)
Row 2:  ▶ Item 1 (selected)
Row 3:    Item 2
Row 4:    Item 3
Row 5:    Item 4
Row 6:    Item 5
Row 7:    Item 6
Row 8:    Item 7
Row 9:    ↓ (scroll indicator)
Row 10: [Hidden status or nothing]
```

### SMALL Layout (11-20 rows) - Termux Full Screen

**Design Rules:**
1. ✅ Minimal title bar (1-2 lines)
2. ✅ Show 10-15 items
3. ✅ Single pane works best
4. ✅ Optional compact status bar
5. ✅ Tight margins

### MEDIUM+ Layout (21+ rows) - Desktop

**Design Rules:**
1. ✅ All standard layouts work
2. ✅ Multi-panel if needed
3. ✅ Full title + status bars

## Step 3: Implement Responsive Layouts

### Detect Terminal Size in Your App

```go
// types.go
type model struct {
    width     int
    height    int
    isMobile  bool  // height <= 15
    isCompact bool  // height <= 10
    // ... other fields
}

// model.go
func (m *model) setSize(width, height int) {
    m.width = width
    m.height = height
    m.isMobile = height <= 15
    m.isCompact = height <= 10

    // Adjust UI based on size
    if m.isCompact {
        m.config.UI.ShowStatus = false  // Save space!
    }
}
```

### Render Different Views

```go
// view.go
func (m model) renderMainContent() string {
    contentWidth, contentHeight := m.calculateLayout()

    if m.isCompact {
        return m.renderCompactView(contentWidth, contentHeight)
    } else if m.isMobile {
        return m.renderMobileView(contentWidth, contentHeight)
    }
    return m.renderDesktopView(contentWidth, contentHeight)
}

func (m model) renderCompactView(width, height int) string {
    // Show only 5-7 items
    visibleCount := min(height-2, 6)

    // Ultra-compact rendering
    // ...
}
```

### Handle Scrolling for Small Screens

```go
// types.go
type model struct {
    offset int  // scroll offset for compact view
    // ...
}

// update_keyboard.go
func (m model) moveDown() (tea.Model, tea.Cmd) {
    if m.cursor < len(m.items)-1 {
        m.cursor++

        // Auto-scroll in compact mode
        if m.isCompact {
            visibleCount := m.height - 2
            if m.cursor >= m.offset+visibleCount {
                m.offset = m.cursor - visibleCount + 1
            }
        }
    }
    return m, nil
}
```

## Step 4: Optimize for Touch (Mouse Support)

Make clickable elements large enough for finger taps:

```go
// Minimum 3×1 for buttons
buttonStyle := lipgloss.NewStyle().
    Padding(0, 2).  // Horizontal padding for width
    Render(" Button ")

// Items should span full width
itemStyle := lipgloss.NewStyle().
    Width(width-4).  // Nearly full width
    Render("▶ " + item)
```

## Mobile-Optimized Template Structure

```
termux-mobile-app/
├── main.go              # Entry point
├── types.go             # Add: isMobile, isCompact, offset
├── model.go             # Detect size, adjust UI config
├── view.go              # 3 views: compact/mobile/desktop
├── update_keyboard.go   # Smart scrolling
├── styles.go            # Compact styles
└── config.go            # Mobile defaults
```

## Best Practices for Termux

### 1. **Test with Keyboard Open**

Always test the worst case (8-10 rows). If it works there, it works everywhere.

### 2. **Provide Size Info**

Show dimensions in compact mode so users know what they're working with:

```go
title := "📱 App (80×10)"  // Show size in compact mode
```

### 3. **Use Scroll Indicators**

```go
if m.offset > 0 {
    content += dimStyle.Render("↑ more above\n")
}
if m.offset+visibleCount < len(m.items) {
    content += dimStyle.Render("↓ more below")
}
```

### 4. **Minimize Visual Noise**

In compact mode:
- No decorative borders
- Minimal icons (or emoji only)
- Short labels
- No padding/margins

### 5. **Support Both Portrait and Landscape**

```go
if m.width > 100 {
    // Landscape mode - can show more horizontally
    return m.renderWideView()
} else {
    // Portrait mode - stack vertically
    return m.renderNarrowView()
}
```

### 6. **Use Efficient Data Structures**

Only load/render what's visible:

```go
// Virtual scrolling
start := m.offset
end := min(start + visibleCount, len(m.items))

for i := start; i < end; i++ {
    // Render only visible items
}
```

## Common Termux Issues & Solutions

### Issue: Too Many Items

**Solution:** Virtual scrolling + pagination

```go
visibleCount := 6  // For compact mode
totalPages := (len(m.items) + visibleCount - 1) / visibleCount
currentPage := m.offset / visibleCount + 1

status := fmt.Sprintf("Page %d/%d", currentPage, totalPages)
```

### Issue: Text Overflows

**Solution:** Truncate smartly

```go
func truncate(s string, maxWidth int) string {
    if len(s) <= maxWidth {
        return s
    }
    return s[:maxWidth-1] + "…"
}
```

### Issue: Can't See Full Content

**Solution:** Horizontal scrolling or word wrap

```go
// Option 1: Wrap
content := lipgloss.NewStyle().
    Width(width).
    Render(text)  // Auto-wraps

// Option 2: Horizontal scroll
content := text[m.scrollX:m.scrollX+width]
```

## Building for Termux

### On Development Machine (Cross-compile)

```bash
# For 64-bit ARM (most modern Android devices)
GOOS=linux GOARCH=arm64 go build -o app

# For 32-bit ARM (older devices)
GOOS=linux GOARCH=arm GOARM=7 go build -o app
```

### Transfer to Termux

```bash
# Via USB/adb
adb push app /sdcard/
# Then in Termux:
# mv /sdcard/app ~/
# chmod +x ~/app

# Or via network
scp app username@phone-ip:~/
```

### Build Directly in Termux

```bash
# In Termux
pkg install golang
go build
```

## Testing Checklist

- [ ] Test with keyboard open (8-10 rows)
- [ ] Test with keyboard closed (35-45 rows)
- [ ] Test in portrait mode
- [ ] Test in landscape mode
- [ ] Test scrolling with many items
- [ ] Test touch/mouse interactions
- [ ] Test with different font sizes
- [ ] Verify no horizontal scrolling needed
- [ ] Check all text is readable
- [ ] Ensure responsive layout switches correctly

## Example Apps to Build

### Perfect for Mobile

1. **Todo List** - Simple, vertical, touch-friendly
2. **Notes** - Text viewing/editing
3. **Package Manager** - List of packages
4. **File Browser** - Navigate filesystem
5. **System Monitor** - Show stats
6. **Quick Launcher** - App menu

### Consider Carefully

❌ **Code Editors** - Too complex for small screens
❌ **Multi-panel Dashboards** - Not enough space
❌ **Tables with many columns** - Horizontal scrolling needed

## Resources

- **Size Detector:** `/home/matt/projects/termux-size-detector`
- **Mobile Demo:** `/home/matt/projects/termux-mobile-demo`
- **Template:** `/home/matt/projects/TUITemplate`

## Quick Start Command

```bash
# Detect your Termux size
cd /home/matt/projects/termux-size-detector
go build && ./termux-size-detector

# Note the dimensions, then create your app
cd /home/matt/projects/TUITemplate
./scripts/new_project.sh

# Remember to implement responsive layouts!
```

---

**Build amazing mobile TUIs!** 📱✨

Termux is powerful - with proper responsive design, your TUI apps can work beautifully on mobile devices!
