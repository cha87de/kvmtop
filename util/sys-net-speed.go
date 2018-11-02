package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// SysNetSpeed reflects network interface speed from /sys/class/net/$iface/speed
type SysNetSpeed struct {
	Value float32
}

// GetSysNetSpeed reads the network device's speed from /sys/class/net
func GetSysNetSpeed(devName string) SysNetSpeed {
	stat := SysNetSpeed{}
	filepath := fmt.Sprint("/sys/class/net/" + devName + "/speed")
	filecontent, _ := ioutil.ReadFile(filepath)
	fmt.Fscan(
		bytes.NewBuffer(filecontent),
		&stat.Value,
	)
	return stat
}
