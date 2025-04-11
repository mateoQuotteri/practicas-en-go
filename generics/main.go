package main

import "fmt"

func main() {

	//v1 := []float64{1.0, 2.0, 3.0}
	// v2 := []int32{1, 2, 3}

	/*fmt.Println("La suma de los valores de v1 es:", Sum01(v1))
	fmt.Println("La suma de los valores de v2 es:", Sum01(v2))

	fmt.Println("La suma de los valores de v1 es:", Sum02(v1))
	fmt.Println("La suma de los valores de v2 es:", Sum02(v2))
	*/
	AnyTipe(1, 8)
	AnyTipe2("Hola", 8)

}

// Generaremos una funcion
// que reciba dos valores de un tipo generico y los sume
func Sum01[N int32 | int64 | int | float32 | float64](v []N) N {
	var total N
	for _, v := range v {
		total += v
	}
	return total
}

type Number interface {
	int32 | int64 | int | float32 | float64
}

// Generaremos una funcion
// que reciba dos valores de un tipo generico y los sume
func Sum02[N Number](v []N) N {
	var total N
	for _, v := range v {
		total += v
	}
	return total
}

//Recibe CUAKQUIER TIPO y devuelve CUAKQUIER TIPO
// LOS DOS VALORES PASADOS POR PARAMRETROS DEBEN SER DEL MMISMO
//TIPO
func AnyTipe[N any](v1, v2 N) {
	fmt.Println("Valores:", v1, v2)
}

//Recibe CUAKQUIER TIPO y devuelve CUAKQUIER TIPO
// las dos variables pueden ser de diferente tipo
// y no tienen que ser del mismo tipo
func AnyTipe2[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println("Valores:", v1, v2)
}
