package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	proxfyne := app.New()
	mainWindow := proxfyne.NewWindow("proxfyne")

	mainWindow.SetContent(widget.NewLabel("Hello World!"))
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}
