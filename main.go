package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nalawade41/relay/runner"
)

func main() {
	ctx := context.Background()

	// We are pushing everything to Run so it can control the errors and main can exit cleanly.
	if err := runner.Run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
