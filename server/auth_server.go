package main

import (
	"encoding/json"
	"fmt"
	"gochat/server/store"
	"gochat/utils"
	"log"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var FileAuth = new(store.FileAuth)
var rounds = 13

func (s *server) handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Invalid request method in signup")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status:"Invalid request method to signup"}`))
	} else {
		var authPacket utils.AuthPacket
		err := json.NewDecoder(r.Body).Decode(&authPacket)
		if err != nil {
			log.Println("Error decoding request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"status":"Invalid request body"}`))
			return
		}
		username := authPacket.Username
		passwd := authPacket.Password
		hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), rounds)
		err = FileAuth.Store(username, string(hashedPasswd))

		if err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status":"db error on signup"}`))
			return
		}
		log.Println("Signed Up")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"signup successful"}`))
		return
	}
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Invalid request method in login")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"Invalid request method to login"}`))
	} else {
		var authPacket utils.AuthPacket
		err := json.NewDecoder(r.Body).Decode(&authPacket)
		if err != nil {
			log.Println("Error decoding request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"status":"Invalid request body"}`))
			return
		}
		username := authPacket.Username
		passwd := authPacket.Password
		storedPasswd, err := FileAuth.GetByUsername(username)
		if err != nil {
			log.Println("login err", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status":"db error or wrong username"}`))
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswd), []byte(passwd))
		if err != nil {
			log.Println("login err", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status":"wrong password"}`))
			return
		}
		log.Println("Logged In")
		//generate 3 random int for session token
		r1 := rand.Intn(1000000)
		r2 := rand.Intn(1000000)
		r3 := rand.Intn(1000000)
		token:= (r1*r2*r3)%1000000000
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"login successful"}`))
		w.Write([]byte(fmt.Sprintf(`{"token":"%d"}`, token)))
		return
	}
}
