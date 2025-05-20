package main

import (
	"fmt"
)

func main() {
	fmt.Println("If anidados")
	var num int = 0
	print("ingrese numero a evaluar : ")
	fmt.Scanln(&num)
	if num == 0 {
		println("El numero es 0")
	} else {
		if num > 0 {
			println("el numero es positivo")
			if num%2 == 0 {
				println("El numero es par")
			} else {
				println("El numero es inpar")
			}
		} else {
			println("el numero es negativo")
		}

	}
}
