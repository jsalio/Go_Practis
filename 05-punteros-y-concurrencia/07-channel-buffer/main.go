package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println()
	var wg sync.WaitGroup

	wg.Add(2)
	ch := make(chan int, 10)

	for num := 1; num <= 10; num++ {
		fmt.Println("Enviando numero : ", num)
		ch <- num
	}
	close(ch)

	for num := range ch {
		println("Recibiendo numero :", num)
	}

}
