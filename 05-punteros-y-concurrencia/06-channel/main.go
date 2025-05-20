package main

import (
	"fmt"
	"sync"
)

func envia_data(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 0; index <= 5; index++ {

		ch <- index
		fmt.Println("Enviado :", index)
	}
	close(ch)
}

func recive_data(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Recibido :", num)
	}
}

func main() {
	println("Usando canales de comunicacion")
	var wg sync.WaitGroup

	wg.Add(2)
	ch := make(chan int)
	go envia_data(ch, &wg)
	go recive_data(ch, &wg)

	wg.Wait()
	fmt.Println("Fin del programa")

}
