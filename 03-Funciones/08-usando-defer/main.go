package main

import (
	"fmt"
	"os"
)

func leerArchivo(nombre string) error {
	archivo, err := os.Open(nombre)
	if err != nil {
		return err
	}
	defer archivo.Close() // Asegura que el archivo se cierre al final
	// Procesar archivo
	return nil
}

func main() {
	fmt.Println("usando defer")
	leerArchivo("")
	defer fmt.Println("Primero en ejecutarse (último defer)")
	defer fmt.Println("Segundo en ejecutarse")
	defer fmt.Println("Tercero en ejecutarse")
	fmt.Println("Ejecutando función principal")
}
