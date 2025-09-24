package git

import (
	"fmt"
	"os"
	"os/exec"
)

// Manager handles git operations
type Manager struct{}

// NewManager creates a new git manager
func NewManager() *Manager {
	return &Manager{}
}

// CommitChanges adds all files and commits with amend
func (g *Manager) CommitChanges(projectDir string) error {
	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Change to project directory
	if err := os.Chdir(projectDir); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Restore original directory on function exit
	defer func() {
		os.Chdir(originalDir)
	}()

	// Git add all files
	if err := g.runCommand("git", "add", "."); err != nil {
		return fmt.Errorf("failed to git add: %w", err)
	}

	// Git commit with amend
	if err := g.runCommand("git", "commit", "--amend", "--no-edit"); err != nil {
		return fmt.Errorf("failed to git commit: %w", err)
	}

	return nil
}

// runCommand executes a git command
func (g *Manager) runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
