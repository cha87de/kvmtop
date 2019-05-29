package profiler

import (
	"encoding/json"
	"fmt"

	"github.com/cha87de/kvmtop/printers"
	"github.com/cha87de/tsprofiler/models"
)

type profilerOutput struct {
	Profile models.TSProfile `json:"profile"`
}

func profileOutput(data models.TSProfile) {
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
