package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Telmate/proxmox-api-go/proxmox"
)

var details = container.NewVBox()

func displayUI(window fyne.Window, c *proxmox.Client) {
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

	ui := container.NewBorder(nil, nil, scrollableAc, nil, details)
	window.SetContent(ui)
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
		vm := vm // Create a new variable to capture the current value
		// "closure" issue
		vmString := fmt.Sprintf("%d - %s", vm.VmID, vm.Name)
		vmButton := widget.NewButton(vmString, func() {
			updateDetails(vm)
		})
		canvas.Add(vmButton)
	}

	return canvas
}

func updateDetails(vm VMDetails) {
	details.RemoveAll()

	vm1stEntry := canvas.NewText(vm.Name, color.Black)
	details.Add(vm1stEntry)

	vm2ndEntry := container.NewBorder(nil, nil,
		canvas.NewText("Status", color.Black), canvas.NewText(vm.Status, color.Black), nil)
	details.Add(vm2ndEntry)

	vm3rdEntry := container.NewBorder(nil, nil,
		canvas.NewText("Node", color.Black), canvas.NewText(vm.Node, color.Black), nil)
	details.Add(vm3rdEntry)

	// todo refresh CPU usage
	cpuUsage := vm.CPU * 100
	vm4thEntry := container.NewBorder(nil, nil,
		canvas.NewText("CPU usage", color.Black), canvas.NewText(strconv.FormatFloat(cpuUsage, 'f', -1, 64), color.Black), nil)
	details.Add(vm4thEntry)
}
