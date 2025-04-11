package main

import (
	"fmt"
	"time"
)

func main() {

	go myProcess("HA")
	go myProcess("BA")
	// Esperar lo suficiente para que las goroutines terminen
	time.Sleep(8 * time.Second)

	myFirstChannel := make(chan string)
	//Con chan defino que en lugar de una variable esto sera un canal

	go myProcessWithChannel("RA", myFirstChannel)
	//Con el canal definido, llamo a la funcion que lo utiliza
	// En este caso, el canal es un string, por lo que el valor enviado debe ser un string

	result := <-myFirstChannel
	fmt.Printf("El resultado del canal es: %s\n", result)
	close(myFirstChannel) // Cierro el canal, ya no lo necesito
}

func myProcessWithChannel(a string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("proceso %s:  - num  %d\n", a, i)
	}
	c <- a // Envio el valor al canal
	// En este caso, el canal es un string, por lo que el valor enviado debe ser un string
}

func myProcess(a string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("proceso %s:  - num  %d\n", a, i)
	}
}
