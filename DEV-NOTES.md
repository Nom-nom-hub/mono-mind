Got it. You want a **master instruction prompt** that tells your AI exactly how to build MonoMind in Go, from scratch, autonomously. Here’s a **complete, actionable instruction set** your AI can follow:

---

# Master AI Instructions – MonoMind CLI (Go)

**Objective:** Build a CLI-first AI agent called **MonoMind** that manages monorepos: parses repos, analyzes dependencies, predicts change impact, orchestrates builds/tests, handles refactors, and manages releases.

---

## 1. Setup Environment

1. Initialize a **Go module**:

   ```bash
   go mod init mono-mind
   ```

2. Install dependencies:

   * Cobra for CLI: `go get github.com/spf13/cobra@latest`
   * Tree-sitter Go bindings for AST parsing (for multi-language support)
   * Logging package (`log` or `zap`)

3. Create project folder structure:

   ```
   cmd/mono/
   internal/analyzer/
   internal/impact/
   internal/build/
   internal/refactor/
   internal/release/
   internal/plugins/
   internal/logger/
   configs/
   scripts/
   ```

---

## 2. Implement CLI

* Use **Cobra** to define root command `mono`.

* Subcommands:

  * `analyze` → parse repo and build dependency graph
  * `impact <file>` → show affected modules/tests
  * `build` → incremental build of impacted modules
  * `refactor` → rename/move functions/modules
  * `release` → version bump, changelog, publish packages

* Each CLI command should call the corresponding **internal module**, not contain logic itself.

---

## 3. Build Core Modules

### 3.1 Analyzer Module

* Recursively scan repo directories.
* Detect modules/packages per language.
* Parse imports/dependencies using **Tree-sitter**.
* Build **directed dependency graph**.
* Store metadata: module name, path, language, dependencies, last modified timestamp.

### 3.2 Impact Module

* Input: changed file/module
* Traverse dependency graph downstream to find all affected modules and tests.
* Output CLI-friendly report or JSON for pipelines.

### 3.3 Build Module

* Rebuild only affected modules.
* Execute builds/tests in parallel with Goroutines.
* Support multiple languages (start with JS or Python).

### 3.4 Refactor Module

* Parse AST using Tree-sitter.
* Support safe rename/move operations.
* Update all references in repo automatically.
* Include **dry-run mode** to preview changes.

### 3.5 Release Module

* Version bump (major/minor/patch)
* Auto-generate changelogs from commits
* Optional: publish to npm, PyPI, Maven, or other package registry.

### 3.6 Plugin Manager

* Load pre/post hooks from plugin folder.
* Execute plugins during build, test, refactor, or release steps.
* Support scripts or compiled binaries.

### 3.7 Logger Module

* Structured logging of all actions: command, status, duration, errors.
* Track metrics like build time, test coverage, refactor success.

---

## 4. Config Files

* `configs/mono.yaml`:

  ```yaml
  modules:
    - name: api
      path: ./src/api
      language: js
    - name: utils
      path: ./src/utils
      language: js

  tests:
    - name: api-tests
      path: ./tests/api
      modules: [api, utils]

  plugins:
    pre-build: ./plugins/pre_build.sh
    post-build: ./plugins/post_build.sh
  ```

---

## 5. Development Guidelines

* Start **single language support** (JS/Python) → expand later.
* Maintain **modular architecture**: CLI calls internal modules.
* Use **concurrency** for file scanning, builds, and tests.
* Implement caching for dependency graph to improve performance.
* Include **dry-run** options for destructive commands.
* Provide **ASCII terminal graphs** of dependencies, optional HTML export.

---

## 6. Recommended Command Usage

```bash
mono analyze          # scan repo and build graph
mono impact src/api/user.js  # show affected modules/tests
mono build            # incremental build of affected modules
mono refactor --rename oldFunc newFunc --dry-run
mono release --bump minor --changelog
```

---

## 7. Deliverables for AI

* Fully working CLI binary
* Modular Go code with internal packages
* Dependency graph + impact analysis working for at least one language
* Incremental build/test orchestration
* Refactor engine with dry-run support
* Release manager with version bump/changelog
* Plugin system with pre/post hooks
* Logging and metrics

