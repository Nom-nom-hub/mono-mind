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

Welcome to MonoMind, your AI-powered assistant for managing complex codebases! This guide will help you get the most out of MonoMind's powerful features.

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
````

## Quick Start

Typical workflow with MonoMind:

```bash
cd /path/to/your/project

mono.exe analyze

mono.exe impact src/api/user.js

mono.exe build

mono.exe test

mono.exe refactor --rename "oldFunctionName:newFunctionName" --file src/utils/helpers.js

mono.exe release --bump minor --changelog
```

## Repository Analysis

### Understanding the Analyzer

The analyzer scans your repository to build a dependency graph. This graph is the foundation for all MonoMind features.

### Running Analysis

```bash
# Analyze the current directory
mono.exe analyze

# Analyze a specific directory
mono.exe analyze --path /path/to/project
```

### Output Formats

* **Tree**: Hierarchical view
* **ASCII**: Box-style diagram
* **Horizontal**: Linear view
* **HTML**: Interactive web page

```bash
mono.exe analyze          # Tree format (default)
mono.exe visualize ascii
mono.exe visualize horizontal
mono.exe visualize html --output graph.html
```

## Impact Analysis

### What is Impact Analysis?

Predicts which modules and tests will be affected by changes to a file. Helps you:

* Focus testing efforts
* Understand scope of changes
* Plan refactorings

### Running Impact Analysis

```bash
mono.exe impact src/components/Button.js
mono.exe impact src/components/Button.js --verbose
```

### Interpreting Results

* Directly affected modules
* Indirectly affected modules (via dependencies)
* Tests that should be run

## Building Projects

### Incremental Builds

Only rebuilds modules that have changed or depend on changed modules.

### Basic Build

```bash
mono.exe build
mono.exe build --dry-run
```

### Parallel Builds

```bash
mono.exe build --parallel
mono.exe build --parallel --max-concurrent 8
```

### Build Configuration

```yaml
build:
  default_command: "npm run build"
  parallel: true
  max_concurrent: 4
```

## Testing

### Intelligent Test Execution

Runs tests only for affected modules, saving time and resources.

### Running Tests

```bash
mono.exe test
mono.exe test --dry-run
mono.exe test --parallel
```

### Test Configuration

```yaml
test:
  default_command: "npm test"
  parallel: true
  max_concurrent: 4
```

## Refactoring

### Safe Refactoring

AST parsing ensures all references are updated correctly.

### Renaming Identifiers

```bash
mono.exe refactor --rename "oldFunction:newFunction" --file src/utils.js
mono.exe refactor --rename "oldFunction:newFunction"
mono.exe refactor --rename "oldFunction:newFunction" --dry-run
```

### Moving Files

```bash
mono.exe refactor --move "src/old/path:src/new/path"
mono.exe refactor --move "src/old/path:src/new/path" --dry-run
```

### Supported Languages

* Go (full AST)
* JavaScript/TypeScript (regex-based)
* Python (regex-based)
* More coming soon

## Release Management

### Automated Releases

Handles version bumping, changelog generation, and publishing.

### Creating Releases

```bash
mono.exe release --bump patch
mono.exe release --bump minor --changelog
mono.exe release --bump major --changelog --publish
```

### Version Bumping

* **Patch**: Bug fixes
* **Minor**: Backward-compatible new features
* **Major**: Breaking changes

### Changelog Generation

```bash
mono.exe release --changelog
mono.exe release --changelog --output CHANGELOG.md
```

## Visualization

### Tree View

```bash
mono.exe visualize tree
```

### ASCII Diagram

```bash
mono.exe visualize ascii
```

### Horizontal View

```bash
mono.exe visualize horizontal
```

### HTML Export

```bash
mono.exe visualize html --output dependencies.html
```

## Configuration

### Configuration File Locations

1. `configs/config.yaml`
2. `config.yaml`
3. `.mono.yaml`

### Sample Configuration

```yaml
log_level: info

analyzer:
  languages: [go, javascript, typescript, python]
  ignore_extensions: [.git, node_modules, vendor, target]

build:
  default_command: "make"
  parallel: true
  max_concurrent: 4

test:
  default_command: "npm test"
  parallel: true
  max_concurrent: 4

release:
  default_bump: "patch"
  changelog_format: "markdown"
```

### Environment Variables

```bash
export MONO_LOG_LEVEL=debug
export MONO_BUILD_MAX_CONCURRENT=8
```

## Plugins

### Extending Functionality

Plugins add features without modifying core code.

### Plugin Directory

`plugins/` in your project.

### Plugin Hooks

* `pre-build`, `post-build`
* `pre-test`, `post-test`
* `pre-release`, `post-release`

### Example Plugin

```bash
#!/bin/bash
echo "Pre-build tasks"
```

### Plugin Configuration

```yaml
plugins:
  pre-build: ["./scripts/prebuild.sh"]
  post-build: ["./scripts/postbuild.sh"]
```

## Best Practices

### Repository Structure

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

* Minimize circular dependencies
* Use explicit imports
* Keep modules focused

### Version Control

* Commit frequently
* Use feature branches
* Tag releases consistently

### Testing

* Write tests for new features
* Run tests regularly
* Use impact analysis

## Troubleshooting

### Common Issues

#### Analyzer Not Finding Modules

* Check file extensions
* Verify ignored directories
* Use `--debug`

#### Builds Failing

* Check build commands
* Verify dependencies
* Use `--dry-run`

#### Tests Not Running

* Ensure test commands are correct
* Verify affected modules
* Use `--verbose`

### Getting Help

```bash
mono.exe --help
mono.exe build --help
mono.exe --debug analyze
```

### Reporting Issues

1. Check GitHub issues
2. Include version info
3. Provide steps to reproduce
4. Include error logs