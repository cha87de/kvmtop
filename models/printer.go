package models

// Printer defines a printer for output
type Printer interface {
	Open()
	Screen(Printable)
	Close()
}
