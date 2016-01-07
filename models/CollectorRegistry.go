package models

type Collector interface {
    Collect(vm VirtualMachine) (string,error)
    CollectDetails(vm VirtualMachine)
}

var(
	collectors []Collector 	
)

func RegisterCollector(collector Collector){
	collectors = append(collectors, collector)
}
func GetCollectors() []Collector {
	return collectors
}