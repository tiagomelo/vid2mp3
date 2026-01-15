// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package ffmpeg

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractAudioFromVideo(t *testing.T) {
	originalExecutor := new(defaultFileSystemExecOpsProvider)

	t.Run("successfully extracts audio from video", func(t *testing.T) {
		defer func() {
			fileSystemExecOpsProvider = originalExecutor
		}()

		fileSystemExecOpsProvider = new(mockOSCommandExecutor)

		err := ExtractAudioFromVideo(context.TODO(), "input.mp4", "output.mp3")
		require.NoError(t, err)
	})

	t.Run("ffmpeg not available", func(t *testing.T) {
		defer func() {
			fileSystemExecOpsProvider = originalExecutor
		}()

		fileSystemExecOpsProvider = &mockOSCommandExecutor{
			toolNotAvailable: true,
		}

		err := ExtractAudioFromVideo(context.TODO(), "input.mp4", "output.mp3")
		require.Equal(t, ErrorFfmpegNotAvailable, err)
	})

	t.Run("error during command execution", func(t *testing.T) {
		defer func() {
			fileSystemExecOpsProvider = originalExecutor
		}()

		errorExecution := errors.New("some execution error")

		errCommandExecution := &ErrorCommandExecution{
			inputFile:  "input.mp4",
			outputFile: "output.mp3",
			err:        errorExecution,
		}
		fileSystemExecOpsProvider = &mockOSCommandExecutor{
			execCommandErr: errorExecution,
		}

		err := ExtractAudioFromVideo(context.TODO(), "input.mp4", "output.mp3")
		require.Equal(t, errCommandExecution, err)
		require.Equal(t, `ffmpeg command execution failed for input file 'input.mp4' and output file 'output.mp3': some execution error`, err.Error())
	})
}

type mockOSCommandExecutor struct {
	execCommandErr   error
	toolNotAvailable bool
}

func (m *mockOSCommandExecutor) ExecCommand(ctx context.Context, name string, arg ...string) (string, error) {
	return "", m.execCommandErr
}

func (m *mockOSCommandExecutor) ToolIsNotAvailable(toolName string) bool {
	return m.toolNotAvailable
}
