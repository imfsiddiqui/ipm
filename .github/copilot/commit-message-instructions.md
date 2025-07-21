# 📝 Commit Message Instructions

All Git commit messages **must** follow the
[Conventional Commits](https://www.conventionalcommits.org/) specification,
**with strict formatting and structure** as defined below.

## ✅ Commit Message Format

### 1. Type Prefix (required, use emoji + type)

Start commit message with one of the following types:

- `🏗️ build` - Changes that affect the build system or external dependencies
- `🔧 chore` - Maintenance tasks not affecting main code
- `👷 ci` - Changes to CI configuration files and scripts
- `📝 docs` - Documentation-only changes
- `✨ feat` - New features
- `🐛 fix` - Bug fixes
- `🚀 perf` - Changes that improve performance
- `♻️ refactor` - Code changes that neither fix bugs nor add features
- `↩️ revert` - Reverting a previous commit
- `🎨 style` - Code style, formatting, missing semicolons, etc.
- `🧪 test` - Adding or updating tests

### 2. Scope (required)

Include a scope that describes what part of the codebase is affected.

**Format:**

```plaintext
type(scope): summary
```

**Example:**

```plaintext
✨ feat(api): add user authentication middleware
```

### 3. Summary Line (required)

- Keep it short and specific
- **Maximum length: 50 characters**
- Should describe **what** was changed, not how

### 4. Detailed Description (required)

After a blank line, write a **detailed explanation** of the change. The body
must:

- Explain **what** was changed
- Explain **why** the change was made
- Explain **how** it was implemented
- Include **edge cases, trade-offs, side effects**, and relevant details
- Mention affected **modules, functions, or dependencies**
- **Line wrapping: 72 characters max per line**

### 5. Footer (optional)

After a blank line, write a **footer** if required. The body must:

- Inform about breaking changes and deprecations (when applicable)
- Reference related issues using `Fixes #<issue-number>` (when applicable)
- Reference pull request that this commit closes `Closes #<pr-number>` or is
  related `Related #<pr-number>` (when applicable)
- **Line wrapping: 72 characters max per line**

### Example Commit Message

```plaintext
🧪 test(api): add edge case tests for null input

Added unit tests in `api/utils.test.js` to validate behavior when
null or undefined inputs are passed to `parseQueryParams`. This
prevents potential runtime exceptions during SSR. Also adjusted
the mock initialization to support legacy test environments.

Fixes #101
```

## 🚫 Non-Compliant Commits

Ensure commits meet all requirements to maintain a clean, consistent, and
automated history. This format should be enforced, non-compliant commits will be
rejected.
