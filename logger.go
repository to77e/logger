package logger

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Error   *log.Logger
	Info    *log.Logger
)

func InitLog() {
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
}
