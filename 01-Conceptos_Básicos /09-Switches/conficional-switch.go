package main

import (
	"fmt"
)

func main() {
	fmt.Println("condicional switch")
	var day string = ""
	print("escriba un numero del 1 al 7 :")
	fmt.Scanln(&day)
	print("sin retorno es :")
	forma_sin_retorno(day)
	println("Con retorno : el dia es :", con_retorno(day))
}

func con_retorno(day string) string {
	switch day {
	case "1":
		return "Lunes"
	case "2":
		return "Martes"
	case "3":
		return "Miercoles"
	case "4":
		return "Jueves"
	case "5":
		return "viernes"
	case "6":
		return "Sabado"
	case "7":
		return "domingo"
	default:
		return "Dia invalido"
	}
}

func forma_sin_retorno(day string) {
	switch day {
	case "1":
		println("el dia es: lunes")

	case "2":
		println("el dia es: martes")

	case "3":
		println("el dia es: miercoles")

	case "4":
		println("el dia es: jueves")

	case "5":
		println("el dia es: viernes")

	case "6":
		println("el dia es: sabado")

	case "7":
		println("el dia es: domingo")

	default:
		println("Dia invalido")
	}
}
