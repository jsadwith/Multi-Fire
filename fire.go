/*

fire.go

Handles URL requests

*/

package main

import (
	"net/http"
	"time"
)

// Make a request to a URL
// TODO: Verify that this correctly forwards headers
func Fire(url string) {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("HTTP_HOST", "value")
	response, err := client.Do(req)

	// If error, most likely the timeout has been reached
	if err != nil {
		// Do nothing
	} else {
		// We don't care about the response so close immediately
		response.Body.Close()
	}
}
