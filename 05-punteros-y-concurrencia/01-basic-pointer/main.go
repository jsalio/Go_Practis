package main

import "fmt"

func updatePointValue(ptr *int) {
	*ptr = 100
}

func main() {
	fmt.Println("Punteros")
	x := 2
	fmt.Println("valor original : ", x)
	fmt.Println("puntero : ", &x)
	updatePointValue(&x)
	fmt.Println("valor actualizado : ", x)
	fmt.Println("puntero : ", &x)

}
