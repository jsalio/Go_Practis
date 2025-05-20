package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("El texto %s no se puede convertir: %v\n", str, err)
		return 0 // O manejar el error de otra forma, según el caso
	}
	return num // Retorna el número convertido si no hay error
}

func main() {

	var input string = ""
	//var num int =0
	println("Conversion de tipos")
	print("Ingrese el numero :")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)

	if err != nil {
		println("introdujo el valor inalido ")
		return
	} else {
		fmt.Printf("Numero %d\n", num)
	}
	value := StringToNumber("125")
	println("num : ", value)
	fmt.Println(StringToNumber("123"))
}
