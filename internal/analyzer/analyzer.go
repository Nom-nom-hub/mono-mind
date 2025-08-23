package analyzer

import (
	"os"
	"path/filepath"
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
	// This is where we would use Tree-sitter or language-specific parsers
	// to extract imports and dependencies
	
	// For now, we'll just identify modules by common patterns
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
	if _, exists := graph.Modules[moduleName]; !exists {
		module := Module{
			Name:         moduleName,
			Path:         filepath.Dir(path),
			Language:     language,
			Dependencies: []string{},
			LastModified: info.ModTime().String(),
		}
		graph.Modules[moduleName] = module
	}
	
	// TODO: Extract actual dependencies using AST parsing
}