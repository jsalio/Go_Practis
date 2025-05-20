package main

import (
	"fmt"
)

type Person struct {
	Name string
	Edad int
}

func (p Person) Description() string {
	return fmt.Sprintf("El nombre de la persona es %s", p.Name)
}

func main() {
	fmt.Println("Structure con metdos")
	newPerson := Person{
		Edad: 16,
		Name: "Luis",
	}
	fmt.Println("Description :", newPerson.Description())
}
