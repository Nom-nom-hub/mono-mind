# MonoMind Documentation

This directory contains the complete documentation suite for MonoMind, optimized for GitHub Pages with professional styling.

## 📚 Documentation Structure

### User Documentation
- --[QUICK_START.md](./QUICK_START.md)-- – Fast path to productivity
- --[USER_GUIDE.md](./USER_GUIDE.md)-- – Complete reference for all features
- --[VISUALIZATION_GUIDE.md](./VISUALIZATION_GUIDE.md)-- – Understanding dependency graphs
- --[PLUGIN_GUIDE.md](./PLUGIN_GUIDE.md)-- – Creating and using plugins
- --[CHANGELOG_GUIDE.md](./CHANGELOG_GUIDE.md)-- – Working with releases
- --[TROUBLESHOOTING.md](./TROUBLESHOOTING.md)-- – Solutions to common problems

### Developer Resources
- --[API.md](./API.md)-- – Technical API reference
- --[DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md)-- – Contributing to MonoMind
- --[DOCUMENTATION_SUMMARY.md](./DOCUMENTATION_SUMMARY.md)-- – Documentation overview

## 🎨 GitHub Pages Styling

This documentation is styled for professional presentation on GitHub Pages.

### Files
- --`style.css`-- – Custom CSS with professional theming
- --`index.html`-- – Main documentation landing page
- --`_layouts/default.html`-- – Jekyll layout template
- --`_config.yml`-- – GitHub Pages configuration

### Features
- --Responsive Design-- – Works on desktop and mobile
- --Syntax Highlighting-- – Proper highlighting for code blocks
- --Professional Typography-- – Clean, readable fonts
- --Navigation-- – Easy navigation between documents
- --Print Styles-- – Optimized for printing

## 🚀 Setting Up GitHub Pages

1. --Enable GitHub Pages-- in repository settings.
2. --Choose source--: Deploy from a branch.
3. --Select branch--: `main` or your default branch.
4. --Set folder--: `/docs` (if using the `docs` folder) or `/` (root).

### Alternative: Root-level Deployment
If deploying from the repository root, move these files:
```bash
# Move to repository root
mv docs/index.html ./
mv docs/style.css ./
mv docs/_layouts ./  # (if using Jekyll)
````

## 📝 Contributing to Documentation

### Writing Guidelines

- Use clear, concise language.
- Include practical examples.
- Keep line length under 120 characters.
- Maintain proper heading hierarchy.
- Test all commands and examples.

### Formatting Standards

- All Markdown files pass `markdownlint` checks.
- Headings have proper spacing.
- Code blocks specify language.
- Tables are properly formatted.
- Links are functional.

### Local Testing

```bash
# Install markdownlint-cli
npm install -g markdownlint-cli

# Check documentation
markdownlint docs/
```

## 🔧 Customization

### Colors

Edit `docs/style.css` to customize theme:

```css
:root {
  --primary-color: #2c3e50;    /- Main headings -/
  --secondary-color: #3498db;  /- Links and accents -/
  --accent-color: #e74c3c;     /- Code keywords -/
}
```

### Layout

Modify `docs/_layouts/default.html` for:

- Navigation structure
- Header/footer content
- Additional CSS/JS includes

## 📄 License

This documentation is part of MonoMind and follows the same license terms.

---

--Need help?-- Check the [troubleshooting guide](./TROUBLESHOOTING.md) or [open an issue](https://github.com/nom-nom-hub/mono-mind/issues).