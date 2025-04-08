package main

import (
	"fmt"
)

func main() {

	numero := 10
	/*
		if numero < 10 {
			fmt.Println("El número es menor que 10")

		} else if numero == 10 {
			fmt.Println("El número es igual a 10")
		} else {
			fmt.Println("El número es mayor que 10")
		}


	*/
	switch numero {
	case 1:
		fmt.Println("El número es uno")
	case 2:
		fmt.Println("El número es dos")
	case 3:
		fmt.Println("El número es tres")
	default:
		fmt.Println("El número no es uno, dos o tres")
	}
}
