package logger

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func InitializeLogger() {
	logger = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime)
}

func Info(message string) {
	logger.Println(message)
}
