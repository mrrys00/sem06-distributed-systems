package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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

	portBuff := make([]byte, 16)
	for {
		_, err := s2.Read(portBuff)
		if err != nil {
			log.Printf("[ERROR] unable to read message: %s", err.Error())
			break
		}
		if len(portBuff) != 0 {
			log.Printf("[MSG] %s", string(portBuff))
			break
		}
	}
	listenUDPPort, err := strconv.Atoi(string(portBuff[:5]))

	sUDPListen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: listenUDPPort,
	})
	if err != nil {
		log.Printf("some listening error : %+v\n", err)
	} else {
		log.Println("no UDP errors")
	}
	defer sUDPListen.Close()

	go handleMessages(s2)
	go handleUDP(sUDPListen)
	go handleConnection(s2, sUDP)

	for {
		//fmt.Print("client!")
		time.Sleep(time.Second)
	}
}

func handleConnection(s2 net.Conn, s *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		if text != "U\n" {
			s2.Write(data)
		} else {
			s.Write([]byte("a chuuj"))
		}
	}
}

func handleMessages(s2 net.Conn) {
	message := make([]byte, 2048)
	for {
		_, err := s2.Read(message)
		if err != nil {
			log.Printf("[ERROR] unable to read message: %s", err.Error())
			break
		}

		log.Printf("[MSG] %s", string(message[:]))
	}
}

func handleUDP(s *net.UDPConn) {
	message := make([]byte, 2048)
	//log.Printf("listening UDP on port %v\n", s.RemoteAddr().String())
	for {
		n, _, err := s.ReadFromUDP(message)
		if err != nil {
			log.Printf("[ERROR] unable to read message: %s", err.Error())
			break
		}

		log.Printf("[MSG] %s", string(message[0:n]))
	}
}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
