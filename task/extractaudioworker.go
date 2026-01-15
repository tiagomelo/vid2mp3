// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package task

import (
	"context"
	"path/filepath"
)

// Params holds the parameters for the extract audio worker.
type Params struct {
	Source    string                     // Source video file path
	Dest      string                     // Destination audio file path
	Extractor func(string, string) error // Extractor function for extracting audio
	Logger    func(string, ...any)       // Logger function for logging messages
	HandleErr func(error)                // Error handler function for handling errors
}

// extractAudioWorker implements the Worker interface
// for extracting audio from video files.
type extractAudioWorker struct {
	params Params
}

// NewExtractAudioWorker creates a new extract audio worker
// with the given parameters.
func NewExtractAudioWorker(params Params) extractAudioWorker {
	return extractAudioWorker{params: params}
}

// Work performs the audio extraction from the video file.
func (w *extractAudioWorker) Work(ctx context.Context) error {
	if err := w.params.Extractor(
		w.params.Source,
		w.params.Dest,
	); err != nil {
		if w.params.HandleErr != nil {
			w.params.HandleErr(err)
		}
		return err
	}
	if w.params.Logger != nil {
		w.params.Logger("Extracted audio from %s and saved to %s",
			filepath.Base(w.params.Source),
			w.params.Dest,
		)
	}
	return nil
}
