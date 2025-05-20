package main

import (
	"fmt"
)

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divicion por cero detectada")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling")
	value, error := div(9, 0)
	if error == nil {
		println("value: ", value)
	} else {
		println("Error :", error.Error())
	}
}
