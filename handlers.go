/*

handlers.go

Handles route requests

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Multi-Fire: View README.txt for usage")
}

// Ignite a fire of all URLs for this kindlingId
func Ignite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kindlingId := vars["kindlingId"]

	fmt.Fprintln(w, "Kindling ID:", kindlingId)
	fmt.Fprintf(w, "URI, %q", r.URL.Path)
}

// Add new set of URLs and return ID
func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URI, %q", r.URL.Path)
}

// Get list of URLs for a certain ID
func Get(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// kindlingId := vars["kindlingId"]

	kindlings := Kindlings{
		Kindling{Urls: "www.foo.com, www.bar.com"},
		Kindling{Urls: "www.google.com, www.blah.com"},
	}

	json.NewEncoder(w).Encode(kindlings)
}
