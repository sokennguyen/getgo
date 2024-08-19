package main

import (
	"fmt"
	"net/http"

	"templgin/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Count struct {
	Count int
}

func render(c *gin.Context, template templ.Component, status int) error {
	c.Header("Content-Type", "text/html")
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func getHome(c *gin.Context) {
	render(c, views.Hello(0), http.StatusOK)
}

func main() {
	count := Count{Count: 0}

	r := gin.Default()
	r.GET("/", getHome)
	r.POST("/", func(c *gin.Context) {
		count.Count++
		fmt.Println("Count:", count.Count)
		render(c, views.Counter(count.Count), http.StatusOK)
	})

	r.Run(":8080")
}
