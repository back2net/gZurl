package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	WINBOX = "C:\\@\\winbox.exe"
	RDP    = "mstsc"
	PUTTY  = "C:\\@\\putty.exe"
	VNC    = "C:\\@\\vncviewer.exe"
)

func main() {

	uri := os.Args[1]
	uri = strings.Replace(uri, "gzurl:", "", -1)
	uri = strings.Replace(uri, "/", "", -1)
	args := strings.Split(uri, "%20")

	ipAddr := getArg(args, 1)
	login := getArg(args, 2)
	password := getArg(args, 3)

	var cmd *exec.Cmd

	switch args[0] {
	case "wbx":
		cmd = exec.Command(WINBOX, ipAddr, login, password)
	case "rdp":
		cmd = exec.Command(RDP, "/v:", ipAddr)
	case "ssh":
		cmd = exec.Command(PUTTY, login+"@"+ipAddr, "-pw", password)
	case "vnc":
		cmd = exec.Command(VNC, ipAddr)
	}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func getArg(args []string, index int) string {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	return args[index]
}
