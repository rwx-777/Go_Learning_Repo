package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	//print HTTP Status
	fmt.Println(r1.Status)

	//Read and display response body
	body, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	r1.Body.Close()´
}
