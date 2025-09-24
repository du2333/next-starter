package nextjs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ProjectManager handles Next.js project operations
type ProjectManager struct {
	config *Config
}

// NewProjectManager creates a new project manager with the given config
func NewProjectManager(config *Config) *ProjectManager {
	return &ProjectManager{
		config: config,
	}
}

// CreateProject creates a new Next.js project with the specified name
func (pm *ProjectManager) CreateProject(projectName string) error {
	cmd := append(pm.config.CreateCommand, projectName)

	execCmd := exec.Command(cmd[0], cmd[1:]...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	return execCmd.Run()
}

// CleanupProject removes unnecessary files and directories
func (pm *ProjectManager) CleanupProject(projectName string) error {
	for _, item := range pm.config.FilesToRemove {
		itemPath := filepath.Join(projectName, item)
		if err := os.RemoveAll(itemPath); err != nil {
			// Don't fail if file doesn't exist
			if !os.IsNotExist(err) {
				return fmt.Errorf("failed to remove %s: %w", itemPath, err)
			}
		}
	}
	return nil
}

// UpdatePageContent replaces the default page.tsx content
func (pm *ProjectManager) UpdatePageContent(projectName string) error {
	pagePath := filepath.Join(projectName, "app", "page.tsx")

	if err := os.WriteFile(pagePath, []byte(pm.config.PageContent), 0644); err != nil {
		return fmt.Errorf("failed to write page.tsx: %w", err)
	}

	return nil
}

// InstallShadcn installs shadcn/ui with the configured theme
func (pm *ProjectManager) InstallShadcn(projectName string) error {
	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Change to project directory
	if err := os.Chdir(projectName); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Restore original directory on function exit
	defer func() {
		os.Chdir(originalDir)
	}()

	// Run shadcn init command
	execCmd := exec.Command(pm.config.ShadcnCommand[0], pm.config.ShadcnCommand[1:]...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	return execCmd.Run()
}

// GetFilesToRemove returns the list of files that will be removed
func (pm *ProjectManager) GetFilesToRemove() []string {
	return pm.config.FilesToRemove
}
