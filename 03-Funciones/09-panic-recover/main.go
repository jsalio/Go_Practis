package main

import (
	"fmt"
)

func main() {
	fmt.Println("Uso de panic and recover")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de panic:", r)
		}
	}()
	panic("Algo salió mal")
}
