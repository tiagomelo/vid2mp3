package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// logEntries sets up the log entries for the application.
func (g *gui) logEntries() *fyne.Container {
	logEntries := widget.NewLabel("")
	logEntries.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(logEntries)
	g.logEntriesLbl = logEntries
	g.logScroll = scroll
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			g.logEntriesLbl.SetText("")
			g.logEntriesLbl.Refresh()
		}),
	)
	return container.NewBorder(
		toolbar,
		container.NewCenter(g.exitButton()),
		nil,
		nil,
		scroll,
	)
}
