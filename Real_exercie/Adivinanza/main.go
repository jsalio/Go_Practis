package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type GameData struct {
	PcNumber        int32
	UserNumber      int32
	UserPoint       int32
	PistasRestantes int32
	Penalty         int32
}

type MatchData string

const (
	IsGreathenThat MatchData = "Is greater than"
	IsLessThat     MatchData = "Is Less Thats"
	IsMatchThat    MatchData = "It can be"
	Match          MatchData = "Congratuliation"
	Loose          MatchData = "You loose"
)

func GenerateRandNum() int32 {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000)
	return int32(number)
}

func IsMatch(game *GameData) (MatchData, error) {
	userNumber := ""
	fmt.Print("Ingresa tu respuesta :")
	_, err := fmt.Scanln(&userNumber)
	if err != nil {
		return "", fmt.Errorf("Invalid input %w", err)
	}
	value, err := ConvertToInt(userNumber)
	if err != nil {
		return "", fmt.Errorf("Invalid input %w", err)
	}

	if int(game.PcNumber) == value {
		game.UserNumber = int32(value)
		return Match, nil
	}
	game.UserNumber = int32(value)
	game.UserPoint -= game.Penalty - 5
	return IsMatchThat, nil
}

func ConvertToInt(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Invalid input %w", err)
	}
	return value, nil
}

func GeneratePista(game *GameData) (MatchData, error) {
	userNumber := ""
	fmt.Print("Ingresa tu posible valor :")
	_, err := fmt.Scanln(&userNumber)
	if err != nil {
		return "", fmt.Errorf("Invalid input %w", err)
	}
	value, err := ConvertToInt(userNumber)
	if err != nil {
		return "", fmt.Errorf("Invalid input %w", err)
	}

	game.UserPoint -= game.Penalty
	if game.PcNumber > int32(value) {
		return IsGreathenThat, nil
	} else if game.PcNumber < int32(value) {
		return IsLessThat, nil
	} else {
		return IsMatchThat, nil
	}
}

func NewGame() GameData {
	return GameData{
		PcNumber:        GenerateRandNum(),
		UserNumber:      0,
		UserPoint:       100,
		PistasRestantes: 5,
		Penalty:         20,
	}
}

func GameOver(game *GameData) error {
	println("Validando el game over")
	if game.UserPoint <= 0 {
		return fmt.Errorf("Perdiste")
	}
	return nil
}

func Menu(game *GameData) (int, error) {
	option := ""
	fmt.Printf("Te quedan %d puntos \n", game.UserPoint)
	fmt.Printf("El numero en que pienso es %d puntos \n", game.PcNumber)
	println("[1] Ingresar respuesta.")
	println("[2] Pedir una pista.")
	println("[3] Salir.")
	println("")
	println("Ingrese una opcion")
	_, err := fmt.Scanln(&option)
	if err != nil {
		return 0, fmt.Errorf("Error de imput %w", err)
	}
	value, err := strconv.Atoi(option)
	if err != nil {
		return 0, fmt.Errorf("Error de imput %w", err)
	}
	return value, nil
}

func main() {
	fmt.Println("Adivina el numero que pienso.")
	const exitOp int = 3
	game := NewGame()
	for {
		if GameOver(&game) != nil {
			println("Terminado")
			return
		}
		menuOp, err := Menu(&game)
		if game.PcNumber == game.UserNumber || menuOp == exitOp || err != nil {
			return
		}
		switch menuOp {
		case 1:
			matchResult, err := IsMatch(&game)
			if err != nil {

			}
			if matchResult == Match {
				println("Finish game")
				return
			}
		case 2:
			match, err := GeneratePista(&game)
			if err != nil {

			}

			fmt.Printf("Pista :%s \n", match)
		default:
			println("Error")
		}
	}
}
