package analyzer

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"mono-mind/internal/logger"
)

// Module represents a module in the monorepo
type Module struct {
	Name          string   `json:"name"`
	Path          string   `json:"path"`
	Language      string   `json:"language"`
	Dependencies  []string `json:"dependencies"`
	LastModified  string   `json:"last_modified"`
}

// RepoGraph represents the dependency graph of the repository
type RepoGraph struct {
	Modules map[string]Module   `json:"modules"`
	Edges   map[string][]string `json:"edges"`
}

// AnalyzeRepo scans the repository and builds a dependency graph
func AnalyzeRepo(rootPath string) (*RepoGraph, error) {
	logger.Info("Starting repository analysis", "path", rootPath)
	
	// Initialize the graph
	graph := &RepoGraph{
		Modules: make(map[string]Module),
		Edges:   make(map[string][]string),
	}
	
	// Walk the directory tree
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories that should be ignored
		if shouldIgnoreDir(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		
		// Process files based on their extension
		if !info.IsDir() {
			processFile(path, info, graph)
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Build edges in the graph based on dependencies
	buildDependencyEdges(graph)
	
	logger.Info("Repository analysis completed", "modules", len(graph.Modules))
	return graph, nil
}

// shouldIgnoreDir checks if a directory should be ignored during analysis
func shouldIgnoreDir(path string) bool {
	// List of directories to ignore
	ignoredDirs := []string{".git", "node_modules", "vendor", "target", "build", "dist"}
	
	for _, dir := range ignoredDirs {
		if filepath.Base(path) == dir {
			return true
		}
	}
	
	return false
}

// processFile processes a file and extracts module information
func processFile(path string, info os.FileInfo, graph *RepoGraph) {
	// Determine language based on file extension
	ext := filepath.Ext(path)
	var language string
	
	switch ext {
	case ".go":
		language = "go"
	case ".js", ".jsx":
		language = "javascript"
	case ".ts", ".tsx":
		language = "typescript"
	case ".py":
		language = "python"
	default:
		// Not a language we're interested in
		return
	}
	
	// Extract module name from path
	moduleName := filepath.Base(filepath.Dir(path))
	
	// Check if we already have this module
	module, exists := graph.Modules[moduleName]
	if !exists {
		module = Module{
			Name:         moduleName,
			Path:         filepath.Dir(path),
			Language:     language,
			Dependencies: []string{},
			LastModified: info.ModTime().String(),
		}
	} else {
		// If module exists but this is a different file, we might need to update dependencies
		// For now, we'll just ensure the language is set correctly
		module.Language = language
	}
	
	// Parse the file to extract dependencies
	dependencies := extractDependencies(path, language)
	module.Dependencies = append(module.Dependencies, dependencies...)
	
	// Update the module in the graph
	graph.Modules[moduleName] = module
}

// extractDependencies parses a file and extracts its dependencies using regex
func extractDependencies(filePath, language string) []string {
	dependencies := []string{}
	
	// Read the file content
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("Failed to open file", "file", filePath, "error", err)
		return dependencies
	}
	defer file.Close()
	
	// Create appropriate regex patterns based on language
	var importPatterns []*regexp.Regexp
	switch language {
	case "go":
		// Go import patterns
		importPatterns = []*regexp.Regexp{
			regexp.MustCompile(`import\s+"([^"]+)"`),                    // Single import
			regexp.MustCompile(`import\s+\(([^)]+)\)`),                 // Multiple imports block
		}
	case "javascript", "typescript":
		// JavaScript/TypeScript import patterns
		importPatterns = []*regexp.Regexp{
			regexp.MustCompile(`import\s+.*from\s+['"]([^'"]+)['"]`),   // ES6 import
			regexp.MustCompile(`const.*=\s+require\(['"]([^'"]+)['"]\)`), // CommonJS require
		}
	case "python":
		// Python import patterns
		importPatterns = []*regexp.Regexp{
			regexp.MustCompile(`import\s+(\w+)`),                       // Simple import
			regexp.MustCompile(`from\s+(\w+)\s+import`),                // From import
		}
	}
	
	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Apply each pattern to the line
		for _, pattern := range importPatterns {
			matches := pattern.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				if len(match) > 1 {
					// Add the dependency (module name)
					dep := match[1]
					// Filter out standard libraries and local imports
					if !isStandardLibrary(dep, language) && !isLocalImport(dep) {
						dependencies = append(dependencies, dep)
					}
				}
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		logger.Error("Failed to read file", "file", filePath, "error", err)
	}
	
	return dependencies
}

// isStandardLibrary checks if a dependency is a standard library (simplified)
func isStandardLibrary(dep, language string) bool {
	// This is a simplified check - in a real implementation, we would have
	// comprehensive lists of standard libraries for each language
	
	switch language {
	case "go":
		// Common Go standard libraries
		stdLibs := []string{"fmt", "os", "io", "net", "http", "strings", "strconv", "time"}
		for _, lib := range stdLibs {
			if dep == lib {
				return true
			}
		}
	case "javascript":
		// Common Node.js built-ins
		builtIns := []string{"fs", "path", "http", "https", "util", "events", "stream"}
		for _, lib := range builtIns {
			if dep == lib {
				return true
			}
		}
	case "python":
		// Common Python standard libraries
		stdLibs := []string{"os", "sys", "json", "re", "datetime", "collections", "itertools"}
		for _, lib := range stdLibs {
			if dep == lib {
				return true
			}
		}
	}
	
	return false
}

// isLocalImport checks if a dependency is a local import
func isLocalImport(dep string) bool {
	// Local imports typically start with . or ..
	return strings.HasPrefix(dep, ".") || strings.HasPrefix(dep, "..")
}

// buildDependencyEdges builds the edges in the dependency graph
func buildDependencyEdges(graph *RepoGraph) {
	// For each module, create edges based on its dependencies
	for moduleName, module := range graph.Modules {
		graph.Edges[moduleName] = append(graph.Edges[moduleName], module.Dependencies...)
	}
}

// GetModuleDependencies returns the dependencies of a specific module
func (graph *RepoGraph) GetModuleDependencies(moduleName string) []string {
	if deps, exists := graph.Edges[moduleName]; exists {
		return deps
	}
	return []string{}
}

// GetDependentModules returns modules that depend on a specific module
func (graph *RepoGraph) GetDependentModules(moduleName string) []string {
	dependents := []string{}
	
	for mod, deps := range graph.Edges {
		for _, dep := range deps {
			if dep == moduleName {
				dependents = append(dependents, mod)
				break
			}
		}
	}
	
	return dependents
}

// PrintGraph prints the dependency graph to the console
func (graph *RepoGraph) PrintGraph() {
	logger.Info("Dependency Graph:")
	
	for moduleName, module := range graph.Modules {
		logger.Info("Module", "name", moduleName, "language", module.Language)
		
		dependencies := graph.GetModuleDependencies(moduleName)
		if len(dependencies) > 0 {
			logger.Info("  Dependencies:")
			for _, dep := range dependencies {
				logger.Info("    -", "dependency", dep)
			}
		}
		
		dependents := graph.GetDependentModules(moduleName)
		if len(dependents) > 0 {
			logger.Info("  Dependents:")
			for _, dep := range dependents {
				logger.Info("    -", "dependent", dep)
			}
		}
	}
}