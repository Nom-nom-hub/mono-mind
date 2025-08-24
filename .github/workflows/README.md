# MonoMind GitHub Actions Workflows

## Overview
This document describes the GitHub Actions workflows configured for the MonoMind project to automate building, testing, and releasing.

## Workflows

### 1. CI (Continuous Integration)
--File--: `.github/workflows/ci.yml`
--Trigger--: Push or pull request to main/master branches

--Jobs--:
- Build the project
- Run tests with coverage
- Lint code with golangci-lint
- Security scan with Gosec
- Validate documentation files

### 2. Build and Test
--File--: `.github/workflows/build-test.yml`
--Trigger--: Push or pull request to main/master branches

--Jobs--:
- Set up Go environment
- Build the MonoMind binary
- Run all tests

### 3. Code Quality
--File--: `.github/workflows/code-quality.yml`
--Trigger--: Push or pull request to main/master branches

--Jobs--:
- Lint Go code with golangci-lint
- Security scanning with Gosec

### 4. Documentation
--File--: `.github/workflows/documentation.yml`
--Trigger--: Push or pull request to main/master branches

--Jobs--:
- Verify all documentation files exist
- Validate Markdown formatting

### 5. Release
--File--: `.github/workflows/release.yml`
--Trigger--: Push of tags matching pattern `v-`

--Jobs--:
- Set up Go environment
- Extract version from tag
- Build binaries for multiple platforms:
  - Linux (AMD64)
  - Windows (AMD64)
  - macOS (AMD64)
  - macOS (ARM64)
- Create GitHub Release with:
  - Built binaries
  - Release notes
  - Download links

## Configuration Files

### Markdown Lint Configuration
--File--: `.markdownlint.json`
Configures rules for Markdown file validation:
- Line length limits
- Heading styles
- List formatting
- Code block formatting

## Workflow Details

### Build Process
1. Uses `actions/checkout@v3` to checkout code
2. Sets up Go environment with `actions/setup-go@v4`
3. Builds binaries with `go build`
4. Runs tests with `go test`

### Testing
- Verbose test output
- Coverage reporting
- All packages tested

### Code Quality
- Multiple linters through golangci-lint
- Security scanning with Gosec
- Markdown validation

### Release Process
1. Triggers on tag push (e.g., `v1.2.3`)
2. Builds cross-platform binaries
3. Creates GitHub Release with assets
4. Generates release notes
5. Provides download links

## Secrets Required
- `GITHUB_TOKEN`: Automatically provided by GitHub Actions for release creation

## Customization
Workflows can be customized by modifying:
- Go version in setup steps
- Build commands and flags
- Test coverage requirements
- Linting rules
- Release asset inclusion
- Trigger branches

## Monitoring
Workflows can be monitored through:
- GitHub Actions tab in the repository
- Email notifications
- Slack/Teams integrations
- Status badges in README

These workflows ensure that MonoMind maintains high code quality, passes all tests, and provides automated releases.