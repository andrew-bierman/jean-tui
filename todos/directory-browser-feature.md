# Directory Browser Feature Implementation Plan

## Feature Overview
Add a directory browser mode that activates when pressing ESC in the main worktree list, allowing users to navigate the filesystem and switch between different repositories or open terminals in non-git directories.

## Key Requirements (from user feedback)
- **ESC behavior**: Return to original repo when ESC is pressed in directory browser
- **Navigation**: Can enter both git and non-git directories
- **Non-git repos**: Show only root worktree entry (no `.workspaces`), with terminal-only access
- **Hidden dirs**: Filter out hidden directories (starting with `.`)
- **UI**: Full screen modal (consistent with other modals)

## Implementation Steps

### 1. Add Directory Browser State (tui/model.go)
- Add `directoryBrowserModal` to `modalType` enum
- Add new struct type `DirEntry` with fields: `Name`, `Path`, `IsGitRepo`
- Add state fields to `Model` struct:
  - `dirBrowserPath string` - current browsing location
  - `dirBrowserEntries []DirEntry` - directories in current path
  - `dirBrowserIndex int` - selected directory cursor
  - `dirBrowserOriginalRepo string` - repo to return to on ESC
- Add message type `directoriesLoadedMsg` with `entries []DirEntry` and `err error`
- Add command function `loadDirectories(path string) tea.Cmd` that:
  - Lists directories using `os.ReadDir()`
  - Filters out files and hidden directories
  - Checks each directory if it's a git repo (has `.git` subdirectory)
  - Returns `directoriesLoadedMsg` with results

### 2. Add ESC Handler in Main View (tui/update.go)
- In `handleMainInput()`, add case for `"esc"`:
  - Get parent directory of current repo: `filepath.Dir(m.repoPath)`
  - Set `m.modal = directoryBrowserModal`
  - Store original repo path in `m.dirBrowserOriginalRepo`
  - Initialize browser state and trigger `loadDirectories()`

### 3. Add Directory Browser Input Handler (tui/update.go)
- Create `handleDirectoryBrowserInput(msg tea.KeyMsg)` function with:
  - **ESC**: Return to original repo (reload model with `m.dirBrowserOriginalRepo`)
  - **Up/k**: Decrease `dirBrowserIndex` (with bounds checking)
  - **Down/j**: Increase `dirBrowserIndex` (with bounds checking)
  - **Enter**:
    - If git repo: Switch to it (reload model with new path)
    - If non-git directory: Navigate into it (trigger `loadDirectories()` with new path)
  - **Backspace/h/left**: Go up one level to parent directory
  - **t**: Open terminal in selected directory (for non-git repos)
- Add case in `handleModalInput()` to route to directory browser handler
- Add handler for `directoriesLoadedMsg` in `Update()` to update state

### 4. Add Directory Browser Renderer (tui/view.go)
- Create `renderDirectoryBrowser()` function that displays:
  - Header with current path
  - List of directories with indicators:
    - `📁` for regular directories
    - `📂 ✓` for git repositories
    - `..` entry for parent directory (always at top)
  - Highlight selected directory
  - Help bar at bottom: "↑↓ navigate | enter select | backspace up | t terminal | esc back"
- Add case in `renderModal()` to call directory browser renderer
- Use full screen layout (consistent with `branchSelectModal` style)

### 5. Handle Non-Git Repository Logic (tui/model.go and git/worktree.go)
- Modify `NewModel()` to detect non-git repositories:
  - Try to initialize `git.Manager`
  - If it fails (not a git repo), set a flag `m.isNonGitRepo = true`
- Add `isNonGitRepo bool` field to `Model` struct
- Modify `git.Manager` to handle non-git repos gracefully:
  - `ListWorktrees()` returns only root directory entry for non-git repos
  - Other git operations return errors/no-ops for non-git repos
- In main view, when `isNonGitRepo == true`:
  - Show only single "root" entry in worktree list
  - Disable git-related keybindings (n, a, d, r, R, K, b, C, P, p)
  - Enable only: enter (switch), t (terminal), esc (back to browser)

### 6. Update Styles (tui/styles.go if needed)
- Add any new color styles for directory browser elements

## Files to Modify
1. **tui/model.go**: State, messages, commands, DirEntry type
2. **tui/update.go**: ESC handler, directory browser input handler, message handlers
3. **tui/view.go**: Directory browser renderer
4. **git/worktree.go**: Non-git repo handling in Manager
5. **tui/styles.go**: Any new styles (optional)

## Testing Strategy
- Test with git repositories (should work as before)
- Test with non-git directories (should show only root entry, terminal-only)
- Test navigation: up/down directory tree
- Test ESC behavior: should return to original repo
- Test terminal opening in non-git directories
- Test hidden directory filtering

## Edge Cases to Handle
- Permission errors when reading directories
- Empty directories (show only `..` parent entry)
- Root directory (`/`) reached (parent of `/` is `/`)
- Symlinks to directories (follow them)
- Very long directory names (truncate in display)

## User Experience Flow

```
Main View (git repo worktree list)
    ↓ [Press ESC]
Directory Browser Modal (parent directory)
    - Shows all subdirectories (non-hidden)
    - Git repos marked with ✓ indicator
    ↓ [Navigate with ↑↓, backspace to go up]
Select a directory
    ↓ [Press Enter]

    If Git Repo:
        New Model initialized → Main View (new repo's worktree list)

    If Non-Git Directory:
        Option A: Navigate into it (show subdirectories)
        Option B: Press 't' to open terminal only
        → Non-Git View (single root entry, terminal-only mode)

    From Non-Git View:
        ↓ [Press ESC]
        Back to Directory Browser

    From Directory Browser:
        ↓ [Press ESC]
        Return to Original Git Repo
```

## Keybinding Summary

### Directory Browser Modal
- `↑/k` - Move up in directory list
- `↓/j` - Move down in directory list
- `enter` - Select directory (switch to git repo, or navigate into directory)
- `backspace/h/left` - Go up to parent directory
- `t` - Open terminal in selected directory
- `esc` - Return to original repository

### Main View (Non-Git Repository Mode)
- `enter` - Switch to directory (opens browser or tmux)
- `t` - Open terminal in root directory
- `esc` - Return to directory browser
- All git keybindings disabled (n, a, d, r, R, K, b, C, P, p, o)
