package main

import (
	"fmt"
)

func main() {
	fmt.Println("If anidados")
	var edad int = 0
	print("Ingrese la edad : ")
	fmt.Scanln(&edad)
	switch {
	case edad <= 12:
		println("Nino")
	case edad > 12 && edad <= 18:
		println("Joven")
	case edad > 18:
		println("Adulto")
	default:
		println("Rango o valor incorrecto")
	}
}
