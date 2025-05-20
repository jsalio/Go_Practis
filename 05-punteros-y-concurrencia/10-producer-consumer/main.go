package main

import (
	"fmt"
	"sync"
	"time"
)

/*
comentario

multilinea
*/
type DataStream struct {
	Identificator []int
	cap           int
	mutex         sync.Mutex
}

func Generator(id int, buffer *DataStream, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		item := id*100 + i // Ítem único basado en ID del productor
		buffer.mutex.Lock()

		// Esperar si el buffer está lleno
		for len(buffer.Identificator) >= buffer.cap {
			buffer.mutex.Unlock()
			time.Sleep(100 * time.Millisecond)
			buffer.mutex.Lock()
		}

		// Agregar ítem al buffer y enviar al canal
		buffer.Identificator = append(buffer.Identificator, item)
		fmt.Printf("Productor %d produjo: %d, Buffer: %v\n", id, item, buffer.Identificator)
		ch <- item

		buffer.mutex.Unlock()
		time.Sleep(200 * time.Millisecond) // Simular tiempo de producción
	}
}

func Consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("Consumidor %d consumió: %d\n", id, item)
		time.Sleep(300 * time.Millisecond) // Simular tiempo de consumo
	}
}

func main() {
	buffer := &DataStream{
		Identificator: make([]int, 0),
		cap:           3,
	}

	ch := make(chan int, buffer.cap)

	var pgw, cgw sync.WaitGroup

	numGenerators := 2

	for i := 1; i <= numGenerators; i++ {
		pgw.Add(1)
		go Generator(i, buffer, ch, &pgw)
	}

	numConsumidores := 2
	for i := 1; i <= numConsumidores; i++ {
		cgw.Add(1)
		go Consumer(i, ch, &cgw)
	}

	pgw.Wait()
	close(ch)

	cgw.Wait()
}
