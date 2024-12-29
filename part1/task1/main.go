package main

import "log/slog"
import "os"

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run program", "error", err)
		os.Exit(1)
	}
}

func run() error {

	return nil
}
