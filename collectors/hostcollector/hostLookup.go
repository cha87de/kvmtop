package hostcollector

import (
	"os"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

func hostLookup(host *models.Host) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	host.AddMetricMeasurement("host_name", models.CreateMeasurement(name))

	if config.Options.Verbose {
		uuid := util.GetSysDmiUUID()
		host.AddMetricMeasurement("host_uuid", models.CreateMeasurement(uuid.Value))
	}
}
