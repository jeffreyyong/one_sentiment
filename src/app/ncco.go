package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// NewNCCO returns the NCCO json
func NewNCCO(eventURL, language, uuid string) ([]byte, error) {
	speech := Speech{
		Context:  []string{"one", "two", "three", "four"},
		Language: language,
		UUID:     []string{uuid},
	}

	talkNCCOOne := TalkNCCO{
		Action: "talk",
		Text:   "Tell us your credit card pin.",
	}

	asrNCCO := ASRNCCO{
		EventURL: []string{eventURL},
		Action:   "input",
		Speech:   speech,
	}

	talkNCCOTwo := TalkNCCO{
		Action: "talk",
		Text:   "Thank you very nuch",
	}

	nccoList := []interface{}{talkNCCOOne, asrNCCO, talkNCCOTwo}

	respBody, err := json.Marshal(nccoList)
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

// TalkNCCO controls the talk of the call
type TalkNCCO struct {
	Action string `json:"action"`
	Text   string `json:"text"`
}

// Speech is the Speech object
type Speech struct {
	Context  []string `json:"context"`
	UUID     []string `json:"uuid"`
	Language string   `json:"language"`
}
