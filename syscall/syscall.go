// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package syscall

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tiagomelo/vid2mp3/syscall/fs"
)

// fsExecOps is a variable that holds the file system ops implementation.
var fsExecOps fs.FileSystemExecOps = fs.OSFileSystemExecOps{}

// ExecCommand executes a command with arguments.
func ExecCommand(ctx context.Context, cmd string, args ...string) (string, error) {
	output, err := fsExecOps.CommandContext(ctx, cmd, args...).CombinedOutput()
	if err != nil {
		return "", errors.WithMessagef(err, "error when executing command [%s] with args %v: output: [%v]", cmd, args, string(output))
	}
	return string(output), nil
}

// ToolIsNotAvailable checks if a tool is not available in the system's PATH.
func ToolIsNotAvailable(toolName string) bool {
	if _, err := fsExecOps.LookPath(toolName); err != nil {
		return true
	}
	return false
}
