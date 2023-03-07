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
	clientsListUDP = make(map[string]net.PacketConn)
	//leaveHandlerChannel, messageHandlerChannel - kanały do komunikacji między gorutynami
	leaveHandlerChannel      = make(chan utils.Message)
	messageHandlerChannel    = make(chan utils.Message)
	quit                     = make(chan struct{})
	messageHandlerUDPChannel = make(chan utils.Message)
)

func main() {
	// z racji na nazwę funkcji z pakietu net - Conn - w moim odczuciu nie ma sensu na siłę wprowadzać
	// pojęcia socketa zatem Java'owy czy Python'owy socket to w nomenklaturze Go jest connection
	connection, err := net.Listen(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	if err != nil {
		panic(err)
	}

	udpServer, err := net.ListenUDP(configuration.TYPEUDP, &configuration.AddressUDP)
	if err != nil {
		panic(err)
	}

	go broadcastActions()
	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleRequest(conn, udpServer)
	}
}

func handleRequest(connection net.Conn, udpServer net.PacketConn) {
	clientsList[connection.RemoteAddr().String()] = connection
	//buf := make([]byte, 1024)
	//_, addr, err := udpServer.ReadFrom(buf)
	//if err != nil {
	//	panic(err)
	//}
	//clientsListUDP[addr.String()] = udpServer

	// obsługa akcji wiadomości
	messageHandlerChannel <- utils.NewMessage(" joined to the chat.", connection)

	// czytaj dane z połączenia
	messageScanner := bufio.NewScanner(connection)
	for messageScanner.Scan() {
		if messageScanner.Text() == "U" {
			messageHandlerUDPChannel <- utils.NewMessage(messageScanner.Text(), connection)
		} else {
			messageHandlerChannel <- utils.NewMessage(": "+messageScanner.Text(), connection)
		}
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
		case msg := <-messageHandlerUDPChannel:
			handleMessageUDP(msg)
		}
	}
}

func handleMessage(msg utils.Message) {
	log.Printf("message not U; message value: %v\n", msg.Text)
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

func handleMessageUDP(msg utils.Message) {
	log.Printf("message U; message value: %v\n", msg.Text)
	//for _, udpChannel := range clientsListUDP {
	//	log.Printf("iterate in clientsListUDP: elem %v\n", udpChannel.RemoteAddr())
	//	if msg.Address == udpChannel.RemoteAddr().String() {
	//		continue
	//	}
	//	_, err := fmt.Fprintln(udpChannel, utils.GetAsciiArt())
	//	if err != nil {
	//		continue
	//	}
	//}
}

func handleLeave(msg utils.Message) {
	for _, conn := range clientsList {
		_, err := fmt.Fprintln(conn, msg.Text)
		if err != nil {
			continue
		}
	}
}
