package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	statusCode, message := handleFileOperation(r)
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func handleFileOperation(r *http.Request) (int, string) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, `{"error": "Only POST method is allowed"}`
	}

	action := r.URL.Query().Get("cmd")
	clientHost := r.URL.Query().Get("hostname")

	if action == "" || clientHost == "" {
		return http.StatusBadRequest, `{"error": "Missing action or hostname parameters"}`
	}

	if action != "add" && action != "rm" {
		return http.StatusBadRequest, `{"error": "Invalid action. Use 'add' or 'rm'"}`
	}

	//filename := r.URL.Query().Get("filename")
	targetHost := filepath.Join(cfg.TargetDir, clientHost)

	mu.Lock()
	defer mu.Unlock()

	switch action {
	case "add":
		if _, err := os.Stat(targetHost); !os.IsNotExist(err) {
			return http.StatusConflict, `{"error": "Host already present in inventory"}`
		}
		err := os.WriteFile(targetHost, []byte("Added by "+clientHost+"\n"), 0644)
		if err != nil {
			return http.StatusInternalServerError, `{"error": "` + err.Error() + `"}`
		}
		return http.StatusOK, `{"message": "Host added successfully"}`

	case "rm":
		if _, err := os.Stat(targetHost); os.IsNotExist(err) {
			return http.StatusNotFound, `{"error": "Host is not in inventory"}`
		}
		err := os.Remove(targetHost)
		if err != nil {
			return http.StatusInternalServerError, `{"error": "` + err.Error() + `"}`
		}
		return http.StatusOK, `{"message": "Host removed successfully"}`
	}

	return http.StatusBadRequest, `{"error": "Unknown action"}`
}
