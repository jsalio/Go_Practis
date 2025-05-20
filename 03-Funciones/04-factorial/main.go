package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	fmt.Printf("El factorial de %d es %d\n", n, factorial(n))
}
