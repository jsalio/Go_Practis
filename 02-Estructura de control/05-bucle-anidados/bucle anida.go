package main

import (
	"fmt"
)

func main() {
	fmt.Println("bucles anidados")
	for i := 1; i <= 10; i++ {
		for num := 1; num <= 12; num++ {
			print(i)
			print(" x ")
			print(num)
			print(" = ")
			println(i * num)
		}
	}
}
