package logs

import (
	"lab01v2/configuration"
	"log"
)

const (
	FATAL = "[FATAL] "
	ERR   = "[ERROR] "
	WARN  = "[WARN] "
	INFO  = "[INFO] "
	DEBUG = "[DEBUG] "
	TRACE = "[TRACE] "
)

func LogFatal(err error, desc string) {
	if configuration.LogLvl >= 1 && err != nil {
		log.Fatalf("%s%s -> %+v\n", FATAL, desc, err.Error())
	}
}

func LogError(err error, desc string) {
	if configuration.LogLvl >= 2 && err != nil {
		log.Fatalf("%s%s -> %+v\n", ERR, desc, err.Error())
	}
}

func LogWarning(err error, desc string) {
	if configuration.LogLvl >= 3 && err != nil {
		log.Panicf("%s%s -> %+v\n", WARN, desc, err.Error())
	}
}

func LogInfo(err error, desc string) {
	if configuration.LogLvl >= 4 && err != nil {
		log.Printf("%s%s -> %+v\n", INFO, desc, err.Error())
	}
}

func LogDebug(err error, desc string) {
	if configuration.LogLvl >= 5 && err == nil {
		log.Printf("%s%s\n", DEBUG, desc)
	}
}

func LogTrace(desc string) {
	if configuration.LogLvl >= 6 {
		log.Printf("%s%s\n", TRACE, desc)
	}
}
