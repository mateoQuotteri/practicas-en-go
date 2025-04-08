package function

import "fmt"

//Display sera publica pq inicia con mayuscula si yo llamaria a Display display no funcionaria
//ya que seria una funcion privada
func Display(myValue int64) {
	fmt.Println("El valor esSSSSSSSSSS:", myValue)
	fmt.Printf("Type: %T, value: %d, addres:%v\n", myValue, myValue, &myValue)

}

func Add(x int, y int) int {
	return x + y
}
