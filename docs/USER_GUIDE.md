# MonoMind User Guide

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Quick Start](#quick-start)
4. [Repository Analysis](#repository-analysis)
5. [Impact Analysis](#impact-analysis)
6. [Building Projects](#building-projects)
7. [Testing](#testing)
8. [Refactoring](#refactoring)
9. [Release Management](#release-management)
10. [Visualization](#visualization)
11. [Configuration](#configuration)
12. [Plugins](#plugins)
13. [Best Practices](#best-practices)
14. [Troubleshooting](#troubleshooting)

## Introduction

Welcome to MonoMind, your AI-powered assistant for managing complex
codebases! This guide will help you get the most out of MonoMind's powerful
features.

MonoMind is designed to help developers:

- Understand complex repository structures
- Predict the impact of code changes
- Optimize build and test processes
- Safely refactor code
- Automate release management

## Installation

### System Requirements
- Windows, macOS, or Linux
- Go 1.19 or higher
- Git (for version control features)

### Installing MonoMind

#### Option 1: Pre-built Binary
Download the latest release from GitHub and place it in your PATH.

#### Option 2: Build from Source
```bash
# Clone the repository
git clone https://github.com/nom-nom-hub/mono-mind.git
cd mono-mind

# Build the binary
go build -o mono.exe ./cmd/mono

# (Optional) Add to PATH
```

## Quick Start

Let's walk through a typical workflow with MonoMind:

```bash
# Navigate to your project directory
cd /path/to/your/project

# Analyze your repository
mono.exe analyze

# Check what would be affected by changing a file
mono.exe impact src/api/user.js

# Build only the affected modules
mono.exe build

# Run tests for affected modules
mono.exe test

# Safely rename a function
mono.exe refactor --rename "oldFunctionName:newFunctionName" --file src/utils/helpers.js

# Create a new release
mono.exe release --bump minor --changelog
```

## Repository Analysis

### Understanding the Analyzer
The analyzer scans your repository to build a dependency graph. This graph is the foundation for all of MonoMind's features.

### Running Analysis
```bash
# Analyze the current directory
mono.exe analyze

# Analyze a specific directory
mono.exe analyze --path /path/to/project
```

### Output Formats
The analyzer supports multiple output formats:
- **Tree**: Hierarchical view of modules and dependencies
- **ASCII**: Box-style diagram
- **Horizontal**: Linear view
- **HTML**: Interactive web page

```bash
# Tree format (default)
mono.exe analyze

# ASCII format
mono.exe visualize ascii

# Horizontal format
mono.exe visualize horizontal

# HTML format
mono.exe visualize html --output graph.html
```

## Impact Analysis

### What is Impact Analysis?
Impact analysis predicts which modules and tests will be affected by changes to a specific file. This helps you:
- Focus testing efforts
- Understand the scope of changes
- Plan refactorings

### Running Impact Analysis
```bash
# Analyze impact of changing a file
mono.exe impact src/components/Button.js

# Analyze impact with verbose output
mono.exe impact src/components/Button.js --verbose
```

### Interpreting Results
The impact analysis shows:
- Directly affected modules
- Indirectly affected modules (through dependencies)
- Tests that should be run

## Building Projects

### Incremental Builds
MonoMind performs incremental builds, only rebuilding modules that have changed or depend on changed modules.

### Basic Build
```bash
# Build affected modules
mono.exe build

# Preview what would be built (dry run)
mono.exe build --dry-run
```

### Parallel Builds
For large projects, you can enable parallel builds:
```bash
# Enable parallel builds
mono.exe build --parallel

# Control concurrency
mono.exe build --parallel --max-concurrent 8
```

### Build Configuration
You can configure build behavior in your config file:
```yaml
build:
  default_command: "npm run build"
  parallel: true
  max_concurrent: 4
```

## Testing

### Intelligent Test Execution
MonoMind runs tests only for modules that are affected by changes, saving time and resources.

### Running Tests
```bash
# Run tests for affected modules
mono.exe test

# Preview test execution (dry run)
mono.exe test --dry-run

# Run tests in parallel
mono.exe test --parallel
```

### Test Configuration
Configure test behavior in your config file:
```yaml
test:
  default_command: "npm test"
  parallel: true
  max_concurrent: 4
```

## Refactoring

### Safe Refactoring
MonoMind uses AST (Abstract Syntax Tree) parsing to safely refactor code, ensuring all references are updated correctly.

### Renaming Identifiers
```bash
# Rename a function in a specific file
mono.exe refactor --rename "oldFunction:newFunction" --file src/utils.js

# Rename across the entire project
mono.exe refactor --rename "oldFunction:newFunction"

# Preview changes (dry run)
mono.exe refactor --rename "oldFunction:newFunction" --dry-run
```

### Moving Files
```bash
# Move a file or directory
mono.exe refactor --move "src/old/path:src/new/path"

# Preview the move
mono.exe refactor --move "src/old/path:src/new/path" --dry-run
```

### Supported Languages
- Go (full AST support)
- JavaScript/TypeScript (regex-based)
- Python (regex-based)
- More languages coming soon

## Release Management

### Automated Releases
MonoMind automates the release process, including version bumping, changelog generation, and publishing.

### Creating Releases
```bash
# Create a patch release
mono.exe release --bump patch

# Create a minor release with changelog
mono.exe release --bump minor --changelog

# Create a major release and publish
mono.exe release --bump major --changelog --publish
```

### Version Bumping
MonoMind follows semantic versioning:
- **Patch**: Backward-compatible bug fixes (1.0.0 → 1.0.1)
- **Minor**: Backward-compatible new features (1.0.0 → 1.1.0)
- **Major**: Breaking changes (1.0.0 → 2.0.0)

### Changelog Generation
MonoMind automatically generates changelogs from Git commit history:
```bash
# Generate changelog for next release
mono.exe release --changelog

# Generate changelog and save to file
mono.exe release --changelog --output CHANGELOG.md
```

## Visualization

### Understanding Dependencies
Visualization helps you understand your project's structure and dependencies.

### Tree View
```bash
# Show dependency tree
mono.exe visualize tree
```

### ASCII Diagram
```bash
# Show ASCII diagram
mono.exe visualize ascii
```

### Horizontal View
```bash
# Show horizontal view
mono.exe visualize horizontal
```

### HTML Export
```bash
# Export to interactive HTML
mono.exe visualize html --output dependencies.html
```

## Configuration

### Configuration File
MonoMind looks for configuration in these locations (in order):
1. `configs/config.yaml`
2. `config.yaml`
3. `.mono.yaml`

### Sample Configuration
```yaml
# Log level
log_level: info

# Analyzer settings
analyzer:
  languages:
    - go
    - javascript
    - typescript
    - python
  ignore_extensions:
    - .git
    - node_modules
    - vendor
    - target

# Build settings
build:
  default_command: "make"
  parallel: true
  max_concurrent: 4

# Test settings
test:
  default_command: "npm test"
  parallel: true
  max_concurrent: 4

# Release settings
release:
  default_bump: "patch"
  changelog_format: "markdown"
```

### Environment Variables
You can also configure MonoMind with environment variables:
```bash
# Set log level
export MONO_LOG_LEVEL=debug

# Set max concurrent builds
export MONO_BUILD_MAX_CONCURRENT=8
```

## Plugins

### Extending Functionality
Plugins allow you to extend MonoMind's functionality without modifying the core code.

### Plugin Directory
Place plugins in the `plugins/` directory in your project.

### Plugin Hooks
- `pre-build`: Runs before builds
- `post-build`: Runs after builds
- `pre-test`: Runs before tests
- `post-test`: Runs after tests
- `pre-release`: Runs before releases
- `post-release`: Runs after releases

### Example Plugin
```bash
#!/bin/bash
# plugins/pre-build.sh
echo "Pre-build tasks"
# Add your custom logic here
```

### Plugin Configuration
Configure plugins in your config file:
```yaml
plugins:
  pre-build:
    - "./scripts/prebuild.sh"
  post-build:
    - "./scripts/postbuild.sh"
```

## Best Practices

### Repository Structure
Organize your monorepo with clear module boundaries:
```
project/
├── packages/
│   ├── api/
│   ├── web/
│   ├── mobile/
│   └── shared/
├── tests/
└── docs/
```

### Dependency Management
- Minimize circular dependencies
- Use explicit imports
- Keep modules focused and small

### Version Control
- Commit frequently with descriptive messages
- Use feature branches for large changes
- Tag releases consistently

### Testing
- Write tests for new features
- Run tests regularly
- Use MonoMind's impact analysis to focus testing

## Troubleshooting

### Common Issues

#### Analyzer Not Finding Modules
- Check that your files have the correct extensions
- Verify that ignored directories are configured correctly
- Run with `--debug` flag for more information

#### Builds Failing
- Check that build commands are configured correctly
- Verify that dependencies are installed
- Run with `--dry-run` to preview builds

#### Tests Not Running
- Ensure test commands are configured
- Check that affected modules are correctly identified
- Run with `--verbose` for more information

### Getting Help
```bash
# Show help for all commands
mono.exe --help

# Show help for a specific command
mono.exe build --help

# Enable debug logging
mono.exe --debug analyze
```

### Reporting Issues
If you encounter problems:
1. Check the GitHub issues page
2. Include version information
3. Provide steps to reproduce
4. Include error messages and logs