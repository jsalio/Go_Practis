package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(500 * time.Millisecond)
		ch1 <- 2
		time.Sleep(500 * time.Millisecond)
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		time.Sleep(700 * time.Millisecond)
		ch2 <- 4
		time.Sleep(700 * time.Millisecond)
		close(ch2)
	}()

	for {
		select {
		case num1, ok := <-ch1:
			if !ok {
				fmt.Println("Canal 1 cerrado, terminando")
				return
			}
			fmt.Println("Recibido canel 1:", num1)
		case num2, ok := <-ch2:
			if !ok {
				fmt.Println("Canal 2 cerrado, terminando")
				return
			}
			fmt.Println("Recibido canel 2:", num2)
		default:
			fmt.Println("Esperando...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}
