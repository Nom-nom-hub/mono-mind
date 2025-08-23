package visualization

import (
	"fmt"
	"mono-mind/internal/analyzer"
)

// PrintASCIIDependencyGraph prints the dependency graph as an ASCII diagram
func PrintASCIIDependencyGraph(graph *analyzer.RepoGraph) {
	fmt.Println("Dependency Graph (ASCII):")
	fmt.Println("========================")
	
	// Print each module and its dependencies
	for moduleName, module := range graph.Modules {
		// Print the module as a box
		fmt.Printf("+-----------------------------+\n")
		fmt.Printf("| %-27s |\n", moduleName)
		fmt.Printf("| %-27s |\n", "("+module.Language+")")
		fmt.Printf("+-----------------------------+\n")
		
		// Print dependencies with arrows
		dependencies := graph.GetModuleDependencies(moduleName)
		if len(dependencies) > 0 {
			for _, dep := range dependencies {
				fmt.Printf("          |\n")
				fmt.Printf("          v\n")
				fmt.Printf("+-----------------------------+\n")
				fmt.Printf("| %-27s |\n", dep)
				fmt.Printf("+-----------------------------+\n")
			}
		} else {
			fmt.Printf("          |\n")
			fmt.Printf("          v\n")
			fmt.Printf("     (no dependencies)\n")
		}
		
		fmt.Println()
	}
}

// PrintHorizontalDependencyGraph prints the dependency graph horizontally
func PrintHorizontalDependencyGraph(graph *analyzer.RepoGraph) {
	fmt.Println("Dependency Graph (Horizontal):")
	fmt.Println("==============================")
	
	// Print all modules in a row
	fmt.Print("Modules: ")
	first := true
	for moduleName := range graph.Modules {
		if !first {
			fmt.Print(" -> ")
		}
		fmt.Print(moduleName)
		first = false
	}
	fmt.Println()
	
	// Print dependencies for each module
	fmt.Println("\nDependencies:")
	for moduleName, module := range graph.Modules {
		dependencies := graph.GetModuleDependencies(moduleName)
		if len(dependencies) > 0 {
			fmt.Printf("  %s (%s):\n", moduleName, module.Language)
			for _, dep := range dependencies {
				fmt.Printf("    -> %s\n", dep)
			}
		}
	}
}