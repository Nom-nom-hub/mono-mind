package release

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"mono-mind/internal/logger"
)

// ReleaseConfig holds configuration for the release process
type ReleaseConfig struct {
	VersionBump string `json:"version_bump"` // major, minor, patch
	Changelog   bool   `json:"changelog"`
	Publish     bool   `json:"publish"`
}

// ReleaseResult holds the result of a release operation
type ReleaseResult struct {
	NewVersion   string   `json:"new_version"`
	Changelog    string   `json:"changelog"`
	Errors       []string `json:"errors"`
	Duration     string   `json:"duration"`
}

// ManageRelease handles version bumping, changelog generation, and publishing
func ManageRelease(config ReleaseConfig) *ReleaseResult {
	logger.Info("Managing release", 
		"version_bump", config.VersionBump, 
		"changelog", config.Changelog,
		"publish", config.Publish)
	
	result := &ReleaseResult{
		NewVersion: "",
		Changelog:  "",
		Errors:     []string{},
	}
	
	// Get the current version
	currentVersion, err := getCurrentVersion()
	if err != nil {
		logger.Error("Failed to get current version", "error", err)
		result.Errors = append(result.Errors, fmt.Sprintf("Failed to get current version: %v", err))
		return result
	}
	
	// Bump the version
	newVersion, err := bumpVersion(currentVersion, config.VersionBump)
	if err != nil {
		logger.Error("Failed to bump version", "error", err)
		result.Errors = append(result.Errors, fmt.Sprintf("Failed to bump version: %v", err))
		return result
	}
	
	result.NewVersion = newVersion
	
	// Generate changelog if requested
	if config.Changelog {
		changelog, err := generateChangelog(currentVersion, newVersion)
		if err != nil {
			logger.Error("Failed to generate changelog", "error", err)
			result.Errors = append(result.Errors, fmt.Sprintf("Failed to generate changelog: %v", err))
		} else {
			result.Changelog = changelog
			// Save changelog to file
			err = saveChangelog(changelog, newVersion)
			if err != nil {
				logger.Error("Failed to save changelog", "error", err)
				result.Errors = append(result.Errors, fmt.Sprintf("Failed to save changelog: %v", err))
			}
		}
	}
	
	// Update version in files
	err = updateVersionInFiles(newVersion)
	if err != nil {
		logger.Error("Failed to update version in files", "error", err)
		result.Errors = append(result.Errors, fmt.Sprintf("Failed to update version in files: %v", err))
	}
	
	// If publish is requested, publish the release
	if config.Publish {
		err = publishRelease(newVersion)
		if err != nil {
			logger.Error("Failed to publish release", "error", err)
			result.Errors = append(result.Errors, fmt.Sprintf("Failed to publish release: %v", err))
		}
	}
	
	logger.Info("Release management completed", 
		"new_version", result.NewVersion,
		"errors", len(result.Errors))
	
	return result
}

// getCurrentVersion gets the current version from Git tags
func getCurrentVersion() (string, error) {
	// Run git command to get the latest tag
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		// If no tags exist, return default version
		return "0.0.0", nil
	}
	
	version := strings.TrimSpace(string(output))
	// Remove 'v' prefix if present
	version = strings.TrimPrefix(version, "v")
	
	return version, nil
}

// bumpVersion bumps the version based on the bump type
func bumpVersion(version, bumpType string) (string, error) {
	// Split version into parts
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format: %s", version)
	}
	
	// Parse version parts
	major := 0
	minor := 0
	patch := 0
	
	_, err := fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		logger.Error("Failed to parse version", "error", err)
		return "", err
	}
	
	// Bump based on type
	switch bumpType {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return "", fmt.Errorf("invalid bump type: %s", bumpType)
	}
	
	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

// generateChangelog generates a changelog from Git history
func generateChangelog(fromVersion, toVersion string) (string, error) {
	// Run git command to get commit history
	var cmd *exec.Cmd
	if fromVersion == "0.0.0" {
		// If this is the first release, get all commits
		cmd = exec.Command("git", "log", "--oneline", "--no-merges")
	} else {
		// Get commits since the last version
		cmd = exec.Command("git", "log", fmt.Sprintf("v%s..HEAD", fromVersion), "--oneline", "--no-merges")
	}
	
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git log: %v", err)
	}
	
	// Format the changelog
	commits := strings.Split(strings.TrimSpace(string(output)), "\n")
	
	var changelog strings.Builder
	changelog.WriteString(fmt.Sprintf("# Changelog\n\n## [%s] - %s\n", toVersion, time.Now().Format("2006-01-02")))
	
	if len(commits) > 0 && commits[0] != "" {
		for _, commit := range commits {
			if commit != "" {
				// Try to categorize commits based on their prefixes
				parts := strings.SplitN(commit, " ", 2)
				if len(parts) == 2 {
					hash := parts[0]
					message := parts[1]
					
					// Add to changelog
					changelog.WriteString(fmt.Sprintf("- %s (%s)\n", message, hash[:7]))
				}
			}
		}
	} else {
		changelog.WriteString("- No changes\n")
	}
	
	return changelog.String(), nil
}

	// saveChangelog saves the changelog to a file
	func saveChangelog(changelog, version string) error {
		filename := "CHANGELOG.md"
	
	// Check if file exists
	var content string
	if _, err := os.Stat(filename); err == nil {
		// File exists, read current content
		data, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		content = string(data)
	}
	
	// Prepend new changelog entry
	newContent := changelog + "\n" + content
	
	// Write to file
	return os.WriteFile(filename, []byte(newContent), 0644)
}

// updateVersionInFiles updates version in relevant files
func updateVersionInFiles(version string) error {
	// In a real implementation, we would update version in:
	// - package.json for Node.js projects
	// - setup.py for Python projects
	// - Cargo.toml for Rust projects
	// - go.mod for Go projects
	// - etc.
	
	logger.Info("Updating version in files", "version", version)
	
	// For now, we'll just log that we would update files
	return nil
}

// publishRelease publishes the release
func publishRelease(version string) error {
	logger.Info("Publishing release", "version", version)
	
	// In a real implementation, we would:
	// 1. Create a Git tag
	// 2. Push the tag to remote repository
	// 3. Publish to package registries (npm, PyPI, etc.)
	
	// Create Git tag
	cmd := exec.Command("git", "tag", "-a", fmt.Sprintf("v%s", version), "-m", fmt.Sprintf("Release version %s", version))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create Git tag: %v", err)
	}
	
	// Push tag
	cmd = exec.Command("git", "push", "origin", fmt.Sprintf("v%s", version))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to push Git tag: %v", err)
	}
	
	return nil
}