package build

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"mono-mind/internal/analyzer"
	"mono-mind/internal/logger"
	"mono-mind/internal/plugins"
)

// BuildConfig holds configuration for the build process
type BuildConfig struct {
	Parallel      bool `json:"parallel"`
	MaxConcurrent int  `json:"max_concurrent"`
	DryRun        bool `json:"dry_run"`
}

// BuildResult holds the result of a build operation
type BuildResult struct {
	ModulesBuilt  []string `json:"modules_built"`
	ModulesSkipped []string `json:"modules_skipped"`
	Errors        []string `json:"errors"`
	Duration      string   `json:"duration"`
}

// IncrementalBuild performs an incremental build of affected modules
func IncrementalBuild(graph *analyzer.RepoGraph, config BuildConfig) *BuildResult {
	logger.Info("Starting incremental build")
	
	// Initialize plugin manager
	pluginManager := plugins.NewPluginManager()
	if err := pluginManager.LoadPluginsFromDir("plugins"); err != nil {
		logger.Error("Failed to load plugins", "error", err)
	}
	
	// Execute pre-build plugins
	if err := pluginManager.ExecuteHook("pre-build"); err != nil {
		logger.Error("Failed to execute pre-build hook", "error", err)
	}
	
	result := &BuildResult{
		ModulesBuilt:  []string{},
		ModulesSkipped: []string{},
		Errors:        []string{},
	}
	
	// In a real implementation, we would:
	// 1. Determine which modules need to be rebuilt based on changes
	// 2. Execute builds in the correct order based on dependencies
	// 3. Handle parallel execution if enabled
	// 4. Collect results and errors
	
	// For demonstration, we'll build all modules
	for moduleName := range graph.Modules {
		if config.DryRun {
			logger.Info("Would build module (dry-run)", "module", moduleName)
			result.ModulesBuilt = append(result.ModulesBuilt, moduleName)
		} else {
			logger.Info("Building module", "module", moduleName)
			
			// Execute build command for the module
			// In a real implementation, this would be language-specific
			err := buildModule(moduleName, graph.Modules[moduleName])
			if err != nil {
				logger.Error("Failed to build module", "module", moduleName, "error", err)
				result.Errors = append(result.Errors, moduleName+": "+err.Error())
			} else {
				result.ModulesBuilt = append(result.ModulesBuilt, moduleName)
			}
		}
	}
	
	// Execute post-build plugins
	if err := pluginManager.ExecuteHook("post-build"); err != nil {
		logger.Error("Failed to execute post-build hook", "error", err)
	}
	
	logger.Info("Incremental build completed", 
		"modules_built", len(result.ModulesBuilt),
		"modules_skipped", len(result.ModulesSkipped))
	
	return result
}

// buildModule executes the build command for a specific module
func buildModule(moduleName string, module analyzer.Module) error {
	// This is a simplified implementation
	// In a real system, we would determine the appropriate build command
	// based on the language and project configuration
	
	var cmd *exec.Cmd
	
	// Validate the module path to prevent directory traversal attacks
	// Clean the path to remove any .. or . components
	cleanPath := filepath.Clean(module.Path)
	
	// Ensure the path is relative and doesn't start with ..
	if filepath.IsAbs(cleanPath) || strings.HasPrefix(cleanPath, "..") {
		return fmt.Errorf("invalid module path: %s", module.Path)
	}
	
	switch module.Language {
	case "go":
		// For Go modules, we might run 'go build'
		cmd = exec.Command("go", "build", "./"+cleanPath)
	case "javascript", "typescript":
		// For JS/TS projects, we might run 'npm run build' or 'yarn build'
		// Check if package.json exists in the module directory
		cmd = exec.Command("npm", "run", "build")
		cmd.Dir = cleanPath
	case "python":
		// For Python projects, we might run a build script
		cmd = exec.Command("python", "setup.py", "build")
		cmd.Dir = cleanPath
	
	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Build failed", "module", moduleName, "output", string(output), "error", err)
		return err
	}
	
	logger.Debug("Build successful", "module", moduleName, "output", string(output))
	return nil
}