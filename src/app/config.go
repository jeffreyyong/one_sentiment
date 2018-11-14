package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config for ASR demo, obtained from the configuration yaml file.
type Config struct {
	VAPIHost     string `yaml:"vapi_host"`
	BearerToken  string `yaml:"bearer_token"`
	SourceNumber string `yaml:"source_number"`
	CallEventURL string `yaml:"call_event_url"`
	AnswerURL    string `yaml:"answer_url"`
	NCCOEventURL string `yaml:"ncco_event_url"`
}

// loadConfig loads the config given the file path
func loadConfig(path string) (*Config, error) {
	cfg := &Config{}

	source, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(source, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.VAPIHost == "" {
		return nil, fmt.Errorf("Missing VAPI host to make a call to")
	}

	if cfg.BearerToken == "" {
		return nil, fmt.Errorf("Missing Bearer token")
	}

	if cfg.SourceNumber == "" {
		return nil, fmt.Errorf("Missing source number")
	}

	if cfg.CallEventURL == "" {
		return nil, fmt.Errorf("Missing call event url")
	}

	if cfg.NCCOEventURL == "" {
		return nil, fmt.Errorf("Missing NCCO event url")
	}

	if cfg.AnswerURL == "" {
		return nil, fmt.Errorf("Missing answer url")
	}

	return cfg, nil
}
