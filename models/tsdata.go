package models

// TSData describes a monitoring item from kvmtop
type TSData struct {
	Host    map[string]string        `json:"host"`
	Domains []map[string]interface{} `json:"domains"`
}
