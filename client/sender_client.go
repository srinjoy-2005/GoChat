package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"gochat/utils"
)

const URL = "http://localhost:8080"

func sendMessage(message utils.Message) ([]byte,error) {
	jsondata,err:=json.Marshal(message)
	if err!=nil {
		return []byte("error in marshalling"), err
	}
	resp,err:=http.Post(URL+"/message","application/json",bytes.NewBuffer(jsondata))
	if err!=nil {
		return []byte("error in sending message"),err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func main() {
	fmt.Println("Welcome to message client")

	ans:=true
	var message utils.Message
	for ans==true{
		fmt.Println("Enter the message to send to server")
		fmt.Scanln(&message.Content)
		fmt.Println("Enter your name")
		fmt.Scanln(&message.Name)
		message.Time = uint32(time.Now().Unix())
		message.Status=0;//sent

		resp, err := sendMessage(message)
		if err != nil {
			fmt.Println("err: ",err)
			fmt.Println((string)(resp))
		} else{
			fmt.Println("response: ",string(resp))
		}
		
		var continu string
		fmt.Println("Do you want to continue sending messages? (y/n)")
		fmt.Scanln(&continu)
		if continu == "n" {
			ans = false
		} 
	}
}	