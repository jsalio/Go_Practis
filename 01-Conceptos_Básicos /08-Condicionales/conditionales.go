package main

import (
	"fmt"
)

func main() {
	fmt.Println("Condicionales")
	var input int = 0
	print("ingrese el numero a evaluar :")
	fmt.Scanln(&input)
	numMod := input % 2
	if numMod == 0 {
		println("Es par.")
	} else {
		println("no es par")
	}
}
