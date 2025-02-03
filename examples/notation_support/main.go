package main

import (
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Chess Notation Examples ===")

	// Example 1: Standard Algebraic Notation (SAN)
	fmt.Println("\n1. Standard Algebraic Notation (SAN)")
	game := chess.NewGame()

	// Make moves using SAN
	sanMoves := []string{
		"e4",    // Pawn to e4
		"e5",    // Pawn to e5
		"Nf3",   // Knight to f3
		"Nc6",   // Knight to c6
		"Bb5",   // Bishop to b5
		"a6",    // Pawn to a6
		"Bxc6",  // Bishop captures on c6
		"dxc6",  // Pawn captures on c6
		"O-O",   // Kingside castling
	}

	fmt.Println("Playing moves in SAN:")
	for _, move := range sanMoves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
		fmt.Printf("  %s\n", move)
	}

	// Example 2: Move String Formats
	fmt.Println("\n2. Move String Formats")
	game = chess.NewGame()

	// Make a move and show different string representations
	if err := game.PushMove("e4", nil); err != nil {
		log.Fatal(err)
	}

	moves := game.ValidMoves()
	if len(moves) > 0 {
		move := moves[0]
		fmt.Printf("Move string: %s\n", move.String())
	}

	// Example 3: Special Moves Notation
	fmt.Println("\n3. Special Moves Notation")
	game = chess.NewGame()

	specialMoves := []string{
		"e4", "e5",    // Regular pawn moves
		"Nf3", "Nc6",  // Knight moves
		"Bb5", "a6",   // Bishop move and pawn move
		"O-O", "O-O",  // Kingside castling for both sides
	}

	fmt.Println("Special moves demonstration:")
	for _, move := range specialMoves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
		fmt.Printf("  %s\n", move)
	}

	// Example 4: Valid Move Generation
	fmt.Println("\n4. Valid Move Generation")
	game = chess.NewGame()

	// Get all valid moves in the starting position
	validMoves := game.ValidMoves()
	fmt.Printf("Valid moves in starting position: %d\n", len(validMoves))
	for i, move := range validMoves {
		if i < 5 { // Show first 5 moves only
			fmt.Printf("  %s\n", move.String())
		}
	}
}
