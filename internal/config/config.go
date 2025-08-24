package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	// Validate the config path to prevent directory traversal attacks
	cleanPath := filepath.Clean(configPath)

	// Ensure the path is absolute and doesn't contain dangerous patterns
	if !filepath.IsAbs(cleanPath) {
		// Convert relative path to absolute
		absPath, err := filepath.Abs(cleanPath)
		if err != nil {
			logger.Error("Failed to get absolute path for config", "path", configPath, "error", err)
			return config, err
		}
		cleanPath = absPath
	}

	// Check for directory traversal patterns
	if strings.Contains(cleanPath, "..") {
		logger.Error("Invalid config path: contains directory traversal", "path", configPath)
		return config, fmt.Errorf("invalid config path: %s", configPath)
	}

	// Check if config file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		logger.Info("Config file not found, using defaults", "path", cleanPath)
		return config, nil
	}

	// Read the config file
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		logger.Error("Failed to read config file", "path", cleanPath, "error", err)
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