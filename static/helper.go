package static

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	"log"
)

type VirtualMachineExtended struct {
		models.VirtualMachine
		StaticData models.Statistic
}

func readItems(vm models.VirtualMachine) (models.Statistic, error) {
	/*
	# virsh dominfo xy
	Id:             4
	Name:           instance-0000012c
	UUID:           5f065a85-9a0f-402c-8470-681155a10021
	OS Type:        hvm
	State:          running
	CPU(s):         2
	CPU time:       15240.3s
	Max memory:     4194304 KiB
	Used memory:    4194304 KiB
	Persistent:     yes
	Autostart:      disable
	Managed save:   no
	Security model: none
	Security DOI:   0
	*/
	staticData, err := util.VirshXDetails("dominfo", vm.Name(), "", 0, 1, nil)
	if err != nil {
		log.Printf("Error while readItems %a", err)
		return models.Statistic{}, err
	}	
	return staticData, nil
}


