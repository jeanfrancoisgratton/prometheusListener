package main

// The structure supports both JSON and YAML, but for now we might not fully cover YAML
type ListenerPayload_s struct {
	Targets []string          `json:"targets" yaml:"targets"`
	Labels  map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

type ListenerPayloads_s []ListenerPayload_s

// The configuration infos needed to run the listener
type Config_s struct {
	//CAcert    string `json:"cacert"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
	Port      uint   `json:"port"`
	TargetDir string `json:"targetdir"`
}
