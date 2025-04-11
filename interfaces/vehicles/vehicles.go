package vehicles

import "fmt"

// Creamos la interfaz Vehicle que define un método Distance
// que debe ser implementado por cualquier tipo que se considere un vehículo.
type Vehicle interface {
	Distance() float64 // Método que calcula la distancia recorrida
}

const (
	CarVehicle         = "CAR"
	MotorcyclyeVehicle = "MOTORCYCLE"
	TruckVehicle       = "TRUCK"
)

func New(vehicle string, time int) (Vehicle, error) {

	// Hacemos un switch para determinar qué tipo de vehículo se está creando
	// y devolvemos el puntero a la estructura correspondiente.
	// Si el tipo de vehículo no es válido, devolvemos un error.
	switch vehicle {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcyclyeVehicle:
		return &Motorcyclye{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("vehicle not found: ", vehicle)
	}
}

// 1era estructura
type Car struct {
	Time int
}

func (c Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

// 2da estructura
type Motorcyclye struct {
	Time int
}

func (c Motorcyclye) Distance() float64 {
	return 120 * (float64(c.Time) / 60)
}

// 3er estructura
type Truck struct {
	Time int
}

func (c Truck) Distance() float64 {
	return 70 * (float64(c.Time) / 60)
}
