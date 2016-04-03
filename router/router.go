/*

router.go

Implements routes from routes.go

*/
package router

import (
	"net/http"

	"github.com/jsadwith/Multi-Fire/util"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	// Create a new Gorilla Mux router
	router := mux.NewRouter().StrictSlash(true)

	// ...with routes defined below
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
