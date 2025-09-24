# Next.js Starter CLI

A Go CLI tool that helps you quickly bootstrap Next.js applications with a standardized setup, eliminating repetitive initialization steps.

## Features

🚀 **Automated Next.js Setup**: Creates a new Next.js project using the latest version  
🧹 **Smart Cleanup**: Removes unnecessary files and folders  
📝 **Custom Template**: Replaces default page with a clean Tailwind CSS template  
🎨 **Shadcn/UI Integration**: Automatically installs and configures shadcn/ui with neutral theme  
📝 **Git Integration**: Automatically commits changes to keep your git history clean  

## Prerequisites

- Go 1.25 or higher
- Node.js and pnpm (required for Next.js and shadcn/ui)
- Git

> **Note**: This tool currently only supports **pnpm** as the package manager. Support for npm and yarn may be added in future versions.

## Installation

### Download Pre-built Binary

Pre-built binaries are automatically generated for every commit and available in [GitHub Releases](../../releases).

### Build from Source

```bash
go build -o next-starter main.go
```

### Usage

```bash
./next-starter <project-name>
```

**Example:**
```bash
./next-starter my-awesome-app
```

## What It Does

1. **Creates Next.js Project**: Runs `pnpx create-next-app@latest --yes [project-name]`
2. **Cleans Up Files**: Removes:
   - `/public` folder
   - `/app/favicon.ico`
   - `README.md`
3. **Updates Template**: Replaces `/app/page.tsx` with:
   ```tsx
   export default function Home() {
     return <h1 className="text-3xl font-bold underline">Hello World</h1>;
   }
   ```
4. **Installs shadcn/ui**: Runs `pnpm dlx shadcn@latest init -b neutral`
5. **Git Commit**: Runs `git add .` and `git commit --amend --no-edit`

## Example Output

```
🚀 Creating Next.js project: my-app
📦 Running: pnpx create-next-app@latest...
🧹 Cleaning up unnecessary files...
   🗑️  Removed: my-app/public
   🗑️  Removed: my-app/app/favicon.ico
   🗑️  Removed: my-app/README.md
📝 Updating page.tsx content...
   ✏️  Updated: my-app/app/page.tsx
🎨 Installing shadcn/ui with neutral theme...
   ✨ Shadcn/ui installed successfully
📝 Committing changes to git...
   📝 Git changes committed
✅ Successfully created and configured Next.js project: my-app
📁 Project location: ./my-app
🏃 Next steps:
   cd my-app
   pnpm dev
```

## Project Structure

```
.
├── main.go                           # CLI entry point
├── go.mod                           # Go module file
├── Makefile                         # Build automation
├── README.md                        # This file
├── requirement.md                   # Original requirements
├── test.sh                          # Test script
├── .gitignore                       # Git ignore rules
└── internal/                        # Private packages
    ├── cli/                         # CLI application logic
    │   └── app.go                   # Main CLI app
    ├── nextjs/                      # Next.js project management
    │   ├── config.go               # Configuration
    │   ├── project.go              # Project operations
    │   └── project_test.go         # Tests
    └── git/                         # Git operations
        └── manager.go              # Git manager
```

### Architecture

The application follows Go best practices with clear separation of concerns:

- **`main.go`**: Minimal entry point, delegates to CLI app
- **`internal/cli`**: CLI interface and user interaction
- **`internal/nextjs`**: Next.js project management logic
- **`internal/git`**: Git operations abstraction
- **Configuration**: Externalized and easily modifiable
- **Testing**: Unit tests for core functionality

## License

MIT License