package test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"mono-mind/internal/analyzer"
	"mono-mind/internal/logger"
)

// TestConfig holds configuration for the test process
type TestConfig struct {
	Parallel      bool `json:"parallel"`
	MaxConcurrent int  `json:"max_concurrent"`
	DryRun        bool `json:"dry_run"`
}

// TestResult holds the result of a test operation
type TestResult struct {
	TestsRun     int      `json:"tests_run"`
	TestsPassed  int      `json:"tests_passed"`
	TestsFailed  int      `json:"tests_failed"`
	Errors       []string `json:"errors"`
	Duration     string   `json:"duration"`
}

// RunTests runs tests for affected modules
func RunTests(graph *analyzer.RepoGraph, config TestConfig) *TestResult {
	logger.Info("Running tests")
	
	result := &TestResult{
		TestsRun:    0,
		TestsPassed: 0,
		TestsFailed: 0,
		Errors:      []string{},
	}
	
	// In a real implementation, we would:
	// 1. Determine which tests to run based on changes
	// 2. Execute tests in the correct order based on dependencies
	// 3. Handle parallel execution if enabled
	// 4. Collect results and errors
	
	// For demonstration, we'll run tests for all modules
	for moduleName := range graph.Modules {
		if config.DryRun {
			logger.Info("Would run tests for module (dry-run)", "module", moduleName)
			result.TestsRun++
			result.TestsPassed++
		} else {
			logger.Info("Running tests for module", "module", moduleName)
			
			// Execute test command for the module
			// In a real implementation, this would be language-specific
			passed, err := runTestsForModule(moduleName, graph.Modules[moduleName])
			if err != nil {
				logger.Error("Failed to run tests for module", "module", moduleName, "error", err)
				result.Errors = append(result.Errors, moduleName+": "+err.Error())
				result.TestsFailed++
			} else if passed {
				result.TestsPassed++
			} else {
				result.TestsFailed++
			}
			result.TestsRun++
		}
	}
	
	logger.Info("Test execution completed", 
		"tests_run", result.TestsRun,
		"tests_passed", result.TestsPassed,
		"tests_failed", result.TestsFailed)
	
	return result
}

// runTestsForModule executes the test command for a specific module
func runTestsForModule(moduleName string, module analyzer.Module) (bool, error) {
	// This is a simplified implementation
	// In a real system, we would determine the appropriate test command
	// based on the language and project configuration
	
	// Validate the module path to prevent directory traversal attacks
	// Clean the path to remove any .. or . components
	cleanPath := filepath.Clean(module.Path)
	
	// Ensure the path is relative and doesn't start with ..
	if filepath.IsAbs(cleanPath) || strings.HasPrefix(cleanPath, "..") {
		return false, fmt.Errorf("invalid module path: %s", module.Path)
	}
	
	var cmd *exec.Cmd
	
	switch module.Language {
	case "go":
		// For Go modules, we might run 'go test'
		cmd = exec.Command("go", "test", "./"+cleanPath+"/...")
	case "javascript", "typescript":
		// For JS/TS projects, we might run 'npm test' or 'yarn test'
		// Check if package.json exists in the module directory
		cmd = exec.Command("npm", "test")
		cmd.Dir = cleanPath
	case "python":
		// For Python projects, we might run 'python -m pytest'
		cmd = exec.Command("python", "-m", "pytest")
		cmd.Dir = cleanPath
	default:
		// Default test command
		cmd = exec.Command("make", "test")
		cmd.Dir = module.Path
	}
	
	// Execute the command
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Tests failed", "module", moduleName, "error", err)
		return false, err
	}
	
	logger.Debug("Tests passed", "module", moduleName)
	return true, nil
}