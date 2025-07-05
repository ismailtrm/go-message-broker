package main

import (
	"encoding/json"
	"fmt"
	"go-message-broker/broker"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/publish", handlePublish)
	http.HandleFunc("/consume", handleConsume)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg broker.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	broker.GlobalQueue.Enqueue(msg)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Message published successfully")
}

func handleConsume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	msg, ok := broker.GlobalQueue.Dequeue()
	if !ok {
		http.Error(w, "No messages in queue", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
