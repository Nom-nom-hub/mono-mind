package main

import (
	"os"
	"mono-mind/internal/config"
	"mono-mind/internal/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfigFromDir(".")
	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize logger based on config
	if cfg.LogLevel == "debug" {
		logger.DebugFlag = true
	}
	
	logger.Init()

	// Execute the root command
	if err := NewRootCmd().Execute(); err != nil {
		logger.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}