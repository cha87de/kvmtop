package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/connector"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/runners"
)

var (
	versionMajor = "2"
	versionMinor string // will be set by go linker
)

func main() {

	/*
		list := []string{"a", "b"}
		fmt.Printf("%+v", list)
		fmt.Printf("%+v", len(list))

		// add at beginning
		end := len(list) - 1
		if end <= 0 {
			end = len(list)
		}
		list = append([]string{"0"}, list[0:end]...)
		fmt.Printf("%+v", list)

		os.Exit(0)
	*/

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
	// close libvirt connection
	err := connector.CloseConnection()
	if err != nil {
		exitcode = 1
	}

	// close printer
	models.Collection.Printer.Close()

	// return exit code
	os.Exit(exitcode)
}
