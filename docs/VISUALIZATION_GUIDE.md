# MonoMind Visualization Guide

## Table of Contents
1. [Introduction](#introduction)
2. [Visualization Types](#visualization-types)
3. [Tree Visualization](#tree-visualization)
4. [ASCII Visualization](#ascii-visualization)
5. [Horizontal Visualization](#horizontal-visualization)
6. [HTML Visualization](#html-visualization)
7. [Customizing Visualizations](#customizing-visualizations)
8. [Interpreting Results](#interpreting-results)
9. [Best Practices](#best-practices)

## Introduction

This guide explains how to use MonoMind's visualization features to understand your repository's structure and dependencies.

## Visualization Types

MonoMind provides several visualization formats to suit different needs:

### Tree View
- Hierarchical representation of modules and dependencies
- Easy to understand nested relationships
- Good for terminal-based workflows

### ASCII View
- Box-style diagram representation
- Visual separation of modules
- Good for simple terminal displays

### Horizontal View
- Linear representation of modules
- Shows relationships in a flow format
- Good for understanding data flow

### HTML View
- Interactive web-based visualization
- Rich formatting and styling
- Can be shared and viewed in browsers
- Good for presentations and documentation

## Tree Visualization

### Basic Usage
```bash
# Generate tree visualization
mono.exe visualize tree

# Or simply
mono.exe visualize
```

### Output Format
```
Dependency Graph:
=================
📁 api (go)
  └─ depends on: shared
  └─ used by: web

📁 web (javascript)
  └─ depends on: api, shared

📁 shared (go)
  └─ used by: api, web

📁 mobile (typescript)
  └─ depends on: shared
```

### Features
- Module names with language indicators
- Clear dependency relationships
- Reverse dependency tracking (used by)
- File type icons

## ASCII Visualization

### Basic Usage
```bash
# Generate ASCII visualization
mono.exe visualize ascii
```

### Output Format
```
Dependency Graph (ASCII):
========================
+-----------------------------+
| api                         |
| (go)                        |
+-----------------------------+
          |
          v
+-----------------------------+
| shared                      |
| (go)                        |
+-----------------------------+

+-----------------------------+
| web                         |
| (javascript)                |
+-----------------------------+
          |
          v
+-----------------------------+
| api                         |
| (go)                        |
+-----------------------------+
```

### Features
- Box-style module representation
- Clear directional arrows
- Language indicators
- Visual separation of components

## Horizontal Visualization

### Basic Usage
```bash
# Generate horizontal visualization
mono.exe visualize horizontal
```

### Output Format
```
Dependency Graph (Horizontal):
==============================
Modules: api -> web -> shared -> mobile

Dependencies:
  api (go):
    -> shared
  web (javascript):
    -> api
    -> shared
  shared (go):
  mobile (typescript):
    -> shared
```

### Features
- Linear module flow
- Compact representation
- Clear dependency listing
- Good for understanding sequential relationships

## HTML Visualization

### Basic Usage
```bash
# Generate HTML visualization
mono.exe visualize html --output dependencies.html
```

### Features
- Interactive web page
- Professional styling with CSS
- Hover effects and visual enhancements
- Can be shared and viewed in any browser
- Printable and presentable
- Can be embedded in documentation

### HTML Output Structure
```html
<!DOCTYPE html>
<html>
<head>
    <title>MonoMind Dependency Graph</title>
    <style>
        /* Professional CSS styling */
    </style>
</head>
<body>
    <div class="container">
        <h1>MonoMind Dependency Graph</h1>
        <div class="module">
            <div class="module-name">api</div>
            <div class="module-language">(go)</div>
            <div class="dependencies">
                <strong>Dependencies:</strong>
                <div class="dependency">shared</div>
            </div>
            <div class="dependents">
                <strong>Used by:</strong>
                <div class="dependent">web</div>
            </div>
        </div>
        <!-- More modules -->
    </div>
</body>
</html>
```

### Customization
You can customize the HTML output by modifying the template in the source code or by post-processing the generated HTML.

## Customizing Visualizations

### Output Files
For HTML visualizations, you can specify the output file:
```bash
# Save to specific file
mono.exe visualize html --output my-project-dependencies.html

# Save to different directory
mono.exe visualize html --output ./docs/dependencies.html
```

### Combining with Other Commands
```bash
# Analyze and visualize in one command
mono.exe analyze && mono.exe visualize

# Generate multiple visualizations
mono.exe visualize tree > tree.txt
mono.exe visualize ascii > ascii.txt
mono.exe visualize horizontal > horizontal.txt
mono.exe visualize html --output graph.html
```

### Automation
Create scripts to automatically generate visualizations:
```bash
#!/bin/bash
# generate-visualizations.sh

echo "Generating dependency visualizations..."

mono.exe visualize tree > docs/tree-view.txt
mono.exe visualize ascii > docs/ascii-view.txt
mono.exe visualize horizontal > docs/horizontal-view.txt
mono.exe visualize html --output docs/dependencies.html

echo "Visualizations generated in docs/ directory"
```

## Interpreting Results

### Understanding Dependencies
- **Dependencies**: Modules that this module depends on
- **Dependents**: Modules that depend on this module
- **Circular Dependencies**: Avoid these when possible
- **Orphan Modules**: Modules with no dependencies or dependents

### Identifying Issues
- **Too many dependencies**: Module may be doing too much
- **Too many dependents**: Module may be a bottleneck
- **Circular dependencies**: Can cause build and testing issues
- **Orphan modules**: May be unused or poorly integrated

### Making Decisions
Use visualizations to:
- Plan refactorings
- Identify optimization opportunities
- Understand impact of changes
- Communicate architecture to team members

## Best Practices

### When to Use Each Visualization

#### Tree View
- Best for: Understanding hierarchical relationships
- Use when: You need to see the nested structure
- Good for: Terminal-based workflows

#### ASCII View
- Best for: Simple visual representation
- Use when: You want clear boxes and arrows
- Good for: Quick terminal viewing

#### Horizontal View
- Best for: Understanding data flow
- Use when: You want to see linear relationships
- Good for: Documentation and simple analysis

#### HTML View
- Best for: Professional presentation
- Use when: You need to share with others
- Good for: Documentation, presentations, and reports

### Regular Visualization Updates
```bash
# Add to your CI/CD pipeline
mono.exe analyze
mono.exe visualize html --output docs/latest-dependencies.html
```

### Version Control
- Commit visualization outputs for important releases
- Compare visualizations over time
- Use visualizations in architectural documentation

### Sharing Visualizations
- Include HTML visualizations in documentation
- Share tree views in team meetings
- Use horizontal views in onboarding materials
- Embed visualizations in project wikis

### Automation Tips
```bash
# Create a Makefile target
visualize:
	mono.exe visualize html --output docs/dependencies.html
	mono.exe visualize tree > docs/dependencies.txt

# Run visualization with analysis
analyze-and-visualize:
	mono.exe analyze
	mono.exe visualize html --output docs/dependencies-$(date +%Y%m%d).html
```

### Integration with Documentation
- Include visualizations in README files
- Add to project wikis
- Use in architectural decision records
- Include in onboarding documentation

By following this guide, you can effectively use MonoMind's visualization features to understand, communicate, and improve your project's architecture.