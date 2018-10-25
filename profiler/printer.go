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
	json, _ := json.Marshal(profilerOutput)
	fmt.Printf("%s\n", json)

	printers.Output(string(json) + "\n")
}
