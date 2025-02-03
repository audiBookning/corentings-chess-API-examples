package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Game Navigation Example ===")
	game := chess.NewGame()

	// Make some moves
	moves := []string{"e4", "e5", "Nf3", "Nc6", "Bb5"}
	for _, move := range moves {
		game.PushMove(move, nil)
	}

	fmt.Printf("Initial position after all moves:\n%v\n", game.Position())

	// Navigate backwards
	game.GoBack()
	fmt.Printf("\nAfter going back one move:\n%v\n", game.Position())

	game.GoBack()
	fmt.Printf("\nAfter going back another move:\n%v\n", game.Position())

	// Navigate forwards
	game.GoForward()
	fmt.Printf("\nAfter going forward one move:\n%v\n", game.Position())

	// Check position in game
	fmt.Printf("\nIs at start: %v\n", game.IsAtStart())
	fmt.Printf("Is at end: %v\n", game.IsAtEnd())

	// Get all positions in game history
	positions := game.Positions()
	fmt.Printf("\nTotal number of positions in game: %d\n", len(positions))
}
