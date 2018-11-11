package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// NewNCCO returns the NCCO json
func NewNCCO(eventURL, language string) ([]byte, error) {
	speech := Speech{
		Context:  []string{"one", "two", "three", "four"},
		Language: language,
	}
	asrNCCO := ASRNCCO{
		EventURL: []string{eventURL},
		Action:   "input",
		Speech:   speech,
	}
	respBody, err := json.Marshal(asrNCCO)
	if err != nil {
		return nil, err
	}
	log.Debug("Response body for NCCO: " + string(respBody))

	return respBody, nil
}

// ASRNCCO is the NCCO object for ASR
type ASRNCCO struct {
	EventURL []string `json:"eventUrl"`
	Action   string   `json:"action"`
	Speech   Speech   `json:"speech"`
}

// Speech is the Speech object
type Speech struct {
	Context  []string `json:"context"`
	UUID     []string `json:"uuid"`
	Language string   `json:"language"`
}
