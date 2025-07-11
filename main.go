package main

import (
	"context"
	"logs-analyser/cmd"
	"os"
)

func main() {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.Execute(ctx, cancel); err != nil {
		cancel()
		os.Exit(1)
	}
}
