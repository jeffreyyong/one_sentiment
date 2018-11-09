package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var host = "localhost:3003"
var vapiURL = "https://api.nexmo.com/v1/calls"
var token string

func getConfig() error {
	if tokenEnv := os.Getenv("TOKEN"); tokenEnv != "" {
		token = tokenEnv
	}

	if hostEnv := os.Getenv("HOST"); hostEnv != "" {
		host = hostEnv
	}

	if vapiURLEnv := os.Getenv("VAPI_ENDPOINT"); vapiURLEnv != "" {
		vapiURL = vapiURLEnv
	}

	var logLevel log.Level
	var err error
	if logLevelEnv := os.Getenv("LOG_LEVEL"); logLevelEnv != "" {
		logLevel, err = log.ParseLevel(logLevelEnv)
		if err != nil {
			logLevel = log.InfoLevel
		}
	} else {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
	return nil
}

func main() {
	err := getConfig()
	if err != nil {
		log.Fatal("Unable to read the configuration", err)
		return
	}

	r := registerRoutes()

	r.Run(":3003")
}
