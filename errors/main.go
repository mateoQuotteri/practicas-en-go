package main

import (
	"errors"
	"fmt"
)

type MyCustomError struct {
	ID string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Error ID: %s", e.ID)
}

func main() {
	//1ra forma de crear un error
	// La función errors.New crea un nuevo error con el mensaje proporcionado.
	/*var err error
	err = errors.New("Este es un error")

	fmt.Println(err) // Imprime el error

	//2da forma de crear un error
	err2 := fmt.Errorf("Este es otro error: %w", err) // %w permite envolver errores
	fmt.Println(err2)

	//3ra forma de crear un error
	erro3 := testError(2)
	erro4 := testError(1)
	//fmt.Println(erro3)
	//fmt.Println(erro4)

	err5 := errors.Join(erro3, erro4) // Combina varios errores en uno solo
	fmt.Println(err5)

	// Comprobando si un erro5 contiene erro3

	fmt.Println(errors.Is(err5, erro3)) // Imprime el error combinado
	*/

	//el defer debe estar arriba de cualquier error que pueda aparecer
	defer func() {
		fmt.Println("Fin del programa")
		r := recover() // Recupera el pánico
		if r != nil {
			fmt.Println("Se produjo un error:", r)
		}
	}()
	//v := 0
	//_ = 5 / v

	panic("Se produjo un pánico") // Genera un pánico
	fmt.Println("Fin del programa principal")

}

func testError(v int) error {
	if v == 1 {
		return &MyCustomError{ID: "12345"}
	} else {
		//retorno un error cualquiera
		return errors.New("Error: El valor no es 1")
	}

}
