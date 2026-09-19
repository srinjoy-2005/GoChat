package main

import (
	"bytes"
	"fmt"
	"gochat/utils"
	"time"
)

const URL = "http://localhost:8080"

func main() {
	fmt.Println("Welcome to message client")

	//auth loop
	for {
		fmt.Println("1. Signup")
		fmt.Println("2. Login")
		fmt.Println("q. Exit")

		var choice string
		fmt.Scanln(&choice)
		
		var username, password string
		// Get username/password
		for {
			fmt.Println("Enter username:")
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

		authPacket := AuthPacket{
			username,
			password,
		}


		if choice == "1" {
			err := sendSignUp(authPacket)
			if err != nil{
				fmt.Println("Signup failed:", err)
				continue
			}

			fmt.Println("Signup successful!")

		} else if choice == "2" {
			token, err := sendLogin(authPacket)
			if err != nil {
				fmt.Println("Login failed:", err)
				continue
			}

			fmt.Println("Login successful!")
			fmt.Println("Token:", token)

			// Now enter messaging loop
			break

		} else if choice == "q" {
			return
		}
	}


	//messaging loop
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