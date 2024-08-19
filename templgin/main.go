package main

import (
	"net/http"

	"templgin/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, template templ.Component, status int) error {
	c.Header("Content-Type", "text/html")
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func getHome(c *gin.Context) {
	render(c, views.Hello("Kien"), http.StatusOK)
}

func main() {
	r := gin.Default()
	r.GET("/", getHome)

	r.Run(":8080")
}
