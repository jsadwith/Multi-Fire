package main

// Kindling model - simple storage for arrays of URLs to fire
type Kindling struct {
    Id        int
    Urls      string
}

// Kindlings - Slice of Kindling
type Kindlings []Kindling