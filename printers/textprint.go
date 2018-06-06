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
	outputOpen()
}

// Screen prints the measurements on the screen
func (printer *TextPrinter) Screen(fields []string, values [][]string) {
	if showheader {
		// iterate over fields
		for _, field := range fields {
			output(fmt.Sprintf("%s\t", field))
		}
		output(fmt.Sprint("\n"))

		// deactivate header
		showheader = false
	}

	// iterate over domains
	for _, domvalue := range values {
		for _, value := range domvalue {
			output(fmt.Sprintf("%s\t", value))
		}
		output(fmt.Sprint("\n"))
	}
}

// Close terminates the printer
func (printer *TextPrinter) Close() {
	outputClose()
}

// CreateText creates a new simple text printer
func CreateText() TextPrinter {
	return TextPrinter{}
}
