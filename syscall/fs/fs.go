// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package fs

import (
	"context"
	"os"
	"os/exec"
)

// ExecCmd interface abstracts the command execution.
type ExecCmd interface {
	// CombinedOutput runs the command and returns its
	// combined standard output and standard error.
	CombinedOutput() ([]byte, error)
}

// FileSystemExecOps interface abstracts file system operations
// related to command execution.
type FileSystemExecOps interface {
	// CommandContext creates a new execCmd for the given command.
	CommandContext(ctx context.Context, name string, arg ...string) ExecCmd

	// LookPath checks if a file is available in the system's PATH.
	LookPath(file string) (string, error)
}

// OSFileSystemExecOps struct implements the fileSystem interface using
// the standard library's os package. This is the real implementation
// that interacts with the actual file system.
type OSFileSystemExecOps struct{}

func (OSFileSystemExecOps) CommandContext(ctx context.Context, name string, arg ...string) ExecCmd {
	return exec.CommandContext(ctx, name, arg...)
}

func (OSFileSystemExecOps) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// DirEntry interface represents a directory entry.
type DirEntry interface {
	// Name returns the name of the directory entry.
	Name() string

	// IsDir returns true if the directory entry is a directory.
	IsDir() bool
}

// FileSystemOps interface abstracts file system operations.
type FileSystemOps interface {
	// ReadDir returns a list of directory entries.
	ReadDir(name string) ([]DirEntry, error)
}

// osFileInfoWrapper wraps the os.FileInfo to implement the fileInfo interface.
type osDirEntryWrapper struct {
	os.DirEntry
}

// Name returns the name of the directory entry.
func (w osDirEntryWrapper) Name() string {
	return w.DirEntry.Name()
}

// IsDir returns true if the directory entry is a directory.
func (w osDirEntryWrapper) IsDir() bool {
	return w.DirEntry.IsDir()
}

// OSFileSystemOps struct implements the FileSystemOps interface using
// the standard library's os package.
type OSFileSystemOps struct{}

// ReadDir reads the directory named by name and returns
// a list of directory entries.
func (OSFileSystemOps) ReadDir(name string) ([]DirEntry, error) {
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}
	result := make([]DirEntry, len(entries))
	for i, entry := range entries {
		result[i] = osDirEntryWrapper{entry}
	}
	return result, nil
}
