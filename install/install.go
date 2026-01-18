package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/coollabsio/jean-tui/config"
	"github.com/coollabsio/jean-tui/internal/branding"
)

// Shell represents a shell type
type Shell string

const (
	Bash Shell = "bash"
	Zsh  Shell = "zsh"
	Fish Shell = "fish"
)

// Detector detects the user's shell and manages installation
type Detector struct {
	Shell   Shell
	RCFile  string
	HomeDir string
}

// NewDetector creates a new shell detector
func NewDetector() (*Detector, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	shell := detectShell()
	rcFile := getRCFile(shell, homeDir)

	return &Detector{
		Shell:   shell,
		RCFile:  rcFile,
		HomeDir: homeDir,
	}, nil
}

// detectShell detects the user's shell from the SHELL environment variable
func detectShell() Shell {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return Bash // Default to bash
	}

	// Extract just the shell name from the full path
	shellName := filepath.Base(shellPath)

	switch shellName {
	case "zsh":
		return Zsh
	case "fish":
		return Fish
	case "bash":
		fallthrough
	default:
		return Bash
	}
}

// getRCFile returns the appropriate rc file path for the given shell
func getRCFile(shell Shell, homeDir string) string {
	return GetRCFileForShell(shell, homeDir)
}

// GetRCFileForShell returns the appropriate rc file path for the given shell
func GetRCFileForShell(shell Shell, homeDir string) string {
	switch shell {
	case Zsh:
		return filepath.Join(homeDir, ".zshrc")
	case Fish:
		return filepath.Join(homeDir, ".config/fish/config.fish")
	case Bash:
		fallthrough
	default:
		return filepath.Join(homeDir, ".bashrc")
	}
}

// GetWrapper returns the appropriate wrapper code for the detected shell
func (d *Detector) GetWrapper() string {
	switch d.Shell {
	case Fish:
		return GetFishWrapper()
	case Zsh:
		fallthrough
	case Bash:
		fallthrough
	default:
		return GetBashZshWrapper()
	}
}

// IsInstalled checks if the integration is already in the rc file
func (d *Detector) IsInstalled() bool {
	content, err := os.ReadFile(d.RCFile)
	if err != nil {
		return false
	}

	return ContainsWrapperMarker(string(content))
}

// Install adds the wrapper to the rc file
func (d *Detector) Install(dryRun bool) error {
	// Check if already installed
	if d.IsInstalled() {
		return fmt.Errorf("%s integration is already installed in %s", branding.CLIName, d.RCFile)
	}

	// Create directory if it doesn't exist (for fish)
	rcDir := filepath.Dir(d.RCFile)
	if err := os.MkdirAll(rcDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", rcDir, err)
	}

	wrapper := d.GetWrapper()

	if dryRun {
		fmt.Printf("Would install %s to %s\n", branding.CLIName, d.RCFile)
		fmt.Printf("Content to be added:\n%s\n", wrapper)
		return nil
	}

	// Create backup
	if _, err := os.Stat(d.RCFile); err == nil {
		backupFile := d.RCFile + ".backup"
		if err := copyFile(d.RCFile, backupFile); err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}
		fmt.Printf("Backup created: %s\n", backupFile)
	}

	// Read existing content
	content, _ := os.ReadFile(d.RCFile)
	existingContent := string(content)

	// Append wrapper to rc file
	newContent := existingContent
	if !strings.HasSuffix(newContent, "\n") && newContent != "" {
		newContent += "\n"
	}
	newContent += "\n" + wrapper + "\n"

	if err := os.WriteFile(d.RCFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write to %s: %w", d.RCFile, err)
	}

	fmt.Printf("✓ %s integration installed to %s\n", branding.CLIName, d.RCFile)
	fmt.Printf("Run: source %s (or restart your terminal)\n", d.RCFile)

	return nil
}

// Update updates an existing installation
func (d *Detector) Update(dryRun bool) error {
	// Check if already installed
	if !d.IsInstalled() {
		return fmt.Errorf("%s integration is not installed in %s. Run '%s init' to install.", branding.CLIName, d.RCFile, branding.CLIName)
	}

	wrapper := d.GetWrapper()

	if dryRun {
		fmt.Printf("Would update %s in %s\n", branding.CLIName, d.RCFile)
		fmt.Printf("New content:\n%s\n", wrapper)
		return nil
	}

	// Read existing content
	content, err := os.ReadFile(d.RCFile)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", d.RCFile, err)
	}

	contentStr := string(content)

	// Remove old integration
	startMarker := GetWrapperStartMarker()
	endMarker := GetWrapperEndMarker()

	startIdx := strings.Index(contentStr, startMarker)
	endIdx := strings.Index(contentStr, endMarker)

	if startIdx == -1 || endIdx == -1 {
		return fmt.Errorf("could not find %s integration markers in %s", branding.CLIName, d.RCFile)
	}

	// Remove everything from start marker to end marker (inclusive)
	newContent := contentStr[:startIdx] + wrapper + "\n" + contentStr[endIdx+len(endMarker):]

	// Clean up extra newlines
	for strings.Contains(newContent, "\n\n\n") {
		newContent = strings.ReplaceAll(newContent, "\n\n\n", "\n\n")
	}

	if err := os.WriteFile(d.RCFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write to %s: %w", d.RCFile, err)
	}

	fmt.Printf("✓ %s integration updated in %s\n", branding.CLIName, d.RCFile)
	fmt.Printf("Run: source %s (or restart your terminal)\n", d.RCFile)

	return nil
}

// Remove removes the integration from the rc file
func (d *Detector) Remove(dryRun bool) error {
	// Check if installed
	if !d.IsInstalled() {
		return fmt.Errorf("%s integration is not installed in %s", branding.CLIName, d.RCFile)
	}

	if dryRun {
		fmt.Printf("Would remove %s from %s\n", branding.CLIName, d.RCFile)
		return nil
	}

	// Read existing content
	content, err := os.ReadFile(d.RCFile)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", d.RCFile, err)
	}

	contentStr := string(content)

	// Remove integration block
	startMarker := GetWrapperStartMarker()
	endMarker := GetWrapperEndMarker()

	startIdx := strings.Index(contentStr, startMarker)
	endIdx := strings.Index(contentStr, endMarker)

	if startIdx == -1 || endIdx == -1 {
		return fmt.Errorf("could not find %s integration markers in %s", branding.CLIName, d.RCFile)
	}

	// Remove the block and the preceding newline
	newContent := contentStr[:startIdx] + contentStr[endIdx+len(endMarker):]

	// Clean up extra newlines
	for strings.Contains(newContent, "\n\n\n") {
		newContent = strings.ReplaceAll(newContent, "\n\n\n", "\n\n")
	}

	if err := os.WriteFile(d.RCFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write to %s: %w", d.RCFile, err)
	}

	fmt.Printf("✓ %s integration removed from %s\n", branding.CLIName, d.RCFile)

	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, content, 0644)
}

// NeedsUpdate checks if the wrapper needs to be updated
// Returns true if:
// - Wrapper is not installed, OR
// - Installed wrapper checksum differs from current template checksum
func (d *Detector) NeedsUpdate(cfg *config.Manager) bool {
	// Not installed = needs update (installation)
	if !d.IsInstalled() {
		return true
	}

	// Calculate current template checksum
	currentChecksum := CalculateWrapperChecksum(d.Shell)

	// Get stored checksum from config
	storedChecksum := cfg.GetWrapperChecksum(string(d.Shell))

	// If no stored checksum or different checksum, needs update
	return storedChecksum == "" || storedChecksum != currentChecksum
}

// AutoUpdate performs an automatic silent update of the wrapper
// This is called on every startup to ensure wrapper is up-to-date
// Returns error only for critical failures; non-critical errors are logged
func (d *Detector) AutoUpdate(cfg *config.Manager) error {
	wrapper := d.GetWrapper()
	currentChecksum := CalculateWrapperChecksum(d.Shell)

	// If not installed, perform installation
	if !d.IsInstalled() {
		// Create directory if it doesn't exist (for fish)
		rcDir := filepath.Dir(d.RCFile)
		if err := os.MkdirAll(rcDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", rcDir, err)
		}

		// Read existing content
		content, _ := os.ReadFile(d.RCFile)
		existingContent := string(content)

		// Append wrapper to rc file
		newContent := existingContent
		if !strings.HasSuffix(newContent, "\n") && newContent != "" {
			newContent += "\n"
		}
		newContent += "\n" + wrapper + "\n"

		if err := os.WriteFile(d.RCFile, []byte(newContent), 0644); err != nil {
			return fmt.Errorf("failed to write to %s: %w", d.RCFile, err)
		}

		// Store checksum
		if err := cfg.SetWrapperChecksum(string(d.Shell), currentChecksum); err != nil {
			// Non-critical: wrapper is installed, just checksum storage failed
			return nil
		}

		return nil
	}

	// Wrapper is installed, perform update
	content, err := os.ReadFile(d.RCFile)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", d.RCFile, err)
	}

	contentStr := string(content)

	// Remove old integration
	startMarker := GetWrapperStartMarker()
	endMarker := GetWrapperEndMarker()

	startIdx := strings.Index(contentStr, startMarker)
	endIdx := strings.Index(contentStr, endMarker)

	if startIdx == -1 || endIdx == -1 {
		return fmt.Errorf("could not find %s integration markers in %s", branding.CLIName, d.RCFile)
	}

	// Remove everything from start marker to end marker (inclusive)
	newContent := contentStr[:startIdx] + wrapper + "\n" + contentStr[endIdx+len(endMarker):]

	// Clean up extra newlines
	for strings.Contains(newContent, "\n\n\n") {
		newContent = strings.ReplaceAll(newContent, "\n\n\n", "\n\n")
	}

	if err := os.WriteFile(d.RCFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write to %s: %w", d.RCFile, err)
	}

	// Store new checksum
	if err := cfg.SetWrapperChecksum(string(d.Shell), currentChecksum); err != nil {
		// Non-critical: wrapper is updated, just checksum storage failed
		return nil
	}

	return nil
}
