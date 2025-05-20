package main

import "fmt"

func executor(a int, b int, local func(int, int) int) int {
	return local(a, b)
}

func sumar(a int, b int) int {
	result := a + b
	return result
}

func restar(a int, b int) int {
	result := a - b
	return result
}

func main() {
	fmt.Println("funcion como parametro")
	fmt.Printf("la suma es %d y la resta es %d \n", executor(8, 6, sumar), executor(8, 6, restar))
}
