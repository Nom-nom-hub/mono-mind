package plugins

import (
	"os"
	"os/exec"
	"path/filepath"
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
		
		// Determine how to execute the plugin based on its extension
		ext := filepath.Ext(pluginPath)
		var cmd *exec.Cmd
		
		switch ext {
		case ".sh":
			// Shell script
			cmd = exec.Command("bash", pluginPath)
		case ".py":
			// Python script
			cmd = exec.Command("python", pluginPath)
		case ".js":
			// JavaScript file
			cmd = exec.Command("node", pluginPath)
		default:
			// Try to execute directly (could be a binary)
			cmd = exec.Command(pluginPath)
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