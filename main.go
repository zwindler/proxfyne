package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func main() {
	proxfyne := app.New()
	mainWindow := proxfyne.NewWindow("proxfyne")

	c, err := createClient()
	if err != nil {
		log.Fatal(err)
	}

	displayMenu(mainWindow, c)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

func displayMenu(window fyne.Window, c *proxmox.Client) {
	vmList, err := getVMs(c)
	if err != nil {
		log.Fatal(err)
	}

	nodes, err := getNodes(c)
	if err != nil {
		log.Fatal(err)
	}

	ac := createAccordion(nodes, vmList)

	scrollableAc := container.NewVScroll(ac)
	columns := container.NewBorder(nil, nil, scrollableAc, nil, nil)

	window.SetContent(columns)
}

func createAccordion(nodes []Node, vmList []VMDetails) fyne.Widget {
	ac := widget.NewAccordion()

	// reorder VMDetails by nodes in a map
	vmMap := make(map[string][]VMDetails)
	for _, vm := range vmList {
		vmMap[vm.Node] = append(vmMap[vm.Node], vm)
	}

	for _, node := range nodes {
		ac.Append(widget.NewAccordionItem(node.Node, createVMList(vmMap[node.Node])))
	}
	ac.MultiOpen = true
	return ac
}

func createVMList(vmList []VMDetails) fyne.CanvasObject {
	canvas := container.NewVBox()

	for _, vm := range vmList {
		vmString := fmt.Sprintf("%d - %s", vm.VmID, vm.Name)
		canvas.Add(widget.NewLabel(vmString))
	}

	return canvas
}
