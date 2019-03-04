package printers

import (
	"fmt"
	"net"
	"os"

	"github.com/cha87de/kvmtop/config"
)

var conn net.Conn
var file *os.File

// OutputOpen establishes the output channel for the printer
func OutputOpen() {
	if config.Options.Output == "tcp" || config.Options.Output == "udp" {
		var err error
		conn, err = net.Dial(config.Options.Output, config.Options.OutputTarget)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open %s connection", config.Options.Output)
			os.Exit(1)
		}
		fmt.Printf("Output will be redirected to %s://%s\n", config.Options.Output, config.Options.OutputTarget)
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

// Output takes a string and prints it to the configured output
func Output(text string) {
	if config.Options.Output == "tcp" || config.Options.Output == "udp" {
		_, err := fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to send to network connection")
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

// OutputClose closes the channel for printing the output
func OutputClose() {
	if config.Options.Output == "tcp" || config.Options.Output == "udp" {
		conn.Close()
	} else if config.Options.Output == "file" {
		file.Close()
	} else {
		// stdout, nothing to close
	}
}
