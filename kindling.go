/*

kindling.go

Kindling Model - Kindling consists of a set of URLs

*/

package main

// Kindling model - simple storage for arrays of URLs to fire
type Kindling struct {
	Id   int    `json:"name"`
	Urls string `json:"urls"`
}

// Kindlings - Slice of Kindling
type Kindlings []Kindling
