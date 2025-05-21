package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func CaptureBinaryOption(label string) (bool, error) {
	option := ""
	fmt.Print(label)
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Errorf("Error de captura %s", err.Error())
		return false, err
	}
	switch option {
	case "yes", "y":
		return true, nil
	case "no", "n":
		return false, nil
	default:
		panic("Error de entrada")
	}
}

func CaptureNumericOption(label string) (int, error) {
	option := ""
	fmt.Print(label)
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Errorf("Error de captura %s", err.Error())
		return 0, err
	}
	return strconv.Atoi(option)
}

func BuildSimbol(sliceSymbols []string) string {
	rand.Seed(time.Now().UnixNano())
	randindex := rand.Intn(len(sliceSymbols))
	return sliceSymbols[randindex]
}

func generateAlphabetUp() []string {
	var alphabet []string
	for i := 0; i < 26; i++ {
		alphabet = append(alphabet, string(rune('A'+i)))
	}

	return alphabet
}

func generateAlphabet() []string {
	var alphabet []string
	for i := 0; i < 26; i++ {
		alphabet = append(alphabet, string(rune('a'+i)))
	}

	return alphabet
}

type CharType uint

const (
	Lowercase CharType = 1 << iota // 0001
	Uppercase                      // 0010
	Number                         // 0100
	Symbol                         // 1000
)

func GetLowerCaseLetter() string {
	sliceAlphabeticDn := generateAlphabet()
	letter := BuildSimbol(sliceAlphabeticDn)
	return letter
}

func GetUpperCaseLetter() string {
	sliceAlphabeticDn := generateAlphabetUp()
	letter := BuildSimbol(sliceAlphabeticDn)
	return letter
}

func GetNumberLetter() string {
	sliceNumber := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	letter := BuildSimbol(sliceNumber)
	return letter
}

func GetSymbol() string {
	sliceSymbols := []string{"!", "@", "#", "%", "^", "&", "*", "(", ":", ")", "_", "="}
	letter := BuildSimbol(sliceSymbols)
	return letter
}

func BuildMap() map[CharType]func() string {
	maps := make(map[CharType]func() string)
	maps[Lowercase] = GetLowerCaseLetter
	maps[Uppercase] = GetUpperCaseLetter
	maps[Number] = GetNumberLetter
	maps[Symbol] = GetSymbol
	return maps
}

func RandomCombo(combinations []CharType) CharType {
	rand.Seed(time.Now().UnixNano())
	return combinations[rand.Intn(len(combinations))]
}

func ExecuteCombination(limite int, operations []CharType) (string, error) {
	functionMap := BuildMap()
	pdw := ""
	for x := 1; x <= limite; x++ {
		ctype := RandomCombo(operations)
		pdw += functionMap[ctype]()
	}
	return pdw, nil
}

func main() {
	fmt.Println("Generador de contrasenas.")
	longitud, err := CaptureNumericOption("Longitud de la contrasena   :")
	if err != nil {
		panic("Error")
	}
	includeUpCase, err1 := CaptureBinaryOption("Incluir Mayusculas (Yes/no) :")
	if err1 != nil {
		panic("Error")
	}
	includeNumber, err2 := CaptureBinaryOption("Incluir numeros (yes/no)    :")
	if err2 != nil {
		panic("Error")
	}
	includeSimbols, err3 := CaptureBinaryOption("Incluir simbolos (yes/no)   :")
	if err3 != nil {
		panic("Error")
	}

	generateNumber, err4 := CaptureNumericOption("Cantidad de contrasenas a generar :")
	if err4 != nil {
		panic("Error")
	}

	for n := 1; n <= generateNumber; n++ {
		pwd, errOut := ExecuteCombination(longitud, buildMapCombination(includeUpCase, includeNumber, includeSimbols))
		if errOut != nil {
			panic("Error saliendo")
		}
		println("Nueva contrasena :", pwd)
	}
}

func buildMapCombination(includeUpCase, includeNumber, includeSimbols bool) []CharType {
	var comb []CharType
	comb = append(comb, Lowercase)
	if includeUpCase {
		comb = append(comb, Uppercase)
	}
	if includeNumber {
		comb = append(comb, Number)
	}
	if includeSimbols {
		comb = append(comb, Symbol)
	}
	return comb
}
