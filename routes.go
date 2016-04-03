/*

routes.go

Implements  Gorilla Mux routes

*/

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	// Create a new Gorilla Mux router
	router := mux.NewRouter().StrictSlash(true)

	// ...with routes defined below
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// Slice of Routes
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Ignite",
		"GET",
		"/ignite/{kindlingId}",
		Ignite,
	},
	Route{
		"Get",
		"GET",
		"/get/{kindlingId}",
		Ignite,
	},
	Route{
		"Add",
		"POST",
		"/add",
		Add,
	},
}
