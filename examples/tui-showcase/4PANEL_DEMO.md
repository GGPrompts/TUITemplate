# 4-Panel Accordion Layout Demo

The TUI Showcase now includes a **4-Panel Accordion Layout** (Tab 12)!

## How to Access

1. Run the showcase:
   ```bash
   cd /home/matt/projects/TUITemplate/examples/tui-showcase
   ./tui-showcase
   ```

2. Navigate to the "4-Panel" tab:
   - Press `Tab` repeatedly until you reach the "4-Panel" tab (tab 12)
   - OR click on the "[ 4-Panel ]" tab button with your mouse

## Layout Structure

```
┌─────────────────────────────────────────────┐
│ HEADER PANEL (25% default, 50% when focused)│
├─────────────────────┬───────────────────────┤
│ LEFT PANEL          │ RIGHT PANEL           │
│ (25% default)       │ (25% default)         │
│ (75% when focused)  │ (75% when focused)    │
├─────────────────────┴───────────────────────┤
│ FOOTER PANEL (25% default, 50% when focused)│
└─────────────────────────────────────────────┘
```

## Keyboard Controls

### Focus Panels
- **1** - Focus LEFT panel (expands to 75% width in middle row)
- **2** - Focus RIGHT panel (expands to 75% width in middle row)
- **3** - Focus FOOTER panel (expands to 50% height)
- **4** - Focus HEADER panel (expands to 50% height)

### Toggle Accordion Mode
- **a** or **A** - Toggle accordion mode ON/OFF
  - ON: Focused panel expands automatically
  - OFF: All panels stay at default size (25%:50%:25% vertical, 50:50 horizontal)

## Mouse Controls

- **Click any panel** to focus it
- **Focused panel** will expand automatically (if accordion mode is ON)
- Borders turn **blue** when focused

## Accordion Behavior

When accordion mode is ON:

### Header Focused (50% height)
```
Header:  50%  (2/4 weight)
Middle:  25%  (1/4 weight)
Footer:  25%  (1/4 weight)
```

### Footer Focused (50% height)
```
Header:  25%  (1/4 weight)
Middle:  25%  (1/4 weight)
Footer:  50%  (2/4 weight)
```

### Left/Right Focused (66% middle row)
```
Header:  16.67%  (1/6 weight)
Middle:  66.67%  (4/6 weight)
Footer:  16.67%  (1/6 weight)

Within middle row:
- Left focused:  75% width (3/4 weight)
- Right focused: 75% width (3/4 weight)
```

## Use Cases

### For Tmuxplexer

This layout is **perfect** for a tmux session manager:

**Header Panel** (expandable to 50%):
- Session statistics
- Quick-launch templates
- Filter/search controls
- Server status

**Left Panel** (expandable to 75%):
- Session list
- Full metadata when expanded
- Scroll through many sessions

**Right Panel** (expandable to 75%):
- Window/pane details
- Visual pane layouts
- Per-pane actions

**Footer Panel** (expandable to 50%):
- Live pane preview (full screen when focused)
- Natural language command input
- Command history
- Quick action buttons

## Key Innovation

This is the **first TUI layout** that allows:
- ✅ **4 panels** (header + middle row + footer)
- ✅ **ALL panels independently expandable**
- ✅ **Header/footer expansion** (not just side panels)
- ✅ **Smart weight distribution** (focused panel gets priority)
- ✅ **Both horizontal AND vertical accordion**

## Testing

Try these sequences:

1. **Test header expansion:**
   - Press `4` → Header expands to 50%
   - Press `1` → Left expands, middle row gets 66%

2. **Test footer expansion:**
   - Press `3` → Footer expands to 50%
   - Press `2` → Right expands, middle row gets 66%

3. **Test accordion toggle:**
   - Press `a` → Turn off accordion
   - Press `1`, `2`, `3`, `4` → Panels stay same size
   - Press `a` → Turn on accordion
   - Press `1` → Left panel expands again

4. **Test mouse:**
   - Click on header → Header expands
   - Click on footer → Footer expands
   - Click on left → Left expands
   - Click on right → Right expands

## Status Messages

Watch the status bar at the bottom - it shows:
- Which panel is focused
- Accordion mode status
- Mouse position
- Current weights

## Next Steps

Once this is tested and working well, we can:
1. Port it to tmuxplexer
2. Add content specific to tmux (sessions, windows, panes, preview)
3. Add tmux-specific interactions
4. Test on Termux mobile

---

**Built with:** TUITemplate dynamic panel system
**Date:** 2025-10-24
