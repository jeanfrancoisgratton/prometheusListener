package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Saves the received JSON payload into the targets directory
func SavePayloadToFile(payload ListenerPayload_s) error {
	if len(payload.Targets) == 0 {
		return fmt.Errorf("No targets specified")
	}

	// Extract first target and remove port if present
	target := payload.Targets[0]
	targetName := strings.Split(target, ":")[0]

	// Define the directory and initial file path
	dir := cfg.TargetDir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	fileName := filepath.Join(dir, targetName+".json")

	//// Check if file exists and add timestamp if necessary
	//if _, err := os.Stat(fileName); err == nil {
	//	// File exists, append date
	//	timestamp := time.Now().Format("2006.01.02")
	//	fileName = filepath.Join(dir, fmt.Sprintf("%s_%s.json", targetName, timestamp))
	//}

	// Marshal payload back to JSON for saving
	jsonData, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling payload to JSON: %v", err)
	}

	// Write the file
	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	log.Printf("Saved payload to %s\n", fileName)
	return nil
}

// Removes the host named in the payload
func RemoveHostFromList(targetHost string) error {
	filename := fmt.Sprintf(filepath.Join(cfg.TargetDir, strings.Split(targetHost, ":")[0]))
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filename)
	}

	// Remove the file
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("error removing file: %v", err)
	}

	log.Printf("Removed file %s\n", filename)
	return nil
}

// Lists all configured targets
func ListTargets() ([]TargetInfo_s, error) {
	hostInfoList := []TargetInfo_s{}
	entries, err := os.ReadDir(cfg.TargetDir)

	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() { // Only list files
			lpayload := ListenerPayload_s{}
			jFile, err := os.ReadFile(filepath.Join(cfg.TargetDir, entry.Name()))
			if err != nil {
				return nil, fmt.Errorf("error reading JSON file: %v", err)
			}
			err = json.Unmarshal(jFile, lpayload)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
			}
			hostInfoList = append(hostInfoList, TargetInfo_s{Filename: entry.Name(), HostInfo: lpayload})
		}
	}
	//return hostInfoList, nil

	// Create response payload
	
	response := ListResponse{Files: files}
	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling response to JSON", http.StatusInternalServerError)
		return
	}
}
