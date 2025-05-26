# Number Guessing Game in Go

This is a console-based number guessing game implemented in Go. The game challenges players to guess a randomly generated number between 0 and 999, providing hints and tracking player performance.

## Game Features

- Random number generation between 0 and 999
- Point system starting at 100 points
- Limited hints (5 available)
- Penalty system for using hints and incorrect guesses
- Attempt counter
- Interactive menu system

## Code Structure

### Data Structures

#### GameData
```go
type GameData struct {
    PcNumber        int32  // The number to be guessed
    UserNumber      int32  // Player's current guess
    UserPoint       int32  // Player's current points
    PistasRestantes int32  // Remaining hints
    Penalty         int32  // Penalty points for using hints
    Attempts        int32  // Number of attempts made
}
```

#### MatchData
```go
type MatchData string
```
Constants for different game states:
- `IsGreaterThan`: Indicates the target number is greater
- `IsLessThan`: Indicates the target number is less
- `IsMatchThat`: Indicates a possible match
- `Match`: Indicates a correct guess
- `Loose`: Indicates game over

### Core Functions

#### GenerateRandNum()
Generates a random number between 0 and 999 using the current time as a seed.

#### IsMatch(game *GameData)
Handles the player's guess:
- Takes user input
- Validates the input
- Compares with the target number
- Updates points and attempts
- Returns appropriate MatchData

#### GeneratePista(game *GameData)
Provides hints to the player:
- Takes user input
- Compares with target number
- Reduces remaining hints
- Applies penalty points
- Returns whether the target is greater or less

#### NewGame()
Initializes a new game with:
- Random target number
- 100 starting points
- 5 available hints
- 20 point penalty
- 0 attempts

#### GameOver(game *GameData)
Checks game over conditions:
- No points remaining
- No hints remaining

#### Menu(game *GameData)
Displays the game menu with options:
1. Enter guess
2. Request a hint
3. Exit

## Game Rules

1. The game starts with 100 points
2. Each hint costs 20 points
3. Each incorrect guess reduces points by 15 (penalty - 5)
4. Players have 5 hints available
5. The game ends when:
   - Player runs out of points
   - Player runs out of hints
   - Player correctly guesses the number
   - Player chooses to exit

## How to Play

1. Run the game
2. Choose an option from the menu:
   - Option 1: Make a guess
   - Option 2: Request a hint
   - Option 3: Exit the game
3. Follow the hints to guess the correct number
4. Try to guess the number with the fewest attempts possible

## Error Handling

The game includes comprehensive error handling for:
- Invalid input
- Invalid numbers
- Invalid menu options
- Game over conditions

## Running the Game

To run the game, ensure you have Go installed and execute:
```bash
go run main.go
``` 