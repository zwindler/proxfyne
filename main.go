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

	toto := make(map[string][]int)
	for _, vm := range vmList {
		toto[vm.Node] = append(toto[vm.Node], vm.VmID)
	}

	for _, node := range nodes {
		ac.Append(widget.NewAccordionItem(node.Node, createVMList(toto[node.Node])))
	}
	ac.MultiOpen = true
	return ac
}

func createVMList(vmList []int) fyne.CanvasObject {
	canvas := container.NewVBox()

	for _, vmId := range vmList {
		vmString := fmt.Sprintf("%d", vmId)
		canvas.Add(widget.NewLabel(vmString))
	}

	return canvas
}
