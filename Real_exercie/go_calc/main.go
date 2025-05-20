package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Determina si se encuantra un operador matematico

si existe retorna true y el operador detectado

input : entrada original del usuario
*/
func is_sep(input string) (bool, string) {
	operationChar := []string{"+", "-", "*", "/", "x"}

	charMap := make(map[string]bool)
	for _, c := range operationChar {
		charMap[c] = true
	}

	var res []string
	for _, r := range input {
		char := string(r)
		if charMap[char] {
			res = append(res, char)
		}
	}

	if len(res) > 1 {
		return false, "multiple operator detected"
	}
	return true, res[0]

}

/*
Funcion para sumar valores

a: valor 1

b: valor 2
*/
func Suma(a int, b int) float32 {

	return float32(a) + float32(b)
}

/*
Funcion para restar valores

a: valor 1

b: valor 2
*/
func Resta(a int, b int) float32 {
	return float32(a) - float32(b)
}

/*
Funcion para multiplicar valores

a: valor 1

b: valor 2
*/
func multi(a int, b int) float32 {
	return float32(a) * float32(b)
}

/*
Funcion para dividir valores

a: valor 1

b: valor 2
*/
func div(a int, b int) float32 {
	if b == 0 {
		panic("Div by zero detected.")
	}
	return float32(a) / float32(b)
}

/*
Separa el operador de los datos a operar

si el operador es valido retorna la funcion de la operacion y en un slice los datos a operar
*/
func operation_split(input string) (func(int, int) float32, []string) {
	isSep, op := is_sep(input)
	if !isSep {
		panic(op)
	}
	split := strings.Split(input, op)

	switch op {
	case "+":
		return Suma, split
	case "-":
		return Resta, split
	case "*", "x":
		return multi, split
	case "/":
		return div, split
	default:
		return func(i1, i2 int) float32 {
			return 0
		}, make([]string, 0)
	}

}

// Convertir a entero una cadena de texto
func castToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return num
}

// Esperar la entrada de un usuario y ejecutar el flujo principal
func wait_Prompt() bool {
	input := ""
	fmt.Print("go-calc-cli (Ingrese una operación como '2 + 3' o 'exit' para salir) >> ")
	fmt.Scanln(&input)
	if input == "exit" {
		return true
	}
	op, values := operation_split(input)

	result := op(castToInt(values[0]), castToInt(values[1]))
	fmt.Printf("Resultado : %1.2f \n", result)
	return false
}

func main() {
	fmt.Println("Calculadora cli")
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recuperado de panic:", r)
				}
			}()
			isExit := wait_Prompt()
			if isExit {
				panic("exit")
			}
		}()

	}
}
