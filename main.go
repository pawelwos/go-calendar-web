package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pawelwos/go-calendar"
)

func main() {

	r := gin.Default()
	r.Static("/css", "./static/css")
	r.LoadHTMLGlob("src/html/*")

	r.GET("/", func(c *gin.Context) {

		var cal = calendar.Create(time.Now().Year(), int(time.Now().Month()))
		head := calendar.GetHead()
		body := cal.GetBody()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"thead": head,
			"tbody": body,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:8080
}
