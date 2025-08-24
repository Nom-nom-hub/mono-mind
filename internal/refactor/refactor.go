package refactor

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"mono-mind/internal/logger"
	"golang.org/x/tools/go/ast/astutil"
)

// RefactorConfig holds configuration for the refactor process
type RefactorConfig struct {
	DryRun   bool   `json:"dry_run"`
	OldName  string `json:"old_name"`
	NewName  string `json:"new_name"`
	FilePath string `json:"file_path"`
}

// RefactorResult holds the result of a refactor operation
type RefactorResult struct {
	FilesChanged []string `json:"files_changed"`
	Errors       []string `json:"errors"`
	Duration     string   `json:"duration"`
}

// Rename performs a safe rename operation across the repository
func Rename(config RefactorConfig) *RefactorResult {
	logger.Info("Performing rename operation", 
		"old_name", config.OldName, 
		"new_name", config.NewName,
		"dry_run", config.DryRun)
	
	result := &RefactorResult{
		FilesChanged: []string{},
		Errors:       []string{},
	}
	
	// If a specific file path is provided, refactor only that file
	if config.FilePath != "" {
		err := refactorFile(config.FilePath, config.OldName, config.NewName, config.DryRun, result)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("Error refactoring file %s: %v", config.FilePath, err))
		}
	} else {
		// Refactor all Go files in the current directory
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			
			// Skip non-Go files
			if !strings.HasSuffix(path, ".go") {
				return nil
			}
			
			// Skip vendor directories
			if strings.Contains(path, "vendor") || strings.Contains(path, "node_modules") {
				return filepath.SkipDir
			}
			
			err = refactorFile(path, config.OldName, config.NewName, config.DryRun, result)
			if err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("Error refactoring file %s: %v", path, err))
			}
			
			return nil
		})
		
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("Error walking directory: %v", err))
		}
	}
	
	logger.Info("Rename operation completed", 
		"files_changed", len(result.FilesChanged),
		"errors", len(result.Errors))
	
	return result
}

// refactorFile refactors a single file
func refactorFile(filePath, oldName, newName string, dryRun bool, result *RefactorResult) error {
	// Parse the file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}
	
	// Apply the refactoring
	changed := false
	astutil.Apply(file, func(c *astutil.Cursor) bool {
		// Check for identifiers that match the old name
		if ident, ok := c.Node().(*ast.Ident); ok {
			if ident.Name == oldName {
				// Replace with the new name
				ident.Name = newName
				changed = true
			}
		}
		return true
	}, nil)
	
	// If no changes were made, return early
	if !changed {
		return nil
	}
	
	// Add file to changed list
	result.FilesChanged = append(result.FilesChanged, filePath)
	
	// If dry run, just report the change
	if dryRun {
		logger.Info("Would refactor file (dry-run)", "file", filePath)
		return nil
	}
	
	// Format and write the modified file
	var buf bytes.Buffer
	err = format.Node(&buf, fset, file)
	if err != nil {
		return fmt.Errorf("error formatting file: %v", err)
	}
	
	err = os.WriteFile(filePath, buf.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	
	logger.Info("Refactored file", "file", filePath)
	return nil
}

// Move moves a file or directory to a new location
func Move(oldPath, newPath string, dryRun bool) *RefactorResult {
	logger.Info("Moving file/directory", 
		"old_path", oldPath, 
		"new_path", newPath,
		"dry_run", dryRun)
	
	result := &RefactorResult{
		FilesChanged: []string{},
		Errors:       []string{},
	}
	
	// If dry run, just report what would be moved
	if dryRun {
		logger.Info("Would move file/directory (dry-run)", "old_path", oldPath, "new_path", newPath)
		result.FilesChanged = append(result.FilesChanged, oldPath)
		return result
	}
	
	// Create the destination directory if it doesn't exist
	destDir := filepath.Dir(newPath)
	err := os.MkdirAll(destDir, 0750)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("Error creating destination directory: %v", err))
		return result
	}
	
	// Move the file or directory
	err = os.Rename(oldPath, newPath)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("Error moving file/directory: %v", err))
		return result
	}
	
	result.FilesChanged = append(result.FilesChanged, oldPath)
	logger.Info("Moved file/directory", "old_path", oldPath, "new_path", newPath)
	
	return result
}