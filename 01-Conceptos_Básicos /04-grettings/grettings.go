package main

import (
	"fmt"
)

func main() {

	var name string = ""

	println("Enter your name :")
	fmt.Scanln(&name)

	println("Welcome", name)
}
