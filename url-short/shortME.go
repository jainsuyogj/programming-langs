package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomVal(num int) string {

	b := make([]rune, num)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenHandle(w http.ResponseWriter, req *http.Request) {

	fmt.Println("In POST handle, shortME")

	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("only post supported"))
		return
	}

	var inp inputRequestJSON
	err := json.NewDecoder(req.Body).Decode(&inp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "Failed to ready body")
		return
	}

	fmt.Println("Printing th decoded inp ", inp.InputURL)

	// TODO: Add a write lock.. as no delete, no need to safeguard read is it?
	lock.Lock()

	defer lock.Unlock()

	// TODO: Looping over the map is expensive, change it

	for _, val := range urlMapper {
		if val == inp.InputURL {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("exists already"))
			//fmt.Println(" We can change this to returrn the same url")
			return
		}
	}
	// Shorten logic
	/*
		1) Get a random string of 11 chars.. improve the logic here, this can lead to more collision if the number of mapped urls is
		full
	*/
	var shortenedURL string
	for randKey := randomVal(11); randKey != ""; {
		if _, ok := urlMapper[randKey]; !ok {
			urlMapper[randKey] = inp.InputURL
			shortenedURL = randKey
			break
		}
		randKey = randomVal(11)
	}

	fmt.Println("\nRandom key is ", shortenedURL)

	outResp := outputRequestJSON{ActualURL: inp.InputURL, ShortenURL: shortenedURL}
	/*
			test, err := json.Marshal(outResp)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Failed to marshal"))
				return
			}

		fmt.Println("Now writing actually ", string(test))
	*/
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outResp)
}
