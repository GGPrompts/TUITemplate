# TUI Research Documentation

This directory contains comprehensive research on TUI (Terminal User Interface) development, tools, and libraries.

## Contents

### Ecosystem Research

**ECOSYSTEM_RESEARCH_2025.md** (989 lines)
- Comprehensive guide to Charm/Bubbletea ecosystem libraries
- Detailed descriptions of 20+ libraries with use cases
- Code examples and integration patterns
- Implementation roadmap for TUI applications
- Performance considerations and best practices

**ECOSYSTEM_QUICK_REFERENCE.md** (96 lines)
- Quick lookup guide for Charm libraries
- Installation commands
- Priority rankings (HIGH/MEDIUM/LOW)
- Library comparison table

### TUI Applications Research

**TUI_APPLICATIONS_RESEARCH.md** (1,249 lines)
- Research on 80+ TUI applications across 8 categories:
  - Image & Drawing Tools
  - Specialized Editors (hex, CSV, JSON, markdown)
  - Media Tools (audio, video, PDF)
  - Development Tools (databases, Git, API testing, Docker)
  - Document & Productivity Tools
  - System & Network Tools
  - File Managers (reference implementations)
- Each tool includes: description, GitHub link, maintenance status, platform support
- Integration strategies and file type detection

**TUI_TOOLS_QUICK_REFERENCE.md** (273 lines)
- Quick reference tables for all 80+ tools
- Installation commands by language/package manager
- 4-phase integration rollout plan
- Tool detection logic examples
- Priority recommendations

## How to Use This Research

### When Building a New TUI App

1. **Start with ECOSYSTEM_QUICK_REFERENCE.md**
   - Identify which Charm libraries you need
   - Install dependencies quickly

2. **Refer to ECOSYSTEM_RESEARCH_2025.md**
   - Read detailed descriptions of each library
   - Review code examples
   - Follow integration patterns

3. **For File Viewer/Editor Apps:**
   - Check TUI_TOOLS_QUICK_REFERENCE.md for existing tools
   - Consider integrating existing tools vs building from scratch
   - Use TUI_APPLICATIONS_RESEARCH.md for detailed tool information

### Research Highlights

#### Top Priority Libraries (from Ecosystem Research)

1. **bubble-table** - Interactive tables with sorting, filtering
2. **go-fzf** - Fuzzy finder for quick search
3. **bubblezone** - Mouse event tracking for clickable UI
4. **Huh** - Official Charm forms library

#### Top Priority TUI Tools (for Integration)

1. **timg** - Universal viewer (images, videos, PDFs)
2. **VisiData** - Data powerhouse (CSV, JSON, Excel, SQLite)
3. **fx** - JSON/YAML/TOML viewer
4. **glow** - Beautiful markdown renderer
5. **harlequin** - SQL IDE in terminal

#### Common Integration Patterns

**Tool Detection Pattern:**
```go
func getAvailableTool(tools []string) string {
    for _, tool := range tools {
        if toolAvailable(tool) {
            return tool
        }
    }
    return ""
}

func toolAvailable(name string) bool {
    _, err := exec.LookPath(name)
    return err == nil
}
```

**File Type Detection Pattern:**
```go
func getToolForFile(path string) string {
    ext := strings.ToLower(filepath.Ext(path))

    switch ext {
    case ".json", ".yaml", ".toml":
        return getAvailableTool([]string{"fx", "jq"})
    case ".csv":
        return getAvailableTool([]string{"visidata", "csv-tui"})
    case ".png", ".jpg", ".gif":
        return getAvailableTool([]string{"timg", "viu"})
    case ".md":
        return getAvailableTool([]string{"glow", "md-tui"})
    default:
        return ""
    }
}
```

## Research Methodology

All research was conducted in October 2025 using:
- Web searches for latest TUI tools and libraries
- GitHub repository analysis (stars, maintenance status, recent commits)
- Documentation review for each library/tool
- Cross-platform compatibility verification
- Active maintenance verification (2024-2025 commits)

## Updates

This research is current as of **October 2025**. For updated information:
- Check library GitHub repositories for latest versions
- Verify tool maintenance status before integration
- Review Charm's official documentation for new releases

## Contributing

If you discover new TUI libraries or tools that should be added:
1. Verify they're actively maintained (commits in last 6 months)
2. Test cross-platform compatibility
3. Document use cases and examples
4. Submit updates to this research

---

**Last Updated:** October 16, 2025
