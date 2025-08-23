package visualization

import (
	"fmt"
	"mono-mind/internal/analyzer"
)

// PrintDependencyGraph prints the dependency graph in a visual format
func PrintDependencyGraph(graph *analyzer.RepoGraph) {
	fmt.Println("Dependency Graph:")
	fmt.Println("=================")
	
	// Print each module and its dependencies
	for moduleName, module := range graph.Modules {
		fmt.Printf("📁 %s (%s)\n", moduleName, module.Language)
		
		// Print dependencies
		dependencies := graph.GetModuleDependencies(moduleName)
		if len(dependencies) > 0 {
			for _, dep := range dependencies {
				fmt.Printf("  └─ depends on: %s\n", dep)
			}
		} else {
			fmt.Printf("  └─ no external dependencies\n")
		}
		
		// Print dependents
		dependents := graph.GetDependentModules(moduleName)
		if len(dependents) > 0 {
			for _, dep := range dependents {
				fmt.Printf("  └─ used by: %s\n", dep)
			}
		}
		
		fmt.Println()
	}
}

// PrintImpactAnalysis prints the impact analysis results in a visual format
func PrintImpactAnalysis(graph *analyzer.RepoGraph, changedFile string) {
	fmt.Printf("Impact Analysis for: %s\n", changedFile)
	fmt.Println("======================")
	
	// Find the module that contains the changed file
	// In a real implementation, we would parse the file path to determine the module
	changedModule := "example" // Placeholder
	
	fmt.Printf("Changed Module: %s\n", changedModule)
	fmt.Println()
	
	// Print affected modules
	dependents := graph.GetDependentModules(changedModule)
	if len(dependents) > 0 {
		fmt.Println("Affected Modules:")
		fmt.Printf("  📁 %s (directly changed)\n", changedModule)
		for _, dep := range dependents {
			fmt.Printf("  📁 %s (depends on %s)\n", dep, changedModule)
		}
	} else {
		fmt.Printf("No modules are affected by changes to %s\n", changedModule)
	}
	
	fmt.Println()
}