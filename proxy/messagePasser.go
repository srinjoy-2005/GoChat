package main

import (
	"encoding/json"
	"fmt"
	"gochat/utils"
	"bytes"
	"log"
	"net/http"
)

var server_URL string= "http://localhost:3001"

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
		w.Write([]byte("[Message received successfully],"))
		message.Status=1;//received here

		jsonPayload, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "Failed to encode message", http.StatusInternalServerError)
			return
		}
		res, err := http.Post(server_URL+"/store", "application/json", bytes.NewBuffer(jsonPayload))	
		if err != nil || res.StatusCode != http.StatusOK {
			w.Write([]byte("[Failed to send to main server]"))
			log.Println("Error sending message to main server:", err)
			return
		}
		log.Printf("headers: #v, status: #v \n", res.Header, res.Status)
		w.Write([]byte("[Message sent to main server successfully]"))

		defer res.Body.Close()

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("Invalid request method - %s", r.Method)
		http.Error(w,msg,http.StatusMethodNotAllowed)
		return
	}
	res, err := http.Get(server_URL+"/receive")
	if (err != nil || res.StatusCode != http.StatusOK) {
		w.Write([]byte("[Failed to get messages from main server]"))
		log.Println("Error getting messages from main server:", err)
		return
	}
	defer res.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[hit /receive ]"))
}

func main(){
	http.HandleFunc("/message",messageHandler)
	http.HandleFunc("/receive",receiveHandler)
	log.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}