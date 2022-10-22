package logger

import (
	"flag"
	logexternal "log"
	"os"
	"time"
)

var (
	infolog, _ = newLogger()
)

type logger struct {
	log *logexternal.Logger
}

// NewLogger creates a new logger.
func newLogger() (logger, error) {
	logpath := "/var/log/app/error.log"
	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		panic(err1)
	}
	Log := logexternal.New(file, "", logexternal.LstdFlags|logexternal.Lshortfile)
	return logger{
		log: Log,
	}, nil
}

func currentDate() string {
	formatDate := "2006-01-02 15:04:05"
	currentTime := time.Now().Format(formatDate)
	return currentTime
}

// Print information
func Println(s string) {
	infolog.log.Println(s)
	logexternal.Println(s)
}
