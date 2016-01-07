package network

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

func statisticEvalFunction(statistics models.Statistic) int64 {
	var sum int64 = 0
	if _, exists := statistics.Values["tx_bytes"]; exists {
		sum = sum + statistics.GetValueAsInt("tx_bytes")	
	}
	if _, exists := statistics.Values["tx_bytes"]; exists {
		sum = sum + statistics.GetValueAsInt("tx_bytes")	
	}
	return sum
}

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
	# virsh domiflist xy
	Interface  Type       Source     Model       MAC
	-------------------------------------------------------
	tap07e88f58-5d bridge     qbr07e88f58-5d virtio      fa:16:3e:63:1c:a9
	*/
	list, err := util.VirshXList("domiflist", vm.Name())
	if err != nil {
		return nil, err
	}	
	return list, nil
}

func readItemStats(vm models.VirtualMachine, vmif models.MeasurementItem) (models.Statistic, error) {
	/*
	# virsh domiflist xy
	Interface  Type       Source     Model       MAC
	-------------------------------------------------------
	tap07e88f58-5d bridge     qbr07e88f58-5d virtio      fa:16:3e:63:1c:a9
	*/
	stat, err := util.VirshXDetails("domifstat", vm.Name(), vmif.Name, 1, 2, statisticEvalFunction)
	if err != nil {
		return models.Statistic{}, err
	}	
	return stat, nil
	
}

