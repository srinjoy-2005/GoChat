package main

import (
	"encoding/json"
	_ "encoding/json"
	"gochat/server/store"
	"gochat/utils"
	"log"
	_ "log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var FileAuth = new(store.FileAuth)
var rounds = 13

func (s *server) handleSignUp(w http.ResponseWriter, r *http.Request){
	if (r.Method!=http.MethodPost){
		log.Println("Invalid request method in signup")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status:"Invalid request method to signup"}`))
	}else{
		var authPacket utils.AuthPacket
		err:= json.NewDecoder(r.Body).Decode(&authPacket)
		if err != nil {
			log.Println("Error decoding request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"status":"Invalid request body"}`))
			return
		}
		username:=authPacket.Username
		passwd:=authPacket.Password
		hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), rounds)
		err=FileAuth.Store(username, string(hashedPasswd))

		if (err!=nil){
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status":"db error on signup"}`))
			return
		}
		log.Println("Signed Up")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":"signup successful"}`))
		return
	}

}
