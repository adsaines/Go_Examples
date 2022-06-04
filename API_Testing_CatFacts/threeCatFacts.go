package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// to properly parse the expected fact
// https://mholt.github.io/json-to-go/ => transforms a json into a usable go object template
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// set the functions to run based on the url being hit, in this case just one function will need run
func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// this determines which functions get run based on which page gets hit by the request
	myRouter.HandleFunc("/", getFacts)
	// finally, instead of passing in nil, we want to pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// get three random cat facts and return them
func getFacts(w http.ResponseWriter, r *http.Request) {
	// get three facts
	// Expansion: take a number of facts to retrieve
	var threeFacts [3]string
	threeFacts[0] = retrieveAFact()
	threeFacts[1] = retrieveAFact()
	threeFacts[2] = retrieveAFact()

	// return the facts
	json.NewEncoder(w).Encode(threeFacts)
}

// retrieve a single cat fact
func retrieveAFact() string {
	// get stuff
	response, err := http.Get("https://catfact.ninja/fact")

	// don't display if there are errors
	if err != nil {
		panic("A cat fact was not returned.")
	}

	// defer the closing of the response body till the end of the method
	defer response.Body.Close()

	// decode and just save the facts portion of the info
	var aFact CatFact
	json.NewDecoder(response.Body).Decode(&aFact)

	return aFact.Fact
}
