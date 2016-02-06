// 1. Talk to user (main.go)

// Live Scraper is a service that receives a request for an Amazon ID,
// Scrapes Amazon's website for information based on that ID, and
// responds back with a JSON Representation of that Resource, or an error
// signifying that it failed to find any data on that Resource.
package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	// HomeHandler handles when a GET request is made simply to localhost:8080,
	// in which case it can only respond with a nil JSON Representation
	r.GET("/", HomeHandler)

	// AmznIDHandler handles when a GET request is made to localhost:8080 with
	// an Amazon ID appended to the end, and responds with a JSON Representation
	// of the request.
	r.GET("/movie/amazon/:id", AmznIDHandler)

	fmt.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "No ID request was made.")
}

func AmznIDHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqID := p.ByName("id")
	url := "http://www.amazon.com/gp/product/" + reqID
	// Make request to Amazon.com to traverse/parse
	amznID, err := ExtractInfo(url)
	if err != nil {
		fmt.Errorf("Expected %T, got %s", amznID, err)
	}
	fmt.Fprintln(w, amznID)
}
