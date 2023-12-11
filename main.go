package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

type Node struct {
	CPU            float64 `json:"cpu"`
	Disk           float64 `json:"disk"`
	ID             string  `json:"id"`
	Level          string  `json:"level"`
	MaxCPU         int     `json:"maxcpu"`
	MaxDisk        float64 `json:"maxdisk"`
	MaxMem         float64 `json:"maxmem"`
	Mem            float64 `json:"mem"`
	Node           string  `json:"node"`
	SSLFingerprint string  `json:"ssl_fingerprint"`
	Status         string  `json:"status"`
	Type           string  `json:"type"`
	Uptime         float64 `json:"uptime"`
	VMList         []VMDetails
}

type moxServer struct {
	serverName string
	VMList     []VMDetails
}

type VMDetails struct {
	ID     int
	VMName string
}

func main() {
	proxfyne := app.New()
	mainWindow := proxfyne.NewWindow("proxfyne")

	apiUrl := os.Getenv("PM_API_URL")
	userID := os.Getenv("PM_USER")
	password := os.Getenv("PM_PASS")
	http_headers := os.Getenv("PM_HTTP_HEADERS")
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	c, err := proxmox.NewClient(apiUrl, nil, http_headers, tlsConfig, "", 5)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Login(userID, password, "")
	if err != nil {
		log.Fatal(err)
	}

	var vmrefs []proxmox.VmRef
	vmlist, err := c.GetVmList()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range vmlist["data"].([]interface{}) {
		vmref := proxmox.VmRef{}
		nodeJSON, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
		}
		err = json.Unmarshal(nodeJSON, &vmref)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
		vmrefs = append(vmrefs, vmref)
	}

	var nodes []Node
	list, err := c.GetNodeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list["data"].([]interface{}) {
		node := Node{}
		nodeJSON, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
		}
		err = json.Unmarshal(nodeJSON, &node)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
		// TODO add VMs on nodes
		nodes = append(nodes, node)
	}

	displayMenu(mainWindow, nodes, vmrefs)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

func displayMenu(window fyne.Window, nodes []Node, vmRefs []proxmox.VmRef) {
	ac := createAccordion(nodes, vmRefs)

	scrollableAc := container.NewVScroll(ac)
	columns := container.NewBorder(nil, nil, scrollableAc, nil, nil)

	window.SetContent(columns)
}

func createAccordion(nodes []Node, vmRefs []proxmox.VmRef) fyne.Widget {
	ac := widget.NewAccordion()

	for _, node := range nodes {
		ac.Append(widget.NewAccordionItem(node.Node, createVMList(node.VMList)))
	}
	ac.MultiOpen = true
	return ac
}

// convertVmList converts the proxmox.VmRef list to VMDetails list
func convertVmList(vmRefs []proxmox.VmRef) []VMDetails {
	var vmDetailsList []VMDetails

	for _, vmRef := range vmRefs {
		currentId := vmRef.VmId()
		vmDetailsList = append(vmDetailsList, VMDetails{ID: currentId, VMName: "toto"})
	}

	return vmDetailsList
}

func createVMList(vmlist []VMDetails) fyne.CanvasObject {
	canvas := container.NewVBox()

	for _, vm := range vmlist {
		vmstring := fmt.Sprintf("%d - %s", vm.ID, vm.VMName)
		canvas.Add(widget.NewLabel(vmstring))
	}

	return canvas
}
