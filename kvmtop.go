package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/connector"
	"github.com/cha87de/kvmtop/runners"
)

var (
	versionMajor = "2"
	versionMinor string // will be set by go linker
)

func main() {
	// handle flags
	initializeFlags()
	if config.Options.Version {
		fmt.Println("kvmtop version " + versionMajor + "." + versionMinor)
		return
	}

	// catch termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		shutdown(0)
	}()

	// connect to libvirt
	connector.Libvirt.ConnectionURI = config.Options.LibvirtURI
	err := connector.InitializeConnection()
	if err != nil {
		fmt.Println("kvmtop will terminate.")
		os.Exit(1)
	}

	// start runners
	runners.InitializeRunners()

	// when runners terminate, shutdown kvmtop
	shutdown(0)
}

func shutdown(exitcode int) {
	// todo close libvirt connection
	err := connector.CloseConnection()
	if err != nil {
		exitcode = 1
	}

	// todo close printer

	os.Exit(exitcode)
}
