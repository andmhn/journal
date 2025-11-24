package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var a = app.NewWithID("journal")
var w = a.NewWindow("journal")

func main() {
	w.Resize(fyne.NewSize(400, 200))
	w.SetMaster()

	w.SetContent(makeUI())
	w.ShowAndRun()

	tidyUp()
}

func makeUI() *fyne.Container {
	count_widget := widget.NewLabel("0")
	count_button := widget.NewButton("count", func() {
		count, _ := strconv.Atoi(count_widget.Text)
		count++
		count_widget.SetText(strconv.Itoa(count))
	})

	reset_button := widget.NewButton("reset", func() {
		count_widget.SetText("0")
	})

	buttons := container.NewHBox(count_button, reset_button)

	pop_up := widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.Resize(fyne.NewSize(300, 100))
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	})

	return container.NewVBox(buttons, count_widget, pop_up, widget.NewEntry())
}

func tidyUp() {
	fmt.Println("Exited")
}
