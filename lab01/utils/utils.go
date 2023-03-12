package utils

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	ASCIIART = "───▄▄▄\n─▄▀░▄░▀▄\n─█░█▄▀░█\n─█░▀▄▄▀█▄█▄▀\n▄▄█▄▄▄▄███▀\n"
)

func PrintMessage(msg string) {
	log.Printf("%s", msg)
	PrintTerminal()
}

func PrintTerminal() {
	fmt.Printf(">> ")
}

func CreateUDPAddr(ip string, port int) *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
}

func GetPortFromTCP(addr net.Addr) (int, error) {
	resArr := strings.Split(addr.String(), ":")
	res, err := strconv.Atoi(resArr[1])

	return res, err
}
