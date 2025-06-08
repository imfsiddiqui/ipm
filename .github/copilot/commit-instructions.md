# Commit Instructions

Generate a Git commit message that strictly follows the [Conventional Commits](https://www.conventionalcommits.org/) format.

Start with a type given in the following list, and optionally include a scope in parentheses.

- `🐛 fix`
- `✨ feat`
- `📦 build`
- `🔧 chore`
- `👷 ci`
- `📝 docs`
- `🎨 style`
- `♻️ refactor`
- `🚀 perf`
- `✅ test`

The first line must be a concise summary of the change, no longer than 50 characters.

After a blank line, include an extremely detailed explanation of what was changed, why it was changed, how it was implemented, and any trade-offs, implications, or related considerations.

- Wrap all lines in the detailed section at 72 characters to ensures optimal readability in terminals, diffs, and version control tools.
- Use clear, technical language.
- Include relevant module names, function names, edge cases handled, and any dependencies affected.
- If the change fixes a bug or addresses an issue, mention the issue number (e.g., `Fixes #123`).

Follow best practices for semantic commit messages.
