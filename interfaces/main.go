package main

import (
	"fmt"

	"github.com/mateoQuotteri/practicas-en-go/interfaces/vehicles"
)

func main() {
	// Creamos un slice de strings que contiene los tipos de vehículos
	// que se pueden crear. Este slice se utiliza para validar los tipos de vehículos
	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "CAR"}

	var d float64
	// Iteramos sobre el slice de tipos de vehículos y creamos una instancia de cada uno
	// utilizando la función New. Luego, llamamos al método Distance de cada vehículo
	for _, v := range vArray {
		fmt.Printf("Creando vehículo: %s\n", v)
		veh, err := vehicles.New(v, 400)

		if err != nil {
			fmt.Println(err)
			// con el continue, pasamos a la siguiente iteración del bucle
			continue
		}
		// Llamamos al método Distance del vehículo creado y lo sumamos a la variable d
		d += veh.Distance()
		fmt.Printf("Distancia recorrida por el vehículo %s: %.2f km\n", v, veh.Distance())
	}

	// Imprimimos la distancia total recorrida por todos los vehículos
	fmt.Printf("Distancia total recorrida por todos los vehículos: %.2f km\n", d)

}
