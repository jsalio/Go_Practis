package main

import (
	"fmt"
)

// aqui go trata a los parametros variadicos como arrglos al ibgual que C# pero aqie no es necesario usar la clave param y especificar el type como array
func promedio(value ...int) int {
	acum := 0
	for _, num := range value {
		acum = acum + num
	}
	return acum / len(value)
}

func main() {
	fmt.Println("Funciones variadicas")
	fmt.Printf("Al promedio es :%d \n", promedio(2, 4, 6, 8, 10, 22))
}
