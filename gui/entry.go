// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package gui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// newEntryWithBinding creates a new entry widget with
// the given placeholder and validator.
// It returns the entry widget and a binding.String
// that can be used to bind the entry's text.
func newEntryWithBinding(
	placeHolder string,
	validator func(input string) error,
) (*widget.Entry, binding.String) {
	binding := binding.NewString()
	w := widget.NewEntryWithData(binding)
	w.SetPlaceHolder(placeHolder)
	w.Validator = validator
	return w, binding
}
