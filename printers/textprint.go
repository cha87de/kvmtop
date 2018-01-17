package printers

import (
	"fmt"

	"github.com/cha87de/kvmtop/models"
)

var showheader = true

// TextPrinter describes the text printer
type TextPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *TextPrinter) Open() {
	// nothing to open
}

// Screen prints the measurements on the screen
func (printer *TextPrinter) Screen(fields []string, values [][]string) {
	if showheader {
		// iterate over fields
		for _, field := range fields {
			fmt.Printf("%s\t", field)
		}
		fmt.Print("\n")

		// deactivate header
		showheader = false
	}

	// iterate over domains
	for _, domvalue := range values {
		for _, value := range domvalue {
			fmt.Printf("%s\t", value)
		}
		fmt.Print("\n")
	}
}

// Close terminates the printer
func (printer *TextPrinter) Close() {
	// nothing to close
}

// CreateText creates a new simple text printer
func CreateText() TextPrinter {
	return TextPrinter{}
}
