package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"gochat/utils"
)

type AuthPacket utils.AuthPacket

func sendSignUp (authPacket AuthPacket) error{
	jsondata,err := json.Marshal(authPacket)
	if (err!=nil){
		return err
	}
	resp, err := http.Post(URL+"/signup", "application/json", bytes.NewBuffer(jsondata))
	if err!=nil || resp.StatusCode != http.StatusOK{
		log.Printf("#v\n", err)
		fmt.Println("Error signing up!")
	}

	return nil
}

func sendLogin (token string){

}