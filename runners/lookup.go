package runners

import (
	"fmt"
	"log"

	"github.com/cha87de/kvmtop/connector"
	libvirt "github.com/libvirt/libvirt-go"
)

func initializeLookup() {
	doms, err := connector.Libvirt.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Fatal("Cannot get list of domains form libvirt.")
	}
	//fmt.Printf("%d running domains:\n", len(doms))

	for _, dom := range doms {
		// query data from libvirt
		name, err := dom.GetName()
		if err != nil {
			continue
		}
		fmt.Println(name)
	}

}
