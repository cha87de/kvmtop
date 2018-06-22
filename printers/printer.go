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
		var err error
		tcpConn, err = net.Dial("tcp", config.Options.OutputTarget)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open tcp connection")
			os.Exit(1)
		}
		fmt.Printf("Output will be redirected to tcp://%s\n", config.Options.OutputTarget)
	} else if config.Options.Output == "file" {
		var err error
		file, err = os.OpenFile(config.Options.OutputTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open output file")
			os.Exit(1)
		}
		fmt.Printf("Output will be redirected to file://%s\n", config.Options.OutputTarget)
	} else {
		// stdout, nothing to open
	}
}

func output(text string) {
	if config.Options.Output == "tcp" {
		_, err := fmt.Fprintf(tcpConn, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to send to tcp connection")
			os.Exit(1)
		}
	} else if config.Options.Output == "file" {
		if _, err := file.WriteString(text); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write to file")
			os.Exit(1)
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
