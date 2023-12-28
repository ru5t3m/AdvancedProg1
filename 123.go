package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct to unmarshal incoming JSON data
type RequestData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/handlepost", handlePost)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON data from the request body
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", requestData.Message)

	// You can perform actions with the received data here

	// Send a response indicating success
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data received successfully")
}
