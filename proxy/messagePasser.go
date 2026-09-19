package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gochat/utils"
	"io"
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
		log.Printf("headers: %v, status: %v \n", res.Header, res.Status)
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("[Failed to get messages from main server]"))
		log.Println("Error getting messages from main server:", err)
		return
	}
	defer res.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[hit /receive ]"))
}

func signupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	res, err := http.Post(server_URL+"/signup", "application/json", r.Body)
	if (err != nil || res.StatusCode != http.StatusOK) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("[Failed to signup with main server]"))
		log.Println("Error signing up with main server:", err)
		return
	}
	defer res.Body.Close()
	log.Println("Signed up successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[Signed up with main server successfully]"))
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	res,err := http.Post(server_URL+"/login", "application/json", r.Body)
	if (err != nil || res.StatusCode != http.StatusOK) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("[Failed to login with main server]"))
		log.Println("Error logging in with main server:", err)
		return
	}
	defer res.Body.Close()
	log.Println("Logged in successfully")
	token, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading login response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("[Failed to read login response]"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func main(){
	http.HandleFunc("/message",messageHandler)
	http.HandleFunc("/receive",receiveHandler)

	http.HandleFunc("/signup",signupHandler)
	http.HandleFunc("/login",loginHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}