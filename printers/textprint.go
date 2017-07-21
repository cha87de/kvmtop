package printers

import "fmt"

// Textprint prints the values separated by tab spaces
func Textprint(values []string) {
	for _, value := range values {
		fmt.Printf("%s\t", value)
	}
	fmt.Print("\n")
}
