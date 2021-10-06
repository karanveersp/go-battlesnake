package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Response Structs

// BattlesnakeInfoResponse represents the basic
// API info and customization parameters.
type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

// HTTP Handlers

// HandleIndex handles the request to "/" with basic API info.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)

	log.Printf("Starting Battlesnake Server at http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
