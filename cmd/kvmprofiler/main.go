package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/connector"
	"github.com/cha87de/kvmtop/profiler"
	"github.com/cha87de/kvmtop/runners"
)

func main() {

	// handle flags
	initializeFlags()

	// connect to libvirt
	connector.Libvirt.ConnectionURI = config.Options.LibvirtURI
	err := connector.InitializeConnection()
	if err != nil {
		fmt.Println("failed to initialize connection to libvirt. kvmprofile will terminate.")
		os.Exit(1)
	}

	// start lookup and collect runners
	var wg sync.WaitGroup
	wg.Add(1) // terminate when first thread terminates
	go runners.InitializeLookup(&wg)
	go runners.InitializeCollect(&wg)
	go profiler.InitializeProfiler(&wg)
	wg.Wait()

}
