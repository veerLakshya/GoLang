package intermediate

import (
	"log"

	"go.uber.org/zap"
)

func zaplogging() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("error in initializing zap logger")
	}

	defer logger.Sync()

	logger.Info("This is an info message")

	logger.Info("User logged in", zap.String("Username", "John Doe"), zap.String("method", "GET"))

}
