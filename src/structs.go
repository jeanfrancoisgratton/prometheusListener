package main

// The structure supports both JSON and YAML, but for now we might not fully cover YAML
type ListenerPayload_s struct {
	Targets []string          `json:"targets" yaml:"targets"`
	Labels  map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// This struct contains the JSON payload and the command it should act upon
type CommandPayload_s struct {
	Command         string            `json:"command"`           // Added JSON tag
	ListenerPayload ListenerPayload_s `json:"prometheus_target"` // Changed to match the JSON key
}

// The configuration infos needed to run the listener
type Config_s struct {
	//CAcert    string `json:"cacert"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
	Port      uint   `json:"port"`
	TargetDir string `json:"targetdir"`
}

type TargetInfo_s struct {
	Filename string            `json:"filename" yaml:"filename"`
	HostInfo ListenerPayload_s `json:"hostinfo" yaml:"hostinfo"`
}
