package senderNATS

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

type Sender struct {
	conn *nats.Conn
}

func Init() (*Sender, error) {
	conn, err := nats.Connect("localhost:4222")
	if err != nil {
		return nil, err
	}
	return &Sender{conn}, nil
}

func (s *Sender) Close() {
	s.conn.Close()
}

func (s *Sender) SendMsg(msg []byte) {
	response, err := s.conn.Request("books", msg, 1*time.Second)

	if err != nil {
		fmt.Printf("Error while send %v\n", err)
		return
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(response.Data, &data)
	fmt.Printf("response %v\n", data)
}
