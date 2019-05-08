package use_case

import (
	"encoding/json"
	"fmt"
)

type Sender interface {
	SendMsg(msg []byte)
}

type SendMsgInteractor struct {
	sender Sender
}

func NewSendMsgInteractor(s Sender) *SendMsgInteractor {
	return &SendMsgInteractor{sender: s}
}

func (s *SendMsgInteractor) GetAllBook() {
	msg := make(map[string]interface{})
	msg["cmd"] = "getAllBooks"
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	s.sender.SendMsg(msgS)
}

func (s *SendMsgInteractor) AddBook(name, author string, year int16) {
	msg := make(map[string]interface{})
	msg["cmd"] = "addBook"
	params := make(map[string]interface{})
	params["year"] = year
	params["author"] = author
	params["name"] = name
	msg["params"] = params
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	s.sender.SendMsg(msgS)
}

func (s *SendMsgInteractor) RemoveBook(id int32) {
	msg := make(map[string]interface{})
	msg["cmd"] = "deleteBook"
	params := make(map[string]interface{})
	params["id"] = id
	msg["params"] = params
	msgS, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	s.sender.SendMsg(msgS)
}