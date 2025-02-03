package main

import (
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Comments Support Examples ===")

	// Example 1: Managing comments in main line
	fmt.Println("\n1. Adding Comments to Main Line")
	game := chess.NewGame()

	// Make moves and add comments
	moves := []struct {
		move    string
		comment string
	}{
		{"e4", "Standard opening"},
		{"e5", "Symmetrical response"},
	}

	for _, m := range moves {
		if err := game.PushMove(m.move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", m.move, err)
			continue
		}
	}

	// Get all comments
	comments := game.Comments()
	fmt.Printf("Game has %d comments\n", len(comments))

	// Example 2: Comments in variations
	fmt.Println("\n2. Comments in Variations")
	game = chess.NewGame()

	// Play main line
	game.PushMove("e4", nil)
	game.PushMove("e5", nil)

	// Go back to add variation
	game.GoBack()

	// Add Sicilian Defense as variation
	game.PushMove("c5", nil)

	// Get comments for the variation
	comments = game.Comments()
	fmt.Printf("Game has %d comments after adding variation\n", len(comments))

	// Example 3: Comments with multiple variations
	fmt.Println("\n3. Comments with Multiple Variations")
	game = chess.NewGame()

	// Play main line
	game.PushMove("e4", nil)
	game.PushMove("e5", nil)
	game.PushMove("Nf3", nil)

	// Go back to add variations
	game.GoBack() // Back to e5
	game.GoBack() // Back to e4

	// Add French Defense
	game.PushMove("e6", nil)

	// Get final comment count
	comments = game.Comments()
	fmt.Printf("Final comment count: %d\n", len(comments))
}
