package branding

import (
	"strings"
	"testing"
)

// TestDefaultValues tests that default branding values are set correctly
func TestDefaultValues(t *testing.T) {
	if CLIName != "jean" {
		t.Errorf("Expected CLIName 'jean', got '%s'", CLIName)
	}

	if SessionPrefix != "jean-" {
		t.Errorf("Expected SessionPrefix 'jean-', got '%s'", SessionPrefix)
	}

	if ConfigDirName != "jean" {
		t.Errorf("Expected ConfigDirName 'jean', got '%s'", ConfigDirName)
	}

	if EnvVarPrefix != "JEAN" {
		t.Errorf("Expected EnvVarPrefix 'JEAN', got '%s'", EnvVarPrefix)
	}
}

// TestGetEnvVar tests the GetEnvVar function
func TestGetEnvVar(t *testing.T) {
	tests := []struct {
		suffix   string
		expected string
	}{
		{"SWITCH_FILE", "JEAN_SWITCH_FILE"},
		{"INIT_ATTEMPTED", "JEAN_INIT_ATTEMPTED"},
		{"WORKSPACE_PATH", "JEAN_WORKSPACE_PATH"},
		{"ROOT_PATH", "JEAN_ROOT_PATH"},
	}

	for _, tt := range tests {
		result := GetEnvVar(tt.suffix)
		if result != tt.expected {
			t.Errorf("GetEnvVar(%q) = %q, want %q", tt.suffix, result, tt.expected)
		}
	}
}

// TestGetDebugLogPath tests the GetDebugLogPath function
func TestGetDebugLogPath(t *testing.T) {
	path := GetDebugLogPath()
	if path != "/tmp/jean-debug.log" {
		t.Errorf("Expected '/tmp/jean-debug.log', got '%s'", path)
	}
}

// TestGetWrapperDebugLogPath tests the GetWrapperDebugLogPath function
func TestGetWrapperDebugLogPath(t *testing.T) {
	path := GetWrapperDebugLogPath()
	if path != "/tmp/jean-wrapper-debug.log" {
		t.Errorf("Expected '/tmp/jean-wrapper-debug.log', got '%s'", path)
	}
}

// TestGetGitDebugLogPath tests the GetGitDebugLogPath function
func TestGetGitDebugLogPath(t *testing.T) {
	path := GetGitDebugLogPath()
	if path != "/tmp/jean-git-debug.log" {
		t.Errorf("Expected '/tmp/jean-git-debug.log', got '%s'", path)
	}
}

// TestGetTmuxConfigMarkerStart tests the GetTmuxConfigMarkerStart function
func TestGetTmuxConfigMarkerStart(t *testing.T) {
	marker := GetTmuxConfigMarkerStart()
	expected := "# === JEAN_TMUX_CONFIG_START_DO_NOT_MODIFY_THIS_LINE ==="
	if marker != expected {
		t.Errorf("Expected '%s', got '%s'", expected, marker)
	}
}

// TestGetTmuxConfigMarkerEnd tests the GetTmuxConfigMarkerEnd function
func TestGetTmuxConfigMarkerEnd(t *testing.T) {
	marker := GetTmuxConfigMarkerEnd()
	expected := "# === JEAN_TMUX_CONFIG_END_DO_NOT_MODIFY_THIS_LINE ==="
	if marker != expected {
		t.Errorf("Expected '%s', got '%s'", expected, marker)
	}
}

// TestGetShellWrapperMarkerStart tests the GetShellWrapperMarkerStart function
func TestGetShellWrapperMarkerStart(t *testing.T) {
	marker := GetShellWrapperMarkerStart()
	expected := "# BEGIN JEAN INTEGRATION"
	if marker != expected {
		t.Errorf("Expected '%s', got '%s'", expected, marker)
	}
}

// TestGetShellWrapperMarkerEnd tests the GetShellWrapperMarkerEnd function
func TestGetShellWrapperMarkerEnd(t *testing.T) {
	marker := GetShellWrapperMarkerEnd()
	expected := "# END JEAN INTEGRATION"
	if marker != expected {
		t.Errorf("Expected '%s', got '%s'", expected, marker)
	}
}

// TestMarkersContainCLIName tests that markers properly use CLIName
func TestMarkersContainCLIName(t *testing.T) {
	// All markers should contain the CLI name in uppercase
	upperName := strings.ToUpper(CLIName)

	startMarker := GetTmuxConfigMarkerStart()
	if !strings.Contains(startMarker, upperName) {
		t.Errorf("Tmux start marker should contain '%s', got '%s'", upperName, startMarker)
	}

	endMarker := GetTmuxConfigMarkerEnd()
	if !strings.Contains(endMarker, upperName) {
		t.Errorf("Tmux end marker should contain '%s', got '%s'", upperName, endMarker)
	}

	shellStart := GetShellWrapperMarkerStart()
	if !strings.Contains(shellStart, upperName) {
		t.Errorf("Shell wrapper start marker should contain '%s', got '%s'", upperName, shellStart)
	}

	shellEnd := GetShellWrapperMarkerEnd()
	if !strings.Contains(shellEnd, upperName) {
		t.Errorf("Shell wrapper end marker should contain '%s', got '%s'", upperName, shellEnd)
	}
}
