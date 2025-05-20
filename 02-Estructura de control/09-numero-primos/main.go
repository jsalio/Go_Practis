package main

import (
	"fmt"
	"math"
)

func esPrimo(n int) bool {
	// Casos base
	if n <= 1 {
		return false // 0 y 1 no son primos
	}
	if n == 2 {
		return true // 2 es primo
	}
	if n%2 == 0 {
		return false // Números pares mayores que 2 no son primos
	}
	// Verificar divisores hasta la raíz cuadrada de n
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false // Si n es divisible por i, no es primo
		}
	}
	return true
}

func main() {
	input := 0

	fmt.Println("Numeros primos")
	fmt.Print("ingrese el numero :")
	fmt.Scanln(&input)

	if esPrimo(input) {
		fmt.Printf("%d es primo\n", input)
	} else {
		fmt.Printf("%d no es primo\n", input)
	}

}
