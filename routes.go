/*

routes.go

Specifies routes

*/

package main

import (
	"net/http"
)

// Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
