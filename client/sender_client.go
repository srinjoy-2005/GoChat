package main

import (
	"bytes"
	"fmt"
	"time"
	"gochat/utils"
)

const URL = "http://localhost:8080"

func main() {
	fmt.Println("Welcome to message client")

	var username, password string
	for {
		fmt.Println("Enter username for signup:")
		fmt.Scanln(&username)
		if len(bytes.TrimSpace([]byte(username))) > 0 {
			break
		}
		fmt.Println("Username cannot be empty, try again.")
	}

	for {
		fmt.Println("Enter password:")
		fmt.Scanln(&password)
		if len(bytes.TrimSpace([]byte(password))) > 0 {
			break
		}
		fmt.Println("Password cannot be empty, try again.")
	}

	var authPacket AuthPacket
	authPacket = AuthPacket{username, password}
	err := sendSignUp(authPacket)
	if err != nil {
		fmt.Println("Signup network error:", err)
		return
	}

	ans:=true
	var message utils.Message
	for ans==true{
		for{
			var content string
			fmt.Println("Enter the message to send to server")
			fmt.Scanln(&content)
			if len(bytes.TrimSpace([]byte(content))) > 0 {
				message.Content=content
				break
			}
			fmt.Println("Empty message, try again")
		}
		for{
			var send_name string
			fmt.Println("Enter the sender name")
			fmt.Scanln(&send_name)
			if len(bytes.TrimSpace([]byte(send_name))) > 0 {
				message.Name=send_name
				break
			}
			fmt.Println("Empty username, try again")
		}
		
		message.Time = uint32(time.Now().Unix())
		message.Status=0;//sent

		resp, err := sendMessage(message)
		if err != nil {
			fmt.Println("err: ",err)
			fmt.Println(string(resp))
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