package cli

import (
	"fmt"
	"os"

	"next-starter/internal/git"
	"next-starter/internal/nextjs"
)

// App represents the CLI application
type App struct {
	projectManager *nextjs.ProjectManager
	gitManager     *git.Manager
}

// NewApp creates a new CLI application
func NewApp() *App {
	config := nextjs.DefaultConfig()
	return &App{
		projectManager: nextjs.NewProjectManager(config),
		gitManager:     git.NewManager(),
	}
}

// Run executes the CLI application
func (app *App) Run(args []string) error {
	if len(args) < 2 {
		app.showUsage(args[0])
		return fmt.Errorf("project name is required")
	}

	projectName := args[1]
	return app.createProject(projectName)
}

// createProject orchestrates the entire project creation process
func (app *App) createProject(projectName string) error {
	fmt.Printf("🚀 Creating Next.js project: %s\n", projectName)

	// Step 1: Create Next.js project
	fmt.Println("📦 Running: pnpx create-next-app@latest...")
	if err := app.projectManager.CreateProject(projectName); err != nil {
		return fmt.Errorf("failed to create Next.js project: %w", err)
	}

	// Step 2: Clean up unnecessary files
	fmt.Println("🧹 Cleaning up unnecessary files...")
	if err := app.projectManager.CleanupProject(projectName); err != nil {
		return fmt.Errorf("failed to cleanup project: %w", err)
	}

	// Show what was removed
	for _, item := range app.projectManager.GetFilesToRemove() {
		fmt.Printf("   🗑️  Removed: %s/%s\n", projectName, item)
	}

	// Step 3: Update page content
	fmt.Println("📝 Updating page.tsx content...")
	if err := app.projectManager.UpdatePageContent(projectName); err != nil {
		return fmt.Errorf("failed to update page content: %w", err)
	}
	fmt.Printf("   ✏️  Updated: %s/app/page.tsx\n", projectName)

	// Step 4: Install shadcn/ui
	fmt.Println("🎨 Installing shadcn/ui with neutral theme...")
	if err := app.projectManager.InstallShadcn(projectName); err != nil {
		return fmt.Errorf("failed to install shadcn: %w", err)
	}
	fmt.Println("   ✨ Shadcn/ui installed successfully")

	// Step 5: Commit git changes
	fmt.Println("📝 Committing changes to git...")
	if err := app.gitManager.CommitChanges(projectName); err != nil {
		return fmt.Errorf("failed to commit git changes: %w", err)
	}
	fmt.Println("   📝 Git changes committed")

	// Success message
	app.showSuccess(projectName)
	return nil
}

// showUsage displays usage information
func (app *App) showUsage(programName string) {
	fmt.Fprintf(os.Stderr, "Usage: %s <project-name>\n", programName)
	fmt.Fprintf(os.Stderr, "Example: %s my-nextjs-app\n", programName)
}

// showSuccess displays success message and next steps
func (app *App) showSuccess(projectName string) {
	fmt.Printf("✅ Successfully created and configured Next.js project: %s\n", projectName)
	fmt.Printf("📁 Project location: ./%s\n", projectName)
	fmt.Printf("🏃 Next steps:\n")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Printf("   pnpm dev\n")
}
