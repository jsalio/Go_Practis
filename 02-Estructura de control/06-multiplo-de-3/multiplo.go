package main

import (
	"fmt"
)

func main() {
	fmt.Println("Multiplo de 3")
	const multiplo = 3
	for limit := 1; limit <= 50; limit++ {
		residuo := limit % multiplo
		if residuo == 0 {
			fmt.Printf("El número %d es múltiplo de 3\n", limit)
		}
	}
}
