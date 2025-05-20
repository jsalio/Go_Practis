package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Función que modifica el campo Age de un struct usando un puntero
func updateAge(person *Person, newAge int) {
	if person == nil {
		println("Invalid ref")
		return
	}
	person.age = newAge // Modifica el campo Age directamente a través del puntero
}

func main() {
	var persona *Person
	fmt.Println("Validar referencias nulas")

	updateAge(persona, 15)
}
