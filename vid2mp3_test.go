// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package vid2mp3

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractAudioFromVideo(t *testing.T) {
	originalExtractor := defaultMp3Extractor

	t.Run("successful extraction", func(t *testing.T) {
		defer func() { defaultMp3Extractor = originalExtractor }()

		defaultMp3Extractor = new(mockMp3Extractor)

		err := ExtractAudioFromVideo(context.Background(), "input.mp4", "output.mp3")
		require.NoError(t, err)
	})

	t.Run("extraction failure", func(t *testing.T) {
		defer func() { defaultMp3Extractor = originalExtractor }()

		expectedErr := errors.New("extraction failed")
		defaultMp3Extractor = &mockMp3Extractor{err: expectedErr}

		err := ExtractAudioFromVideo(context.Background(), "input.mp4", "output.mp3")
		require.Equal(t, expectedErr, err)
	})
}

type mockMp3Extractor struct {
	err error
}

func (m *mockMp3Extractor) ExtractAudioFromVideo(ctx context.Context, inputVideoPath, outputAudioPath string) error {
	return m.err
}
