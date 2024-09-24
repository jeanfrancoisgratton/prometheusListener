package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SavePayloadToFile saves the received JSON payload into the /tmp/targets directory
func SavePayloadToFile(targets []string, payload CommandPayload_s) error {
	if len(targets) == 0 {
		return fmt.Errorf("no targets specified")
	}

	// Extract first target and remove port if present
	target := targets[0]
	targetName := strings.Split(target, ":")[0]

	// Define the directory and initial file path
	dir := "/tmp/targets"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	fileName := filepath.Join(dir, targetName+".json")

	// Check if file exists and add timestamp if necessary
	if _, err := os.Stat(fileName); err == nil {
		// File exists, append date
		timestamp := time.Now().Format("2006.01.02")
		fileName = filepath.Join(dir, fmt.Sprintf("%s_%s.json", targetName, timestamp))
	}

	// Marshal payload back to JSON for saving
	jsonData, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling payload to JSON: %v", err)
	}

	// Write the file
	if err := ioutil.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	log.Printf("Saved payload to %s\n", fileName)
	return nil
}

// Handler for incoming JSON POST requests
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the incoming JSON payload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON into the CommandPayload struct
	var payload CommandPayload_s
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Process the payload (write to file)
	err = SavePayloadToFile(payload.ListenerPayload.Targets, payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save payload: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ListenerPayload received and saved."))
}
