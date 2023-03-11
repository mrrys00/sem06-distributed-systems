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
	id            int
}

func NewUser(id int, c net.Conn) *User {
	return &User{
		connection: c,
		id:         id,
	}
}

func (u *User) Handle() {
	defer u.connection.Close()
	defer u.exit()
	for {
		message := make([]byte, 2048)
		readLen, err := u.connection.Read(message)
		if err != nil {
			log.Printf("[ERROR] unabel to received msg from %v\n", u.connection.RemoteAddr())
			log.Printf("error %+v", err)
			u.exit()
			break
		}
		log.Printf("[INFO] msg received from %v\n", u.connection.RemoteAddr())

		users.SendMsg(u, string(message[:readLen]))
	}
}

func (u *User) exit() {
	users.RemoveUser(u)
}

func (u *User) SendMsg(msg string) {
	if u.connection == nil {
		return
	}

	log.Printf("message to be sent: %+v, : len: %+v\n", msg, len(msg))
	_, err := u.connection.Write([]byte(msg))
	if err != nil {
		//log
	}
}

func (u *User) SendMsgUDP(msg string, port int) {
	if u.connection == nil {
		return
	}

	log.Printf("trying to send UDP message : %+v", port)
	s, err := net.ResolveUDPAddr("udp", "localhost:"+strconv.Itoa(port))
	if err == nil {
		log.Printf("resolved address : %v", s.Port)
		sUDP, err := net.DialUDP("udp", nil, s)
		checkError(err)
		_, err = sUDP.Write([]byte(msg))
		defer sUDP.Close()
	}

	if err != nil {
		log.Printf("error %+v\n", err)
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

		tempMsg := "from: " + strconv.Itoa(from.id) + " :: " + msg
		u.SendMsg(tempMsg)
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
		log.Printf("compare ports : from %v : to %v", from.Port, port)
		if port == from.Port {
			log.Print("ignore port : %v\n", port)
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
		user := NewUser(getUPDPort(connection.RemoteAddr()), connection)
		users.AddUser(user)
		configurationMsg := strconv.Itoa(getUPDPort(user.connection.RemoteAddr()))
		log.Printf("user connected on : %v\n", configurationMsg)
		user.SendMsg(configurationMsg)
		go user.Handle()
	}
}

func handleUDP(s *net.UDPConn) {
	for {
		message := make([]byte, 2048)
		readed, addr, err := s.ReadFromUDP(message)
		checkError(err)

		fmt.Printf("UDP received (data count %d) %v\n", readed, addr)
		if readed > 0 {
			log.Printf("recived message: \n%v\n", string(message[:readed]))
			users.SendMsgUDP(addr, string(message[:]))
		}
	}
}
func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
