package handlers

import (
	"lab01v2/logs"
	"lab01v2/utils"
	"net"
)

func HandleTCPIncomingMsg(sTCP net.Conn) {
	message := make([]byte, 2048)
	for {
		readed, err := sTCP.Read(message)
		logs.LogWarning(err, "Unable to handle incoming TCP message")

		utils.PrintMessage(string(message[:readed]))
	}
}

func sendViaTCP(sTCP net.Conn, data []byte) {
	_, err := sTCP.Write(data)
	logs.LogError(err, "Cannot send message via TCP")
}
