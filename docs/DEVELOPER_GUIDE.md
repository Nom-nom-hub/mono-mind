# MonoMind Developer Guide

## Table of Contents

1. [Introduction](#introduction)
2. [Project Structure](#project-structure)
3. [Architecture](#architecture)
4. [Development Setup](#development-setup)
5. [Code Style](#code-style)
6. [Testing](#testing)
7. [Adding New Features](#adding-new-features)
8. [Working with Modules](#working-with-modules)
9. [Plugin Development](#plugin-development)
10. [Documentation](#documentation)
11. [Release Process](#release-process)

## Introduction

This guide is for developers who want to contribute to MonoMind or extend its
functionality. It covers the project structure, architecture, and development
practices.

## Project Structure

```bash
mono-mind/
├── cmd/
│   └── mono/              # CLI entry point
│       ├── main.go
│       └── root.go
├── configs/               # Default configuration files
│   └── config.yaml
├── docs/                  # Documentation
├── example/               # Example projects
├── internal/              # Internal packages
│   ├── analyzer/          # Repository analysis
│   ├── build/             # Build orchestration
│   ├── config/            # Configuration management
│   ├── impact/            # Impact analysis
│   ├── logger/            # Logging utilities
│   ├── plugins/           # Plugin system
│   ├── refactor/          # Code refactoring
│   ├── release/           # Release management
│   ├── test/              # Test orchestration
│   └── visualization/     # Data visualization
├── plugins/               # Example plugins
├── scripts/               # Utility scripts
└── go.mod                 # Go module definition
```

## Architecture

### Core Principles

1. **Modularity**: Each feature is implemented as a separate module
2. **CLI-First**: All functionality is accessible through the command line
3. **Extensibility**: Plugins can extend functionality without core changes
4. **Language Agnostic**: Supports multiple programming languages
5. **Performance**: Optimized for large repositories

### Data Flow

1. **Analyzer** scans the repository and builds a dependency graph
2. **Impact Engine** uses the graph to predict change effects
3. **Build/Test Orchestrator** executes targeted operations
4. **Refactor Engine** safely modifies code
5. **Release Manager** handles versioning and publishing
6. **Visualization** presents data in multiple formats
7. **Plugin System** extends functionality through hooks

### Key Components

- **CLI Layer**: Cobra-based command interface
- **Core Modules**: Independent packages for each feature
- **Configuration**: YAML-based configuration system
- **Logging**: Structured logging with Logrus
- **Plugin System**: Hook-based extensibility

## Development Setup

### Prerequisites

- Go 1.19 or higher
- Git
- Node.js (for JavaScript examples)
- Python (for Python examples)

### Getting Started

```bash
# Fork and clone the repository
git clone https://github.com/nom-nom-hub/mono-mind.git
cd mono-mind

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build
go build -o mono.exe ./cmd/mono

# Run
./mono.exe --help
```

### Development Workflow

1. Create a feature branch
2. Make changes
3. Add tests
4. Update documentation
5. Run tests
6. Commit changes
7. Push and create pull request

## Code Style

### Go Coding Standards
Follow the official Go coding standards:
- Use `gofmt` for formatting
- Write clear, concise comments
- Use meaningful variable names
- Keep functions focused and small

### Naming Conventions
- **Packages**: lowercase, single word (e.g., `analyzer`)
- **Functions**: CamelCase (e.g., `AnalyzeRepo`)
- **Variables**: camelCase (e.g., `repoGraph`)
- **Constants**: UPPER_SNAKE_CASE (e.g., `MAX_CONCURRENT`)

### Error Handling
- Always handle errors explicitly
- Use descriptive error messages
- Wrap errors with context when appropriate
- Return errors early when possible

### Logging
- Use the internal logger package
- Include relevant context in log messages
- Use appropriate log levels (debug, info, warn, error)
- Avoid logging sensitive information

## Testing

### Test Structure

Tests are organized by module:

```bash
internal/
├── analyzer/
│   ├── analyzer.go
│   └── analyzer_test.go
└── build/
    ├── build.go
    └── build_test.go
```

### Writing Tests

```go
// Example test
func TestAnalyzeRepo(t *testing.T) {
    // Setup
    tempDir := createTempRepo()
    defer os.RemoveAll(tempDir)

    // Execute
    graph, err := analyzer.AnalyzeRepo(tempDir)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, graph)
    assert.Greater(t, len(graph.Modules), 0)
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific module tests
go test ./internal/analyzer
```

### Test Best Practices

- Write tests for new functionality
- Use table-driven tests for multiple cases
- Mock external dependencies
- Test edge cases and error conditions
- Keep tests fast and focused

## Adding New Features

### Feature Development Process
1. **Design**: Plan the feature and API
2. **Implement**: Write the code
3. **Test**: Add comprehensive tests
4. **Document**: Update documentation
5. **Review**: Get code review

### Creating a New Module
1. Create a new directory in `internal/`
2. Implement the core functionality
3. Add a public API
4. Write tests
5. Add CLI command integration

### Example Module Structure
```go
// internal/example/example.go
package example

import (
    "mono-mind/internal/logger"
)

// Config holds configuration for the example module
type Config struct {
    Enabled bool `json:"enabled"`
}

// Result holds the result of example operations
type Result struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}

// Process performs the example operation
func Process(config Config) *Result {
    logger.Info("Processing example", "enabled", config.Enabled)
    
    result := &Result{
        Success: true,
        Message: "Example processed successfully",
    }
    
    return result
}
```

## Working with Modules

### Module Interface
Each module should follow a consistent pattern:
- Configuration struct
- Result struct
- Main processing function
- Helper functions as needed

### Dependency Management
Modules should:
- Have minimal dependencies
- Use interfaces for external dependencies
- Avoid circular dependencies
- Handle errors gracefully

### Concurrency
For performance-critical modules:
- Use goroutines for parallel processing
- Limit concurrent operations
- Handle context cancellation
- Use sync primitives when needed

## Plugin Development

### Plugin System Overview
Plugins are external programs executed at specific hooks:
- Pre/Post build
- Pre/Post test
- Pre/Post release

### Creating Plugins
Plugins can be written in any language:
```bash
#!/bin/bash
# plugins/pre-build.sh
echo "Pre-build plugin"
# Your plugin logic here
```

```python
#!/usr/bin/env python3
# plugins/post-test.py
print("Post-test plugin")
# Your plugin logic here
```

### Plugin Best Practices
- Handle errors gracefully
- Use exit codes appropriately
- Log important information
- Be fast and efficient
- Document plugin behavior

## Documentation

### Documentation Structure
- **README.md**: Project overview
- **docs/README.md**: User documentation
- **docs/API.md**: API reference
- **docs/USER_GUIDE.md**: User guide
- **docs/DEVELOPER_GUIDE.md**: Developer guide (this document)

### Writing Documentation
- Use clear, simple language
- Include examples
- Keep it up to date
- Organize logically
- Use consistent formatting

### Code Comments
- Comment exported functions and types
- Explain complex logic
- Document assumptions
- Keep comments accurate

## Release Process

### Versioning
MonoMind follows semantic versioning:
- **Major**: Breaking changes
- **Minor**: New features
- **Patch**: Bug fixes

### Release Steps
1. Update version in code
2. Generate changelog
3. Create Git tag
4. Build binaries
5. Create GitHub release
6. Update documentation

### Pre-release Checklist
- [ ] All tests pass
- [ ] Documentation is up to date
- [ ] Code is reviewed
- [ ] Version is updated
- [ ] Changelog is generated
- [ ] Binaries build successfully

### Post-release
- [ ] Verify release on GitHub
- [ ] Update version in development
- [ ] Announce release
- [ ] Monitor for issues