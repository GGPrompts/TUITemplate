# TFE (Terminal File Explorer) Integration Strategy

**Date:** 2025-10-20
**Purpose:** Guide for integrating a Project Manager TUI with TFE

## Executive Summary

TFE is an exceptionally well-architected terminal file manager with **built-in support for launching external TUI applications**. It already integrates lazygit, lazydocker, htop, and lnav. Your Project Manager can integrate in the same way with minimal effort.

**Key Finding:** TFE is not just compatible with your PM - it's specifically designed to work with tools like it.

---

## TFE Overview

| Aspect | Detail |
|--------|--------|
| **Language** | Go 1.24+ |
| **Framework** | Bubbletea (TUI) + Lipgloss (styling) + Bubbles (components) |
| **Architecture** | 14 focused modules, ~4,500 LOC |
| **Purpose** | Modern Midnight Commander alternative |
| **Key Feature** | Built-in app launcher for external TUI tools |

### Core Features
- Dual-pane file browser with live syntax-highlighted preview
- Context menu (right-click or F2)
- Keyboard shortcuts (F1-F12, vim-style navigation)
- Full mouse support
- 3 view modes: List, Detail, Tree
- 3 display layout modes: Single-pane, Dual-pane, Full preview
- Fuzzy search (Ctrl+P)
- Favorites system
- Trash/Recycle bin
- Prompts library

---

## Integration Approaches (Ranked by Complexity)

### 1. Context Menu Launch (EASIEST - 2-3 hours)

**What:** Add "📋 Project Manager" to right-click menu

**How:**
```go
// In context_menu.go, add menu item if tool available:
if editorAvailable("pm") {
    items = append(items, contextMenuItem{"📋 Project Manager", "launch_pm"})
}

// Add handler:
case "launch_pm":
    return m, openTUITool("pm", m.currentPath)
```

**Pros:**
- Minimal code changes (1 file)
- Zero risk of breaking existing features
- Launches PM with full directory context
- Users already know right-click = options

**Cons:**
- PM runs independently, no state sharing

**Perfect for:** Quick integration, keeping PM standalone

---

### 2. PM Toggle View Mode (MEDIUM - 4-6 hours)

**What:** Press Ctrl+Shift+P to show embedded PM dashboard alongside file browser

**How:**
- Add new `viewProjectManager` mode to types.go
- Handle Ctrl+Shift+P in update_keyboard.go
- Create renderer function in project_manager.go
- Add to view.go dispatcher

**Pros:**
- Integrated experience
- Can share state with file browser
- Users can toggle between views
- Seamless workflow

**Cons:**
- More code to maintain
- PM logic now in TFE codebase

**Perfect for:** Tight integration, seamless UX

---

### 3. PM as Side Panel (COMPLEX - 8-12 hours)

**What:** Like dual-pane mode but with PM on right instead of preview

**Layout:**
```
┌──────────────────────┬──────────────────┐
│   File Browser       │   Project Info   │
│                      │   Tasks          │
│   src/               │   Contributors   │
│   • main.go          │   Status         │
│                      │   ...            │
└──────────────────────┴──────────────────┘
```

**Pros:**
- Ultimate integration
- See files and projects simultaneously
- Full state sharing
- Professional UX

**Cons:**
- Complex layout management
- More keyboard handling needed
- Couples PM into TFE

**Perfect for:** Full-featured IDE-like experience

---

## Architecture Patterns Used by TFE

### Pattern 1: Message-Driven Updates

TFE uses Bubbletea's message-passing architecture:

```
User Input → Message → Update Handler → Model Change → Render
```

**For PM:** Add new message types without touching core dispatch logic.

### Pattern 2: View Mode Dispatcher

TFE supports multiple view modes via a simple switch:

```go
func (m model) View() string {
    switch m.viewMode {
    case viewSinglePane:
        return m.renderSinglePane()
    case viewDualPane:
        return m.renderDualPane()
    case viewFullPreview:
        return m.renderFullPreview()
    }
}
```

**For PM:** Add `viewProjectManager` case and render your PM UI.

### Pattern 3: External App Launching

TFE has proven pattern for launching and managing child processes:

```go
func openTUITool(tool, dir string) tea.Cmd {
    c := exec.Command(tool)
    c.Dir = dir  // Pass working directory
    
    return tea.Sequence(
        tea.ClearScreen,
        tea.ExecProcess(c, func(err error) tea.Msg {
            return editorFinishedMsg{err}
        }),
    )
}
```

**For PM:** Use this exact pattern to launch your PM.

### Pattern 4: Dialog System

TFE includes dialogs for input, confirmation, and messages:

```go
m.showDialog = true
m.dialog.dialogType = dialogInput
m.dialog.title = "Enter text:"
```

**For PM:** Extend with new dialog types like `dialogSelectProject`.

### Pattern 5: Status Messages

Automatic feedback system with auto-clear:

```go
m.setStatusMessage("Operation completed", false)  // Green, auto-disappears
m.setStatusMessage("Error occurred", true)        // Red, auto-disappears
```

**For PM:** Use same system for task creation, project updates, etc.

---

## File-by-File Implementation Guide

### Minimal Change (Context Menu)

**Edit:** `context_menu.go` only

Lines to modify:
1. Line ~70: Add detection
2. Line ~260: Add handler
3. Lines ~100-102: Add menu item

### Medium Change (PM View Mode)

**Edit:** 4 files

1. **types.go** - Add viewProjectManager constant
2. **context_menu.go** - Same as minimal
3. **update_keyboard.go** - Add Ctrl+Shift+P handler
4. **view.go** - Add dispatcher case
5. **project_manager.go** (NEW) - Render function

### Complex Change (Side Panel)

**Edit:** 6+ files

- All above PLUS:
- model.go - Layout calculations
- styles.go - New PM styles
- update_keyboard.go - Pane switching
- render_file_list.go - Adjust widths

---

## Critical Implementation Details

### 1. Working Directory Context

TFE always passes `c.Dir = path` to external apps:

```go
return m, openTUITool("pm", m.currentPath)  // PM launches in this directory
```

Your PM should:
- Accept working directory as context
- Read/write project files from that directory
- Use relative paths where possible

### 2. Terminal State Management

TFE handles terminal cleanup automatically:
- Clears screen before launching
- Restores state after
- No manual ANSI codes needed

Your PM just needs to:
- Use Bubbletea or similar
- Don't worry about terminal cleanup
- Let parent process handle restoration

### 3. Error Handling

TFE captures exit codes and shows feedback:

```go
func(err error) tea.Msg {
    return editorFinishedMsg{err}
}
```

If PM exits with error, TFE shows status message.

### 4. State Persistence

TFE patterns for persistence:
- Favorites: stored in JSON file
- Trash: separate directory
- Command history: in-memory + optional save

Your PM can follow same patterns.

---

## Integration Checklist

### Pre-Integration

- [ ] PM builds as standalone `pm` command
- [ ] PM accepts directory as argument or uses CWD
- [ ] PM uses Bubbletea or compatible TUI framework
- [ ] PM handles terminal cleanup properly

### Integration Steps

#### Option A: Context Menu (Easiest)

- [ ] 1. Backup context_menu.go
- [ ] 2. Add editorAvailable("pm") check in getContextMenuItems()
- [ ] 3. Add menu item for PM
- [ ] 4. Add case "launch_pm" in executeContextMenuAction()
- [ ] 5. Test: Right-click in TFE, see PM in menu
- [ ] 6. Test: Click PM menu item, launches PM

#### Option B: Toggle View (Medium)

- [ ] 1. All steps from Option A
- [ ] 2. Add viewProjectManager to types.go
- [ ] 3. Create project_manager.go with renderProjectManager()
- [ ] 4. Add case to view.go dispatcher
- [ ] 5. Add Ctrl+Shift+P handler to update_keyboard.go
- [ ] 6. Test: Toggle view mode with Ctrl+Shift+P

#### Option C: Side Panel (Complex)

- [ ] 1. All steps from Option B
- [ ] 2. Modify model.go layout calculations
- [ ] 3. Add PM styles to styles.go
- [ ] 4. Update render_file_list.go widths
- [ ] 5. Handle pane switching in update_keyboard.go
- [ ] 6. Comprehensive testing of all interactions

---

## Success Criteria

### Minimal Integration
- Right-click shows PM menu item
- PM launches in correct directory
- PM exits cleanly, returns to TFE
- File list refreshes after PM closes

### Medium Integration
- Ctrl+Shift+P toggles PM view
- PM view shows project information
- Can switch between file browser and PM
- No errors or terminal corruption

### Full Integration
- Side-by-side layout works smoothly
- Can interact with both panes
- PM state syncs with file browser (optional)
- Performance remains responsive

---

## Risk Assessment

| Change Level | Risk | Mitigation |
|-------------|------|-----------|
| Context Menu | Very Low | Only touching context_menu.go |
| View Toggle | Low | Not touching core logic, just adding mode |
| Side Panel | Medium | Changes layout code, needs testing |

**Rollback:** Easy - either revert code or disable PM feature with conditional check.

---

## Testing Strategy

1. **Unit test** - PM launches and exits without errors
2. **Integration test** - PM works from TFE context menu
3. **Regression test** - Existing TFE features still work
4. **User test** - Natural workflow is smooth

---

## Deployment

Once integration works:

1. PM needs to be available as `pm` command
   ```bash
   go install github.com/youruser/pm@latest
   # Or provide installation instructions
   ```

2. TFE will automatically detect it:
   ```go
   if editorAvailable("pm") {
       // Show in menu
   }
   ```

3. Users get seamless experience:
   ```
   Right-click in TFE → "📋 Project Manager" → PM launches
   ```

---

## Documentation Needed

Once implemented, document:

1. **For Users:**
   - How to launch PM from TFE
   - What happens when PM exits
   - How to create tasks for files

2. **For Developers:**
   - Integration approach chosen
   - Code patterns used
   - How to extend further

---

## Future Enhancements

After basic integration works, consider:

1. **File-Task Linking** - Files can reference related tasks
2. **Commit Integration** - Auto-link tasks to git commits
3. **Status Sync** - Task status updates shown in file browser
4. **Search Bridge** - Search files or tasks from unified UI
5. **Dual Editing** - Edit file while viewing related tasks

---

## Conclusion

TFE is **specifically designed for integration** with external tools. Your PM fits naturally into its architecture. The path is clear:

1. **Shortest Path:** Add to context menu (2-3 hours)
2. **Better UX:** Add as toggle view (4-6 hours)
3. **Full Integration:** Side panel (8-12 hours)

Start with #1, iterate to #2 or #3 based on user feedback.

---

## References

All documentation saved to:
- `TFE_EXPLORATION_ANALYSIS.md` - Complete deep-dive (680+ lines)
- `TFE_QUICK_REFERENCE.md` - Quick start guide
- `TFE_CODE_PATTERNS.md` - Copy-paste ready code examples

TFE source: `/home/matt/projects/TFE`

