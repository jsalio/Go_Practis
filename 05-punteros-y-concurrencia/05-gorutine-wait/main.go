package main

import (
	"fmt"
	"sync"
)

func print_number(wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 1; index <= 10; index++ {
		fmt.Printf("Valor actual: %d\n", index)
	}
}

func main() {
	fmt.Println("gorutines with wait")
	//Ejemplo de gorutines avanzado usando buenas practicas

	var wg sync.WaitGroup

	wg.Add(1)
	go print_number(&wg)
	wg.Wait()
	fmt.Println("Fin del programa")

}
