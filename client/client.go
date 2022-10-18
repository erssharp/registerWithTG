package main

import (
	"log"
	"net/http"
)

func main() {
	client := http.Client{}

	resp, err := http.Get("http://127.0.0.1:8000/get")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(resp.Body)
	req := new(http.Request)
	req.RemoteAddr = "http://localhost:8080"
	req.RequestURI = "/employee/get"

	client.Do(req)
}
