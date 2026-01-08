package gui

import "fyne.io/fyne/v2/widget"

// exitButton creates a button that allows the user to exit the application.
func (g *gui) exitButton() *widget.Button {
	exitButton := widget.NewButton("exit", func() {
		g.fyneApp.Quit()
	})
	exitButton.Importance = widget.DangerImportance

	return exitButton
}
