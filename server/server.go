package main

import (
	"encoding/json"
	"gochat/utils"
	"log"
	"net/http";
	"gochat/server/store"
)

const server_URL string = "localhost:3001"

var FileStore = new(store.FileStore)

// func store(message utils.Message) {
// 	store.Store(message)
// 	log.Println("Message stored successfully")
// }

func storeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message utils.Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		message.Status = 2 //received at server
		FileStore.Store(message)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[Message stored successfully]"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main(){
	http.HandleFunc("/", storeHandler)
	log.Println("Starting server on :3001")
	log.Fatal(http.ListenAndServe(server_URL, nil))	
}