package senderTCP

import (
	"bufio"
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

func (s Sender) Close() {
	err := s.conn.Close()
	if err != nil {
		fmt.Printf("Error while disconnect %v\n", err)
	}
}

func (s Sender) SendMsg(msg []byte) {
	_, err := s.conn.Write(append(msg, byte('\n')))
	if err != nil {
		fmt.Printf("Error while send %v\n", err)
		return
	}
	response, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error while read %v\n", err)
		return
	}
	fmt.Printf("response %v\n", response)
	if err == io.EOF {
		fmt.Printf("Warning: server close connect!\n")
	}
}
