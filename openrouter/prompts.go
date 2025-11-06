package openrouter

// Default AI prompts for commit messages, branch names, and PR content
// These can be overridden by user-customized prompts in the config

const (
	// DefaultCommitPrompt generates a one-line conventional commit message from git context
	// Placeholders: {status}, {diff}, {branch}, {log}
	DefaultCommitPrompt = `## Context

- Current git status: {status}
- Current git diff (staged and unstaged changes): {diff}
- Current branch: {branch}
- Recent commits: {log}

## Your task

Based on the above changes, create a single git commit with a proper Conventional Commits message and push to git. Never add that this commit is generated with claude code in the end.`

	// DefaultBranchNamePrompt generates a semantic branch name from git diff
	// The {diff} placeholder will be replaced with the actual git diff
	DefaultBranchNamePrompt = `Generate a short, semantic git branch name for these changes.

Return ONLY the branch name (lowercase, kebab-case, max 40 characters). No explanations or markdown.

Examples: fix-login-bug, feat-dark-theme, refactor-api-client

Git diff:
{diff}`

	// DefaultPRPrompt generates a PR title and release notes style description from git diff
	// The {diff} placeholder will be replaced with the actual git diff
	DefaultPRPrompt = `Generate a pull request title and release notes style description for these changes.

Return ONLY valid JSON in this format (no markdown, no extra text):
{"title": "...", "description": "..."}

Requirements:
- title: CRITICAL - MUST be 72 characters or less (hard limit). Present tense, user-friendly summary.
- description: Required. Release notes in markdown format following the structure below.

Description Format (use markdown):
## What's Changed

### Security & Fixes
- Brief user-facing description
- Another fix if applicable

### Improvements
- Enhancement description
- Another improvement if applicable

Important Guidelines:
- Use simple, user-friendly language (no technical jargon)
- Keep each item to ONE short line (max ~80 characters)
- Group changes logically by category
- Only include categories that have relevant changes
- Focus on user-facing benefits, not implementation details
- Skip internal refactoring or minor tweaks unless significant

Example JSON Response:
{"title": "Add dark mode support and improve performance", "description": "## What's Changed\n\n### Improvements\n- New dark mode theme with automatic system preference detection\n- Reduced initial load time by optimizing image loading"}

Git diff:
{diff}`
)

// GetDefaultCommitPrompt returns the default commit message prompt
func GetDefaultCommitPrompt() string {
	return DefaultCommitPrompt
}

// GetDefaultBranchNamePrompt returns the default branch name prompt
func GetDefaultBranchNamePrompt() string {
	return DefaultBranchNamePrompt
}

// GetDefaultPRPrompt returns the default PR content prompt
func GetDefaultPRPrompt() string {
	return DefaultPRPrompt
}
