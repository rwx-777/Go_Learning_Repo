// Go Cheat Sheet
package main

import (
	"fmt"
	"time"
)

func main() {
	//primitive Data types
	var x = "Hello World\n"
	y := "Hello World\n"

	fmt.Printf(x)
	fmt.Printf(y)

	var z = int(42)
	w := int(42)

	fmt.Printf("%d\n%d\n", z, w)

	//Slices and Maps ugh
	var mySlice = make([]string, 0)
	mySlice = append(mySlice, "hek")

	var myMap = make(map[string]string)
	myMap["some key"] = "some string"

	//Pointers
	var count = int(42)
	var ptr = &count
	fmt.Println(ptr)  //prints memory address of count
	fmt.Println(*ptr) //prints the value of count
	*ptr = 100        //writes new value into count
	fmt.Println(count)

	//Structs
	var Person101 = new(Person)
	Person101.Name = "Gary"
	Person101.Age = 21
	Person101.SayHello()

	//Interfaces
	var Person2 = new(Person)
	Person2.Name = "Dweeb"
	Greet(Person2)

	//Control structures
	if z == 1 {
		fmt.Println("Z is equal to 1")
	} else {
		fmt.Println("Z is not equal to 1")
	}

	var hek = "foo"

	//Switch statment
	switch hek {
	case "foo":
		fmt.Println("found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	//For Loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Concurrency
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	//end of main
}

//Person is a struct
type Person struct {
	Name string
	Age  int
}

//SayHello to annoying go lint
func (p *Person) SayHello() {
	//like dis
	fmt.Println("hello", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func f() {
	fmt.Println("f function")
}
