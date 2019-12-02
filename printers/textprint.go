package printers

import (
	"fmt"

	"kvmtop/models"
)

var showheader = true

// TextPrinter describes the text printer
type TextPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *TextPrinter) Open() {
	OutputOpen()
}

// Screen prints the measurements on the screen
func (printer *TextPrinter) Screen(printable models.Printable) {
	fields := printable.DomainFields
	values := printable.DomainValues

	if showheader {
		// iterate over fields
		for _, field := range fields {
			Output(fmt.Sprintf("%s\t", field))
		}
		Output(fmt.Sprint("\n"))

		// deactivate header
		showheader = false
	}

	// iterate over domains
	for _, domvalue := range values {
		for _, value := range domvalue {
			Output(fmt.Sprintf("%s\t", value))
		}
		Output(fmt.Sprint("\n"))
	}
}

// Close terminates the printer
func (printer *TextPrinter) Close() {
	OutputClose()
}

// CreateText creates a new simple text printer
func CreateText() TextPrinter {
	return TextPrinter{}
}
