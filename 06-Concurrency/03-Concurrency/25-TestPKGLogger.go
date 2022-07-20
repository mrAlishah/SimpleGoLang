package main

import (
	"database/sql"
	"os"

	//go get -u github.com/mgutz/logxi/v1
	log "github.com/mgutz/logxi/v1"
)

// create package variable for Logger interface
var logger log.Logger

func main() {
	// use default logger
	who := "mario"

	log.Info("Hello", "who", who)

	// create a logger with a unique identifier which
	// can be enabled from environment variables
	logger = log.New("pkg")
	logger.Error("sdfsdfsd")
	// specify a writer, use NewConcurrentWriter if it is not concurrent
	// safe
	modelLogger := log.NewLogger(log.NewConcurrentWriter(os.Stdout), "models")

	_, err := sql.Open("postgres", "dbname=testdb")
	if err != nil {
		modelLogger.Error("Could not open database", "err", err)
	}

	fruit := "apple"
	languages := []string{"go", "javascript"}
	if log.IsDebug() {
		// use key-value pairs after message
		logger.Debug("OK", "fruit", fruit, "languages", languages)
	}
}
