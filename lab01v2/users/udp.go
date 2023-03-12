package users

import (
	"fmt"
	"lab01v2/logs"
	"net"
	"strconv"
)

func (u *User) HandleUDP(listenUDP *net.UDPConn, usersList *Users) {
	defer u.connectionUDP.Close()
	defer u.exit(usersList)

	for {
		message := make([]byte, 2048)
		readed, _, err := listenUDP.ReadFromUDP(message)
		if err != nil {
			logs.LogWarning(err, "Cannot handle user UDP communication")
			u.exit(usersList)
			break
		}

		usersList.SendMsgUDP(u, string(message[:readed]))
	}
}

func (u *User) SendMsgUDP(msg string) {
	if u.connectionUDP == nil {
		return
	}

	logs.LogDebug(nil, fmt.Sprintf("UDP Message to be sent: %+v, : len: %+v\n", msg, len(msg)))

	_, err := u.connectionUDP.Write([]byte(msg))
	logs.LogWarning(err, "UDP message did not send correctly")
}

func (us *Users) SendMsgUDP(from *User, msg string) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	for _, u := range us.users {
		if u.id == from.id {
			continue
		}

		if u == nil {
			continue
		}

		tempMsg := "from: " + strconv.Itoa(from.id) + " ::\n" + msg
		u.SendMsgUDP(tempMsg)
	}
}
