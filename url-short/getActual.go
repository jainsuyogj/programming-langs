package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println(" \nIn handle GET")

	fmt.Println(" Path is ", req.URL.Path)

	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET supported"))
		return
	}

	id := strings.TrimPrefix(req.URL.Path, "/x/")

	if shortURL, exists := urlMapper[id]; exists {
		test := outputRequestJSON{shortURL, id}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		respJSON, _ := json.Marshal(test)
		w.Write(respJSON)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("with id not found" + id))
	}
}
