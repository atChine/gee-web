package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	//r := gee.New()
	//r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	//	for k, v := range req.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//
	//r.Run(":9999")

	r := gee.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.POST("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
