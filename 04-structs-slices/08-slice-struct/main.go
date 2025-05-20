package main

import (
	"fmt"
)

type Products struct {
	nombre string
	price  float64
}

func main() {
	fmt.Println("slice con struct")
	products := []Products{
		{
			nombre: "martillo",
			price:  12.78,
		},
		{
			nombre: "clavo",
			price:  1.75,
		},
	}
	total := Total_Price(products)
	fmt.Printf("Total es: %f \n", total)

}

func Total_Price(products []Products) float64 {
	var total float64 = 0
	for _, produt := range products {
		total += produt.price
	}
	return total
}
