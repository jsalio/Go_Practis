package main

import (
	"fmt"
)

func main() {
	var number = 1 //implicit declaration
	fmt.Println("implicit int:", number)

	var mynumber int = 12 // explicit declaration
	fmt.Println("Int:", mynumber)

	var aInt8 int8 = 127 // explicit decalration for a 8 bit number
	fmt.Println("int 8:", aInt8)

	var aint16 int16 = 15 // explicit decalration for a 16 bit number
	fmt.Println("int 16:", aint16)

	var aInt32 int32 = 50000 // explicit decalration for a 32 bit number
	fmt.Println("int 32:", aInt32)

	var aInt64 int64 = 50000 // explicit decalration for a 64 bit number
	fmt.Println("int 64:", aInt64)

	var myString string = "hello" // explicit decalration for string dta
	fmt.Println("Esto es un texto:", myString)

	var mybool bool = true // explicit decalration for boolean data
	fmt.Println("Esto es un bolean:", mybool)

	var myfloat float32 = 3.14
	println("Esto es un valor de punto flotante de 32 bits", myfloat)

	// var aint16E int16 = 37000 // generate error for unsupport content in data type
}
