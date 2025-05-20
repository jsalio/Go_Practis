package main

import (
	"fmt"
)

func main() {
	fmt.Println("Trabajndo conm map")
	var dataMap map[string]int = make(map[string]int)
	dataMap["alex"] = 12
	dataMap["ass"] = 2
	dataMap["xina"] = 21
	dataMap["olef"] = 8
	dataMap["xela"] = 10

	for key, value := range dataMap {
		fmt.Printf("Key {%s} , value {%d} \n", key, value)
	}
}
