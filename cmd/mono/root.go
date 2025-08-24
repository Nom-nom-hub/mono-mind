package main

import (
	"strings"
	"github.com/spf13/cobra"
	"mono-mind/internal/analyzer"
	"mono-mind/internal/build"
	"mono-mind/internal/impact"
	"mono-mind/internal/logger"
	"mono-mind/internal/refactor"
	"mono-mind/internal/release"
	"mono-mind/internal/test"
	"mono-mind/internal/visualization"
)

// NewRootCmd creates the root command for the mono CLI
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mono",
		Short: "MonoMind is an AI-powered monorepo management tool",
		Long: `MonoMind is an AI-powered development assistant designed to autonomously 
manage monorepos and complex codebases. It combines code analysis, build orchestration, 
testing, refactoring, and release management into a single intelligent CLI-driven interface.`,
	}

	// Add subcommands
	rootCmd.AddCommand(newAnalyzeCmd())
	rootCmd.AddCommand(newImpactCmd())
	rootCmd.AddCommand(newBuildCmd())
	rootCmd.AddCommand(newRefactorCmd())
	rootCmd.AddCommand(newReleaseCmd())
	rootCmd.AddCommand(newTestCmd())
	rootCmd.AddCommand(newVisualizeCmd())

	// Add global flags
	rootCmd.PersistentFlags().BoolVar(&logger.DebugFlag, "debug", false, "Enable debug logging")

	return rootCmd
}

// Subcommand functions will be implemented in their respective files
func newAnalyzeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze the repository and build dependency graph",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Analyzing repository...")
			// Get current directory as the root path
			rootPath := "."
			graph, err := analyzer.AnalyzeRepo(rootPath)
			if err != nil {
				logger.Error("Failed to analyze repository", "error", err)
				return
			}
			
			// Print the dependency graph using visualization
			visualization.PrintDependencyGraph(graph)
			
			logger.Info("Analysis complete", "modules", len(graph.Modules))
		},
	}
	return cmd
}

func newImpactCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "impact [file]",
		Short: "Show affected modules/tests for a change",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			logger.Info("Analyzing impact for file", "file", file)
			
			// Get current directory as the root path
			rootPath := "."
			
			// First, analyze the repo to get the dependency graph
			graph, err := analyzer.AnalyzeRepo(rootPath)
			if err != nil {
				logger.Error("Failed to analyze repository", "error", err)
				return
			}
			
			// Perform impact analysis
			result := impact.AnalyzeImpact(graph, file)
			
			// Print impact analysis using visualization
			// For now, we'll pass the graph since our impact result doesn't have the right structure
			visualization.PrintImpactAnalysis(graph, file)
			
			logger.Info("Impact analysis completed", 
				"affected_modules", len(result.AffectedModules),
				"affected_tests", len(result.AffectedTests))
		},
	}
	return cmd
}

func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Incremental build based on changes",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Building affected modules...")
			
			// Get current directory as the root path
			rootPath := "."
			
			// First, analyze the repo to get the dependency graph
			graph, err := analyzer.AnalyzeRepo(rootPath)
			if err != nil {
				logger.Error("Failed to analyze repository", "error", err)
				return
			}
			
			// Configure build
			config := build.BuildConfig{
				Parallel:      true,
				MaxConcurrent: 4,
				DryRun:        false,
			}
			
			// Perform incremental build
			result := build.IncrementalBuild(graph, config)
			
			logger.Info("Build completed", 
				"modules_built", len(result.ModulesBuilt),
				"errors", len(result.Errors))
		},
	}
	return cmd
}

func newRefactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refactor",
		Short: "Safe rename/move across modules",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Refactoring code...")
			
			// Get flags
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			rename, _ := cmd.Flags().GetString("rename")
			move, _ := cmd.Flags().GetString("move")
			file, _ := cmd.Flags().GetString("file")
			
			// Configure refactor
			config := refactor.RefactorConfig{
				DryRun: dryRun,
			}
			
			// Perform refactor based on flags
			var result *refactor.RefactorResult
			if rename != "" {
				// Parse rename format: oldName:newName
				names := strings.Split(rename, ":")
				if len(names) != 2 {
					logger.Error("Invalid rename format. Use oldName:newName")
					return
				}
				config.OldName = names[0]
				config.NewName = names[1]
				config.FilePath = file
				result = refactor.Rename(config)
			} else if move != "" {
				// Parse move format: oldPath:newPath
				paths := strings.Split(move, ":")
				if len(paths) != 2 {
					logger.Error("Invalid move format. Use oldPath:newPath")
					return
				}
				result = refactor.Move(paths[0], paths[1], dryRun)
			} else {
				logger.Info("No refactor operation specified. Use --rename or --move")
				return
			}
			
			logger.Info("Refactor completed", 
				"files_changed", len(result.FilesChanged),
				"errors", len(result.Errors))
		},
	}
	
	// Add flags
	cmd.Flags().Bool("dry-run", false, "Preview changes without applying them")
	cmd.Flags().String("rename", "", "Rename identifier (format: oldName:newName)")
	cmd.Flags().String("move", "", "Move file/directory (format: oldPath:newPath)")
	cmd.Flags().String("file", "", "Specific file to refactor (for rename operations)")
	
	return cmd
}

func newReleaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "Version bump, changelog, publish packages",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Managing release...")
			
			// Get flags
			bump, _ := cmd.Flags().GetString("bump")
			changelog, _ := cmd.Flags().GetBool("changelog")
			publish, _ := cmd.Flags().GetBool("publish")
			
			// Configure release
			config := release.ReleaseConfig{
				VersionBump: bump,
				Changelog:   changelog,
				Publish:     publish,
			}
			
			// Manage release
			result := release.ManageRelease(config)
			
			logger.Info("Release completed", 
				"new_version", result.NewVersion,
				"errors", len(result.Errors))
		},
	}
	
	// Add flags
	cmd.Flags().String("bump", "patch", "Version bump type (major, minor, patch)")
	cmd.Flags().Bool("changelog", false, "Generate changelog")
	cmd.Flags().Bool("publish", false, "Publish the release")
	
	return cmd
}

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Run tests for affected modules",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Running tests...")
			
			// Get flags
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			parallel, _ := cmd.Flags().GetBool("parallel")
			
			// Get current directory as the root path
			rootPath := "."
			
			// First, analyze the repo to get the dependency graph
			graph, err := analyzer.AnalyzeRepo(rootPath)
			if err != nil {
				logger.Error("Failed to analyze repository", "error", err)
				return
			}
			
			// Configure test
			config := test.TestConfig{
				Parallel:      parallel,
				MaxConcurrent: 4,
				DryRun:        dryRun,
			}
			
			// Run tests
			result := test.RunTests(graph, config)
			
			logger.Info("Test execution completed", 
				"tests_run", result.TestsRun,
				"tests_passed", result.TestsPassed,
				"tests_failed", result.TestsFailed,
				"errors", len(result.Errors))
		},
	}
	
	// Add flags
	cmd.Flags().Bool("dry-run", false, "Preview test execution without running tests")
	cmd.Flags().Bool("parallel", true, "Run tests in parallel")
	
	return cmd
}

func newVisualizeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "visualize [type]",
		Short: "Visualize the repository structure",
		Long:  "Visualize the repository structure in different formats",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Get visualization type (default to tree)
			visType := "tree"
			if len(args) > 0 {
				visType = args[0]
			}
			
			// Get output file flag
			outputFile, _ := cmd.Flags().GetString("output")
			
			logger.Info("Visualizing repository", "type", visType, "output", outputFile)
			
			// Get current directory as the root path
			rootPath := "."
			
			// Analyze the repo to get the dependency graph
			graph, err := analyzer.AnalyzeRepo(rootPath)
			if err != nil {
				logger.Error("Failed to analyze repository", "error", err)
				return
			}
			
			// If output file is specified and type is html, generate HTML
			if outputFile != "" && visType == "html" {
				err = visualization.PrintHTMLDependencyGraph(graph, outputFile)
				if err != nil {
					logger.Error("Failed to generate HTML visualization", "error", err)
					return
				}
				logger.Info("HTML visualization saved", "file", outputFile)
				return
			}
			
			// Print the appropriate visualization
			switch visType {
			case "tree":
				visualization.PrintDependencyGraph(graph)
			case "ascii":
				visualization.PrintASCIIDependencyGraph(graph)
			case "horizontal":
				visualization.PrintHorizontalDependencyGraph(graph)
			default:
				logger.Error("Unknown visualization type", "type", visType)
				logger.Info("Available types: tree, ascii, horizontal, html")
			}
		},
	}
	
	// Add flags
	cmd.Flags().String("output", "", "Output file for HTML visualization")
	
	return cmd
}