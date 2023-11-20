package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type moxServer struct {
	serverName string
	VMList     []VMDetails
}

type VMDetails struct {
	id     int
	VMName string
}

var (
	demo = []moxServer{
		{serverName: "server1", VMList: []VMDetails{
			{id: 100, VMName: "vm1-100"},
			{id: 101, VMName: "vm2-101"},
		}},
		{serverName: "server2", VMList: []VMDetails{
			{id: 200, VMName: "one"},
			{id: 201, VMName: "two"},
			{id: 202, VMName: "three"},
			{id: 203, VMName: "four"},
			{id: 204, VMName: "five"},
			{id: 205, VMName: "six"},
			{id: 206, VMName: "seven"},
			{id: 207, VMName: "eight"},
		}},
	}
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

	scrollableAc := container.NewVScroll(ac)

	window.SetContent(scrollableAc)
}

func createAccordion() fyne.Widget {
	ac := widget.NewAccordion()

	for _, server := range demo {
		ac.Append(widget.NewAccordionItem(server.serverName, createVMList(server.VMList)))
	}
	ac.MultiOpen = true
	return ac
}

func createVMList(vmlist []VMDetails) fyne.CanvasObject {
	canvas := container.NewVBox()

	for _, vm := range vmlist {
		vmstring := fmt.Sprintf("%d - %s", vm.id, vm.VMName)
		canvas.Add(widget.NewLabel(vmstring))
	}

	return canvas
}
