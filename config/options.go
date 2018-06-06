package config

// Options holds set of runtime configuration parameters
var Options struct {
	Version    bool   `short:"v" long:"version" description:"Show version"`
	Frequency  int    `short:"f" long:"frequency" description:"Frequency (in seconds) for collecting metrics" default:"1"`
	Runs       int    `short:"r" long:"runs" description:"Amount of collection runs" default:"-1"`
	LibvirtURI string `short:"c" long:"connection" description:"connection uri to libvirt daemon" default:"qemu:///system"`
	ProcFS     string `long:"procfs" description:"path to the proc filesystem" default:"/proc"`
	Verbose    bool   `long:"verbose" description:"Verbose output, adds more detailed fields"`

	EnableCPU  bool `long:"cpu" description:"enable cpu metrics"`
	EnableMEM  bool `long:"mem" description:"enable memory metrics"`
	EnableDISK bool `long:"disk" description:"enable disk metrics"`
	EnableNET  bool `long:"net" description:"enable network metrics"`
	EnableIO   bool `long:"io" description:"enable io metrics (requires root)"`

	PrintBatch bool   `short:"b" long:"batch" description:"[DEPRECATED: use --printer=text instead] use simple output e.g. for scripts"`
	Printer    string `short:"p" long:"printer" description:"the output printer to use (valid printers: ncurses, text, json)" default:"ncurses"`

	Output       string `short:"o" long:"output" description:"the output channel to send printer output (valid output: stdout, file, tcp)" default:"stdout"`
	OutputTarget string `long:"target" description:"for output 'file' the location, for 'tcp' the url to the tcp server"`
}
