package plugins

import (
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
		// TODO: Execute the plugin binary or script
		// This would involve running the plugin as a subprocess
	}
	
	return nil
}