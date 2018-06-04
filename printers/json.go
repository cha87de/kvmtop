package printers

import (
	"fmt"

	"github.com/cha87de/kvmtop/models"
)

// JSONPrinter describes the json printer
type JSONPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *JSONPrinter) Open() {
	// nothing to open
}

// Screen prints the measurements on the screen
func (printer *JSONPrinter) Screen(fields []string, values [][]string) {
	fmt.Printf("[")
	for _, domvalue := range values {
		fmt.Printf("{")
		for i, value := range domvalue {
			fmt.Printf("\"%s\": \"%s\",", fields[i], value)
		}
		fmt.Printf("},")
	}
	fmt.Printf("]\n")
}

// Close terminates the printer
func (printer *JSONPrinter) Close() {
	// nothing to close
}

// CreateJSON creates a new simple text printer
func CreateJSON() JSONPrinter {
	return JSONPrinter{}
}
