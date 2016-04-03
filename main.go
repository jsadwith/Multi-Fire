/*

Multi-Fire

Allows for multiple concurrent URLs to be fired off from a
single client connection in order to decrease the processing
responsibilities of the client.

See README.md for usage.

*/

package main

import (
	"log"
	"net/http"

	"github.com/jsadwith/Multi-Fire/model"
	"github.com/jsadwith/Multi-Fire/router"
)

func main() {

	// Create SQLite tables if they don't exist
	model.CreateTableKindling()
	model.CreateTableTwig()

	// Register routes, ensuring trailing slashes redirect to route - /route/ -> /route
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
