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
}

func myProcess(a string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("proceso %s:  - num  %d\n", a, i)
	}
}
