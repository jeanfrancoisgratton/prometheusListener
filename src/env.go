package main

import (
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
	"path/filepath"
)

type Config_s struct {
	//CAcert    string `json:"cacert"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
	Port      uint   `json:"port"`
	TargetDir string `json:"targetdir"`
}

func loadConfig() (Config_s, *cerr.CustomError) {
	var payload Config_s

	rcFile := filepath.Join("/etc", "prometheus", "prometheusListener.json")
	_, err := os.Stat(rcFile)
	// We need to create the environment file if it does not exist
	if os.IsNotExist(err) {
		f := fmt.Sprintf("Configuration file %s not found", rcFile)
		panic(f)
	}

	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return Config_s{}, &cerr.CustomError{Title: "Error reading the file", Message: err.Error()}
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return Config_s{}, &cerr.CustomError{Title: "Error unmarshalling JSON", Message: err.Error()}
	} else {
		return payload, nil
	}
}

func (cs Config_s) SaveEnvironmentFile() *cerr.CustomError {

	jStream, err := json.MarshalIndent(cs, "", "  ")
	if err != nil {
		return &cerr.CustomError{Title: err.Error(), Fatality: cerr.Fatal}
	}
	rcFile := filepath.Join("/etc", "prometheus", "prometheusListener.json")
	if err = os.WriteFile(rcFile, jStream, 0644); err != nil {
		return &cerr.CustomError{Title: "Unable to write JSON file", Message: err.Error(), Fatality: cerr.Fatal}
	}

	return nil
}

func setup() *cerr.CustomError {
	cfg := Config_s{}

	cfg.Cert = hf.GetStringValFromPrompt("Enter the path to your SSL certificate: ")
	cfg.Key = hf.GetStringValFromPrompt("Enter the path to its key: ")
	cfg.Port = uint(hf.GetIntValFromPrompt("Enter the port the listener should listen on: "))
	cfg.TargetDir = hf.GetStringValFromPrompt("Enter the path where the hostnames should be added/removed: ")

	return cfg.SaveEnvironmentFile()
}
