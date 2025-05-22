package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// isPrime verifica si un número es primo
func isPrime(x int) bool {
	if x <= 1 {
		return false // 0 y 1 no son primos
	}
	if x == 2 {
		return true // 2 es primo
	}
	if x%2 == 0 {
		return false // Si es par y mayor que 2, no es primo
	}
	// Probar divisores impares hasta sqrt(x)
	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		if x%i == 0 {
			return false // Si es divisible por i, no es primo
		}
	}
	return true
}

func calculateRange(min int, max int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i <= max; i++ {
		if isPrime(i) {
			channel <- i
		}
	}
}

func ViewResults(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sliceValues []int
	for value := range channel {
		sliceValues = append(sliceValues, value)
	}
	sort.Ints(sliceValues)
	fmt.Printf("Números primos encontrados: %v\n", sliceValues)
	fmt.Printf("Cantidad de primos: %d\n", len(sliceValues))
}

func ProcessUserSearch(limit int) {
	start := time.Now()
	channel := make(chan int)
	limitPerProcess := 100 //limit / runtime.NumCPU() //100
	splitIn := (limit + limitPerProcess - 1) / limitPerProcess

	var wg sync.WaitGroup
	var calcWG sync.WaitGroup

	wg.Add(1)
	go ViewResults(channel, &wg)

	for x := 0; x < splitIn; x++ {
		calcWG.Add(1)
		min := x*limitPerProcess + 1
		max := min + limitPerProcess - 1
		if max > limit {
			max = limit
		}
		go calculateRange(min, max, channel, &calcWG)
	}

	calcWG.Wait()

	close(channel)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}

func main() {
	fmt.Println("Prime number search")
	limit := 0
	fmt.Print("Ingrese el limite a buscar (0-1000) :")
	_, err := fmt.Scanln(&limit)
	if err != nil {
		fmt.Errorf("Valor invalido %s", err.Error())
		return
	}

	ProcessUserSearch(limit)
}
