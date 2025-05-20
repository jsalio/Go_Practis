package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Función que modifica el campo Age de un struct usando un puntero
func updateAge(person *Person, newAge int) {
	person.age = newAge // Modifica el campo Age directamente a través del puntero
}

func update(person Person, age int) {
	person.age = age
}

func main() {
	fmt.Println()

	p := Person{
		name: "Joe",
		age:  75,
	}
	fmt.Println("Antes de modificar:", p) // Output: Antes de modificar: {Alice 25}

	// Pasamos un puntero al struct a la función
	updateAge(&p, 30)

	fmt.Println("Después de modificar:", p) // Output: Después de modificar: {Al
}
