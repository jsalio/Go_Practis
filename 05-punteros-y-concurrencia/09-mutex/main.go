package main

import (
	"fmt"
	"sync"
)

// Struct thga represents a person entity
type Person struct {
	//name of preson
	Name string
	//Age of person
	Age int
	//Propertie for prevent multiple manipulation in concurrency scennario
	mu sync.Mutex
}

// increase age of the current person
// the age to increase is a ref of person created
func (p *Person) BirthDay() {
	p.mu.Lock() // in this line protect the current memory reference for others threats
	p.Age++
	p.mu.Unlock() // in this line release the memory ref for others threats
}

// Get te age of the current person
func (p *Person) GetAge() int {
	p.mu.Lock()
	defer p.mu.Unlock() // ensure that the lock every release the ref
	return p.Age
}

func main() {
	fmt.Println("Mutex samples")
	persona := &Person{
		Name: "Raul",
		Age:  3,
	}
	var gw sync.WaitGroup

	for i := 3; i <= 18; i++ {
		gw.Add(1)
		go func() {
			defer gw.Done()
			persona.BirthDay()
			persona.mu.Lock()
			fmt.Printf("%s tiene %d. Memory address object [%p] struct data address[%p] \n", persona.Name, persona.Age, &persona, &persona.Age)
			persona.mu.Unlock()
		}()

	}

	gw.Wait()

	fmt.Println("Edad final:", persona.GetAge())

}
