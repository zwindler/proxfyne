package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	proxfyne := app.New()
	mainWindow := proxfyne.NewWindow("proxfyne")

	// TODO add popup if env vars are not populated

	c, err := proxfyneCreateClient()
	if err != nil {
		log.Fatal(err)
	}

	displayUI(mainWindow, c)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}
