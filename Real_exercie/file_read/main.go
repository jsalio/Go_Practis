package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FrecuencyWordCount(texto string, wordMap map[string]int) {
	for _, word := range strings.Split(texto, " ") {
		_, exists := wordMap[word]
		if !exists {
			wordMap[word] = 1
		} else {
			wordMap[word]++
		}
	}
}

func WordCount(texto string, counter *int) {
	var splitLeng = strings.Split(texto, " ")
	*counter += len(splitLeng)
}

func CharaterCounter(texto string, counter *int) {
	var textLeng = len(texto)
	*counter += textLeng
}

func ReadFile(filepath string) (wordCount int, characterCount int, frequencyWord map[string]int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, 0, make(map[string]int), nil
		}
		return 0, 0, make(map[string]int), fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()
	lineReader := bufio.NewScanner(file)
	var wordFCounter = make(map[string]int)
	var wordCounter int
	var charCounter int

	for lineReader.Scan() {
		FrecuencyWordCount(lineReader.Text(), wordFCounter)
		WordCount(lineReader.Text(), &wordCounter)
		CharaterCounter(lineReader.Text(), &charCounter)
	}
	wordCount = wordCounter
	characterCount = charCounter
	frequencyWord = wordFCounter
	return
}

func main() {
	fmt.Println("Read file")
	fmt.Print("Ingrese el nombre o la ruta del archivo :")
	filePath := ""
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		println("La entrada es incorrecta")
		return
	}
	wordCount, characterCount, frequencyWord, err := ReadFile(filePath)
	if err != nil {
		return
	}
	fmt.Printf("Contenido :\n Palabras: [%d] \n Caracteres [%d] \n,", wordCount, characterCount)
	for word, freq := range frequencyWord {
		fmt.Printf("Palabra : [%s] frequencia [%d] \n", word, freq)
	}
}
