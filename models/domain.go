package models

// Domain defines a domain in libvirt
type Domain struct {
	*Measurable
	Name string
	UUID string
	PID  int
}
