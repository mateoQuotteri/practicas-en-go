package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`        // Etiqueta JSON para el campo Id
	Age      int    `json:"age"`       // Etiqueta JSON para el campo Age
	Name     string `json:"name"`      // Etiqueta JSON para el campo Name
	LastName string `json:"last_name"` // Etiqueta JSON para el campo LastName
	Email    string `json:"email"`     // Etiqueta JSON para el campo Email

}

// Display es un método que devuelve una representación en cadena del usuario
func (u User) Display() {
	v, _ := json.Marshal(u) // Convierte el struct a JSON

	fmt.Println(string(v))
}

// COMO UN SETTER EN JAVA EL ASTERISCO ES FUNDAMENTAL PARA QUE EL ATRIBUTO SE MODIFIQUE "GLOBALMENTE"
func (u *User) setName(name string) {
	u.Name = name
}

func main() {
	user := User{
		Id:       1,
		Age:      30,
		Name:     "Mateo",
		LastName: "Quotteri",
		Email:    "quotteri@gmail.com",
	}

	//fmt.Println("User struct:", user)

	//v, err := json.Marshal(user) // Convierte el struct a JSON

	//fmt.Println("JSON:", string(v), "error:", err)

	user.Display()
	user.setName("Tobias")
	user.Display()

}
