package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	proxfyne := app.New()
	mainWindow := proxfyne.NewWindow("proxfyne")

	displayMenu(mainWindow)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

func displayMenu(window fyne.Window) {
	ac := createAccordion()

	window.SetContent(ac)
}

func createAccordion() fyne.Widget {
	ac := widget.NewAccordion(
		widget.NewAccordionItem("server1", widget.NewLabel("vm1")),
		widget.NewAccordionItem("server2", widget.NewLabel("vm2")),
	)
	ac.MultiOpen = true
	return ac
}
