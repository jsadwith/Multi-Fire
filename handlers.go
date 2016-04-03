/*

handlers.go

Handles route requests

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Multi-Fire: View README.txt for usage")

	w.Header().Set("Content-Type", "application/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Ignite a fire of all URLs for this kindlingId
func Ignite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	kindlingId, err := strconv.Atoi(vars["kindlingId"])
	if err != nil {
		panic(err)
	}

	kindling := GetKindling(kindlingId)

	// Fire off each of the Twigs
	for _, twig := range kindling.Twigs {
		go Fire(twig.Url)
	}

	w.Header().Set("Content-Type", "image/gif; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

// Add new set of URLs and return ID
func Add(w http.ResponseWriter, r *http.Request) {
	var twigs Twigs

	// Open request body
	// Set limit to prevent large JSON POSTs
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// First, we need to create Kindling. Then we can add the Twigs
	kindlingId := AddKindling()

	// Unmarshal URLs into Twigs
	if err := json.Unmarshal(body, &twigs); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Add Twigs
	for _, twig := range twigs {
		// log.Printf("%+v\n", twig.Url)
		twig.KindlingId = kindlingId
		AddTwig(twig)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// Return entity as JSON
	if err := json.NewEncoder(w).Encode(kindlingId); err != nil {
		panic(err)
	}
}

// Get list of URLs for a certain ID
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	kindlingId, err := strconv.Atoi(vars["kindlingId"])
	if err != nil {
		panic(err)
	}

	kindling := GetKindling(kindlingId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(kindling)
}
