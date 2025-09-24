package nextjs

// Config holds the configuration for Next.js project creation
type Config struct {
	// Template content for page.tsx
	PageContent string

	// Files and directories to remove after project creation
	FilesToRemove []string

	// Create command and arguments
	CreateCommand []string

	// Shadcn installation command and arguments
	ShadcnCommand []string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		PageContent: `export default function Home() {
  return <h1 className="text-3xl font-bold underline">Hello World</h1>;
}
`,
		FilesToRemove: []string{
			"public",
			"app/favicon.ico",
			"README.md",
		},
		CreateCommand: []string{"pnpx", "create-next-app@latest", "--yes"},
		ShadcnCommand: []string{"pnpm", "dlx", "shadcn@latest", "init", "-b", "neutral"},
	}
}
