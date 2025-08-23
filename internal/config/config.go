package config

import (
	"os"
	"path/filepath"
	"mono-mind/internal/logger"
	
	"gopkg.in/yaml.v2"
)

// Config represents the application configuration
type Config struct {
	LogLevel string `yaml:"log_level"`
	Analyzer AnalyzerConfig `yaml:"analyzer"`
	Build BuildConfig `yaml:"build"`
	Release ReleaseConfig `yaml:"release"`
}

// AnalyzerConfig represents the analyzer configuration
type AnalyzerConfig struct {
	Languages []string `yaml:"languages"`
	IgnoreExtensions []string `yaml:"ignore_extensions"`
}

// BuildConfig represents the build configuration
type BuildConfig struct {
	DefaultCommand string `yaml:"default_command"`
	Parallel bool `yaml:"parallel"`
	MaxConcurrent int `yaml:"max_concurrent"`
}

// ReleaseConfig represents the release configuration
type ReleaseConfig struct {
	DefaultBump string `yaml:"default_bump"`
	ChangelogFormat string `yaml:"changelog_format"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		LogLevel: "info",
		Analyzer: AnalyzerConfig{
			Languages: []string{"go", "javascript", "typescript", "python"},
			IgnoreExtensions: []string{".git", ".svn", "node_modules", "vendor", "target", "build", "dist", ".DS_Store"},
		},
		Build: BuildConfig{
			DefaultCommand: "make",
			Parallel: true,
			MaxConcurrent: 4,
		},
		Release: ReleaseConfig{
			DefaultBump: "patch",
			ChangelogFormat: "markdown",
		},
	}
}

// LoadConfig loads the configuration from a file
func LoadConfig(configPath string) (*Config, error) {
	// Use default config as base
	config := DefaultConfig()
	
	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Info("Config file not found, using defaults", "path", configPath)
		return config, nil
	}
	
	// Read the config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Error("Failed to read config file", "path", configPath, "error", err)
		return config, err
	}
	
	// Parse the YAML
	err = yaml.Unmarshal(data, config)
	if err != nil {
		logger.Error("Failed to parse config file", "path", configPath, "error", err)
		return config, err
	}
	
	logger.Info("Loaded configuration", "path", configPath)
	return config, nil
}

// LoadConfigFromDir loads the configuration from a directory
func LoadConfigFromDir(dirPath string) (*Config, error) {
	// Look for config file in the directory
	configPath := filepath.Join(dirPath, "configs", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Try alternative names
		alternatives := []string{
			filepath.Join(dirPath, "configs", "config.yml"),
			filepath.Join(dirPath, "config.yaml"),
			filepath.Join(dirPath, "config.yml"),
			filepath.Join(dirPath, ".mono.yaml"),
			filepath.Join(dirPath, ".mono.yml"),
		}
		
		for _, alt := range alternatives {
			if _, err := os.Stat(alt); err == nil {
				configPath = alt
				break
			}
		}
	}
	
	return LoadConfig(configPath)
}