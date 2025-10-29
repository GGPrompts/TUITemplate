#!/bin/bash

# new_project.sh - Create a new TUI project from TUITemplate
# Usage: ./scripts/new_project.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
print_header() {
    echo -e "${BLUE}╔════════════════════════════════════════════════════════╗${NC}"
    echo -e "${BLUE}║         TUITemplate - New Project Generator          ║${NC}"
    echo -e "${BLUE}╚════════════════════════════════════════════════════════╝${NC}"
    echo ""
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${YELLOW}→${NC} $1"
}

# Get TUITemplate directory
TEMPLATE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

print_header

# Prompt for project details
echo -e "${YELLOW}Enter project details:${NC}\n"

# App name
read -p "App name (lowercase, hyphenated, e.g., json-viewer): " APP_NAME
if [ -z "$APP_NAME" ]; then
    print_error "App name is required"
    exit 1
fi

# Validate app name format
if ! [[ "$APP_NAME" =~ ^[a-z][a-z0-9-]*$ ]]; then
    print_error "App name must be lowercase and hyphenated (e.g., my-app)"
    exit 1
fi

# Title
read -p "App title (display name, e.g., JSON Viewer): " APP_TITLE
if [ -z "$APP_TITLE" ]; then
    APP_TITLE=$(echo "$APP_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')
fi

# Description
read -p "Description: " DESCRIPTION
if [ -z "$DESCRIPTION" ]; then
    DESCRIPTION="A TUI application built with Bubbletea"
fi

# Author
read -p "Author name: " AUTHOR
if [ -z "$AUTHOR" ]; then
    AUTHOR=$(git config user.name 2>/dev/null || echo "")
fi
if [ -z "$AUTHOR" ]; then
    AUTHOR="Your Name"
fi

# License
read -p "License (default: MIT): " LICENSE
if [ -z "$LICENSE" ]; then
    LICENSE="MIT"
fi

# Layout
echo -e "\n${YELLOW}Choose layout type:${NC}"
echo "  1) single       - Single full-screen pane"
echo "  2) dual_pane    - Side-by-side panes (recommended)"
echo "  3) multi_panel  - 3+ panels"
echo "  4) tabbed       - Tabbed interface"
read -p "Layout (1-4, default: 2): " LAYOUT_CHOICE

case "$LAYOUT_CHOICE" in
    1) LAYOUT="single" ;;
    3) LAYOUT="multi_panel" ;;
    4) LAYOUT="tabbed" ;;
    *) LAYOUT="dual_pane" ;;
esac

# Components
echo -e "\n${YELLOW}Available components:${NC}"
echo "  panel, list, input, dialog, menu, status, preview, table"
read -p "Components (comma-separated, or 'all'): " COMPONENTS
if [ -z "$COMPONENTS" ]; then
    COMPONENTS="panel,list,status"
fi

# Package name (for go module)
PACKAGE_NAME=$(echo "$APP_NAME" | tr '-' '_')

# Project directory
PROJECT_DIR="$(dirname "$TEMPLATE_DIR")/$APP_NAME"

echo ""
echo -e "${BLUE}Project Configuration:${NC}"
echo "  Name:        $APP_NAME"
echo "  Title:       $APP_TITLE"
echo "  Description: $DESCRIPTION"
echo "  Author:      $AUTHOR"
echo "  License:     $LICENSE"
echo "  Layout:      $LAYOUT"
echo "  Components:  $COMPONENTS"
echo "  Location:    $PROJECT_DIR"
echo ""

read -p "Create project? (y/n): " CONFIRM
if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
    print_info "Cancelled"
    exit 0
fi

echo ""
print_info "Creating project..."

# Check if directory exists
if [ -d "$PROJECT_DIR" ]; then
    print_error "Directory $PROJECT_DIR already exists"
    exit 1
fi

# Create project directory
mkdir -p "$PROJECT_DIR"
print_success "Created directory: $PROJECT_DIR"

# Copy and process template files
print_info "Processing template files..."

for tmpl_file in "$TEMPLATE_DIR"/template/*.tmpl; do
    filename=$(basename "$tmpl_file" .tmpl)
    output_file="$PROJECT_DIR/$filename"

    # Read template file
    content=$(<"$tmpl_file")

    # Replace template variables
    content="${content//\{\{.AppName\}\}/$APP_NAME}"
    content="${content//\{\{.AppTitle\}\}/$APP_TITLE}"
    content="${content//\{\{.PackageName\}\}/$PACKAGE_NAME}"
    content="${content//\{\{.Description\}\}/$DESCRIPTION}"
    content="${content//\{\{.Author\}\}/$AUTHOR}"
    content="${content//\{\{.License\}\}/$LICENSE}"

    # Write output file
    echo "$content" > "$output_file"
    print_success "Created: $filename"
done

# Copy components
if [ "$COMPONENTS" = "all" ]; then
    print_info "Copying all components..."
    cp -r "$TEMPLATE_DIR/components" "$PROJECT_DIR/"
    print_success "Copied all components"
else
    print_info "Copying selected components..."
    mkdir -p "$PROJECT_DIR/components"

    IFS=',' read -ra COMP_ARRAY <<< "$COMPONENTS"
    for comp in "${COMP_ARRAY[@]}"; do
        comp=$(echo "$comp" | xargs) # trim whitespace
        if [ -d "$TEMPLATE_DIR/components/$comp" ]; then
            cp -r "$TEMPLATE_DIR/components/$comp" "$PROJECT_DIR/components/"
            print_success "Copied component: $comp"
        else
            print_error "Component not found: $comp"
        fi
    done
fi

# Copy lib directory
print_info "Copying utility libraries..."
cp -r "$TEMPLATE_DIR/lib" "$PROJECT_DIR/"
print_success "Copied lib/"

# Copy .claude directory (bubbletea skill)
if [ -d "$TEMPLATE_DIR/.claude" ]; then
    print_info "Copying bubbletea skill..."
    cp -r "$TEMPLATE_DIR/.claude" "$PROJECT_DIR/"
    print_success "Copied .claude/skills/bubbletea/"
fi

# Initialize Go module
print_info "Initializing Go module..."
cd "$PROJECT_DIR"

# Create go.mod from template
cat > go.mod << EOF
module github.com/$AUTHOR/$APP_NAME

go 1.24.0

require (
	github.com/charmbracelet/bubbletea v1.3.10
	github.com/charmbracelet/lipgloss v1.1.1
	github.com/charmbracelet/bubbles v0.21.0
	gopkg.in/yaml.v3 v3.0.1
)
EOF

print_success "Created go.mod"

# Run go mod tidy
if command -v go &> /dev/null; then
    go mod tidy &> /dev/null && print_success "Ran go mod tidy" || print_error "go mod tidy failed"
else
    print_error "Go not found in PATH"
fi

# Create example config
print_info "Creating example config..."
mkdir -p "$HOME/.config/$APP_NAME"

cat > "$HOME/.config/$APP_NAME/config.yaml.example" << EOF
# $APP_TITLE Configuration

theme: "dark"
keybindings: "default"

layout:
  type: "$LAYOUT"
  split_ratio: 0.5
  show_divider: true

ui:
  show_title: true
  show_status: true
  show_line_numbers: false
  mouse_enabled: true
  show_icons: true
  icon_set: "nerd_font"

performance:
  lazy_loading: true
  cache_size: 100
  async_operations: true

logging:
  enabled: false
  level: "info"
  file: "~/.local/share/$APP_NAME/debug.log"
EOF

print_success "Created example config"

# Create README.md
print_info "Creating README..."

cat > README.md << EOF
# $APP_TITLE

$DESCRIPTION

## Features

- 🚀 Fast and lightweight
- 🎨 Beautiful TUI interface
- ⚙️ Configurable (see ~/.config/$APP_NAME/config.yaml.example)
- 🖱️ Mouse and keyboard support

## Installation

\`\`\`bash
go install github.com/$AUTHOR/$APP_NAME@latest
\`\`\`

Or build from source:

\`\`\`bash
git clone https://github.com/$AUTHOR/$APP_NAME.git
cd $APP_NAME
go build
\`\`\`

## Usage

\`\`\`bash
$APP_NAME
\`\`\`

### Keyboard Shortcuts

- \`q\` - Quit
- \`?\` - Help
- \`↑/↓\` or \`k/j\` - Navigate
- \`Enter\` - Select
- \`Space\` - Toggle

## Configuration

Copy the example config:

\`\`\`bash
cp ~/.config/$APP_NAME/config.yaml.example ~/.config/$APP_NAME/config.yaml
\`\`\`

Then edit \`~/.config/$APP_NAME/config.yaml\` to customize.

## Development

This project was built with [TUITemplate](https://github.com/yourname/TUITemplate).

### Project Structure

\`\`\`
$APP_NAME/
├── main.go              # Entry point
├── types.go             # Type definitions
├── model.go             # Model initialization
├── update.go            # Message dispatcher
├── update_keyboard.go   # Keyboard handling
├── update_mouse.go      # Mouse handling
├── view.go              # View rendering
├── styles.go            # Lipgloss styles
├── config.go            # Configuration
├── components/          # UI components
└── lib/                 # Utility libraries
\`\`\`

### Building

\`\`\`bash
go build
\`\`\`

### Testing

\`\`\`bash
go test ./...
\`\`\`

## License

$LICENSE

## Author

$AUTHOR
EOF

print_success "Created README.md"

# Create .gitignore
cat > .gitignore << EOF
# Binaries
$APP_NAME
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
EOF

print_success "Created .gitignore"

# Summary
echo ""
echo -e "${GREEN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║              Project Created Successfully!             ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo ""
echo "  cd $PROJECT_DIR"
echo "  go run .                    # Run the app"
echo "  go build                    # Build binary"
echo ""
echo -e "${BLUE}Customize your app:${NC}"
echo ""
echo "  1. Edit types.go          - Define your data structures"
echo "  2. Edit model.go          - Initialize your state"
echo "  3. Edit view.go           - Implement your UI"
echo "  4. Edit update_keyboard.go - Add keyboard shortcuts"
echo "  5. Edit update_mouse.go   - Add mouse interactions"
echo ""
echo -e "${BLUE}Documentation:${NC}"
echo ""
echo "  README.md                 - Project overview"
echo "  $TEMPLATE_DIR/USAGE.md    - Detailed usage guide"
echo "  $TEMPLATE_DIR/docs/       - Research and examples"
echo ""
echo -e "${YELLOW}Happy coding! 🚀${NC}"
echo ""
