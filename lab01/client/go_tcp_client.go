package main

import (
	"io"
	"lab01/configuration"
	"log"
	"net"
	"os"
)

var serverDown = make(chan struct{})

func main() {
	connection, err := net.Dial(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	if err != nil {
		panic(err)
	}
	//udpClient, err := net.Dial(configuration.TYPEUDP, configuration.HOST+":"+configuration.PORT)
	//if err != nil {
	//	panic(err)
	//}

	go func() {
		_, err := io.Copy(os.Stdout, connection)
		if err != nil {
			log.Print(err)
		}
		log.Println("finito")
		// wyślij sygnał o zakończeniu działania do głównej gorutyny
		serverDown <- struct{}{}
	}()

	clientAction(connection, os.Stdin)
	//clientActionUDP(udpClient, os.Stdin)

	connection.Close()
	//udpClient.Close()
	// oczekiwanie na zakończenie działającej "w tle" gorutyny z obsługą połączenia
	<-serverDown
}

func clientAction(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		panic(err)
	}
}

//func clientActionUDP(dst io.Writer, src io.Reader) {
//	if _, err := io.Copy(dst, src); err != nil {
//		panic(err)
//	}
//}
