package main

import (
	"fmt"
)

func capturaDatoUsuario() string {
	input := ""
	print("Escriba un nombre :")
	fmt.Scanln(&input)
	return input
}

func main() {
	fmt.Println("Slices avanzados")
	slice := []string{}
	for {
		input := capturaDatoUsuario()
		if input == "salir" {

			break
		}
		slice = append(slice, input)
		fmt.Printf("nuevo tamano : %d \n", len(slice))
		continue
	}
}
