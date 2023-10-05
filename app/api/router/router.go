package router

import (
	"fmt"
	"net/http"
)

func ApiHandler() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Welcome to Homepage")
	})

	http.HandleFunc("/movies", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Welcome to Movies Page")
	})
}
