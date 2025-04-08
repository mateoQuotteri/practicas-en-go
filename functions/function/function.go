package function

import "fmt"

func Display(myValue int) {
	fmt.Println("El valor es:", myValue)
	fmt.Printf("Type: %T, value: %d, addres:%v\n", myValue, myValue, &myValue)

}
