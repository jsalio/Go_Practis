package main

import (
	"fmt"
)

type Person struct {
	Name string
	Edad int
}

func main() {
	fmt.Println("Uso de struc basico")

	newPerson := Person{
		Edad: 16,
		Name: "Luis",
	}

	fmt.Printf("El nombre es %s i du edad es %d \n", newPerson.Name, newPerson.Edad)
}
