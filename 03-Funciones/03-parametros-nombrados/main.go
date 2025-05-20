package main

import "fmt"

// Función con valores de retorno nombrados para calcular área y perímetro
func calcularRectangulo(largo, ancho float64) (area float64, perimetro float64) {
	area = largo * ancho
	perimetro = 2 * (largo + ancho)
	return // Retorna area y perimetro automáticamente
}

func main() {
	fmt.Println("Valores multiples.")
	area, perimetro := calcularRectangulo(3, 9)

	fmt.Printf("El area es : %f. el perimetro es :%f \n", area, perimetro)
}
