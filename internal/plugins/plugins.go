package plugins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"mono-mind/internal/logger"
)

// PluginManager handles loading and executing plugins
type PluginManager struct {
	Plugins map[string][]string // hook name -> list of plugin paths
}

// NewPluginManager creates a new plugin manager
func NewPluginManager() *PluginManager {
	return &PluginManager{
		Plugins: make(map[string][]string),
	}
}

// validatePluginPath validates that a plugin path is safe to execute
func validatePluginPath(pluginPath, pluginsDir string) (string, error) {
	// Clean the path to remove any .. or . components
	cleanPath := filepath.Clean(pluginPath)

	// Get absolute path
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Ensure the path doesn't contain dangerous patterns
	if strings.Contains(absPath, "..") {
		return "", fmt.Errorf("plugin path contains directory traversal: %s", pluginPath)
	}

	// If a plugins directory is specified, ensure the plugin is within it
	if pluginsDir != "" {
		pluginsAbsDir, err := filepath.Abs(pluginsDir)
		if err != nil {
			return "", fmt.Errorf("failed to get plugins directory absolute path: %v", err)
		}

		// Check if the plugin path is within the plugins directory
		relPath, err := filepath.Rel(pluginsAbsDir, absPath)
		if err != nil || strings.HasPrefix(relPath, "..") {
			return "", fmt.Errorf("plugin path is outside allowed directory: %s", pluginPath)
		}
	}

	// Validate file extension
	ext := filepath.Ext(absPath)
	allowedExts := []string{".sh", ".py", ".js", ".bash", ".zsh"}
	isAllowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isAllowed = true
			break
		}
	}

	// For files without extensions, check if they're executable (but this is risky)
	if ext == "" {
		// Only allow if explicitly configured - for now, reject
		return "", fmt.Errorf("plugin must have a valid extension: %s", pluginPath)
	}

	if !isAllowed {
		return "", fmt.Errorf("plugin has disallowed extension: %s", ext)
	}

	return absPath, nil
}

// RegisterPlugin registers a plugin for a specific hook
func (pm *PluginManager) RegisterPlugin(hook, pluginPath string) {
	if _, exists := pm.Plugins[hook]; !exists {
		pm.Plugins[hook] = []string{}
	}
	pm.Plugins[hook] = append(pm.Plugins[hook], pluginPath)
	logger.Debug("Registered plugin", "hook", hook, "plugin", pluginPath)
}

// ExecuteHook executes all plugins registered for a hook
func (pm *PluginManager) ExecuteHook(hook string) error {
	logger.Debug("Executing hook", "hook", hook)
	
	plugins, exists := pm.Plugins[hook]
	if !exists {
		logger.Debug("No plugins registered for hook", "hook", hook)
		return nil
	}
	
	for _, pluginPath := range plugins {
		logger.Info("Executing plugin", "hook", hook, "plugin", pluginPath)
		
		// Check if the plugin file exists
		if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
			logger.Error("Plugin file not found", "plugin", pluginPath)
			continue
		}
		
		// Validate the plugin path for security
		// This prevents command injection and directory traversal attacks
		validatedPath, err := validatePluginPath(pluginPath, "plugins")
		if err != nil {
			logger.Error("Plugin path validation failed", "plugin", pluginPath, "error", err)
			continue
		}

		// Determine how to execute the plugin based on its extension
		ext := filepath.Ext(validatedPath)
		var cmd *exec.Cmd

		switch ext {
		case ".sh", ".bash", ".zsh":
			// Shell script
			cmd = exec.Command("bash", validatedPath) // #nosec G204 -- Path validated by validatePluginPath()
		case ".py":
			// Python script
			cmd = exec.Command("python", validatedPath) // #nosec G204 -- Path validated by validatePluginPath()
		case ".js":
			// JavaScript file
			cmd = exec.Command("node", validatedPath) // #nosec G204 -- Path validated by validatePluginPath()
		default:
			// This should not happen due to validation, but handle gracefully
			logger.Error("Unexpected plugin extension after validation", "plugin", validatedPath, "extension", ext)
			continue
		}
		
		// Execute the plugin
		output, err := cmd.CombinedOutput()
		if err != nil {
			logger.Error("Plugin execution failed", "plugin", pluginPath, "error", err, "output", string(output))
		} else {
			logger.Info("Plugin executed successfully", "plugin", pluginPath, "output", string(output))
		}
	}
	
	return nil
}

// LoadPluginsFromDir loads all plugins from a directory
func (pm *PluginManager) LoadPluginsFromDir(dirPath string) error {
	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		logger.Debug("Plugin directory not found", "directory", dirPath)
		return nil
	}
	
	// Walk the directory to find plugin files
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if info.IsDir() {
			return nil
		}
		
		// Determine the hook from the file name
		// Convention: filename format is "hookname.ext"
		hook := filepath.Base(path)
		if ext := filepath.Ext(hook); ext != "" {
			hook = hook[:len(hook)-len(ext)]
		}
		
		// Register the plugin
		pm.RegisterPlugin(hook, path)
		
		return nil
	})
}