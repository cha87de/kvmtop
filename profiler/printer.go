package profiler

import (
	"encoding/json"
	"fmt"

	"github.com/cha87de/kvmtop/printers"
	"github.com/cha87de/tsprofiler/spec"
)

type profilerOutput struct {
	Profile spec.TSProfile `json:"profile"`
}

func profileOutput(data spec.TSProfile) {
	// text := printTransitMatrix(data.TXMatrix)
	profilerOutput := profilerOutput{
		Profile: data,
	}
	json, err := json.Marshal(profilerOutput)
	// fmt.Printf("%s\n", json)
	if err != nil {
		fmt.Printf("Error while marshaling profiler output (%+v): %+v", profilerOutput, err)
	} else {
		printers.Output(string(json) + "\n")
	}
}
