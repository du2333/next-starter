#!/bin/bash

# Test script for next-starter CLI
# This demonstrates how to use the tool

echo "🧪 Testing next-starter CLI"
echo

echo "1. Building the application..."
make build

echo
echo "2. Testing help message..."
./next-starter

echo
echo "3. Example usage (not executed):"
echo "   ./next-starter test-project"
echo
echo "This would:"
echo "   - Create a new Next.js project named 'test-project'"
echo "   - Remove unnecessary files (/public, /app/favicon.ico, README.md)"
echo "   - Replace /app/page.tsx with custom content"
echo "   - Install shadcn/ui with neutral theme"
echo "   - Commit changes to git"

echo
echo "✅ Test script complete!"
echo "📚 To use the tool: ./next-starter <your-project-name>"