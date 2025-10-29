package tui

import "github.com/charmbracelet/lipgloss"

// Theme color variables - initialized with Matrix theme, updated by ApplyTheme()
var (
	// Colors - initialized with Matrix theme
	primaryColor   = lipgloss.Color("#00FF41") // Bright Matrix green
	secondaryColor = lipgloss.Color("#008F11") // Medium green
	accentColor    = lipgloss.Color("#00FF41") // Bright Matrix green for highlights
	warningColor   = lipgloss.Color("#AAFF00") // Yellow-green for warnings
	successColor   = lipgloss.Color("#00FF41") // Bright green for success
	errorColor     = lipgloss.Color("#FF0000") // Red for errors
	mutedColor     = lipgloss.Color("#00AA00") // Medium green for muted text
	bgColor        = lipgloss.Color("#000000") // Pure black background
	fgColor        = lipgloss.Color("#00FF41") // Bright green text
)

// Style variables - mutable, rebuilt by ApplyTheme()
var (
	// Base styles
	baseStyle lipgloss.Style

	// Panel styles
	panelStyle       lipgloss.Style
	activePanelStyle lipgloss.Style

	titleStyle lipgloss.Style

	// List item styles
	selectedItemStyle    lipgloss.Style
	normalItemStyle      lipgloss.Style
	currentWorktreeStyle lipgloss.Style

	// Detail styles
	detailKeyStyle   lipgloss.Style
	detailValueStyle lipgloss.Style

	// Help/Status bar
	helpStyle   lipgloss.Style
	statusStyle lipgloss.Style
	errorStyle  lipgloss.Style

	// Modal styles
	modalStyle       lipgloss.Style
	modalTitleStyle  lipgloss.Style
	inputLabelStyle  lipgloss.Style

	buttonStyle              lipgloss.Style
	selectedButtonStyle      lipgloss.Style
	cancelButtonStyle        lipgloss.Style
	selectedCancelButtonStyle lipgloss.Style

	// Delete button styles (red for danger)
	deleteButtonStyle         lipgloss.Style
	selectedDeleteButtonStyle  lipgloss.Style
	disabledButtonStyle        lipgloss.Style

	// Notification styles
	successNotifStyle lipgloss.Style
	errorNotifStyle   lipgloss.Style
	warningNotifStyle lipgloss.Style
	infoNotifStyle    lipgloss.Style
)

// InitStyles initializes styles with the Matrix theme on startup
func InitStyles() {
	// Apply the default Matrix theme
	ApplyTheme("matrix")
}
