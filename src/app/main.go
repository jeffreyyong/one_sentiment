package main

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

var host string

func getConfig() error {
	host = os.Getenv("HOST")
	if host == "" {
		return errors.New("Can't get hostname")
	}

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
