package visualization

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"mono-mind/internal/analyzer"
)

// HTMLTemplate is the template for HTML visualization
const HTMLTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>MonoMind Dependency Graph</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background-color: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            text-align: center;
        }
        .module {
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 15px;
            margin: 10px 0;
            background-color: #f9f9f9;
        }
        .module-name {
            font-weight: bold;
            font-size: 18px;
            color: #007acc;
        }
        .module-language {
            color: #666;
            font-style: italic;
        }
        .dependencies, .dependents {
            margin: 10px 0;
        }
        .dependency, .dependent {
            margin: 5px 0;
            padding: 5px;
            background-color: #e9e9e9;
            border-radius: 3px;
        }
        .no-dependencies {
            color: #999;
            font-style: italic;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>MonoMind Dependency Graph</h1>
        {{range $moduleName, $module := .Modules}}
        <div class="module">
            <div class="module-name">{{$moduleName}}</div>
            <div class="module-language">({{$module.Language}})</div>
            
            <div class="dependencies">
                <strong>Dependencies:</strong>
                {{if $module.Dependencies}}
                    {{range $dep := $module.Dependencies}}
                    <div class="dependency">{{$dep}}</div>
                    {{end}}
                {{else}}
                    <div class="no-dependencies">No external dependencies</div>
                {{end}}
            </div>
            
            <div class="dependents">
                <strong>Used by:</strong>
                {{$dependents := getDependents $moduleName}}
                {{if $dependents}}
                    {{range $dep := $dependents}}
                    <div class="dependent">{{$dep}}</div>
                    {{end}}
                {{else}}
                    <div class="no-dependencies">Not used by any other modules</div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>
`

// PrintHTMLDependencyGraph prints the dependency graph as an HTML file
func PrintHTMLDependencyGraph(graph *analyzer.RepoGraph, filename string) error {
	// Validate the filename to prevent directory traversal attacks
	// Clean the path and ensure it's a simple filename
	cleanFilename := filepath.Clean(filename)
	
	// Ensure the filename doesn't contain path separators
	if filepath.Dir(cleanFilename) != "." {
		return fmt.Errorf("invalid filename: %s", filename)
	}
	
	// Ensure the filename has an html extension
	if filepath.Ext(cleanFilename) != ".html" {
		cleanFilename += ".html"
	}
	
	// Create a map of module names to their dependents
	dependentsMap := make(map[string][]string)
	for moduleName := range graph.Modules {
		dependents := graph.GetDependentModules(moduleName)
		dependentsMap[moduleName] = dependents
	}
	
	// Create a function to get dependents for a module
	funcMap := template.FuncMap{
		"getDependents": func(moduleName string) []string {
			return dependentsMap[moduleName]
		},
	}
	
	// Parse the template
	tmpl, err := template.New("dependencyGraph").Funcs(funcMap).Parse(HTMLTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}
	
	// Create the output file
	file, err := os.Create(cleanFilename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	
	// Execute the template
	err = tmpl.Execute(file, graph)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	
	return nil
}