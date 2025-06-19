package juju

import (
	"fmt"
	"os"
)

const (
	debugLogCommand = "debug-log"
)

type WriteDebugLogOptions struct {
	FileName string
	Replay   bool
}

func (j Juju) WriteDebugLog(opts *WriteDebugLogOptions) error {
	args := []string{debugLogCommand}
	if opts.FileName == "" {
		return fmt.Errorf("file name is required")
	}

	if opts.Replay {
		args = append(args, "--replay")
	}

	args = append(args, "--no-tail")

	output, err := j.Runner.Run(args...)
	if err != nil {
		return fmt.Errorf("failed to get logs: %w", err)
	}

	if err := os.WriteFile(opts.FileName, output, 0644); err != nil {
		return fmt.Errorf("failed to write debug log to file %s: %w", opts.FileName, err)
	}

	return nil
}
