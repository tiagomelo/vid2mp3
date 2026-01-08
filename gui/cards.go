package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// videoInputFileSelectionCard creates a card for video
// input file selection.
func (g *gui) videoInputFileSelectionCard() *widget.Card {
	const (
		chooseButtonLabel            = "choose..."
		videoFilePathPlaceHolder     = "path/to/file"
		videoFilePath                = "video file path"
		videoFilesDirectoryPath      = "video files directory path"
		videoFilesDirPathPlaceHolder = "path/to/directory"
	)

	videoFilePathChooseButton := widget.NewButton(chooseButtonLabel, func() {
	})
	videoFilePathChooseButton.Importance = widget.HighImportance

	videoFilePathEntry, videoFilePathEntryBinding := newEntryWithBinding(
		videoFilePathPlaceHolder,
		validateNotEmpty,
	)
	g.videoFilePathEntryBinding = videoFilePathEntryBinding

	videoSelectionRadioButtons := newHorizontalRadioGroup(
		[]string{videoFilePath, videoFilesDirectoryPath},
		0,
		func(selected string) {
			switch selected {
			case videoFilePath:
				videoFilePathEntry.SetText("")
				videoFilePathEntry.SetPlaceHolder(videoFilePathPlaceHolder)
				videoFilePathChooseButton.OnTapped = func() {
					dialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.ShowError(err, g.fyneWindow)
							return
						}
						if read == nil {
							return
						}
						g.videoFilePathEntryBinding.Set(read.URI().Path())
					}, g.fyneWindow)

					dialog.SetFilter(storage.NewExtensionFileFilter([]string{
						".png", ".jpg", ".jpeg", ".webp",
						".tif", ".tiff", ".avif", ".heic"},
					))

					dialog.Show()
				}
			case videoFilesDirectoryPath:
				videoFilePathEntry.SetText("")
				videoFilePathEntry.SetPlaceHolder(videoFilesDirPathPlaceHolder)
				videoFilePathChooseButton.OnTapped = func() {
					dialog := dialog.NewFolderOpen(
						func(list fyne.ListableURI, err error) {
							if err != nil {
								dialog.ShowError(err, g.fyneWindow)
								return
							}
							if list == nil {
								return
							}
							g.videoFilePathEntryBinding.Set(list.Path())
						},
						g.fyneWindow,
					)

					dialog.Show()
				}
			}
		},
	)

	videoSelectionContainer := container.NewGridWithRows(
		2,
		videoSelectionRadioButtons,
		container.NewGridWithColumns(2,
			container.NewStack(videoFilePathEntry),
			container.NewStack(videoFilePathChooseButton)),
	)

	return widget.NewCard("source", "", videoSelectionContainer)
}

// outputFileSelectionCard creates a card for output file selection.
func (g *gui) outputFileSelectionCard() *widget.Card {
	return nil
}

// cards creates the main cards for the GUI.
func (g *gui) cards() *fyne.Container {
	spacer := container.NewVBox()
	spacer.Resize(fyne.NewSize(0, 16))

	return container.NewVBox(
		spacer,
		g.videoInputFileSelectionCard(),
		spacer,
		g.outputFileSelectionCard(),
		spacer,
		container.NewCenter(g.exitButton()),
		spacer,
	)
}
