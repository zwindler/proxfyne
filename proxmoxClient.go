package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

type VMDetails struct {
	CPU       float64 `json:"cpu"`
	Disk      float64 `json:"disk"`
	DiskRead  float64 `json:"diskread"`
	DiskWrite float64 `json:"diskwrite"`
	ID        string  `json:"id"`
	MaxCPU    int     `json:"maxcpu"`
	MaxDisk   float64 `json:"maxdisk"`
	MaxMem    float64 `json:"maxmem"`
	Mem       float64 `json:"mem"`
	Name      string  `json:"name"`
	NetIn     float64 `json:"netin"`
	NetOut    float64 `json:"netout"`
	Node      string  `json:"node"`
	Status    string  `json:"status"`
	Template  int     `json:"template"`
	Type      string  `json:"type"`
	Uptime    float64 `json:"uptime"`
	VmID      int     `json:"vmid"`
}

func createClient() (c *proxmox.Client, err error) {
	apiUrl := os.Getenv("PM_API_URL")
	userID := os.Getenv("PM_USER")
	password := os.Getenv("PM_PASS")
	http_headers := os.Getenv("PM_HTTP_HEADERS")
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	c, err = proxmox.NewClient(apiUrl, nil, http_headers, tlsConfig, "", 5)
	if err != nil {
		return
	}
	err = c.Login(userID, password, "")
	if err != nil {
		return
	}
	return
}

func getVMs(c *proxmox.Client) (vms []VMDetails, err error) {
	vmList, err := c.GetVmList()
	if err != nil {
		return
	}

	for _, item := range vmList["data"].([]interface{}) {
		vmInfos := VMDetails{}
		nodeJSON, err := json.Marshal(item)
		if err != nil {
			return []VMDetails{}, fmt.Errorf("error unMarshalling JSON: %w", err)
		}
		err = json.Unmarshal(nodeJSON, &vmInfos)
		if err != nil {
			return []VMDetails{}, fmt.Errorf("error unMarshalling JSON: %w", err)
		}
		vms = append(vms, vmInfos)
	}
	return
}

func getNodes(c *proxmox.Client) (nodes []Node, err error) {
	list, err := c.GetNodeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list["data"].([]interface{}) {
		node := Node{}
		nodeJSON, err := json.Marshal(item)
		if err != nil {
			return nodes, fmt.Errorf("error unMarshalling JSON: %w", err)
		}
		err = json.Unmarshal(nodeJSON, &node)
		if err != nil {
			return nodes, fmt.Errorf("error unMarshalling JSON: %w", err)
		}
		nodes = append(nodes, node)
	}
	return
}
