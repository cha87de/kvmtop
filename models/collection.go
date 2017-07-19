package models

// Collection of domains and other stuff
var Collection struct {
	Domains map[string]Domain
}

// Domain defines a domain in libvirt
type Domain struct {
	Name string
	UUID string
}
