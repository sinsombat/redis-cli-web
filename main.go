package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Command struct {
	// Connection string `json:"connection"`
	// Password   string `json:"password"`
	Command string `json:"command"`
}

func executeRedisCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cmd Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if cmd.Command == "" {
		http.Error(w, "No command provided", http.StatusBadRequest)
		return
	}

	exec.Command("redis-cli", strings.Fields(cmd.Command)...).Output()

	result, err := exec.Command("redis-cli", strings.Fields(cmd.Command)...).Output()
	if err != nil {
		http.Error(w, "Error executing command: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": string(result)})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/execute", executeRedisCommand)

	port := "5000"
	fmt.Println("Server listening on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
