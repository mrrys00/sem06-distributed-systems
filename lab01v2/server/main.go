package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type User struct {
	connection net.Conn
}

func NewUser(c net.Conn) *User {
	return &User{
		connection: c,
	}
}

func (u *User) Handle() {
	defer u.connection.Close()
	defer u.exit()
	for {
		message := make([]byte, 2048)
		_, err := u.connection.Read(message)
		if err != nil {
			log.Printf("[ERROR] unabel to received msg from %v\n", u.connection.RemoteAddr())
			continue
		}
		log.Printf("[INFO] msg received from %v\n", u.connection.RemoteAddr())

		users.SendMsg(u, string(message))
	}
}

func (u *User) exit() {
	users.RemoveUser(u)
}

func (u *User) SendMsg(msg string) {
	if u.connection == nil {
		return
	}

	_, err := u.connection.Write([]byte(msg))
	if err != nil {
		//log
	}
}

type Users struct {
	users     []*User
	userMutex sync.Mutex
}

func (us *Users) AddUser(u *User) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	us.users = append(us.users, u)
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

func (us *Users) SendMsg(from *User, msg string) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()
	for _, u := range us.users {
		if u == from {
			continue
		}

		if u == nil {
			continue
		}

		u.SendMsg(msg)
	}
}

var users Users

func main() {
	sUDP, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 8080,
	})
	checkError(err)

	s2, err := net.Listen("tcp", "localhost:8080")
	checkError(err)
	defer sUDP.Close()
	defer s2.Close()

	go handleUDP(sUDP)
	go handleConnection(s2)

	for {
		//fmt.Print("alive!")
		time.Sleep(time.Second)
	}
}

func handleConnection(s net.Listener) {
	for {
		connection, err := s.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("Connected %v\n", connection)
		user := NewUser(connection)
		users.AddUser(user)
		go user.Handle()
	}
}

func handleUDP(s *net.UDPConn) {
	for {
		message := make([]byte, 20)
		readed, addr, err := s.ReadFromUDP(message)
		checkError(err)

		fmt.Printf("UDP received (data count %d) %v", readed, addr)
	}
}
func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
