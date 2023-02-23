package main

import (
	"fmt"
	"go_Gee/gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "%s", req.URL.Path)
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	log.Fatal(r.Run(":8080"))
}
