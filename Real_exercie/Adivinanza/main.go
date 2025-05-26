// Package main implements a number guessing game where players try to guess
// a randomly generated number between 0 and 999.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GameData holds all the game state information
type GameData struct {
	PcNumber        int32 // The randomly generated number to be guessed
	UserNumber      int32 // The current number guessed by the user
	UserPoint       int32 // Current points of the user (starts at 100)
	PistasRestantes int32 // Number of hints remaining (starts at 5)
	Penalty         int32 // Points deducted for using hints (20 points)
	Attempts        int32 // Counter for total number of attempts made
}

// MatchData represents different possible game states as string constants
type MatchData string

// Game state constants
const (
	IsGreaterThan MatchData = "Is greater than" // Indicates the target number is higher
	IsLessThan    MatchData = "Is less than"    // Indicates the target number is lower
	IsMatchThat   MatchData = "It can be"       // Indicates a possible match
	Match         MatchData = "Congratulations" // Indicates a correct guess
	Loose         MatchData = "You lose"        // Indicates game over
)

// GenerateRandNum generates a random number between 0 and 999
// using the current time as a seed for better randomization
func GenerateRandNum() int32 {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000)
	return int32(number)
}

// IsMatch handles the player's guess attempt
// Parameters:
//   - game: Pointer to the current game state
//
// Returns:
//   - MatchData: The result of the guess attempt
//   - error: Any error that occurred during the process
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

	game.Attempts++ // Increment attempt counter
	if int(game.PcNumber) == value {
		game.UserNumber = int32(value)
		return Match, nil
	}
	game.UserNumber = int32(value)
	game.UserPoint -= game.Penalty - 5 // Deduct points for incorrect guess
	return IsMatchThat, nil
}

// ConvertToInt converts a string input to an integer
// Parameters:
//   - input: The string to convert
//
// Returns:
//   - int: The converted integer
//   - error: Any error that occurred during conversion
func ConvertToInt(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	return value, nil
}

// GeneratePista provides a hint to the player about the target number
// Parameters:
//   - game: Pointer to the current game state
//
// Returns:
//   - MatchData: The hint result
//   - error: Any error that occurred during the process
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

	game.PistasRestantes--         // Decrease remaining hints
	game.Attempts++                // Increment attempt counter
	game.UserPoint -= game.Penalty // Apply penalty for using hint
	if game.PcNumber > int32(value) {
		return IsGreaterThan, nil
	} else if game.PcNumber < int32(value) {
		return IsLessThan, nil
	} else {
		return IsMatchThat, nil
	}
}

// NewGame initializes a new game with default values
// Returns:
//   - GameData: A new game state with initialized values
func NewGame() GameData {
	return GameData{
		PcNumber:        GenerateRandNum(),
		UserNumber:      0,
		UserPoint:       100,
		PistasRestantes: 5,
		Penalty:         20,
		Attempts:        0,
	}
}

// GameOver checks if the game should end based on current conditions
// Parameters:
//   - game: Pointer to the current game state
//
// Returns:
//   - error: Error if game should end, nil if game should continue
func GameOver(game *GameData) error {
	if game.UserPoint <= 0 {
		return fmt.Errorf("you lose: no points left")
	}
	if game.PistasRestantes <= 0 {
		return fmt.Errorf("you lose: no hints left")
	}
	return nil
}

// Menu displays the game menu and handles user input
// Parameters:
//   - game: Pointer to the current game state
//
// Returns:
//   - int: The selected menu option
//   - error: Any error that occurred during the process
func Menu(game *GameData) (int, error) {
	option := ""
	fmt.Printf("You have %d points, %d hints remaining, and %d attempts made\n", game.UserPoint, game.PistasRestantes, game.Attempts)
	// fmt.Printf("The number I'm thinking of is %d\n", game.PcNumber) // Debug line, remove in production
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

// main is the entry point of the program
// It initializes the game and runs the main game loop
func main() {
	fmt.Println("Guess the number I'm thinking of!")
	const exitOp int = 3
	game := NewGame()
	for {
		// Check for game over conditions
		if err := GameOver(&game); err != nil {
			fmt.Printf("Game over: %v. You made %d attempts.\n", err, game.Attempts)
			return
		}

		// Get menu selection
		menuOp, err := Menu(&game)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Check for win or exit conditions
		if game.PcNumber == game.UserNumber || menuOp == exitOp {
			if game.PcNumber == game.UserNumber {
				fmt.Printf("You won in %d attempts!\n", game.Attempts)
			} else {
				fmt.Printf("Exiting game. You made %d attempts.\n", game.Attempts)
			}
			return
		}

		// Handle menu options
		switch menuOp {
		case 1: // Make a guess
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
		case 2: // Request a hint
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
