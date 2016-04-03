/*

Multi-Fire

Allows for multiple concurrent URLs to be fired off from a
single client connection in order to decrease the processing
responsibilities of the client.

See README.md for usage.

*/

package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"

    // "github.com/jsadwith/Multi-Fire/models"
    "github.com/gorilla/mux"
)

func main() {
    // Register routes, ensuring trailing slashes redirect to route - /route/ -> /route
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/ignite/{kindlingId}", Ignite)
    router.HandleFunc("/get/{kindlingId}", Get)
    router.HandleFunc("/add", Add)

    log.Fatal(http.ListenAndServe(":8080", router))
}

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Multi-Fire: View README.txt for usage");
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
