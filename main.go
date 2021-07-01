package main

import (
	"flag"
	"fmt"

	"github.com/gazzenger/winssh-pageant/pageant"
	"github.com/lxn/win"
)

var (
	sshPipe       = flag.String("sshpipe", `\\.\pipe\openssh-ssh-agent`, "Named pipe for Windows OpenSSH agent")
	noPageantPipe = flag.Bool("no-pageant-pipe", false, "Toggle pageant named pipe proxying")
)

func main() {
	flag.Parse()

	// Start a proxy/redirector for the pageant named pipes
	if !*noPageantPipe {
		go pageant.PipeProxy()
	}

	pageant.SshPipe = sshPipe

	pageantWindow := pageant.CreatePageantWindow()
	if pageantWindow == 0 {
		fmt.Println(fmt.Errorf("CreateWindowEx failed: %v", win.GetLastError()))
		return
	}

	// main message loop
	var msg win.MSG
	for win.GetMessage(&msg, 0, 0, 0) > 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}
