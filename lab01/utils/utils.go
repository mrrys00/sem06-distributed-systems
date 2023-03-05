package utils

import "net"

type Message struct {
	Text    string
	Address string
}

func NewMessage(msg string, conn net.Conn) Message {
	addr := conn.RemoteAddr().String()
	return Message{
		Text:    addr + msg,
		Address: addr,
	}
}
