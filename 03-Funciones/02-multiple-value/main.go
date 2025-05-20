package main

import "fmt"

func valor_multiple(num int) (int, int) {
	return num * num, num * num * num
}

func main() {
	fmt.Println("Valores multiples.")
	q, c := valor_multiple(3)

	fmt.Printf("El cuadrado es : %d. el cubo es :%d \n", q, c)
}
