package sender

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Sender struct {
	conn net.Conn
}

func Init() (*Sender, error) {
	conn, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		return nil, err
	}
	return &Sender{conn}, nil
}

func (s Sender)Close() {
	err := s.conn.Close()
	if err != nil {
		fmt.Printf("Error while disconnect %v\n", err)
	}
}

func (s Sender)GetAllBook(){
	msg := make (map[string]interface{})
	msg["cmd"] = "getAllBooks"
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	s.sendMsg(msgS)
}

func (s Sender)AddBook(name, author string, year int16){
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
	s.sendMsg(msgS)
}

func (s Sender)RemoveBook(id int32) {
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
	s.sendMsg(msgS)
}

func (s Sender)sendMsg(msg []byte) {
	_, err := s.conn.Write(append(msg, byte('\n')))
	if err != nil {
		fmt.Printf("Error while send %v\n", err)
		return
	}
	response, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil && err != io.EOF{
		fmt.Printf("Error while read %v\n", err)
		return
	}
	fmt.Printf("response %v\n", response)
	if err == io.EOF{
		fmt.Printf("Warning: server close connect!\n")
	}
}
