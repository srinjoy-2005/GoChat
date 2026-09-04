package main

import (
	"encoding/json"
	"fmt";
	"net/http";
	"time";
	"bytes"
)

const URL = "http://localhost:8080"

func sendMessage(message Message, ) string {
	jsondata,err:=json.Marshal(message)
	if err!=nil {
		return "error in marshalling"
	}
	resp,err:=http.Post(URL,"application/json",bytes.NewBuffer(jsondata))
	if err!=nil {
		return "error in sending message"
	}
	resp.Body.Close()
	return "message sent successfully"
}

func main() {
	fmt.Println("Welcome to message client")

	ans:=true
	var message Message 
	for ans==true{
		fmt.Println("Enter the message to send to server")
		fmt.Scanln(&message.content)
		fmt.Println("Enter your name")
		fmt.Scanln(&message.name)
		message.time = uint32(time.Now().Unix())

		
		var continu string
		fmt.Println("Do you want to continue sending messages? (y/n)")
		fmt.Scanln(&continu)
		if continu == "n" {
			ans = false
		} 
	}
}	