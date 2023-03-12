package users

import (
	"net"
	"sync"
)

type User struct {
	connection    net.Conn
	connectionUDP *net.UDPConn
	id            int
}

type Users struct {
	users     []*User
	userMutex sync.Mutex
}

func NewUser(id int, c net.Conn, cUDP *net.UDPConn) *User {
	return &User{
		connection:    c,
		connectionUDP: cUDP,
		id:            id,
	}
}
