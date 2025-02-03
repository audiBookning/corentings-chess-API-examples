package main

import (
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
	"github.com/corentings/chess/v2/opening"
)

func main() {
	fmt.Println("=== Opening Book Examples ===")

	// Example 1: Finding a specific opening
	fmt.Println("\n1. Finding a Specific Opening (French Defense)")
	game := chess.NewGame()
	game.PushMove("e4", nil)
	game.PushMove("e6", nil)

	book := opening.NewBookECO()
	if opening := book.Find(game.Moves()); opening != nil {
		fmt.Printf("Opening Found:\n")
		fmt.Printf("  Name: %s\n", opening.Title())
		fmt.Printf("  Current Position:\n%v\n", game.Position())
	}

	// Example 2: Finding possible variations
	fmt.Println("\n2. Finding Opening Variations (Scandinavian Defense)")
	game = chess.NewGame()
	game.PushMove("e4", nil)
	game.PushMove("d5", nil)

	possibilities := book.Possible(game.Moves())
	fmt.Printf("Found %d possible variations:\n", len(possibilities))
	for _, o := range possibilities {
		fmt.Printf("  - %s\n", o.Title())
	}

	// Example 3: Exploring multiple opening lines
	fmt.Println("\n3. Exploring Multiple Opening Lines (Ruy Lopez Variations)")
	game = chess.NewGame()
	moves := []string{"e4", "e5", "Nf3", "Nc6", "Bb5"}
	for _, move := range moves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
	}

	// Find the opening and its variations
	if opening := book.Find(game.Moves()); opening != nil {
		fmt.Printf("Main Line:\n")
		fmt.Printf("  Name: %s\n", opening.Title())
		
		// Try some common variations
		variations := [][]string{
			{"a6"},            // Morphy Defense
			{"Nf6"},          // Berlin Defense
			{"Bc5"},          // Classical Defense
			{"Nd4"},          // Bird's Defense
		}

		fmt.Printf("\nCommon Variations:\n")
		for _, variation := range variations {
			// Create a new game for each variation
			varGame := chess.NewGame()
			// Play main line
			for _, move := range moves {
				varGame.PushMove(move, nil)
			}
			// Play variation move
			for _, move := range variation {
				if err := varGame.PushMove(move, nil); err != nil {
					continue
				}
			}
			// Find the variation in the opening book
			if varOpening := book.Find(varGame.Moves()); varOpening != nil {
				fmt.Printf("  - %s\n", varOpening.Title())
			}
		}
	}

	// Example 4: Early opening transitions
	fmt.Println("\n4. Opening Transitions")
	game = chess.NewGame()
	transitionMoves := []string{"d4", "Nf6", "c4"} // Start with d4
	for _, move := range transitionMoves {
		if err := game.PushMove(move, nil); err != nil {
			log.Printf("Error making move %s: %v\n", move, err)
			continue
		}
		// Check opening at each move
		if opening := book.Find(game.Moves()); opening != nil {
			fmt.Printf("After %s: %s\n", move, opening.Title())
		}
	}
}

func main2() {
	//game := chess.NewGame()

}
