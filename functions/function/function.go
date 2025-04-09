package function

import (
	"errors"
	"fmt"
)

// Display sera publica pq inicia con mayuscula si yo llamaria a Display display no funcionaria
// ya que seria una funcion privada
func Display(myValue int64) {
	fmt.Println("El valor esSSSSSSSSSS:", myValue)
	fmt.Printf("Type: %T, value: %d, addres:%v\n", myValue, myValue, &myValue)

}

func Add(x int, y int) int {
	return x + y
}

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("Division cero no permitida")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("Hubo un error! :(")
	}
}

// FUNCION QUE RETORNA DOS VALORES X E Y, AMBOS ENTEROS
func Split(v int) (x, y int) {
	x = v / 2
	y = v - x
	return x, y

}

// FUNCION QUE RETORNA UN VALOR FLOAT64,
// QUE ES LA SUMA DE TODOS LOS VALORES PASADOS COMO PARAMETRO
// sin la necesidad de saber cuantos valores pasaremos
func MSum(values ...float64) (sum float64) {
	for _, v := range values {
		sum += v
	}
	return sum
}
