# MonoMind Changelog

## Table of Contents
1. [Introduction](#introduction)
2. [Changelog Format](#changelog-format)
3. [Versioning](#versioning)
4. [Generating Changelogs](#generating-changelogs)
5. [Manual Changelog Entries](#manual-changelog-entries)
6. [Best Practices](#best-practices)

## Introduction

This document explains how MonoMind handles changelogs and how you can work with them in your projects.

## Changelog Format

MonoMind generates changelogs in Markdown format following the Keep a Changelog standard:

```markdown
# Changelog

## [1.2.0] - 2025-01-15
### Added
- New feature X
- New feature Y

### Changed
- Improved performance of Z
- Updated documentation

### Fixed
- Bug in module A
- Issue with configuration loading

## [1.1.0] - 2025-12-01
### Added
- Initial release
```

### Sections
- **Added**: New features
- **Changed**: Changes in existing functionality
- **Deprecated**: Soon-to-be removed features
- **Removed**: Removed features
- **Fixed**: Bug fixes
- **Security**: Security vulnerabilities

## Versioning

MonoMind follows Semantic Versioning (SemVer):

### Version Format
`MAJOR.MINOR.PATCH`

### Version Bumps
- **MAJOR**: Breaking changes
- **MINOR**: Backward compatible new features
- **PATCH**: Backward compatible bug fixes

### Examples
- `1.0.0` → `1.0.1` (patch)
- `1.0.1` → `1.1.0` (minor)
- `1.1.0` → `2.0.0` (major)

## Generating Changelogs

### Automatic Generation
MonoMind can automatically generate changelogs from Git commit history:

```bash
# Generate changelog for next release
mono.exe release --changelog

# Generate changelog with specific version
mono.exe release --changelog --version 1.2.0
```

### Commit Message Guidelines
For best changelog generation, follow these commit message conventions:

```
feat: Add new authentication module
fix: Resolve issue with database connections
chore: Update dependencies
docs: Improve API documentation
```

### Supported Commit Types
- `feat`: New features
- `fix`: Bug fixes
- `chore`: Maintenance tasks
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Test changes

## Manual Changelog Entries

### Adding Custom Entries
You can add custom entries to your changelog:

```markdown
## [1.2.0] - 2023-01-15
### Added
- New authentication system
- Custom entry: Special holiday feature

### Changed
- Improved API performance
```

### Editing Changelogs
You can manually edit the `CHANGELOG.md` file to:
- Add missing entries
- Correct information
- Reorganize sections
- Add more details

## Best Practices

### Writing Good Changelog Entries
- Use clear, concise language
- Focus on what changed for users
- Include issue numbers when relevant
- Keep entries in chronological order
- Use consistent formatting

### Commit Message Best Practices
- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit first line to 72 characters
- Reference issues and pull requests

### Version Release Best Practices
- Update changelog before tagging releases
- Keep a "Unreleased" section for upcoming changes
- Link to commits or issues when possible
- Include release dates
- Follow consistent formatting

### Example Good Changelog Entry
```markdown
## [1.2.0] - 2023-01-15
### Added
- User authentication system (#123)
- API rate limiting (#145)

### Changed
- Improved database connection handling (#130)
- Updated dependencies to latest versions

### Fixed
- Memory leak in image processing (#135)
- Incorrect error messages in user registration (#140)
```

### Example Poor Changelog Entry
```markdown
## [1.2.0] - 2023-01-15
### Added
- Stuff
- More stuff

### Fixed
- Things
```