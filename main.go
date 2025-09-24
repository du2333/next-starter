package main

import (
	"fmt"
	"os"

	"next-starter/internal/cli"
)

func main() {
	app := cli.NewApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}
}
