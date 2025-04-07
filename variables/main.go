package main

import (
	"fmt"
	"unsafe"
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

	var my8BitsUnitVar uint8 = 255 // uint8 puede almacenar valores de 0 a 255
	fmt.Println("my8BitsUnitVar:", my8BitsUnitVar)
	fmt.Println("type de my8BitsUnitVar:", fmt.Sprintf("%T", my8BitsUnitVar)) // Imprime el tipo de la variable
	fmt.Println("Espacio en memoria de my8BitsUnitVar:", &my8BitsUnitVar)     // Imprime la dirección de memoria de la variable
	fmt.Println(unsafe.Sizeof(my8BitsUnitVar) * 8)                            // Imprime el tamaño de la variable en bytes
}
