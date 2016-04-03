/*

logger.go

Simple web logger from http://thenewstack.io/make-a-restful-json-api-go/

*/

package util

import (
	"log"
	"net/http"
	"time"
)

// Logger to log all HTTP requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
