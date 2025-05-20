package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice basico")

	slice := []string{"Hola", "Mundo", "desde", "GO"}
	for _, text := range slice {
		fmt.Print(" ", text)
	}
}
