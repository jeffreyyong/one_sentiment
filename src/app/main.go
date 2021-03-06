package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var host = "localhost:3003"
var vapiURL = "https://api.nexmo.com/v1/calls"
var token string
var config *Config

var (
	configFilePathPtr = flag.String("c", "client-config.yaml", "Path to configuration file")
)

func parseCommandArgs() error {
	flag.Parse()

	if *configFilePathPtr == "" {
		errMsg := "Error: You have to provide the file path to the config"
		log.Error(errMsg)
		flag.Usage()
		return errors.New(errMsg)
	}
	return nil
}

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

	configFilePath := "config.yaml"
	if err := parseCommandArgs(); err == nil {
		configFilePath = "config.yaml"
		configFilePath = *configFilePathPtr
	}

	// Get from config file
	cfg, err := loadConfig(configFilePath)
	config = cfg
	if err != nil {
		log.Error("Failed to load config: " + fmt.Sprintf("%v", err))
		os.Exit(1)
	}
	log.Debug("Config " + fmt.Sprintf("%+v", cfg))
	return nil
}

func main() {
	err := getConfig()
	if err != nil {
		log.Fatal("Unable to read the configuration", err)
		return
	}

	r := registerRoutes()

	r.Run(config.Addr)
}
