package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

}
