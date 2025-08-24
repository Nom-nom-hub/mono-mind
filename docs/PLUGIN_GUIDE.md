# MonoMind Plugin Development Guide

## Table of Contents

1. [Introduction](#introduction)  
2. [Plugin System Overview](#plugin-system-overview)  
3. [Plugin Hooks](#plugin-hooks)  
4. [Creating Plugins](#creating-plugins)  
5. [Plugin Types](#plugin-types)  
6. [Plugin Configuration](#plugin-configuration)  
7. [Best Practices](#best-practices)  
8. [Examples](#examples)  
9. [Distribution](#distribution)  

## Introduction

This guide explains how to create and use plugins with MonoMind. Plugins allow you to extend MonoMind's functionality without modifying the core codebase.

## Plugin System Overview

### How Plugins Work

1. Plugins are external programs or scripts  
2. They are executed at specific hooks during MonoMind operations  
3. Plugins receive context information through environment variables  
4. Plugins can modify the environment or filesystem  
5. Plugin output is logged by MonoMind  

### Plugin Directory

By default, MonoMind looks for plugins in the `plugins/` directory in your project root.

### Plugin Discovery

- Plugins are automatically discovered based on filename  
- Filename format: `hookname.extension`  
- Example: `pre-build.sh`, `post-test.py`  

## Plugin Hooks

### Available Hooks

- `pre-analyze`: Before repository analysis
- `post-analyze`: After repository analysis
- `pre-build`: Before builds
- `post-build`: After builds
- `pre-test`: Before tests
- `post-test`: After tests
- `pre-refactor`: Before refactoring
- `post-refactor`: After refactoring
- `pre-release`: Before releases
- `post-release`: After releases

### Hook Execution Order

1. `pre-*` hooks execute before the main operation
2. Main operation executes
3. `post-*` hooks execute after the main operation

### Hook Context

Plugins receive context information through:

- Environment variables  
- Command line arguments  
- Standard input (for complex data)  

## Creating Plugins

### Basic Plugin Structure

```bash
#!/bin/bash
# plugins/pre-build.sh

# Log plugin execution
echo "[$(date)] Pre-build plugin executing"

# Access environment variables
echo "Working directory: $PWD"
echo "MonoMind version: $MONO_VERSION"

# Perform plugin logic
# ...

# Exit with appropriate code
exit 0
````

### Required Elements

1. **Shebang** (for script plugins): `#!/bin/bash`, `#!/usr/bin/env python3`, etc.
2. **Executable permissions**: `chmod +x plugin-file`
3. **Proper exit codes**: 0 for success, non-zero for errors
4. **Error handling**: Handle errors gracefully

### Environment Variables

Plugins have access to these environment variables:

- `MONO_VERSION`: Current MonoMind version
- `MONO_HOOK`: Current hook name
- `MONO_WORKDIR`: Working directory
- `MONO_CONFIG`: Path to config file

## Plugin Types

### Shell Scripts

```bash
#!/bin/bash
# plugins/post-build.sh

echo "Build completed at $(date)"
```

### Python Scripts

```python
#!/usr/bin/env python3
# plugins/pre-test.py

import os
import sys
from datetime import datetime

print(f"[{datetime.now()}] Pre-test plugin executing")
print(f"Working directory: {os.getcwd()}")

# Your plugin logic here
```

### JavaScript Files

```javascript
#!/usr/bin/env node
// plugins/post-release.js

console.log(`[${new Date().toISOString()}] Post-release plugin executing`);
console.log(`Working directory: ${process.cwd()}`);

// Your plugin logic here
```

### Compiled Binaries

Any executable binary can be used as a plugin:

```go
// plugins/custom-plugin.go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Custom plugin executing")
    fmt.Printf("Working directory: %s\n", os.Getenv("PWD"))
}
```

Compile with: `go build -o plugins/custom-plugin plugins/custom-plugin.go`

## Plugin Configuration

### Plugin Directory Configuration

You can configure the plugin directory in your config file:

```yaml
plugins:
  directory: "./custom-plugins"
```

### Specific Plugin Configuration

Configure individual plugins:

```yaml
plugins:
  pre-build:
    - "./scripts/prebuild.sh"
    - "./scripts/validate.sh"
  post-build:
    - "./scripts/postbuild.sh"
```

### Plugin Parameters

Pass parameters to plugins through environment variables:

```yaml
plugins:
  pre-build:
    - path: "./scripts/prebuild.sh"
      env:
        BUILD_TARGET: "production"
        VERBOSE: "true"
```

## Best Practices

### Error Handling

```bash
#!/bin/bash
# plugins/pre-build.sh

set -e  # Exit on error

# Check prerequisites
if ! command -v npm &> /dev/null; then
    echo "Error: npm is not installed"
    exit 1
fi

# Your plugin logic
echo "Pre-build checks passed"
```

### Logging

```python
#!/usr/bin/env python3
# plugins/post-test.py

import sys
from datetime import datetime

def log(message, level="INFO"):
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    print(f"[{timestamp}] [{level}] {message}")
    sys.stdout.flush()

log("Post-test plugin starting")
# Your plugin logic here
log("Post-test plugin completed")
```

### Resource Management

- Clean up temporary files
- Limit resource usage
- Handle timeouts gracefully
- Use non-blocking operations when possible

### Security

- Validate input
- Avoid executing untrusted code
- Limit file system access
- Use secure permissions

## Examples

### Pre-Build Validation Plugin

```bash
#!/bin/bash
# plugins/pre-build-validate.sh

echo "Validating build environment..."

# Check Node.js version
if ! command -v node &> /dev/null; then
    echo "Error: Node.js is not installed"
    exit 1
fi

NODE_VERSION=$(node --version)
echo "Node.js version: $NODE_VERSION"

# Check required files
if [ ! -f "package.json" ]; then
    echo "Error: package.json not found"
    exit 1
fi

echo "Build environment validation passed"
```

### Post-Test Reporting Plugin

```python
#!/usr/bin/env python3
# plugins/post-test-report.py

import json
import os
import sys

def main():
    print("Generating test report...")
    
    # Read test results (example)
    test_results = {
        "tests_run": 42,
        "tests_passed": 40,
        "tests_failed": 2,
        "coverage": "85%"
    }
    
    # Write report
    with open("test-report.json", "w") as f:
        json.dump(test_results, f, indent=2)
    
    print(f"Test report generated: {test_results['tests_passed']}/{test_results['tests_run']} tests passed")

if __name__ == "__main__":
    main()
```

### Pre-Release Version Check Plugin

```javascript
#!/usr/bin/env node
// plugins/pre-release-check.js

const fs = require('fs');
const { execSync } = require('child_process');

console.log('Checking release readiness...');

// Check if working directory is clean
try {
    const status = execSync('git status --porcelain', { encoding: 'utf-8' });
    if (status.trim() !== '') {
        console.error('Error: Working directory is not clean');
        process.exit(1);
    }
} catch (error) {
    console.error('Error checking git status:', error.message);
    process.exit(1);
}

console.log('Release check passed');
```

## Distribution

### Sharing Plugins

1. **GitHub Gists**: Share single-file plugins
2. **GitHub Repositories**: Share complex plugin collections
3. **Package Managers**: Distribute through npm, PyPI, etc.
4. **Direct Download**: Host plugins on your website

### Plugin Collections

Create collections of related plugins:

```
my-monomind-plugins/
├── build/
│   ├── pre-build-validate.sh
│   └── post-build-notify.sh
├── test/
│   ├── pre-test-setup.sh
│   └── post-test-report.py
└── release/
    ├── pre-release-check.js
    └── post-release-announce.sh
```

### Installation Script

Provide an installation script for your plugins:

```bash
#!/bin/bash
# install-plugins.sh

mkdir -p plugins
curl -o plugins/pre-build.sh https://example.com/plugins/pre-build.sh
curl -o plugins/post-build.sh https://example.com/plugins/post-build.sh
chmod +x plugins/*.sh

echo "Plugins installed successfully"
```

### Documentation

When distributing plugins, include:

- README with usage instructions
- Example configuration
- Troubleshooting guide
- License information