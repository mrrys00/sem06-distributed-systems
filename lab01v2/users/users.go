package users

import (
	"fmt"
	"lab01v2/logs"
	"net"
	"strconv"
)

func (u *User) HandleTCP(usersList *Users) {
	defer u.connection.Close()
	defer u.exit(usersList)

	for {
		message := make([]byte, 2048)
		readLen, err := u.connection.Read(message)
		if err != nil {
			logs.LogWarning(err, "Cannot handle user TCP communication")
			u.exit(usersList)
			break
		}

		usersList.SendMsgTCP(u, string(message[:readLen]))
	}
}

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

func (u *User) exit(usersList *Users) {
	logs.LogTrace("Go to remove user")
	usersList.RemoveUser(u)
}

func (u *User) SendMsgTCP(msg string) {
	if u.connection == nil {
		return
	}

	logs.LogDebug(nil, fmt.Sprintf("TCP Message to be sent: %+v, : len: %+v\n", msg, len(msg)))

	_, err := u.connection.Write([]byte(msg))
	logs.LogWarning(err, "TCP message did not send correctly")
}

func (u *User) SendMsgUDP(msg string) {
	if u.connectionUDP == nil {
		return
	}

	logs.LogDebug(nil, fmt.Sprintf("UDP Message to be sent: %+v, : len: %+v\n", msg, len(msg)))

	_, err := u.connectionUDP.Write([]byte(msg))
	logs.LogWarning(err, "UDP message did not send correctly")
}

func (us *Users) AddUser(u *User) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	us.users = append(us.users, u)
	logs.LogDebug(nil, "New user added")
}

func (us *Users) RemoveUser(u *User) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	for i, user := range us.users {
		if user != u {
			continue
		}

		us.users = append(us.users[0:i], us.users[i:]...)
	}
}

func (us *Users) SendMsgTCP(from *User, msg string) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	for _, u := range us.users {
		if u == from {
			continue
		}

		if u == nil {
			continue
		}

		tempMsg := "from: " + strconv.Itoa(from.id) + " :: " + msg
		u.SendMsgTCP(tempMsg)
	}
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
