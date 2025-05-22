package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// CharType represents the type of character in a password
type CharType uint

const (
	Lowercase CharType = 1 << iota // 0001
	Uppercase                      // 0010
	Number                         // 0100
	Symbol                         // 1000
)

// Constants for password generation
const (
	MinPasswordLength = 8
	MaxPasswordLength = 128
)

var (
	lowercaseLetters = generateAlphabet(false)
	uppercaseLetters = generateAlphabet(true)
	numbers          = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols          = []string{"!", "@", "#", "%", "^", "&", "*", "(", ":", ")", "_", "=", "+", "-", "[", "]", "{", "}", "|", "\\", "/", "?", "<", ">", "~", "`"}
)

// CaptureBinaryOption captures a yes/no input from the user
func CaptureBinaryOption(label string) (bool, error) {
	option := ""
	fmt.Print(label)
	_, err := fmt.Scanln(&option)
	if err != nil {
		return false, fmt.Errorf("error capturing input: %w", err)
	}
	switch option {
	case "yes", "y":
		return true, nil
	case "no", "n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input: %s", option)
	}
}

// CaptureNumericOption captures a numeric input from the user
func CaptureNumericOption(label string) (int, error) {
	option := ""
	fmt.Print(label)
	_, err := fmt.Scanln(&option)
	if err != nil {
		return 0, fmt.Errorf("error capturing input: %w", err)
	}
	num, err := strconv.Atoi(option)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	return num, nil
}

// BuildSymbol returns a random symbol from the given slice
func BuildSymbol(sliceSymbols []string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(sliceSymbols))))
	if err != nil {
		return "", fmt.Errorf("error generating random number: %w", err)
	}
	return sliceSymbols[n.Int64()], nil
}

// generateAlphabet generates a slice of letters (uppercase or lowercase)
func generateAlphabet(uppercase bool) []string {
	var alphabet []string
	start := 'a'
	if uppercase {
		start = 'A'
	}
	for i := 0; i < 26; i++ {
		alphabet = append(alphabet, string(rune(int(start)+i)))
	}
	return alphabet
}

// GetLowerCaseLetter returns a random lowercase letter
func GetLowerCaseLetter() (string, error) {
	return BuildSymbol(lowercaseLetters)
}

// GetUpperCaseLetter returns a random uppercase letter
func GetUpperCaseLetter() (string, error) {
	return BuildSymbol(uppercaseLetters)
}

// GetNumberLetter returns a random number
func GetNumberLetter() (string, error) {
	return BuildSymbol(numbers)
}

// GetSymbol returns a random symbol
func GetSymbol() (string, error) {
	return BuildSymbol(symbols)
}

// BuildMap creates a map of character types to their generator functions
func BuildMap() map[CharType]func() (string, error) {
	maps := make(map[CharType]func() (string, error))
	maps[Lowercase] = GetLowerCaseLetter
	maps[Uppercase] = GetUpperCaseLetter
	maps[Number] = GetNumberLetter
	maps[Symbol] = GetSymbol
	return maps
}

// RandomCombo returns a random character type from the given combinations
func RandomCombo(combinations []CharType) (CharType, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(combinations))))
	if err != nil {
		return 0, fmt.Errorf("error generating random number: %w", err)
	}
	return combinations[n.Int64()], nil
}

// ExecuteCombination generates a password of the specified length using the given character types
func ExecuteCombination(limite int, operations []CharType) (string, error) {
	functionMap := BuildMap()
	pdw := ""
	for x := 1; x <= limite; x++ {
		ctype, err := RandomCombo(operations)
		if err != nil {
			return "", fmt.Errorf("error selecting character type: %w", err)
		}
		char, err := functionMap[ctype]()
		if err != nil {
			return "", fmt.Errorf("error generating character: %w", err)
		}
		pdw += char
	}
	return pdw, nil
}

func main() {
	fmt.Println("Password Generator")

	longitud, err := CaptureNumericOption("Password length: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if longitud < MinPasswordLength || longitud > MaxPasswordLength {
		fmt.Printf("Password length must be between %d and %d\n", MinPasswordLength, MaxPasswordLength)
		return
	}

	includeUpCase, err := CaptureBinaryOption("Include uppercase letters? (yes/no): ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	includeNumber, err := CaptureBinaryOption("Include numbers? (yes/no): ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	includeSymbols, err := CaptureBinaryOption("Include symbols? (yes/no): ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	combinations := buildMapCombination(includeUpCase, includeNumber, includeSymbols)
	if len(combinations) == 0 {
		fmt.Println("Error: At least one character type must be selected")
		return
	}

	generateNumber, err := CaptureNumericOption("Number of passwords to generate: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\nGenerated Passwords:")
	for n := 1; n <= generateNumber; n++ {
		pwd, err := ExecuteCombination(longitud, combinations)
		if err != nil {
			fmt.Printf("Error generating password: %v\n", err)
			continue
		}
		fmt.Printf("%d. %s\n", n, pwd)
	}
}

// buildMapCombination creates a slice of character types based on user preferences
func buildMapCombination(includeUpCase, includeNumber, includeSymbols bool) []CharType {
	var comb []CharType
	comb = append(comb, Lowercase)
	if includeUpCase {
		comb = append(comb, Uppercase)
	}
	if includeNumber {
		comb = append(comb, Number)
	}
	if includeSymbols {
		comb = append(comb, Symbol)
	}
	return comb
}
