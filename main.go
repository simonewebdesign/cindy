package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Get example.com
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
		log.Printf("Error reading body: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("BODY: %q", body)
}
