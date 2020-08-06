package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "executing /a\n")
	case "/b":
		fmt.Fprint(w, "executing /b\n")
	case "/c":
		fmt.Fprint(w, "Executing /c\n")
	default:
		http.Error(w, "404 Not Found\n", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8080", &r)
}
