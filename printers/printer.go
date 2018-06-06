package printers

import (
	"fmt"
	"net"
	"os"

	"github.com/cha87de/kvmtop/config"
)

var tcpConn net.Conn
var file *os.File

func outputOpen() {
	if config.Options.Output == "tcp" {
		tcpConn, _ = net.Dial("tcp", config.Options.OutputTarget)
		fmt.Printf("Output will be redirected to tcp://%s\n", config.Options.OutputTarget)
	} else if config.Options.Output == "file" {
		var err error
		file, err = os.OpenFile(config.Options.OutputTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Output will be redirected to file://%s\n", config.Options.OutputTarget)
	} else {
		// stdout, nothing to open
	}
}

func output(text string) {
	if config.Options.Output == "tcp" {
		fmt.Fprintf(tcpConn, text)
	} else if config.Options.Output == "file" {
		if _, err := file.WriteString(text); err != nil {
			panic(err)
		}
	} else {
		// print to stdout
		fmt.Print(text)
	}
}

func outputClose() {
	if config.Options.Output == "tcp" {
		tcpConn.Close()
	} else if config.Options.Output == "file" {
		file.Close()
	} else {
		// stdout, nothing to close
	}
}
