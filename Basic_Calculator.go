package main

import (
	"fmt"
)

//Go Calculator

var number1, number2 int

var operator string = ""
var result int = 0

func addition() {
	result = number1 + number2
}

func subtraction() {
	result = number1 - number2
}

func mulitply() {
	result = number1 * number2
}

func division() {
	result = number1 / number2
}

func main2() {

	fmt.Printf("Enter first number: ")
	fmt.Scanf("%d", &number1)

	fmt.Printf("Enter operator: ")
	fmt.Scanf("%s", &operator)

	fmt.Printf("Enter second number: ")
	fmt.Scanf("%d", &number2)

	if operator == "+" {
		addition()
	} else if operator == "-" {
		subtraction()
	} else if operator == "*" {
		mulitply()
	} else if operator == "/" {
		division()
	} else {
		fmt.Println("Invalid Operator!")
	}

	fmt.Println("Your result is: ")
	fmt.Printf("%d\n", result)

}
