package disk

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

type VirtualMachineExtended struct {
		models.VirtualMachine
		Items []models.MeasurementItem
		// key: name of Item
		Statistic map[string] models.Statistic
}

func (vmx *VirtualMachineExtended) lookupStats() (map[string] models.Statistic, error) {
	stats := make(map[string] models.Statistic)
	// loop through vm's interfaces
	for _, vmif := range vmx.Items {
		ifstat, err := readItemStats(vmx.VirtualMachine, vmif)
		if err != nil{
			continue;
		}
		stats[vmif.Name] = ifstat
	}
	return stats, nil
}

func readItems(vm models.VirtualMachine) ([]models.MeasurementItem, error) {
	/*
	# virsh domblklist xy
	Target     Source
	------------------------------------------------
	vda        /var/lib/nova/instances/9e703375-7ca7-4135-9534-b7aaeb5a14e5/disk
	*/
	list, err := util.VirshXList("domblklist", vm.Name())
	if err != nil {
		return nil, err
	}	
	return list, nil
}

func evalDiskSize(map[string] int64) int64{
	return 0
}

func readItemStats(vm models.VirtualMachine, vmif models.MeasurementItem) (models.Statistic, error) {
	/*
	Capacity:       21474836480
	Allocation:     483991552
	Physical:       483991552
	*/
	stat1, err1 := util.VirshXDetails("domblkinfo", vm.Name(), vmif.Name, 0, 1, nil)
	if err1 != nil {
		return models.Statistic{}, err1
	}
	
	/*
	hda rd_req 23316
	hda rd_bytes 461506990
	hda wr_req 0
	hda wr_bytes 0
	hda flush_operations 0
	hda rd_total_times 1514874067
	hda wr_total_times 0
	hda flush_total_times 0
	*/
	stat2, err2 := util.VirshXDetails("domblkstat", vm.Name(), vmif.Name, 1, 2, nil)
	if err2 != nil {
		return models.Statistic{}, err2
	}
	
	// copy stats from stat1 to stat2
	for k, v := range stat1.Values {
    	stat2.Values[k] = v
	}
	
	return stat2, nil
}


