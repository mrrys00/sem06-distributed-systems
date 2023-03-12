package handlers

import (
	"fmt"
	"lab01/configuration"
	"lab01/logs"
	"lab01/utils"
	"net"
	"strconv"
)

func HandleChannelUDP(sConfig net.Conn, clientID *int) *net.UDPConn {
	portBuff := make([]byte, 16)
	var readed int
	var err error
	for {
		readed, err = sConfig.Read(portBuff)
		logs.LogError(err, "Cannot read TCP data")
		if len(portBuff) != 0 {
			logs.LogDebug(err, fmt.Sprintf("UDP channel created correctly on port: %s", string(portBuff[:readed])))
			break
		}
	}

	listenUDPPort, err := strconv.Atoi(string(portBuff[:readed]))
	logs.LogWarning(err, "Cannot convert port configuration to int")
	*clientID = listenUDPPort

	logs.LogDebug(err, fmt.Sprintf("ClientID assigned as: %v", *clientID))
	sUDPListen, err := net.ListenUDP(configuration.TYPEUDP, utils.CreateUDPAddr(configuration.HOST, listenUDPPort))
	logs.LogWarning(err, "Cannot setup additional UDP channel")

	logs.LogDebug(err, "UDP channel setup correctly")

	return sUDPListen
}

func HandleUDPIncomingMsg(sUDP *net.UDPConn) {
	message := make([]byte, 2048)
	for {
		readed, _, err := sUDP.ReadFromUDP(message)
		logs.LogWarning(err, "Unable to handle incoming UDP message")

		utils.PrintMessage(string(message[:readed]))
	}
}

func sendViaUDP(sUDP *net.UDPConn) {
	_, err := sUDP.Write([]byte(utils.ASCIIART))
	logs.LogError(err, "Cannot send message via UDP")
}
