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
	Attempts        int32 // Contador de intentos
}

type MatchData string

const (
	IsGreaterThan MatchData = "Is greater than"
	IsLessThan    MatchData = "Is less than"
	IsMatchThat   MatchData = "It can be"
	Match         MatchData = "Congratulations"
	Loose         MatchData = "You lose"
)

func GenerateRandNum() int32 {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000)
	return int32(number)
}

func IsMatch(game *GameData) (MatchData, error) {
	userNumber := ""
	fmt.Print("Enter your guess: ")
	_, err := fmt.Scanln(&userNumber)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	value, err := ConvertToInt(userNumber)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	game.Attempts++ // Incrementar contador de intentos
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
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	return value, nil
}

func GeneratePista(game *GameData) (MatchData, error) {
	userNumber := ""
	fmt.Print("Enter your possible value: ")
	_, err := fmt.Scanln(&userNumber)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	value, err := ConvertToInt(userNumber)
	if err != nil {
		return "", fmt.Errorf("invalid number: %w", err)
	}

	game.PistasRestantes--
	game.Attempts++ // Incrementar contador de intentos
	game.UserPoint -= game.Penalty
	if game.PcNumber > int32(value) {
		return IsGreaterThan, nil
	} else if game.PcNumber < int32(value) {
		return IsLessThan, nil
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
		Attempts:        0, // Inicializar contador de intentos
	}
}

func GameOver(game *GameData) error {
	if game.UserPoint <= 0 {
		return fmt.Errorf("you lose: no points left")
	}
	if game.PistasRestantes <= 0 {
		return fmt.Errorf("you lose: no hints left")
	}
	return nil
}

func Menu(game *GameData) (int, error) {
	option := ""
	fmt.Printf("You have %d points, %d hints remaining, and %d attempts made\n", game.UserPoint, game.PistasRestantes, game.Attempts)
	fmt.Printf("The number I'm thinking of is %d\n", game.PcNumber) // Para depuración, puede eliminarse en producción
	fmt.Println("[1] Enter guess")
	fmt.Println("[2] Request a hint")
	fmt.Println("[3] Exit")
	fmt.Println("")
	fmt.Print("Enter an option: ")
	_, err := fmt.Scanln(&option)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}
	value, err := strconv.Atoi(option)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	if value < 1 || value > 3 {
		return 0, fmt.Errorf("invalid option: must be 1, 2, or 3")
	}
	return value, nil
}

func main() {
	fmt.Println("Guess the number I'm thinking of!")
	const exitOp int = 3
	game := NewGame()
	for {
		if err := GameOver(&game); err != nil {
			fmt.Printf("Game over: %v. You made %d attempts.\n", err, game.Attempts)
			return
		}
		menuOp, err := Menu(&game)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		if game.PcNumber == game.UserNumber || menuOp == exitOp {
			if game.PcNumber == game.UserNumber {
				fmt.Printf("You won in %d attempts!\n", game.Attempts)
			} else {
				fmt.Printf("Exiting game. You made %d attempts.\n", game.Attempts)
			}
			return
		}
		switch menuOp {
		case 1:
			matchResult, err := IsMatch(&game)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			if matchResult == Match {
				fmt.Printf("You won in %d attempts!\n", game.Attempts)
				return
			}
			fmt.Printf("Hint: %s\n", matchResult)
		case 2:
			if game.PistasRestantes <= 0 {
				fmt.Println("Error: No hints remaining")
				continue
			}
			match, err := GeneratePista(&game)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("Hint: %s\n", match)
		default:
			fmt.Println("Error: Invalid option")
		}
	}
}
