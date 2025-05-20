package main

import "fmt"

func fibonacci(n int) []int {
	// Manejar casos especiales
	if n < 0 {
		return []int{}
	}
	if n == 0 {
		return []int{0}
	}
	// Inicializar el slice con los primeros dos términos
	fib := []int{0, 1}
	// Generar los términos restantes
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	// Devolver solo los primeros n términos
	return fib[:n]
}

func main() {
	n := 10 // Número de términos
	resultado := fibonacci(n)
	fmt.Printf("Los primeros %d términos de la serie de Fibonacci son: %v\n", n, resultado)
}
