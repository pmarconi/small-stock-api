package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LogsDirpath = "logs"
)

func InitMainLogger() error {
	err := os.Mkdir(LogsDirpath, 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetLogFile() *os.File {
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v%v%v.log", year, int(month), day)
	filePath, err := os.OpenFile(LogsDirpath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return filePath
}

func Info() *log.Logger {
	getFilePath := GetLogFile()
	return log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Warning() *log.Logger {
	getFilePath := GetLogFile()
	return log.New(getFilePath, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Error() *log.Logger {
	getFilePath := GetLogFile()
	return log.New(getFilePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Fatal() *log.Logger {
	getFilePath := GetLogFile()
	return log.New(getFilePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// utils.Info().Println("This is Info an log message")
// utils.Warning().Println("This is a Fatal Error log message")
// utils.Error().Println("This is a Warning log message")
// utils.Fatal().Println("This is an log message")
