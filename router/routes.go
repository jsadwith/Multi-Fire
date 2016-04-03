/*

routes.go

Specifies routes

*/

package router

import (
	"net/http"

	"github.com/jsadwith/Multi-Fire/api"
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
		api.Index,
	},
	Route{
		"Ignite",
		"GET",
		"/ignite/{kindlingId}",
		api.Ignite,
	},
	Route{
		"Get",
		"GET",
		"/get/{kindlingId}",
		api.Get,
	},
	Route{
		"Add",
		"POST",
		"/add",
		api.Add,
	},
}
