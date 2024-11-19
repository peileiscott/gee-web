package main

import (
	"log"
	"net/http"

	"github.com/peileiscott/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Welcome to home page</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// GET /hello?name=[name]
		c.String(http.StatusOK, "Hello %s\n", c.Query("name"))
	})

	r.POST("/login", func(c *gee.Context) {
		// POST /login username=[username]&password=[password]
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	if err := r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}
