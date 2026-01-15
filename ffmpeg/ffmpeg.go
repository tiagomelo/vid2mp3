// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package ffmpeg

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/tiagomelo/vid2mp3/syscall"
)

// ErrorCommandExecution represents an error that occurs
// during the execution of the ffmpeg command.
type ErrorCommandExecution struct {
	inputFile  string
	outputFile string
	err        error
}

// Error returns the error message for ErrorCommandExecution.
func (e *ErrorCommandExecution) Error() string {
	return errors.WithMessagef(e.err, "ffmpeg command execution failed for input file '%s' and output file '%s'", e.inputFile, e.outputFile).Error()
}

// ffmpegToolName is the name of the ffmpeg tool.
const ffmpegToolName = "ffmpeg"

// ErrorFfmpegNotAvailable is returned when ffmpeg is not available in the system's PATH.
var ErrorFfmpegNotAvailable = fmt.Errorf("%s is not available", ffmpegToolName)

// fileSystemExecOpsProvider is a variable that holds the function
// that executes a command with arguments.
var fileSystemExecOpsProvider fileSystemExecOps = &defaultFileSystemExecOpsProvider{}

// fileSystemExecOps defines an interface for executing OS commands.
type fileSystemExecOps interface {
	ExecCommand(ctx context.Context, name string, arg ...string) (string, error)

	ToolIsNotAvailable(toolName string) bool
}

// defaultFileSystemExecOpsProvider is the default implementation of fileSystemExecOps.
type defaultFileSystemExecOpsProvider struct{}

// ExecCommand executes a command with arguments.
func (d *defaultFileSystemExecOpsProvider) ExecCommand(ctx context.Context, name string, arg ...string) (string, error) {
	return syscall.ExecCommand(ctx, name, arg...)
}

// ToolIsNotAvailable checks if a tool is not available in the system's PATH.
func (d *defaultFileSystemExecOpsProvider) ToolIsNotAvailable(toolName string) bool {
	return syscall.ToolIsNotAvailable(toolName)
}

// ExtractAudioFromVideo extracts audio from a video file using ffmpeg.
// ffmpeg must be installed and available in the system's PATH.
func ExtractAudioFromVideo(ctx context.Context, inputVideoPath, outputAudioPath string) error {
	if fileSystemExecOpsProvider.ToolIsNotAvailable(ffmpegToolName) {
		return ErrorFfmpegNotAvailable
	}

	if _, err := fileSystemExecOpsProvider.ExecCommand(
		ctx,
		ffmpegToolName,
		"-i",
		inputVideoPath,
		"-q:a",
		"0",
		"-map",
		"a",
		outputAudioPath,
	); err != nil {
		return &ErrorCommandExecution{
			inputFile:  inputVideoPath,
			outputFile: outputAudioPath,
			err:        err,
		}
	}
	return nil
}
