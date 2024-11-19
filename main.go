package main

import (
	"log"
	"net/http"

	"github.com/peileiscott/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})

	r.POST("/ping", func(c *gee.Context) {
		c.JSON(http.StatusCreated, gee.H{
			"message": "pong",
		})
	})

	r.GET("/html", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	if err := r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}
