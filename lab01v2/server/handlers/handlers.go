package handlers

import (
	"fmt"
	"lab01v2/configuration"
	"lab01v2/logs"
	. "lab01v2/users"
	"lab01v2/utils"
	"net"
	"strconv"
)

func HandleNewConnection(sTCP net.Listener, sUDP *net.UDPConn, usersList *Users) {
	for {
		connection, err := sTCP.Accept()
		logs.LogWarning(err, "Cannot accept clients connection")

		clientPort, err := utils.GetPortFromTCP(connection.RemoteAddr())
		logs.LogWarning(err, "Cannot read port to UDP channel configuration")

		remoteAddr, err := net.ResolveUDPAddr(configuration.TYPEUDP, configuration.HOST+":"+strconv.Itoa(clientPort))
		logs.LogWarning(err, "Cannot resolve UDP address")

		connectionUDP, err := net.DialUDP(configuration.TYPEUDP, nil, remoteAddr)
		logs.LogWarning(err, "Cannot setup UDP channel")

		logs.LogDebug(err, "UDP connection created correctly")

		user := NewUser(clientPort, connection, connectionUDP)
		usersList.AddUser(user)
		configurationMsg := strconv.Itoa(clientPort)
		logs.LogDebug(nil, fmt.Sprintf("New client connected on port: %v", configurationMsg))
		user.SendMsgTCP(configurationMsg)

		go user.HandleTCP(usersList)
		go user.HandleUDP(sUDP, usersList)
	}
}
