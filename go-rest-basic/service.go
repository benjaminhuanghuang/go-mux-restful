package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	router := mux.NewRouter()
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "dddd")
	})

	// Start server
	const port string = ":8964"

	http.ListenAndServe(port, router)
}
