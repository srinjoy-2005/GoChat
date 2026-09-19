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
	if err!=nil{
		return err
	}
	if resp.StatusCode != http.StatusOK{
		log.Printf("%+v\n", err)
		return fmt.Errorf("Error signing up with main server")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}

func sendLogin (authPacket AuthPacket) (string, error) {
	jsondata,err := json.Marshal(authPacket)
	if (err!=nil){
		return "", err
	}
	resp, err := http.Post(URL+"/login", "application/json", bytes.NewBuffer(jsondata))
	if err!=nil || resp.StatusCode != http.StatusOK{
		log.Printf("%+v\n", err)
		return "", fmt.Errorf("Error logging in with main server: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading login response: %v", err)
	}
	return string(body), nil
}