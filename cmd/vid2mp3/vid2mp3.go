// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/tiagomelo/vid2mp3/ffmpeg"
)

// run executes the main logic of the vid2mp3 command.
func run(args []string, log *slog.Logger) error {
	if len(args) < 1 {
		return errors.New("usage: vid2mp3 <input-video> [output.mp3]")
	}

	ctx := context.Background()
	defer log.InfoContext(ctx, "completed")

	input := strings.TrimSpace(args[0])
	if input == "" {
		return errors.New("input video path is required")
	}

	var output string
	if len(args) >= 2 {
		output = strings.TrimSpace(args[1])
	}
	if output == "" {
		output = defaultOutputPath(input)
	}

	log.InfoContext(ctx, "extracting audio from video",
		slog.String("input_file", input),
		slog.String("output_file", output),
	)

	if err := ffmpeg.ExtractAudioFromVideo(ctx, input, output); err != nil {
		return err
	}

	log.InfoContext(ctx, "audio extracted successfully")
	return nil
}

// defaultOutputPath generates a default output path
// for the MP3 file based on the input video file path.
func defaultOutputPath(inputPath string) string {
	dir := filepath.Dir(inputPath)
	base := filepath.Base(inputPath)

	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	if name == "" {
		name = "output"
	}

	return filepath.Join(dir, fmt.Sprintf("%s.mp3", name))
}

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	if err := run(os.Args[1:], log); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
