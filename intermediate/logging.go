package main

import (
	"log"
	"os"
)

func loggingg() {
	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message. ")

	// Log Flags
	log.SetFlags(log.Ldate)
	log.Println("This is a log message with only date. ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and fileName. ")

	infoLogger.Println("This is an info message")
	warnLogger.Println("this is a warn logger")
	errorLogger.Println("this is a error log")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Could not create or open file.")
		return
	}

	defer file.Close()

	// writing in a file or anywhere we want
	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("this is a debug logger")
	infoLogger1.Println("this is a info logger")
	warnLogger1.Println("this is a warring")
	errorLogger1.Println("this is an error log")

	// common logging packages - logrus, zap

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
