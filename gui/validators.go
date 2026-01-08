// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package gui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// validateNumber is a simple validator that checks if
// the input string can be converted to an integer.
func validateNumber(s string) error {
	if err := validateNoSpaces(s); err != nil {
		return err
	}
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("%s is not a valid number", s)
	}
	return nil
}

// validateNotEmpty is a simple validator that checks if the input string is empty.
func validateNotEmpty(s string) error {
	if s == "" {
		return errors.New("cannot be blank")
	}
	return nil
}

// validateNoSpaces checks that the input does not contain any space characters.
func validateNoSpaces(input string) error {
	if strings.Contains(input, " ") {
		return errors.New("cannot contain spaces")
	}
	return validateNotEmpty(input)
}

// validateImagePath checks if the image path entry binding is set and
// validates that the path is not empty.
func (g *gui) validateImagePath() error {
	if g.videoFilePathEntryBinding == nil {
		return fmt.Errorf("binding not set")
	}
	val, err := g.videoFilePathEntryBinding.Get()
	if err != nil {
		return errors.Wrap(err, "failed to get video file path binding value")
	}
	return validateNotEmpty(val)
}
