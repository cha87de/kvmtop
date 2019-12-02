package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"kvmtop/config"
)

// GetCmdLine reads the cmdline for a process from /proc
func GetCmdLine(pid int) string {
	filepath := fmt.Sprint(config.Options.ProcFS, "/", strconv.Itoa(pid), "/cmdline")
	filecontent, _ := ioutil.ReadFile(filepath)
	return string(filecontent)
}

// GetProcessList reads and returns all PIDs from the proc filesystem
func GetProcessList() []int {
	files, err := ioutil.ReadDir(config.Options.ProcFS)
	if err != nil {
		log.Fatal(err)
	}

	var processes []int
	for _, f := range files {
		// is it a folder?
		if !f.IsDir() {
			continue
		}
		// is the name a number?
		if pid, err := strconv.Atoi(f.Name()); err == nil {
			processes = append(processes, pid)
		}
	}

	return processes
}
