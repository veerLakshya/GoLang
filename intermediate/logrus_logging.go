package main

import "github.com/sirupsen/logrus"

func logruslogging() {

	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// example
	log.Info("This is an info message")
	log.Warn("This is a warning")
	log.Error("This is an error log")

	// logging with fields
	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in ")
}
