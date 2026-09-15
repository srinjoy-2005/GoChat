package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"gochat/utils"
)

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