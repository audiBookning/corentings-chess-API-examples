package main

import (
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Game Outcomes Examples ===")

	// Example 1: Checkmate Detection
	fmt.Println("\n1. Checkmate Detection (Fool's Mate)")
	game := chess.NewGame()
	foolsMate := []string{"f3", "e5", "g4", "Qh4"}
	for _, move := range foolsMate {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	fmt.Printf("Position after Fool's Mate:\n%v\n", game.Position())
	fmt.Printf("Game Outcome: %v\n", game.Outcome())
	fmt.Printf("Game Method: %v\n", game.Method())

	// Example 2: Stalemate Detection
	fmt.Println("\n2. Stalemate Detection")
	game = chess.NewGame()
	stalemateMoves := []string{
		"e3", "a5", "Qh5", "Ra6", "Qxa5", "h5", "h4", "Rah6",
		"Qxc7", "f6", "Qxd7+", "Kf7", "Qxb7", "Qd3", "Qxb8", "Qh7",
		"Qxc8", "Kg6", "Qe6",
	}
	for _, move := range stalemateMoves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	fmt.Printf("Position after stalemate sequence:\n%v\n", game.Position())
	fmt.Printf("Game Outcome: %v\n", game.Outcome())
	fmt.Printf("Game Method: %v\n", game.Method())

	// Example 3: Draw Conditions
	fmt.Println("\n3. Draw Conditions")
	game = chess.NewGame()

	// Make some moves to demonstrate threefold repetition
	moves := []string{"Nf3", "Nf6", "Ng1", "Ng8", "Nf3", "Nf6"}
	for _, move := range moves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	fmt.Printf("\nGame state after repetition:\n")
	fmt.Printf("Game Outcome: %v\n", game.Outcome())
	fmt.Printf("Game Method: %v\n", game.Method())

	// Example 4: Game Termination Methods
	fmt.Println("\n4. Game Termination Methods")
	game = chess.NewGame()

	// Make some moves
	moves = []string{"e4", "e5", "Nf3", "Nc6"}
	for _, move := range moves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	// White resigns
	game.Resign(chess.White)

	fmt.Printf("Current game state:\n")
	fmt.Printf("Outcome: %v\n", game.Outcome())
	fmt.Printf("Method: %v\n", game.Method())
	fmt.Printf("Is game over: %v\n", game.Outcome() != chess.NoOutcome)
}
