package main

import (
	"fmt"
)

func main() {
	input := 0
	fmt.Println("Validacioon de entrada")

	for {
		fmt.Print("Escriba un numero del uno al 100 :")
		fmt.Scanln(&input)
		if input >= 1 && input <= 100 {
			break
		} else {
			println("Valor invalido")
			continue
		}
	}
}
