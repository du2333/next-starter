package nextjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config == nil {
		t.Fatal("Expected config to not be nil")
	}

	if len(config.FilesToRemove) == 0 {
		t.Error("Expected files to remove to be configured")
	}

	if config.PageContent == "" {
		t.Error("Expected page content to be configured")
	}

	if len(config.CreateCommand) == 0 {
		t.Error("Expected create command to be configured")
	}

	if len(config.ShadcnCommand) == 0 {
		t.Error("Expected shadcn command to be configured")
	}
}

func TestProjectManagerCleanup(t *testing.T) {
	// Create a temporary project directory
	tempDir := t.TempDir()
	projectName := "test-project"
	projectPath := filepath.Join(tempDir, projectName)

	// Create the project directory
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		t.Fatal(err)
	}

	// Create files that should be removed
	publicDir := filepath.Join(projectPath, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatal(err)
	}

	readmePath := filepath.Join(projectPath, "README.md")
	if err := os.WriteFile(readmePath, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Change to temp directory
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tempDir)

	// Test cleanup
	config := DefaultConfig()
	pm := NewProjectManager(config)

	if err := pm.CleanupProject(projectName); err != nil {
		t.Fatalf("Cleanup failed: %v", err)
	}

	// Verify files are removed
	if _, err := os.Stat(publicDir); !os.IsNotExist(err) {
		t.Error("Expected public directory to be removed")
	}

	if _, err := os.Stat(readmePath); !os.IsNotExist(err) {
		t.Error("Expected README.md to be removed")
	}
}
