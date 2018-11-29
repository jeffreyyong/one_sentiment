package main

import (
	"context"
	"fmt"
	"os"
	"time"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

	log "github.com/sirupsen/logrus"
)

// ISentimentAnalysis is the sentiment analsysis interface
type ISentimentAnalysis interface {
	Analyse(text string) (*Result, error)
}

// SentimentAnalysisService contains the Google Client for doing sentiment analysis
type SentimentAnalysisService struct {
	googleLanguageClient *language.Client
}

// NewSentimentAnalysisService initialises a new instance of sentiment analysis service
func NewSentimentAnalysisService() (*SentimentAnalysisService, error) {
	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/demo/google_credentials.json"); err != nil {
		return nil, err
	}

	sas := &SentimentAnalysisService{}
	backgroundContext := context.Background()
	ctx, cancel := context.WithTimeout(backgroundContext, 5*time.Second)
	defer cancel()
	googleClient, err := language.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	sas.googleLanguageClient = googleClient
	return sas, nil
}

// Analyse analyses the text
func (sas *SentimentAnalysisService) Analyse(text string) (*Result, error) {

	var finalResult *Result

	ctx := context.Background()
	entityResult, err := sas.googleLanguageClient.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})

	if err != nil {
		log.Error("Failed to analyse entity" + fmt.Sprintf("%v", err))
		return nil, err
	}

	for _, entityResult := range entityResult.Entities {
		entity := Entity{entityResult.Name, entityResult.GetSalience()}
		finalResult.Entities = append(finalResult.Entities, entity)
	}

	sentimentResult, err := sas.googleLanguageClient.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})

	if err != nil {
		log.Error("Failed to analyse sentiment" + fmt.Sprintf("%v", err))
		return nil, err
	}

	sentiment := computeSentiment(sentimentResult.DocumentSentiment.GetScore())

	finalResult.Sentiment = sentiment

	return finalResult, nil
}

func computeSentiment(sentimentScore float32) string {

	var sentiment string

	switch {
	case sentimentScore < -0.9:
		sentiment = "very negative"
	case sentimentScore < -0.6:
		sentiment = "medium negative"
	case sentimentScore < -0.3:
		sentiment = "slighly negative"
	case sentimentScore < -0.0:
		sentiment = "neutral"
	case sentimentScore < 0.3:
		sentiment = "slightly positive"
	case sentimentScore < 0.6:
		sentiment = "medium positive"
	case sentimentScore < 0.9:
		sentiment = "very positive"
	}

	return sentiment
}
