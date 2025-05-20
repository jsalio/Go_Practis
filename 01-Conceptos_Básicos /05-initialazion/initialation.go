package main

import (
	"fmt"
)

func main() {
	number := 1 //implicit declaration
	fmt.Println("implicit int:", number)

	mynumber := 12 // explicit declaration
	fmt.Println("Int:", mynumber)

	myString := "hello" // explicit decalration for string dta
	fmt.Println("Esto es un texto:", myString)

	mybool := true // explicit decalration for boolean data
	fmt.Println("Esto es un bolean:", mybool)

	myfloat := 3.14
	println("Esto es un valor de punto flotante de 32 bits", myfloat)

	// var aint16E int16 = 37000 // generate error for unsupport content in data type
}
