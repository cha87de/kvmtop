package connector

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

// Libvirt provides access to the current libvirt connection
var Libvirt struct {
	ConnectionURI string
	Connection    *libvirt.Connect
}

// InitializeConnection connect to Libvirt.connectionURI
func InitializeConnection() error {
	conn, err := libvirt.NewConnect(Libvirt.ConnectionURI)
	if err != nil {
		log.Printf("Failed to connect to libvirt. %+v", err)
		return err
	}
	Libvirt.Connection = conn
	return nil
}

// CloseConnection close connection to libvirt
func CloseConnection() error {
	_, err := Libvirt.Connection.Close()
	if err != nil {
		log.Printf("Failed to close connection to libvirt. %+v", err)
		return err
	}
	return nil
}
