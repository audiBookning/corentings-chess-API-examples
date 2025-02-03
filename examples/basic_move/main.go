package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	// Create a new game
	game := chess.NewGame()

	// Make some moves
	moves := []string{"e4", "e5", "Nf3", "Nc6"}
	for _, move := range moves {
		err := game.PushMove(move, nil)
		if err != nil {
			fmt.Printf("Error making move %s: %v\n", move, err)
			return
		}
		fmt.Printf("Made move: %s\n", move)
	}

	// Print current position
	fmt.Printf("\nCurrent position:\n%v\n", game.Position())

	// Get valid moves
	validMoves := game.ValidMoves()
	fmt.Printf("\nValid moves: %d\n", len(validMoves))

	// Check game status
	if game.Outcome() != chess.NoOutcome {
		fmt.Printf("Game ended: %s by %s\n", game.Outcome(), game.Method())
	}
}
