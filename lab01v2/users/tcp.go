package users

import (
	"fmt"
	"lab01v2/logs"
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

func (u *User) SendMsgTCP(msg string) {
	if u.connection == nil {
		return
	}

	logs.LogDebug(nil, fmt.Sprintf("TCP Message to be sent: %+v, : len: %+v\n", msg, len(msg)))

	_, err := u.connection.Write([]byte(msg))
	logs.LogWarning(err, "TCP message did not send correctly")
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
