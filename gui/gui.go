package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// gui represents the graphical user interface for the app.
type gui struct {
	fyneApp    fyne.App
	fyneWindow fyne.Window

	// ui components.
	appTabs                   *container.AppTabs
	logEntriesLbl             *widget.Label
	logScroll                 *container.Scroll
	videoFilePathEntryBinding binding.String
}

// New creates a new instance of the GUI for the app.
func New() *gui {
	const appTitle = "vid2mp3"
	a := app.NewWithID("info.tiagomelo.vid2mp3")
	w := a.NewWindow(appTitle)
	g := &gui{
		fyneApp:    a,
		fyneWindow: w,
	}
	g.setupUI()
	return g
}

// setupUI initializes the user interface components of the app.
func (g *gui) setupUI() {
	const (
		width  = 680
		height = 300
	)

	// set window dimensions.
	g.fyneWindow.Resize(fyne.NewSize(width, height))

	// setup the ui components.
	g.fyneWindow.SetContent(g.setupContent())

	// quit application when the window is closed.
	g.fyneWindow.SetCloseIntercept(func() {
		g.fyneApp.Quit()
	})
}

// setupContent creates and returns the main content of the GUI.
func (g *gui) setupContent() fyne.CanvasObject {
	g.appTabs = container.NewAppTabs(
		container.NewTabItem("configuration", g.cards()),
		container.NewTabItem("execution log", g.logEntries()),
	)
	g.appTabs.SetTabLocation(container.TabLocationTop)
	return g.appTabs
}

// Run starts the GUI application and enters the main event loop.
func (g *gui) Run() {
	g.fyneWindow.ShowAndRun()
}
