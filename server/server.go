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
			log.Println("Error decoding request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"Invalid request body"}`))
			return
		}
		message.Status = 2 //received at server
		FileStore.Store(message)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Message stored successfully"}`))
	} else {
		log.Println("Invalid request method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Invalid request method"}`))
	}
}

func main(){
	store.InitDB() 
	defer store.DB.Close() 

	http.HandleFunc("/", storeHandler)
	log.Println("Starting server on :3001")
	log.Fatal(http.ListenAndServe(server_URL, nil))	
}