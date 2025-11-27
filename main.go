package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	pref fyne.Preferences
}

func (app *App) GetCount() string {
	return app.pref.String("count")
}

func (app *App) Increment() string {
	value, _ := strconv.Atoi(app.GetCount())
	value++
	countStr := strconv.Itoa(value)
	app.SetCount(countStr)
	return countStr
}

func (app *App) SetCount(value string) {
	app.pref.SetString("count", value)
}

type UI struct {
	app    App
	window fyne.Window
}

func (ui *UI) makeUI() *fyne.Container {
	app := ui.app
	count_widget := widget.NewLabel(app.pref.String("count"))
	count_button := widget.NewButton("count", func() {
		ui.change_count(count_widget, app.Increment())
	})

	reset_button := widget.NewButton("reset", func() {
		ui.change_count(count_widget, "0")
	})

	buttons := container.NewHBox(count_button, reset_button)

	///-----------------------------------
	entry := newNumericalEntry()

	submitEntry := func() {
		if entry.Text != "" {
			ui.change_count(count_widget, entry.Text)
			entry.SetText("")
		}
	}
	set_button := widget.NewButton("set count", submitEntry)

	set_button.Disable()
	entry.OnSubmitted = func(s string) {
		submitEntry()
	}
	entry.OnChanged = func(s string) {
		if s != "" {
			set_button.Enable()
		} else {
			set_button.Disable()
		}
	}

	entry_container := container.NewVBox(entry, set_button)
	///---------------------------------

	return container.NewVBox(buttons, count_widget, entry_container)
}

func (ui *UI) change_count(count_widget *widget.Label, countStr string) {
	count_widget.SetText(countStr)
	ui.app.SetCount(countStr)
}

func main() {
	var journalApp = app.NewWithID("journal")
	var window = journalApp.NewWindow("journal")
	app := &App{pref: journalApp.Preferences()}
	ui := &UI{app: *app, window: window}

	window.Resize(fyne.NewSize(400, 200))
	window.SetMaster()
	window.SetContent(ui.makeUI())

	journalApp.Settings().SetTheme(theme.DarkTheme())

	window.ShowAndRun()
}
