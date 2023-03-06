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
	clientsList    = make(map[string]net.Conn)
	clientsListUDP = make(map[string]net.Conn)
	//leaveHandlerChannel, messageHandlerChannel - kanały do komunikacji między gorutynami
	leaveHandlerChannel   = make(chan utils.Message)
	messageHandlerChannel = make(chan utils.Message)
	quit                  = make(chan struct{})
	//messageHandlerUDPChannel = make(chan utils.Message)
)

func main() {
	// z racji na nazwę funkcji z pakietu net - Conn - w moim odczuciu nie ma sensu na siłę wprowadzać
	// pojęcia socketa zatem Java'owy czy Python'owy socket to w nomenklaturze Go jest connection
	connection, err := net.Listen(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	if err != nil {
		log.Fatal(err)
	}

	udpServer, err := net.ListenUDP(configuration.TYPEUDP, &configuration.AddressUDP)
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
		//buf := make([]byte, 1024)
		//_, addr, err := udpServer.ReadFrom(buf)
		//if err != nil {
		//	log.Println(err)
		//	continue
		//}
		go handleRequest(conn, udpServer)
	}
}

func handleRequest(connection net.Conn, udpServer net.PacketConn) {
	clientsList[connection.RemoteAddr().String()] = connection
	//clientsListUDP[net.UDPAddr{}]

	// obsługa akcji wiadomości
	messageHandlerChannel <- utils.NewMessage(" joined to the chat.", connection)

	messageScanner := bufio.NewScanner(connection)
	for messageScanner.Scan() {
		messageHandlerChannel <- utils.NewMessage(": "+messageScanner.Text(), connection)
	}

	// usuwanie zakończonego połączenia z klientem gdy ten się rozłączy
	delete(clientsList, connection.RemoteAddr().String())
	// obsługa akcji opuszczenia czatu
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
	if msg.Text == "U" {
		for _, updChannel := range clientsListUDP {
			if msg.Address == updChannel.RemoteAddr().String() {
				continue
			}
			_, err := fmt.Fprintln(updChannel, msg.Text)
			if err != nil {
				continue
			}
		}
		return
	}
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
