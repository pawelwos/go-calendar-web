package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pawelwos/go-calendar"
)

func addDate(y int, m int) string {
	if m < 12 {
		return fmt.Sprintf("%v/%v", y, (m + 1))
	} else {
		return fmt.Sprintf("%v/%v", (y + 1), 1)
	}
}

func subDate(y int, m int) string {
	if m > 1 {
		return fmt.Sprintf("%v/%v", y, (m - 1))
	} else {
		return fmt.Sprintf("%v/%v", (y - 1), 12)
	}
}

func main() {

	r := gin.Default()
	r.Static("/css", "./static/css")
	r.SetFuncMap(template.FuncMap{
		"addDate": addDate,
		"subDate": subDate,
	})
	r.LoadHTMLGlob("src/html/*")

	r.GET("/", func(c *gin.Context) {
		year := time.Now().Year()
		month := int(time.Now().Month())
		var cal = calendar.Create(year, month)
		head := calendar.GetHead()
		body := cal.GetBody()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"thead": head,
			"tbody": body,
			"year":  year,
			"month": month,
		})
	})

	r.GET("/:year/:month", func(c *gin.Context) {
		year, _ := strconv.Atoi(c.Param("year"))
		month, _ := strconv.Atoi(c.Param("month"))
		var cal = calendar.Create(year, month)
		head := calendar.GetHead()
		body := cal.GetBody()
		fmt.Printf("%v, %v", year, month)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"thead": head,
			"tbody": body,
			"year":  year,
			"month": month,
		})
	})
	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:8080
}
