package handlers

import (
	"bufio"
	"lab01/logs"
	"lab01/utils"
	"net"
	"os"
)

func HandleMsgSend(sTCP net.Conn, sUDP *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		utils.PrintTerminal()
		text, err := reader.ReadString('\n')
		logs.LogError(err, "Cannot read user input")

		data := []byte(text)
		if text != "U\n" {
			sendViaTCP(sTCP, data)
		} else {
			sendViaUDP(sUDP)
		}
	}
}
