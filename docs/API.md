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

This document provides detailed API documentation for MonoMind's internal
modules. These APIs can be used to extend MonoMind's functionality or
integrate it into other tools.

## Module Structure

MonoMind is organized into several internal modules:

```bash
internal/
├── analyzer/     # Repository analysis and dependency graph
├── build/        # Build orchestration
├── config/       # Configuration management
├── impact/       # Impact analysis
├── logger/       # Logging utilities
├── plugins/      # Plugin system
├── refactor/     # Code refactoring
├── release/      # Release management
├── test/         # Test orchestration
└── visualization/ # Data visualization
```

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
func AnalyzeRepo(rootPath string) (*RepoGraph, error)
```

Scans the repository and builds a dependency graph.

Parameters:

- `rootPath`: Path to the repository root

Returns:

- `*RepoGraph`: The dependency graph
- `error`: Any error that occurred

#### GetModuleDependencies

```go
func (graph *RepoGraph) GetModuleDependencies(moduleName string) []string
```

Returns the dependencies of a specific module.

Parameters:

- `moduleName`: Name of the module

Returns:

- `[]string`: List of dependencies

#### GetDependentModules

```go
func (graph *RepoGraph) GetDependentModules(moduleName string) []string
```

Returns modules that depend on a specific module.

Parameters:

- `moduleName`: Name of the module

Returns:

- `[]string`: List of dependent modules

## Build API

### Types

#### BuildConfig

```go
type BuildConfig struct {
    Parallel      bool `json:"parallel"`
    MaxConcurrent int  `json:"max_concurrent"`
    DryRun        bool `json:"dry_run"`
}
```

#### BuildResult

```go
type BuildResult struct {
    ModulesBuilt   []string `json:"modules_built"`
    ModulesSkipped []string `json:"modules_skipped"`
    Errors         []string `json:"errors"`
    Duration       string   `json:"duration"`
}
```

### Functions

#### IncrementalBuild

```go
func IncrementalBuild(graph *analyzer.RepoGraph, config BuildConfig) *BuildResult
```

Performs an incremental build of affected modules.

Parameters:

- `graph`: Dependency graph from analyzer
- `config`: Build configuration

Returns:

- `*BuildResult`: Result of the build operation

## Impact API

### Types

#### ImpactResult

```go
type ImpactResult struct {
    ChangedFile      string   `json:"changed_file"`
    AffectedModules  []string `json:"affected_modules"`
    AffectedTests    []string `json:"affected_tests"`
    Conflicts        []string `json:"conflicts"`
}
```

### Functions

#### AnalyzeImpact

```go
func AnalyzeImpact(graph *analyzer.RepoGraph, changedFile string) *ImpactResult
```

Analyzes the impact of a file change on the repository.

Parameters:

- `graph`: Dependency graph from analyzer
- `changedFile`: Path to the changed file

Returns:

- `*ImpactResult`: Result of the impact analysis

## Refactor API

### Types

#### RefactorConfig
```go
type RefactorConfig struct {
    DryRun   bool   `json:"dry_run"`
    OldName  string `json:"old_name"`
    NewName  string `json:"new_name"`
    FilePath string `json:"file_path"`
}
```

#### RefactorResult
```go
type RefactorResult struct {
    FilesChanged []string `json:"files_changed"`
    Errors       []string `json:"errors"`
    Duration     string   `json:"duration"`
}
```

### Functions

#### Rename
```go
func Rename(config RefactorConfig) *RefactorResult
```
Performs a safe rename operation across the repository.

Parameters:
- `config`: Refactor configuration

Returns:
- `*RefactorResult`: Result of the refactor operation

#### Move
```go
func Move(oldPath, newPath string, dryRun bool) *RefactorResult
```
Moves a file or directory to a new location.

Parameters:
- `oldPath`: Current path of the file/directory
- `newPath`: New path for the file/directory
- `dryRun`: Whether to perform a dry run

Returns:
- `*RefactorResult`: Result of the move operation

## Release API

### Types

#### ReleaseConfig
```go
type ReleaseConfig struct {
    VersionBump string `json:"version_bump"` // major, minor, patch
    Changelog   bool   `json:"changelog"`
    Publish     bool   `json:"publish"`
}
```

#### ReleaseResult
```go
type ReleaseResult struct {
    NewVersion   string   `json:"new_version"`
    Changelog    string   `json:"changelog"`
    Errors       []string `json:"errors"`
    Duration     string   `json:"duration"`
}
```

### Functions

#### ManageRelease
```go
func ManageRelease(config ReleaseConfig) *ReleaseResult
```
Handles version bumping, changelog generation, and publishing.

Parameters:
- `config`: Release configuration

Returns:
- `*ReleaseResult`: Result of the release operation

## Test API

### Types

#### TestConfig
```go
type TestConfig struct {
    Parallel      bool `json:"parallel"`
    MaxConcurrent int  `json:"max_concurrent"`
    DryRun        bool `json:"dry_run"`
}
```

#### TestResult
```go
type TestResult struct {
    TestsRun     int      `json:"tests_run"`
    TestsPassed  int      `json:"tests_passed"`
    TestsFailed  int      `json:"tests_failed"`
    Errors       []string `json:"errors"`
    Duration     string   `json:"duration"`
}
```

### Functions

#### RunTests
```go
func RunTests(graph *analyzer.RepoGraph, config TestConfig) *TestResult
```
Runs tests for affected modules.

Parameters:
- `graph`: Dependency graph from analyzer
- `config`: Test configuration

Returns:
- `*TestResult`: Result of the test execution

## Visualization API

### Functions

#### PrintDependencyGraph
```go
func PrintDependencyGraph(graph *analyzer.RepoGraph)
```
Prints the dependency graph in a visual format.

Parameters:
- `graph`: Dependency graph from analyzer

#### PrintASCIIDependencyGraph
```go
func PrintASCIIDependencyGraph(graph *analyzer.RepoGraph)
```
Prints the dependency graph as an ASCII diagram.

Parameters:
- `graph`: Dependency graph from analyzer

#### PrintHorizontalDependencyGraph
```go
func PrintHorizontalDependencyGraph(graph *analyzer.RepoGraph)
```
Prints the dependency graph horizontally.

Parameters:
- `graph`: Dependency graph from analyzer

#### PrintHTMLDependencyGraph
```go
func PrintHTMLDependencyGraph(graph *analyzer.RepoGraph, filename string) error
```
Prints the dependency graph as an HTML file.

Parameters:
- `graph`: Dependency graph from analyzer
- `filename`: Output file name

Returns:
- `error`: Any error that occurred

## Plugin API

### Types

#### PluginManager
```go
type PluginManager struct {
    Plugins map[string][]string // hook name -> list of plugin paths
}
```

### Functions

#### NewPluginManager
```go
func NewPluginManager() *PluginManager
```
Creates a new plugin manager.

Returns:
- `*PluginManager`: New plugin manager instance

#### RegisterPlugin
```go
func (pm *PluginManager) RegisterPlugin(hook, pluginPath string)
```
Registers a plugin for a specific hook.

Parameters:
- `hook`: Hook name
- `pluginPath`: Path to the plugin

#### ExecuteHook
```go
func (pm *PluginManager) ExecuteHook(hook string) error
```
Executes all plugins registered for a hook.

Parameters:
- `hook`: Hook name

Returns:
- `error`: Any error that occurred

#### LoadPluginsFromDir
```go
func (pm *PluginManager) LoadPluginsFromDir(dirPath string) error
```
Loads all plugins from a directory.

Parameters:
- `dirPath`: Path to the plugins directory

Returns:
- `error`: Any error that occurred