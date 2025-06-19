package juju

import (
	"fmt"
	"os"
)

const (
	debugLogCommand = "debug-log"
)

type PrintDebugLogOptions struct {
	Replay bool
}

func (j Juju) PrintDebugLog(opts *PrintDebugLogOptions) error {
	args := []string{debugLogCommand}

	if opts.Replay {
		args = append(args, "--replay")
	}

	args = append(args, "--no-tail")

	output, err := j.Runner.Run(args...)
	if err != nil {
		return fmt.Errorf("failed to get logs: %w", err)
	}

	if _, err := os.Stdout.Write(output); err != nil {
		return fmt.Errorf("failed to write logs to stdout: %w", err)
	}

	return nil
}
