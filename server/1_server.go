package main

import (
	"encoding/json"
	"fmt"
	"gochat/server/store"
	"gochat/utils"
	"log"
	"net/http"
)

var URL string = ":3001"
type server struct{
	addr string
	router *http.ServeMux
}

var FileStore = new(store.FileStore)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello"))
}

func (s *server) handleStore(w http.ResponseWriter, r *http.Request){
	var message utils.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"Invalid request body"}`))
		return
	}
	message.Status=2 //receive at reciver
	FileStore.Store(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"Message stored successfully"}`))
}

func (s *server) handleReceive (w http.ResponseWriter, r *http.Request){
	FileStore.GetAll()
	log.Println("gets messages")
	w.WriteHeader(http.StatusOK)
}

func (s *server) handleOnline(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		log.Println("Invalid request method in online")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"Invalid request method to online"}`))
		return
	}
	size:=len(InMemoryTokenStore.tokens)
	sizeStr := fmt.Sprintf("%d", size)
	log.Println("online:", size)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"online users count: ` + (sizeStr) + `"}`))
}



func main(){
	store.InitDB()

	//all interfaces can request
	s:=&server{
		addr:URL,
		router: http.NewServeMux(),
	}
	s.router.HandleFunc("POST /store", s.handleStore)
	s.router.HandleFunc("GET /receive", s.handleReceive)
	s.router.HandleFunc("POST /signup", s.handleSignUp)
	s.router.HandleFunc("POST /login", s.handleLogin)
	s.router.HandleFunc("GET /online", s.handleOnline)
	log.Println("Starting server on :3001")

	http.ListenAndServe(s.addr, s.router)
}