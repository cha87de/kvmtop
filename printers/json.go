package printers

import (
	"fmt"
	"strconv"

	"github.com/cha87de/kvmtop/models"
)

// JSONPrinter describes the json printer
type JSONPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *JSONPrinter) Open() {
	outputOpen()
}

// Screen prints the measurements on the screen
func (printer *JSONPrinter) Screen(fields []string, values [][]string) {
	output(fmt.Sprintf("["))
	for i, domvalue := range values {
		if i > 0 {
			output(fmt.Sprintf(","))
		}
		output(fmt.Sprintf("{"))
		for j, value := range domvalue {
			if j > 0 {
				output(fmt.Sprintf(","))
			}

			// but """ only for strings
			if _, err := strconv.ParseInt(value, 10, 64); err == nil {
				output(fmt.Sprintf("\"%s\": %s", fields[j], value))
			} else if _, err := strconv.ParseFloat(value, 64); err == nil {
				output(fmt.Sprintf("\"%s\": %s", fields[j], value))
			} else {
				output(fmt.Sprintf("\"%s\": \"%s\"", fields[j], value))
			}
		}
		output(fmt.Sprintf("}"))
	}
	output(fmt.Sprintf("]\n"))
}

// Close terminates the printer
func (printer *JSONPrinter) Close() {
	outputClose()
}

// CreateJSON creates a new simple text printer
func CreateJSON() JSONPrinter {
	return JSONPrinter{}
}
