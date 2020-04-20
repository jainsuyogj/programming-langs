package main

import (
	"math/rand"
	"net/http"
	"time"
)

//https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986

func main() {
	// We will start with very basic implementation
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/shortMe", shortenHandle)
	http.HandleFunc("/x/", handleGet)
	http.ListenAndServe(":8080", nil)
}
