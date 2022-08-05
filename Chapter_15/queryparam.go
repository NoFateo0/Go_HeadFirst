package main

import (
	"log"
	"net/http"
)

// getParameter returns the first value associated with a
// named query parameter from an http.Request.
func getParameter(request *http.Request, parameterName string) string {
	query := request.URL.Query()
	return query[parameterName][0]
}

// YOUR CODE HERE: Set up a handler function.
func queryHandler(writer http.ResponseWriter, request *http.Request) {
	text := getParameter(request, "text")
	body := []byte("<h1>" + text + "</h1>")
	_, err := writer.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// YOUR CODE HERE: Set up your function to handle
	// requests for the "/big" path.
	// Then start an HTTP server on port 8080.

	http.HandleFunc("/big", queryHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
