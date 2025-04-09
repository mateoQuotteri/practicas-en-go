package main

import (
	"fmt"
	//	"unsafe"
	"github.com/mateoQuotteri/practicas-en-go/functions/function"
)

func main() {
	/*var myIntVar int64 = 12
	function.Display(myIntVar)

	v := function.Add(2, 3)
	fmt.Println("El resultado de la suma es:", v)*/

	value, err := function.Calc(function.SUM, 20.12, 34)
	fmt.Println("value: ", value, "- err: ", err)

	value, err = function.Calc(function.SUB, 20.12, 34)

	xVal, yVal := function.Split(10)
	fmt.Println("xVal:", xVal, "yVal:", yVal)

	val45 := function.MSum(100, 200, 300, 400, 500)
	fmt.Println("Suma se MSum:", val45)

	//var myIntVar int
	//myIntVar = -12
	//fmt.Println("Enter an integer:", myIntVar)

	// := "Hello, World!"
	//fmt.Println("Enter a string:", myString)
	//fmt.Println("myString espacio en memoria:", &myString)

	//{
	//const myConstVar = 3.14 // Constantes no pueden ser cambiadas
	//fmt.Println("mi constante:", myConstVar)

	// myConstVar = 2.71 // Esto generará un error de compilación

	//}

	//var my8BitsUnitVar uint8 = 255 // uint8 puede almacenar valores de 0 a 255
	//fmt.Println("my8BitsUnitVar:", my8BitsUnitVar)
	//fmt.Println("type de my8BitsUnitVar:", fmt.Sprintf("%T", my8BitsUnitVar)) // Imprime el tipo de la variable
	//fmt.Println("Espacio en memoria de my8BitsUnitVar:", &my8BitsUnitVar)     // Imprime la dirección de memoria de la variable
	//fmt.Println(unsafe.Sizeof(my8BitsUnitVar) * 8)                            // Imprime el tamaño de la variable en bytes

	/*var my8BitsUintVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n",
		my8BitsUintVar,
		my8BitsUintVar,
		unsafe.Sizeof(my8BitsUintVar),
		unsafe.Sizeof(my8BitsUintVar)*8)

	var myFloat32Var float32 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var,
		unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myStringVar string = "Hello, World!"
	fmt.Printf("mi valor es: %s\n", myStringVar)*/

}
