package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var results map[string]string

func registerRoutes() *gin.Engine {

	r := gin.Default()

	results = make(map[string]string)
	// Serve HTML/JS page
	r.LoadHTMLGlob("templates/**/*.html")
	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/result/:id", func(c *gin.Context) {
		id := c.Param("id")
		if result, ok := results[id]; ok {
			c.JSON(http.StatusOK, struct {
				Result string `json:"result"`
			}{Result: result})
		}
	})

	r.POST("/asr", func(c *gin.Context) {
		var number Number
		c.Bind(&number)
		uuid := "test"
		results["test"] = "salut test"

		c.JSON(http.StatusOK, struct {
			UUID string `json:"uuid"`
		}{UUID: uuid})
	})

	// VAPI endpoints
	r.POST("/callback", func(c *gin.Context) {
		id := c.Param("id")
		results[id] = "test result"

		c.JSON(http.StatusOK, struct{}{})
	})

	r.GET("/ncco", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct{}{})
	})

	r.Static("/public", "./public")
	return r
}
