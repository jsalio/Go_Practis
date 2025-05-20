package main

import (
	"fmt"
)

func increment() func() int {
	value := 0
	return func() int {
		value++
		return value
	}
}

func main() {
	fmt.Println("Closures")
	next := increment()
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
	fmt.Printf(" %d \n", next())
}
