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

	{
		const myConstVar = 3.14 // Constantes no pueden ser cambiadas
		fmt.Println("mi constante:", myConstVar)

		// myConstVar = 2.71 // Esto generará un error de compilación

	}

}
