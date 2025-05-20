package main

import (
	"fmt"
)

func main() {
	fmt.Println("Work with infinite loop")
	numLimit := 0
	for {
		if numLimit >= 5 {
			break
		}
		println("El numero es :", numLimit)
		numLimit++
	}
}
