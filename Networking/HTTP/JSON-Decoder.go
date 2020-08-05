package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON response
// {Message:"All is good","Status":"Success"}

type Status struct {
	Message string
	Status  string
}

func main() {
	res, err := http.Post("http://IP:PORT/ping", "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
