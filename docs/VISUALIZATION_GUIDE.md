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
MonoMind provides several visualization formats:

### Tree View
- Hierarchical representation of modules and dependencies
- Easy to understand nested relationships
- Ideal for terminal workflows

### ASCII View
- Box-style diagram
- Visual separation of modules
- Good for simple terminal displays

### Horizontal View
- Linear module representation
- Shows relationships in a flow format
- Good for understanding data flow

### HTML View
- Interactive web-based visualization
- Rich formatting and styling
- Shareable in browsers
- Useful for presentations and documentation

## Tree Visualization

### Basic Usage
```bash
# Generate tree visualization
mono.exe visualize tree

# Or simply
mono.exe visualize
````

### Output Example

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
- Reverse dependency tracking (`used by`)
- File type icons

## ASCII Visualization

### Basic Usage

```bash
mono.exe visualize ascii
```

### Output Example

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
mono.exe visualize horizontal
```

### Output Example

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
- Good for sequential relationships

## HTML Visualization

### Basic Usage

```bash
mono.exe visualize html --output dependencies.html
```

### Features

- Interactive web page
- Professional CSS styling
- Hover effects and visual enhancements
- Shareable and viewable in any browser
- Printable and presentable
- Embeddable in documentation

### HTML Output Structure

```html
<!DOCTYPE html>
<html>
<head>
    <title>MonoMind Dependency Graph</title>
    <style>
        /- Professional CSS styling -/
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

Modify the HTML template in the source code or post-process the generated HTML.

## Customizing Visualizations

### Output Files

```bash
mono.exe visualize html --output my-project-dependencies.html
mono.exe visualize html --output ./docs/dependencies.html
```

### Combining with Other Commands

```bash
mono.exe analyze && mono.exe visualize
mono.exe visualize tree > tree.txt
mono.exe visualize ascii > ascii.txt
mono.exe visualize horizontal > horizontal.txt
mono.exe visualize html --output graph.html
```

### Automation Example

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

- --Dependencies--: Modules this module depends on
- --Dependents--: Modules that depend on this module
- --Circular Dependencies--: Avoid when possible
- --Orphan Modules--: Modules with no dependencies or dependents

### Identifying Issues

- Too many dependencies → Module may be overloaded
- Too many dependents → Module may be a bottleneck
- Circular dependencies → Can break build/test processes
- Orphan modules → May be unused or poorly integrated

### Making Decisions

Use visualizations to:

- Plan refactorings
- Identify optimization opportunities
- Understand change impact
- Communicate architecture to the team

## Best Practices

### When to Use Each Visualization

#### Tree View

- Best for hierarchical relationships
- Terminal-based workflows

#### ASCII View

- Best for simple boxes and arrows
- Quick terminal viewing

#### Horizontal View

- Best for linear data flow
- Documentation and analysis

#### HTML View

- Best for professional presentations
- Documentation, reports, and team sharing

### Regular Visualization Updates

```bash
mono.exe analyze
mono.exe visualize html --output docs/latest-dependencies.html
```

### Version Control

- Commit visualizations for key releases
- Compare outputs over time
- Include in architectural docs

### Sharing Visualizations

- Embed HTML visualizations in documentation
- Present tree views in meetings
- Use horizontal views for onboarding
- Include visualizations in project wikis

### Automation Tips

```bash
# Makefile example
visualize:
	mono.exe visualize html --output docs/dependencies.html
	mono.exe visualize tree > docs/dependencies.txt

analyze-and-visualize:
	mono.exe analyze
	mono.exe visualize html --output docs/dependencies-$(date +%Y%m%d).html
```

### Integration with Documentation

- Include visualizations in README files
- Add to project wikis
- Use in architectural decision records
- Include in onboarding documentation