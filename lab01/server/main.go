package main

import (
	"fmt"
	"lab01/configuration"
	"lab01/logs"
	"lab01/server/handlers"
	. "lab01/users"
	"lab01/utils"
	"net"
	"strconv"
	"time"
)

var usersList Users

func main() {
	portInt, err := strconv.Atoi(configuration.PORT)
	logs.LogFatal(err, "Cannot convert port number to int")
	sUDP, err := net.ListenUDP(configuration.TYPEUDP, utils.CreateUDPAddr(configuration.HOST, portInt))
	logs.LogFatal(err, fmt.Sprintf("Cannot listen UDP on provided port: %v", portInt))

	sTCP, err := net.Listen(configuration.TYPE, configuration.HOST+":"+configuration.PORT)
	logs.LogFatal(err, fmt.Sprintf("Cannot listen TCP on provided port: %v", portInt))

	defer sUDP.Close()
	defer sTCP.Close()

	go handlers.HandleNewConnection(sTCP, sUDP, &usersList)

	for {
		logs.LogTrace(fmt.Sprintf("Server running ..."))
		time.Sleep(time.Second * 30)
	}
}
