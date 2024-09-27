package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	switch payload.Command {
	case "add":
		perr := SavePayloadToFile(payload.ListenerPayload)
		if perr != nil {
			http.Error(w, fmt.Sprintf("Failed to save payload: %v", perr), http.StatusInternalServerError)
			return
		}
	case "rm":
		perr := RemoveHostFromList(payload.ListenerPayload.Targets[0])
		if perr != nil {
			http.Error(w, fmt.Sprintf("Unable to remove host %s: %v",
				payload.ListenerPayload.Targets[0], perr), http.StatusInternalServerError)
		}
	case "ls":
		hostsList, perr := ListTargets()
		if perr != nil {
			http.Error(w, fmt.Sprintf("Could not list hosts: %v", perr), http.StatusInternalServerError)
		}
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ListenerPayload received and saved."))
}
