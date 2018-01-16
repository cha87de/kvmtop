package config

// Options holds set of runtime configuration parameters
var Options struct {
	Version    bool   `short:"v" long:"version" description:"Show version"`
	Frequency  int    `short:"f" long:"frequency" description:"Frequency (in seconds) for collecting metrics" default:"1"`
	Runs       int    `short:"r" long:"runs" description:"Amount of collection runs" default:"-1"`
	LibvirtURI string `short:"c" long:"connection" description:"connection uri to libvirt daemon" default:"qemu:///system"`

	EnableCPU  bool `long:"cpu" description:"enable cpu metrics"`
	EnableMEM  bool `long:"mem" description:"enable memory metrics"`
	EnableDISK bool `long:"disk" description:"enable disk metrics"`
	EnableNET  bool `long:"net" description:"enable network metrics"`

	PrintBatch bool `short:"b" long:"batch" description:"use simple output e.g. for scripts"`
}
