# MonoMind API Documentation

## Table of Contents

1. [Overview](#overview)
2. [Module Structure](#module-structure)
3. [Analyzer API](#analyzer-api)
4. [Build API](#build-api)
5. [Impact API](#impact-api)
6. [Refactor API](#refactor-api)
7. [Release API](#release-api)
8. [Test API](#test-api)
9. [Visualization API](#visualization-api)
10. [Plugin API](#plugin-api)

## Overview

This document provides detailed API documentation for MonoMind's internal modules.  
These APIs can be used to extend MonoMind's functionality or integrate it into other tools.

## Module Structure

MonoMind is organized into several internal modules:

```bash
internal/
├── analyzer/      # Repository analysis and dependency graph
├── build/         # Build orchestration
├── config/        # Configuration management
├── impact/        # Impact analysis
├── logger/        # Logging utilities
├── plugins/       # Plugin system
├── refactor/      # Code refactoring
├── release/       # Release management
├── test/          # Test orchestration
└── visualization/ # Data visualization
````

## Analyzer API

### Types

#### Module

```go
type Module struct {
    Name          string   `json:"name"`
    Path          string   `json:"path"`
    Language      string   `json:"language"`
    Dependencies  []string `json:"dependencies"`
    LastModified  string   `json:"last_modified"`
}
```

#### RepoGraph

```go
type RepoGraph struct {
    Modules map[string]Module   `json:"modules"`
    Edges   map[string][]string `json:"edges"`
}
```

### Functions

#### AnalyzeRepo

```go
func AnalyzeRepo(rootPath string) (-RepoGraph, error)
```

Scans the repository and builds a dependency graph.

Parameters:

- `rootPath`: Path to the repository root

Returns:

- `-RepoGraph`: The dependency graph
- `error`: Any error that occurred

#### GetModuleDependencies

```go
func (graph -RepoGraph) GetModuleDependencies(moduleName string) []string
```

Returns the dependencies of a specific module.

Parameters:

- `moduleName`: Name of the module

Returns:

- `[]string`: List of dependencies

#### GetDependentModules

```go
func (graph -RepoGraph) GetDependentModules(moduleName string) []string
```

Returns modules that depend on a specific module.

Parameters:

- `moduleName`: Name of the module

Returns:

- `[]string`: List of dependent modules