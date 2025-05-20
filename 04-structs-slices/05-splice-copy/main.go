package main

import (
	"fmt"
)

func main() {
	fmt.Println("copiando slices")
	slice := []string{"1", "2", "3", "4", "5"}
	spliceCopy := make([]string, len(slice))
	copy(spliceCopy, slice)
	spliceCopy[3] = "0"
	for n := 0; n <= len(slice)-1; n++ {
		fmt.Printf("original :%s vs  copia :%s \n", slice[n], spliceCopy[n])
	}
}
