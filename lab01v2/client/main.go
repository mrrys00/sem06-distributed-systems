package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	s, err := net.ResolveUDPAddr("udp", "localhost:8080")
	sUDP, err := net.DialUDP("udp", nil, s)
	checkError(err)

	s2, err := net.Dial("tcp", "localhost:8080")
	checkError(err)
	defer sUDP.Close()
	defer s2.Close()

	// w rutyne wpierdolić czytanie tekstu
	//go handleUDP(sUDP)
	go handleConnection(s2)
	go handleMessages(s2)

	for {
		//fmt.Print("client!")
		time.Sleep(time.Second)
	}
}

func handleConnection(s2 net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("tcp>> ")
		text, _ := reader.ReadString('\n')
		s2.Write([]byte(text + "\n"))
	}
}

func handleMessages(s2 net.Conn) {
	message := make([]byte, 2048)
	for {

		_, err := s2.Read(message)
		if err != nil {
			log.Printf("[ERROR] unable to read message: %s", err.Error())
		}

		log.Printf("[MSG] %s", string(message))
	}
}

func handleUDP(s *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("udp>> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err := s.Write(data)
		checkError(err)
	}
}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
