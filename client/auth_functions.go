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

type AuthPacket utils.AuthPacket

func sendSignUp (authPacket AuthPacket) error{
	jsondata,err := json.Marshal(authPacket)
	if (err!=nil){
		return err
	}
	resp, err := http.Post(URL+"/signup", "application/json", bytes.NewBuffer(jsondata))
	if err!=nil || resp.StatusCode != http.StatusOK{
		log.Printf("#+v \n", err)
		fmt.Println("Error signing up!")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}

func sendLogin (token string){

}