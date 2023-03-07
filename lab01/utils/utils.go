package utils

import (
	"net"
	"os"
	"path/filepath"
)

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

func GetAsciiArt() string {
	absPath, _ := filepath.Abs("./utils/files/asciiart.txt")
	content, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	return string(content[:])
}
