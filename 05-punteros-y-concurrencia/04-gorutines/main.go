package main

import (
	"fmt"
	"sync"
	"time"
)

func Print_number() {
	index := 0
	for {
		if index == 10 {
			break
		}
		index++
		fmt.Println("Valor actual : %d", index)
		time.Sleep(100 * time.Millisecond) //Simula una tarea cada 1 millisegundo
	}
}

func print_number(wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 1; index <= 10; index++ {
		fmt.Printf("Valor actual: %d\n", index)
	}
}

func main() {
	fmt.Println("gorutines")
	go Print_number() //esto es una gorutine que es un proceso que se ejecuta en paralelo en otro hilo
	println("Fin del hilo")
	time.Sleep(1 * time.Second) //simula un processo en el hilo princical que dura un segundo

	//Ejemplo de gorutines avanzado usando buenas practicas

	var wg sync.WaitGroup

	wg.Add(1)
	go print_number(&wg)
	wg.Wait()
	fmt.Println("Fin del programa")

}
