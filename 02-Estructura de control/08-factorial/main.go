package main

import (
	"fmt"
)

func factorial(limit int) int {
	if limit < 0 {
		return -1
	}

	if limit == 0 || limit == 1 {
		return 1
	}

	resultado := 1
	for n := 1; n <= limit; n++ {
		resultado *= n
	}
	return resultado
}

func main() {
	input := 0

	fmt.Println("Factorial de un numero")
	fmt.Print("ingrese el numero a buscar el factorial :")
	fmt.Scanln(&input)
	// var acum int = 0

	fmt.Printf("El factorial de %d es %d\n", input, factorial(input))

}
