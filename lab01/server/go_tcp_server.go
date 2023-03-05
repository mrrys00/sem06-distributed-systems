package main

import (
	"bufio"
	"fmt"
	"lab01/configuration"
	"lab01/utils"
	"log"
	"net"
)

var (
	//clientsList - mapa rutyn (wątków na sterydach) port:połączenie
	// https://go.dev/tour/concurrency/1
	clientsList = make(map[string]net.Conn)
	//leaveHandlerChannel, messageHandlerChannel - kanały do komunikacji między gorutynami
	leaveHandlerChannel   = make(chan utils.Message)
	messageHandlerChannel = make(chan utils.Message)
)

func main() {
	// z racji na nazwę funkcji z pakietu net - Conn - w moim odczuciu nie ma sensu na siłę wprowadzać
	// pojęcia socketa zatem Java'owy czy Python'owy socket to w nomenklaturze Go jest connection
	connection, err := net.Listen(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	if err != nil {
		log.Fatal(err)
	}

	go broadcastActions()
	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(connection net.Conn) {
	clientsList[connection.RemoteAddr().String()] = connection

	// obsługa akcji wiadomości, obsługa akcji opuszczenia czatu
	messageHandlerChannel <- utils.NewMessage(" joined to the chat.", connection)

	messageScanner := bufio.NewScanner(connection)
	for messageScanner.Scan() {
		messageHandlerChannel <- utils.NewMessage(": "+messageScanner.Text(), connection)
	}

	delete(clientsList, connection.RemoteAddr().String())
	// usuwanie zakończonego połączenia z klientem gdy ten się rozłączy
	leaveHandlerChannel <- utils.NewMessage(" has left the chat.", connection)

	connection.Close()
}

func broadcastActions() {
	for {
		select {
		case msg := <-messageHandlerChannel:
			handleMessage(msg)
		case msg := <-leaveHandlerChannel:
			handleLeave(msg)
		}
	}
}

func handleMessage(msg utils.Message) {
	for _, connection := range clientsList {
		if msg.Address == connection.RemoteAddr().String() {
			continue
		}
		_, err := fmt.Fprintln(connection, msg.Text)
		if err != nil {
			continue
		}
	}
}

func handleLeave(msg utils.Message) {
	for _, conn := range clientsList {
		_, err := fmt.Fprintln(conn, msg.Text)
		if err != nil {
			continue
		}
	}
}
