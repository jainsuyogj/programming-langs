package main

import "sync"

//1) Keep a map of shortenedUrl to mappedUrl
//2) Keep a set of mappedUrl for duplicates.

//https://stackoverflow.com/questions/36167200/how-safe-are-golang-maps-for-concurrent-read-write-operations
var urlMapper = map[string]string{}
var lock = sync.RWMutex{}

// The Reqest DTO
type inputRequestJSON struct {
	InputURL string
}

type outputRequestJSON struct {
	ActualURL  string
	ShortenURL string
}
