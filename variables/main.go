package main

import (
	"fmt"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("Enter an integer:", myIntVar)

	myString := "Hello, World!"
	fmt.Println("Enter a string:", myString)
	fmt.Println("myString espacio en memoria:", &myString)
}
