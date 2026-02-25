package utils

import (
	"context"
	"os/exec"
	"time"
)

func RunCommand(bin string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, args...)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
