package function

import "fmt"

func Display(myValue int64) {
	fmt.Println("El valor esSSSSSSSSSS:", myValue)
	fmt.Printf("Type: %T, value: %d, addres:%v\n", myValue, myValue, &myValue)

}
