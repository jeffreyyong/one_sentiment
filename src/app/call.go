package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// CallRequest is an object that will be marshalled to JSON when
// making a call to VAPI
type CallRequest struct {
	Destinations []Destination `json:"to"`
	AnswerURL    []string      `json:"answer_url"`
	EventURL     []string      `json:"event_url"`
	Sources      []Source      `json:"from"`
}

// CallResponse is an object that holds the call response
type CallResponse struct {
	UUID           string `json:"uuid"`
	Status         string `json:"started"`
	Direction      string `json:"direction"`
	ConversationID string `json:"conversation_id"`
}

// Destination is the destination object
type Destination struct {
	CallType string `json:"type"`
	Number   string `json:"number"`
}

// Source is the source object
type Source struct {
	CallType string `json:"type"`
	Number   string `json:"number"`
}

// Response is the HTTP response that contains the status code and response body
type Response struct {
	StatusCode int
	Body       []byte
}

// Caller is a caller instance
type Caller struct {
	vapiHost          string
	destinationNumber string
	sourceNumber      string
	eventURL          string
	answerURL         string
	bearerToken       string
}

// NewCaller initialises a new caller instance
func NewCaller(number Number) *Caller {
	caller := &Caller{
		vapiHost:          config.VAPIHost,
		destinationNumber: number.Destination,
		sourceNumber:      config.SourceNumber,
		eventURL:          config.CallEventURL,
		answerURL:         config.AnswerURL,
		bearerToken:       config.BearerToken,
	}
	return caller
}

// Call makes a call
func (c *Caller) Call() (string, error) {
	url := fmt.Sprintf("http://%s/v1/calls", c.vapiHost)

	reqBody, err := c.buildCallReqBody()
	if err != nil {
		log.Error("Failed to build call object as request body: " + fmt.Sprintf("%v", err))
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error("Failed to construct a POST request calling VAPI: " + fmt.Sprintf("%v", err))
		return "", err
	}

	resp, err := performRequest(req, c.bearerToken)
	if err != nil {
		log.Error("Failed to perform a POST request calling VAPI: " + fmt.Sprintf("%v", err))
		return "", err
	}

	var uuid string
	var callResp CallResponse
	if resp.StatusCode == 201 {
		if err = json.Unmarshal(resp.Body, &callResp); err != nil {
			log.Error("Failed to unmashal to Call Response " + fmt.Sprintf("%v", err))
			return "", err
		}
		uuid = callResp.UUID
	} else {
		log.Errorf("Failed to make a call, status code: %v, resp body: %v", resp.StatusCode, string(resp.Body))
		return "", fmt.Errorf("Failed to make a call")
	}
	return uuid, nil
}

func (c *Caller) buildCallReqBody() ([]byte, error) {
	destinations := []Destination{{"phone", c.destinationNumber}}
	sources := []Source{{"phone", c.sourceNumber}}

	call := CallRequest{
		Destinations: destinations,
		Sources:      sources,
		EventURL:     []string{c.eventURL},
		AnswerURL:    []string{c.answerURL},
	}
	reqBody, err := json.Marshal(call)
	log.Debug("Request body for the call: " + string(reqBody))
	if err != nil {
		return nil, err
	}

	return reqBody, nil
}

func performRequest(req *http.Request, bearerToken string) (*Response, error) {
	bearerToken = "Bearer " + bearerToken
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("HTTP client failed to make request: " + fmt.Sprintf("%v", err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: " + fmt.Sprintf("%v", err))
		return nil, err
	}
	return &Response{StatusCode: resp.StatusCode, Body: body}, nil
}
