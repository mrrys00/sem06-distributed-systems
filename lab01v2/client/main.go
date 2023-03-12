package main

import (
	"fmt"
	"lab01v2/client/handlers"
	"lab01v2/configuration"
	"lab01v2/logs"
	"lab01v2/utils"
	"net"
	"time"
)

var (
	clientID int
)

func main() {
	s, err := net.ResolveUDPAddr(configuration.TYPEUDP, configuration.HOST+":"+configuration.PORT)
	sUDP, err := net.DialUDP(configuration.TYPEUDP, nil, s)
	logs.LogFatal(err, "Cannot connect to UDP socket")

	sTCP, err := net.Dial(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	logs.LogFatal(err, "Cannot connect to TCP socket")

	sUDPListen := handlers.HandleChannelUDP(sTCP, &clientID)

	defer sUDPListen.Close()
	defer sTCP.Close()
	defer sUDP.Close()

	go handlers.HandleTCPIncomingMsg(sTCP)
	go handlers.HandleUDPIncomingMsg(sUDPListen)
	go handlers.HandleMsgSend(sTCP, sUDP)

	for {
		logs.LogTrace(fmt.Sprintf("I'm client %v", clientID))
		utils.PrintTerminal()
		time.Sleep(time.Second * 30)
	}
}
