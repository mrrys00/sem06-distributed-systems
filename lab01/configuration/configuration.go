package configuration

const (
	HOST          = "localhost"
	PORT          = "8080"
	TYPE          = "tcp"
	TYPEUDP       = "udp"
	MULTICASTHOST = "230.1.1.1"
	MULTICASTPORT = "42345"

	LOGLEVELNONE  = 0
	LOGLEVELFATAL = 1
	LOGLEVELERROR = 2
	LOGLEVELWARN  = 3
	LOGLEVELINFO  = 4
	LOGLEVELDEBUG = 5
	LOGLEVELTRACE = 6
)

var LogLvl = LOGLEVELTRACE
