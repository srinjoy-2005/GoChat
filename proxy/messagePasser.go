package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"gochat/utils"
)

const server_URL ="http://localhost:3001"

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message utils.Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return 
		}
		fmt.Printf("Received message: %#v\n", message)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message received successfully"))

		res, err := http.Post(server_URL, "application/json", r.Body)
		if err != nil {
			http.Error(w, "Error forwarding message to server", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main(){
	http.HandleFunc("/message",messageHandler)
	log.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}