package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	connection    net.Conn
	connectionUDP *net.UDPConn
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
			log.Printf("error %+v", err)
			break
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

func (u *User) SendMsgUDP(msg string, port int) {
	if u.connection == nil {
		return
	}

	s, err := net.ResolveUDPAddr("udp", "localhost:"+string(port))
	log.Printf("resolved address : %v", s.Port)
	sUDP, err := net.DialUDP("udp", nil, s)
	checkError(err)
	defer sUDP.Close()

	_, err = sUDP.Write([]byte(msg))
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

func getUPDPort(addr net.Addr) int {
	resArr := strings.Split(addr.String(), ":")
	res, err := strconv.Atoi(resArr[1])
	if err != nil {
		log.Fatalf("get udp addr error %+v \n", err)
	}
	return res
}

func (us *Users) SendMsgUDP(from *net.UDPAddr, msg string) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	log.Println("sendMsgUDP works here")
	for _, u := range us.users {
		port := getUPDPort(u.connection.RemoteAddr())
		log.Printf("run for port : %v", port)
		if port == from.Port {
			continue
		}

		if u == nil {
			continue
		}

		u.SendMsgUDP(msg, port)
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
		if readed > 0 {
			log.Printf("works here\n")
			users.SendMsgUDP(addr, string(message))
		}
	}
}
func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
