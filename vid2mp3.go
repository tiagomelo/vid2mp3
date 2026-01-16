// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package vid2mp3

import (
	"context"

	"github.com/tiagomelo/vid2mp3/ffmpeg"
)

// mp3Extractor defines an interface for extracting audio from video files.
type mp3Extractor interface {
	ExtractAudioFromVideo(ctx context.Context, inputVideoPath, outputAudioPath string) error
}

// ffmpegExtractor is an implementation of mp3Extractor that uses ffmpeg.
type ffmpegExtractor struct{}

// ExtractAudioFromVideo extracts audio from a video file using ffmpeg.
func (f *ffmpegExtractor) ExtractAudioFromVideo(ctx context.Context, inputVideoPath, outputAudioPath string) error {
	return ffmpeg.ExtractAudioFromVideo(ctx, inputVideoPath, outputAudioPath)
}

// defaultMp3Extractor is the default implementation of mp3Extractor.
var defaultMp3Extractor mp3Extractor = &ffmpegExtractor{}

// ExtractAudioFromVideo extracts audio from a video file.
func ExtractAudioFromVideo(ctx context.Context, inputVideoPath, outputAudioPath string) error {
	return defaultMp3Extractor.ExtractAudioFromVideo(ctx, inputVideoPath, outputAudioPath)
}
