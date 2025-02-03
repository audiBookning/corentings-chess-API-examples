package main

import (
	"fmt"
	"strings"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Comprehensive Chess Package Example ===\n")

	// 1. Creating games in different ways
	fmt.Println("1. Different ways to create a game:")
	
	// 1.1 From starting position
	gameDefault := chess.NewGame()
	fmt.Printf("Default game position:\n%v\n", gameDefault.Position())

	// 1.2 From FEN string
	fenGame, err := chess.FEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	if err != nil {
		fmt.Printf("Error creating game from FEN: %v\n", err)
		return
	}
	gameFEN := chess.NewGame(fenGame)
	fmt.Printf("\nGame from FEN position:\n%v\n", gameFEN.Position())

	// 1.3 From PGN
	pgnStr := `[Event "Example Game"]
[Site "GitHub.com"]
[Date "2025.02.03"]
[Round "1"]
[White "Player 1"]
[Black "Player 2"]
[Result "*"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 *`

	pgnReader := strings.NewReader(pgnStr)
	if gameFunc, err := chess.PGN(pgnReader); err == nil {
		gamePGN := chess.NewGame(gameFunc)
		fmt.Printf("\nGame from PGN:\n%v\n", gamePGN.Position())
	}

	// 2. Advanced Game Management
	fmt.Println("\n2. Advanced Game Management:")
	game := chess.NewGame()

	// 2.1 Playing moves with variations
	moves := []string{"e4", "e5", "Nf3", "Nc6", "Bb5"}
	fmt.Println("Main line moves:")
	for _, move := range moves {
		if err := game.PushMove(move, nil); err != nil {
			fmt.Printf("Error making move %s: %v\n", move, err)
			return
		}
		fmt.Printf("%s ", move)
	}
	fmt.Println()

	// 2.2 Add variations
	game.GoBack() // Go back to position after Nc6
	fmt.Println("\nAdding variation after Nc6:")
	if err := game.PushMove("Bc4", &chess.PushMoveOptions{ForceMainline: false}); err == nil {
		fmt.Println("Added variation: Bc4 (Italian Game)")
	}

	// Return to main line
	game.NavigateToMainLine()
	
	// 2.3 Position Analysis
	fmt.Println("\n3. Position Analysis:")
	fmt.Printf("Current FEN: %s\n", game.FEN())
	fmt.Printf("Total positions in game: %d\n", len(game.Positions()))
	fmt.Printf("Current turn: %v\n", game.Position().Turn())

	// 2.4 Move Generation and Validation
	fmt.Println("\n4. Move Generation and Validation:")
	validMoves := game.ValidMoves()
	fmt.Printf("Number of valid moves: %d\n", len(validMoves))
	fmt.Println("Sample of possible moves:")
	for i, move := range validMoves {
		if i < 5 { // Show first 5 moves
			fmt.Printf("%s ", move.String())
		}
	}
	fmt.Println("...")

	// 2.5 Game Control and Status
	fmt.Println("\n5. Game Control and Status:")
	draws := game.EligibleDraws()
	fmt.Printf("Available draw methods: %d\n", len(draws))
	if len(draws) > 0 {
		fmt.Println("Draw methods available:")
		for _, method := range draws {
			fmt.Printf("- %v\n", method)
		}
	}

	// 2.6 Tag Pairs Management
	fmt.Println("\n6. Tag Pairs Management:")
	game.AddTagPair("Event", "Comprehensive Example")
	game.AddTagPair("Site", "GitHub.com")
	game.AddTagPair("White", "Player 1")
	game.AddTagPair("Black", "Player 2")
	game.AddTagPair("TimeControl", "5+3")
	
	fmt.Println("Game Tags:")
	fmt.Printf("Event: %s\n", game.GetTagPair("Event"))
	fmt.Printf("Site: %s\n", game.GetTagPair("Site"))
	fmt.Printf("Time Control: %s\n", game.GetTagPair("TimeControl"))

	// 2.7 Game Export
	fmt.Println("\n7. Game Export:")
	fmt.Printf("Final PGN:\n%s\n", game.String())

	// 2.8 Game Outcome
	fmt.Println("\n8. Game Outcome:")
	if game.Outcome() != chess.NoOutcome {
		fmt.Printf("Game ended: %s by %s\n", game.Outcome(), game.Method())
	} else {
		fmt.Println("Game is ongoing")
		
		// Demonstrate resignation
		game.Resign(chess.Black)
		fmt.Printf("After Black resigns: %s by %s\n", game.Outcome(), game.Method())
	}
}
