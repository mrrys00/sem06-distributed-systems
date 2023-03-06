package configuration

import "net"

const (
	HOST          = "localhost"
	PORT          = "8080"
	TYPE          = "tcp"
	TYPEUDP       = "udp"
	MULTICASTHOST = "230.1.1.1"
	MULTICASTPORT = "42345"
)

var (
	AddressUDP = net.UDPAddr{
		Port: 8080,
		IP:   net.IP{127, 0, 0, 1},
	}
)
