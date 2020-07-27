package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	Bar string
	Baz string
}

func main22() {
	var f = foo{"joe junior", "Hello Shabado"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	fmt.Println(b)
	json.Unmarshal(b, &f)
	fmt.Println(f)
}
