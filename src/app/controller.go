package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var results map[string]*Result
var languages map[string]string

func registerRoutes() *gin.Engine {

	r := gin.Default()

	languages = make(map[string]string)
	results = make(map[string]*Result)
	// Serve HTML/JS page
	r.LoadHTMLGlob(config.TemplatePath + "/**/*.html")
	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/result/:id", func(c *gin.Context) {
		id := c.Param("id")
		if result, ok := results[id]; ok {
			c.JSON(http.StatusOK, struct {
				Result *Result `json:"result"`
			}{Result: result})
		}
	})

	r.POST("/asr", func(c *gin.Context) {
		var number Number
		err := c.BindJSON(&number)
		if err != nil {
			log.Error("Can't read number", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		caller := NewCaller(number)
		uuid, err := caller.Call()
		log.Debug("UUID of call: " + fmt.Sprintf("%v", uuid))

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		languages[uuid] = number.Language
		caller.uuid = uuid

		c.JSON(http.StatusOK, struct {
			UUID string `json:"uuid"`
		}{UUID: uuid})
	})

	// VAPI endpoints
	r.POST("/callback", func(c *gin.Context) {
		var cb callback
		c.BindJSON(&cb)
		cbResults := cb.Speech.Results

		if len(cbResults) < 1 {
			log.Error("No result in callback")
		} else {
			sas, err := NewSentimentAnalysisService()
			if err != nil {
				log.Error("Failed to initialise sentiment analysis service: " + fmt.Sprintf("%v", err))
				return
			}
			result, err := sas.Analyse(cbResults[0].Text)
			if err != nil {
				log.Error("Failed to analyse result: " + fmt.Sprintf("%v", err))
			}
			result.Word = cbResults[0].Text
			result.Language = languages[cb.UUID]

			results[cb.UUID] = result
		}

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/ncco", func(c *gin.Context) {
		uuid := c.Query("uuid")
		ncco, err := NewNCCO(config.NCCOEventURL, languages[uuid], uuid)
		if err != nil {
			log.Error("Failed to initialise NCCO " + fmt.Sprintf("%v", err))
			c.JSON(http.StatusInternalServerError, nil)
		}

		log.Debug("NCCO " + string(ncco))
		c.Data(http.StatusOK, "application/json; charset=utf-8", ncco)
	})

	r.Static("/public", config.PublicFiles)
	return r
}
