package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Chess Components Examples ===")

	// Example 1: Square Handling
	fmt.Println("\n1. Square Handling")
	
	// Create squares using constants
	a1 := chess.A1
	h8 := chess.H8
	fmt.Printf("Corner squares: %s and %s\n", a1, h8)

	// Create square from file and rank
	file := chess.FileE
	rank := chess.Rank4
	e4 := chess.NewSquare(file, rank)
	fmt.Printf("Created square: %s (File: %v, Rank: %v)\n", e4, e4.File(), e4.Rank())

	// Example 2: Position Handling
	fmt.Println("\n2. Position Handling")
	
	// Create positions
	pos := chess.StartingPosition()
	fmt.Printf("Starting position:\n%v\n", pos)

	// Position information
	fmt.Printf("Turn: %v\n", pos.Turn())
	fmt.Printf("Valid moves: %d\n", len(pos.ValidMoves()))
	fmt.Printf("Halfmove clock: %d\n", pos.HalfMoveClock())

	// Create a game to demonstrate position updates
	game := chess.NewGame()
	
	// Make some moves
	moves := []string{"e4", "e5", "Nf3"}
	for _, move := range moves {
		err := game.PushMove(move, nil)
		if err != nil {
			fmt.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	// Get current position
	currentPos := game.Position()
	fmt.Printf("\nPosition after moves:\n%v\n", currentPos)

	// Example 3: Position Hashing
	fmt.Println("\n3. Position Hashing")
	
	// Create hasher
	hasher := chess.NewZobristHasher()

	// Hash starting position
	startFEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	startHash, err := hasher.HashPosition(startFEN)
	if err == nil {
		fmt.Printf("Starting position hash: %s\n", startHash)
	}

	// Hash after 1.e4
	afterE4FEN := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	e4Hash, err := hasher.HashPosition(afterE4FEN)
	if err == nil {
		fmt.Printf("Position after 1.e4 hash: %s\n", e4Hash)
	}
}
