package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os/exec"
)

var errCommandNotFound = errors.New("command not found")

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
}

func main() {
	http.HandleFunc("/api/cmd", handleCommand)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetCommand(w, r)
	case http.MethodPost:
		handlePostCommand(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	if command == "" {
		http.Error(w, "Missing 'command' query parameter", http.StatusBadRequest)
		return
	}

	executeCommandAndSendResponse(w, command)
}

func handlePostCommand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var cmdReq CommandRequest
	err = json.Unmarshal(body, &cmdReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	command := cmdReq.Command
	if command == "" {
		http.Error(w, "Missing 'command' in request body", http.StatusBadRequest)
		return
	}

	executeCommandAndSendResponse(w, command)
}

func executeCommandAndSendResponse(w http.ResponseWriter, command string) {
	output, err := executeCommand(command)
	if err != nil {
		if err == errCommandNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := CommandResponse{
		Output: output,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 127 {
				return "", errCommandNotFound
			}
		}
		return "", err
	}
	return string(output), nil
}
