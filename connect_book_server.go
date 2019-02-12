package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func getAllBook(){
	msg := make (map[string]interface{})
	msg["cmd"] = "getAllBooks"
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	sendMsg(msgS)
}

func addBook(name, author string, year int16){
	msg := make (map[string]interface{})
	msg["cmd"] = "addBook"
	params := make (map[string]interface{})
	params["year"] = year
	params["author"] = author
	params["name"] = name
	msg["params"] = params
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	sendMsg(msgS)
}

func removeBook(id int32) {
	msg := make (map[string]interface{})
	msg["cmd"] = "deleteBook"
	params := make (map[string]interface{})
	params["id"] = id
	msg["params"] = params
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	sendMsg(msgS)
}

func sendMsg(msg []byte) {
	conn, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		fmt.Printf("Error while connect %v\n", err)
		return
	}
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Printf("Error while send %v\n", err)
		return
	}
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF{
		fmt.Printf("Error while read %v\n", err)
		return
	}
	fmt.Printf("response %v\n", response)
}
