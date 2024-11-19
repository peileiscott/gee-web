package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/peileiscott/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	if err := r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}
