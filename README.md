# jean - AI-Powered Git Worktree TUI with Claude Code Support

## What is jean?

jean is a terminal user interface (TUI) that makes working with Git worktrees effortless. Instead of managing multiple branches and worktrees manually, jean gives you a clean, keyboard-driven interface to:

- **Manage worktrees** instantly with single keystrokes
- **Automate workflows** with AI-generated commit messages, branch names, and PR descriptions
- **Handle GitHub PRs** without leaving the terminal
- **Open editors** directly on worktrees
- **Maintain persistent Claude CLI sessions** for each branch

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and designed for developers who juggle multiple feature branches simultaneously.

> **Why "jean"?** Named after Jean-Claude Van Damme - combining "Jean" (the CLI) with "Claude" (AI integration). A perfect martial arts kick of productivity!

## Features

- **Git Worktree Management** - Create, switch, and delete worktrees with single keystrokes
- **AI-Powered Workflow** - Auto-generate commit messages, branch names, and PR content (11+ AI models)
- **GitHub PR Automation** - Create draft PRs, browse PRs, merge with strategy selection
- **Tmux Sessions** - Persistent Claude CLI and terminal sessions per worktree
- **5 Themes** - Matrix, Coolify, Dracula, Nord, Solarized with dynamic switching
- **Multi-Editor Support** - Open worktrees in VS Code, Cursor, Neovim, Vim, Sublime, Atom, or Zed
- **Branch Management** - Rename branches, checkout, change base branch, pull from base
- **Debug Logging** - Enable logs for troubleshooting

## Installation

### With Go (Recommended)

```bash
go install github.com/andrew-bierman/jean-tui@latest
jean init  # Set up shell integration
```

### From Source

```bash
git clone https://github.com/andrew-bierman/jean-tui
cd jean-tui
go build -o jean
sudo mv jean /usr/local/bin/
jean init
```

### Quick Install Script

```bash
curl -fsSL https://github.com/andrew-bierman/jean-tui/raw/main/install.sh | bash
```

## Prerequisites

- **Git**: For worktree operations
- **tmux**: For session management (`brew install tmux` on macOS, `sudo apt install tmux` on Linux)
- **GitHub CLI**: For PR operations (`brew install gh` on macOS, `sudo apt install gh` on Linux)

## Quick Start

Run in any Git repository:

```bash
cd /path/to/your/repo
jean
```

Or test on another repo without navigating:

```bash
jean -path /path/to/other/repo
```

## Keybindings Quick Reference

### Navigation & Core
| Key | Action |
|-----|--------|
| `↑`/`↓` or `j`/`k` | Navigate worktrees |
| `Enter` | Switch to worktree (Claude session) |
| `t` | Open terminal session |
| `q` | Quit |

### Worktree Management
| Key | Action |
|-----|--------|
| `n` | Create new worktree |
| `a` | Create from existing branch |
| `d` | Delete worktree |
| `o` | Open in editor |
| `r` | Refresh (fetch + auto-pull) |

### Git Operations
| Key | Action |
|-----|--------|
| `b` | Change base branch |
| `B` | Rename branch |
| `K` | Checkout branch |
| `c` | Commit (with AI) |
| `p` | Push to remote |
| `u` | Update from base |

### GitHub & PRs
| Key | Action |
|-----|--------|
| `P` | Create draft PR |
| `N` | Create worktree from PR |
| `L` | Local merge (worktree → base) |
| `v` | View PR in browser |
| `M` | Merge PR |
| `g` | Open repo in browser |

### Application
| Key | Action |
|-----|--------|
| `e` | Select editor |
| `s` | Settings menu |
| `S` | Manage tmux sessions |
| `h` | Help modal |

## Configuration

### User Config
Settings stored in `~/.config/jean/config.json` per repository:
- **Base branch** - Default branch for new worktrees
- **Editor** - Preferred IDE (code, cursor, nvim, vim, subl, atom, zed)
- **Theme** - Visual theme (press `s` → Theme to change)
- **AI Settings** - OpenRouter API key, model selection, feature toggles
- **Debug logs** - Enable logging to `/tmp/jean-debug.log`

### Tmux Configuration

Press `s` → Tmux Config to install an opinionated tmux configuration with:
- Mouse support & scrolling
- 10k line scrollback buffer
- 256 color + true color support
- Shift+Left/Right to switch windows
- Ctrl+D to detach
- Better pane borders and status bar

### Setup Scripts

Automatically run commands when creating new worktrees. Create `jean.json` in your repository root:

```json
{
  "scripts": {
    "setup": "npm install && cp $JEAN_ROOT_PATH/.env ."
  }
}
```

**Environment variables available:**
- `JEAN_WORKSPACE_PATH` - Path to the newly created worktree
- `JEAN_ROOT_PATH` - Path to the repository root directory
- `JEAN_BRANCH` - Current branch name

The setup script runs automatically for every new worktree (created with `n` or `a` keys). Script failures are shown as warnings and won't block worktree creation.

## Workflows

### Create Draft PR (Single Command)
Press `P` to:
1. Auto-commit changes
2. Rename branch with AI (optional)
3. Generate PR title/description with AI (optional)
4. Create draft PR
5. Store PR URL

### Push with Smart Naming
Press `p` to:
1. Check for uncommitted changes
2. Auto-commit if needed
3. Rename random branches with AI
4. Push to remote

### Session Management

Both Claude and terminal sessions can coexist for the same worktree:
- `Enter` creates Claude session (`jean-<branch>`)
- `t` creates terminal session (`jean-<branch>-terminal`)
- Detach anytime with `Ctrl+B D`
- View all sessions with `S`

## Themes

5 built-in themes available (press `s` → Theme):
1. **Matrix** - Green terminal aesthetic
2. **Coolify** - Purple/violet theme
3. **Dracula** - Pink/cyan theme
4. **Nord** - Blue/cyan theme
5. **Solarized** - Blue/teal theme

## Development

### Build & Test
```bash
# Run locally
go run main.go

# Build binary
go build -o jean

# Test with custom path
go run main.go -path /path/to/test/repo

# Check version
./jean --version
```

### Custom Start Commands

Jean supports building with custom CLI names and agent commands for forks or specialized deployments. This allows you to create variants like `ralph-tui`, `opencode`, or any custom command name while keeping all jean functionality.

**Using Makefile (recommended):**
```bash
# Build with custom name
make build-custom NAME=ralph-tui

# Build with blank terminal (worktree management only, no agent)
make build-blank
```

**Build with custom branding (manual):**
```bash
go build -ldflags "\
  -X github.com/andrew-bierman/jean-tui/internal/branding.CLIName=myapp \
  -X github.com/andrew-bierman/jean-tui/internal/branding.SessionPrefix=myapp- \
  -X github.com/andrew-bierman/jean-tui/internal/branding.ConfigDirName=myapp \
  -X github.com/andrew-bierman/jean-tui/internal/branding.EnvVarPrefix=MYAPP" \
  -o myapp
```

**What gets customized:**
| Variable | Default | Example Custom | Affects |
|----------|---------|----------------|---------|
| `CLIName` | `jean` | `opencode` | CLI name, help text, messages |
| `SessionPrefix` | `jean-` | `opencode-` | Tmux session names |
| `ConfigDirName` | `jean` | `opencode` | Config directory (`~/.config/<name>`) |
| `EnvVarPrefix` | `JEAN` | `OPENCODE` | Environment variables |
| `AgentCommand` | `claude` | `ralph-tui` or `` | Command in agent window |
| `AgentWindowName` | `claude` | `agent` or `shell` | Tmux window name |

**Blank terminal mode (worktree-only):**
```bash
# Build with no agent - just worktree management with blank terminals
go build -ldflags "\
  -X github.com/andrew-bierman/jean-tui/internal/branding.AgentCommand= \
  -X github.com/andrew-bierman/jean-tui/internal/branding.AgentWindowName=shell" \
  -o worktree-tui
```

**Custom agent command (e.g., ralph-tui):**
```bash
go build -ldflags "\
  -X github.com/andrew-bierman/jean-tui/internal/branding.CLIName=ralph-launcher \
  -X github.com/andrew-bierman/jean-tui/internal/branding.AgentCommand=ralph-tui \
  -X github.com/andrew-bierman/jean-tui/internal/branding.AgentWindowName=ralph" \
  -o ralph-launcher
```

This is useful for:
- **Blank terminal mode**: Worktree management only, start any agent you want manually
- **Custom agents**: Launch `ralph-tui`, `aider`, or any other AI agent instead of Claude
- **Team deployments**: Custom naming with specific agent configurations
- **Forks**: Distinct identities while tracking upstream jean

### Project Structure
- `main.go` - CLI entry point
- `tui/` - Bubble Tea TUI (model, update, view, styles)
- `git/` - Git worktree operations
- `session/` - Tmux session management
- `config/` - Configuration management
- `github/` - GitHub PR operations
- `openrouter/` - AI integration

For detailed architecture and development guides, see [CLAUDE.md](./CLAUDE.md).

## Platform Support

- **Linux** ✅ Full support
- **macOS** ✅ Full support
- **Windows** ⚠️ WSL2 required

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Contributing

Contributions are welcome! Submit a Pull Request.

## License

MIT

## Acknowledgments

- Inspired by [git-worktree-tui](https://github.com/FredrikMWold/git-worktree-tui)
- Built with [Charm](https://charm.sh/) ecosystem
