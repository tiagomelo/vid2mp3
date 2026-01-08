// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package gui

import "fyne.io/fyne/v2/widget"

// newHorizontalRadioGroup creates a horizontal radio group with the given options,
// selected index, and a callback function that is called when the selection changes.
func newHorizontalRadioGroup(
	options []string,
	selectedIndex int,
	onChanged func(string),
) *widget.RadioGroup {
	radioGroup := &widget.RadioGroup{}

	radioGroup.OnChanged = func(selected string) {
		if selected == "" {
			radioGroup.SetSelected(options[selectedIndex])
		}
		onChanged(selected)
	}

	radioGroup.Options = options
	radioGroup.Horizontal = true
	radioGroup.SetSelected(options[selectedIndex])

	return radioGroup
}
