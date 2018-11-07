package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")
	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/result/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "100" {
			c.HTML(http.StatusOK, "result.html", nil)
			return
		}
	})

	r.POST("/asr/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "call" {
			var number Number
			c.Bind(&number)
			// TODO: Call VAPI

			c.Redirect(http.StatusMovedPermanently, "/result/100")
		}
	})

	r.Static("/public", "./public")
	return r
}
