// Package branding provides configurable CLI naming for custom forks.
//
// All values can be overridden at build time using ldflags:
//
//	go build -ldflags "-X github.com/andrew-bierman/jean-tui/internal/branding.CLIName=ralph-tui \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.SessionPrefix=ralph- \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.ConfigDirName=ralph \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.EnvVarPrefix=RALPH"
//
// Example for creating a custom fork named "opencode":
//
//	go build -ldflags "-X github.com/andrew-bierman/jean-tui/internal/branding.CLIName=opencode \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.SessionPrefix=opencode- \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.ConfigDirName=opencode \
//	                   -X github.com/andrew-bierman/jean-tui/internal/branding.EnvVarPrefix=OPENCODE"
package branding

import (
	"fmt"
	"strings"
)

// These variables can be overridden at build time using ldflags.
// See package documentation for examples.
var (
	// CLIName is the name of the CLI command (e.g., "jean", "ralph-tui", "opencode")
	CLIName = "jean"

	// SessionPrefix is the prefix for tmux session names (e.g., "jean-", "ralph-", "opencode-")
	SessionPrefix = "jean-"

	// ConfigDirName is the name of the config directory under ~/.config/ (e.g., "jean", "ralph", "opencode")
	ConfigDirName = "jean"

	// EnvVarPrefix is the prefix for environment variables (e.g., "JEAN", "RALPH", "OPENCODE")
	EnvVarPrefix = "JEAN"

	// AgentCommand is the command to run in the agent window.
	// Default is "claude" which runs Claude CLI with appropriate flags.
	// Set to "" (empty) for a blank terminal - useful for worktree-only mode.
	// Set to any command (e.g., "ralph-tui", "aider", "cursor") to run that instead.
	AgentCommand = "claude"

	// AgentWindowName is the name of the agent window in tmux (e.g., "claude", "agent", "ai")
	AgentWindowName = "claude"
)

// GetEnvVar returns the full environment variable name with the configured prefix.
// Example: GetEnvVar("SWITCH_FILE") returns "JEAN_SWITCH_FILE" by default.
func GetEnvVar(suffix string) string {
	return fmt.Sprintf("%s_%s", EnvVarPrefix, suffix)
}

// GetDebugLogPath returns the debug log path for the CLI.
// Example: returns "/tmp/jean-debug.log" by default.
func GetDebugLogPath() string {
	return fmt.Sprintf("/tmp/%s-debug.log", CLIName)
}

// GetWrapperDebugLogPath returns the wrapper debug log path for the CLI.
// Example: returns "/tmp/jean-wrapper-debug.log" by default.
func GetWrapperDebugLogPath() string {
	return fmt.Sprintf("/tmp/%s-wrapper-debug.log", CLIName)
}

// GetGitDebugLogPath returns the git debug log path for the CLI.
// Example: returns "/tmp/jean-git-debug.log" by default.
func GetGitDebugLogPath() string {
	return fmt.Sprintf("/tmp/%s-git-debug.log", CLIName)
}

// GetTmuxConfigMarkerStart returns the start marker for tmux config.
func GetTmuxConfigMarkerStart() string {
	upper := strings.ToUpper(CLIName)
	return fmt.Sprintf("# === %s_TMUX_CONFIG_START_DO_NOT_MODIFY_THIS_LINE ===", upper)
}

// GetTmuxConfigMarkerEnd returns the end marker for tmux config.
func GetTmuxConfigMarkerEnd() string {
	upper := strings.ToUpper(CLIName)
	return fmt.Sprintf("# === %s_TMUX_CONFIG_END_DO_NOT_MODIFY_THIS_LINE ===", upper)
}

// GetShellWrapperMarkerStart returns the start marker for shell wrapper integration.
func GetShellWrapperMarkerStart() string {
	upper := strings.ToUpper(CLIName)
	return fmt.Sprintf("# BEGIN %s INTEGRATION", upper)
}

// GetShellWrapperMarkerEnd returns the end marker for shell wrapper integration.
func GetShellWrapperMarkerEnd() string {
	upper := strings.ToUpper(CLIName)
	return fmt.Sprintf("# END %s INTEGRATION", upper)
}

// IsAgentEnabled returns true if an agent command is configured.
// Returns false if AgentCommand is empty (blank terminal mode).
func IsAgentEnabled() bool {
	return AgentCommand != ""
}

// IsClaudeAgent returns true if the default Claude agent is configured.
func IsClaudeAgent() bool {
	return AgentCommand == "claude"
}
