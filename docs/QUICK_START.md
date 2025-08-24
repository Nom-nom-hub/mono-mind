# MonoMind Quick Start Guide

## What is MonoMind?

MonoMind is an AI-powered development assistant that helps you manage complex codebases. It automates repository analysis, impact prediction, build orchestration, testing, refactoring, and release management.

## Installation

### Prerequisites

- Go 1.19 or higher  
- Git (for version control features)  

### Install MonoMind

```bash
# Clone the repository
git clone https://github.com/nom-nom-hub/mono-mind.git
cd mono-mind

# Build the binary
go build -o mono.exe ./cmd/mono

# (Optional) Add to PATH
````

## First Steps

### 1. Navigate to Your Project

```bash
cd /path/to/your/project
```

### 2. Analyze Your Repository

```bash
mono.exe analyze
```

Output:

```
Dependency Graph:
=================
📁 api (go)
  └─ depends on: shared
  └─ used by: web

📁 web (javascript)
  └─ depends on: api, shared

📁 shared (go)
  └─ used by: api, web
```

### 3. Check Impact of Changes

```bash
mono.exe impact src/api/user.js
```

Output:

```
Impact Analysis for: src/api/user.js
======================
Changed Module: api

Affected Modules:
  📁 api (directly changed)
  📁 web (depends on api)
```

### 4. Build Affected Modules

```bash
mono.exe build
```

### 5. Run Tests for Affected Modules

```bash
mono.exe test
```

### 6. Refactor Code Safely

```bash
mono.exe refactor --rename "oldFunction:newFunction" --file src/utils.js
```

### 7. Create a Release

```bash
mono.exe release --bump minor --changelog
```

## Key Commands

| Command         | Description                                |
| --------------- | ------------------------------------------ |
| `analyze`       | Scan repository and build dependency graph |
| `impact <file>` | Show affected modules for a file change    |
| `build`         | Incremental builds of affected modules     |
| `test`          | Run tests for affected modules             |
| `refactor`      | Safely rename or move code                 |
| `release`       | Manage version bumps and releases          |
| `visualize`     | Visualize repository structure             |

## Configuration

Create `configs/config.yaml`:

```yaml
log_level: info
analyzer:
  languages:
    - go
    - javascript
    - typescript
    - python
build:
  parallel: true
  max_concurrent: 4
```

## Plugins

Create `plugins/pre-build.sh`:

```bash
#!/bin/bash
echo "Pre-build checks..."
# Add your custom logic here
```

Make it executable:

```bash
chmod +x plugins/pre-build.sh
```

## Next Steps

1. --Read the Full Documentation--: Check `docs/` directory
2. --Explore Advanced Features--: Try HTML visualization, AST refactoring
3. --Customize Configuration--: Tailor MonoMind to your workflow
4. --Write Plugins--: Extend functionality with custom scripts
5. --Contribute--: Help improve MonoMind on GitHub

## Getting Help

```bash
# Show all commands
mono.exe --help

# Show help for a specific command
mono.exe build --help

# Enable debug output
mono.exe --debug analyze
```

Start using MonoMind today to simplify your development workflow!