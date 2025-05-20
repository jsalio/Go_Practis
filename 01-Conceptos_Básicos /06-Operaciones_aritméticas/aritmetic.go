package main

import (
	"fmt"
)

func main() {
	num_one := 0
	num_two := 0

	print("Enter number #1 :")
	fmt.Scanln(&num_one)
	print("Enter number #1 :")
	fmt.Scanln(&num_two)

	println("Sum   :", num_one+num_two)
	println("Rest  :", num_one-num_two)
	println("Multi :", num_one*num_two)
	println("Div   :", num_one/num_two)
}
