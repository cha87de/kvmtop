package static

import (
	"github.com/cha87de/kvmtop/models"
	"fmt"
	"log"
	"bytes"
)

var (
	// store static information about VMs
	vms map[string] VirtualMachineExtended
)

type StaticCollector struct {

}

func (collector StaticCollector) Collect(vm models.VirtualMachine) (string, error) {
		vmx := vms[vm.Name()]
		// nothing to do, just look up the osUuid and print it as a result
		vmOsUuid := vmx.StaticData.Values["UUID"]
		result := fmt.Sprintf("%s", vmOsUuid)
		return result, nil
}

func (collector StaticCollector) CollectDetails(vm models.VirtualMachine) {
	staticData, err := readItems(vm)
	if err != nil {
		log.Printf("Error while CollectDetails in StaticCollector %a", err)
		return
	}
	if vmx, exists := vms[vm.Name()]; exists {
		vmx.StaticData = staticData
		vms[vm.Name()] = vmx
	}else{
		vms[vm.Name()] = VirtualMachineExtended{vm, staticData}
	}

}

func DefineFlags() {
	//flag.BoolVar(&CPU_EACH, "cpu-each", CPU_EACH, "CPU each")
}

func PrintHeader(buffer *bytes.Buffer) {
	buffer.WriteString("UUID\t")	
}

func Initialize() {
	vms = make(map[string] VirtualMachineExtended)
	models.RegisterCollector(StaticCollector{})
}
