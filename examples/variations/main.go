package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	// Create a new game
	game := chess.NewGame()

	// Play main line
	mainLine := []string{"e4", "e5", "Nf3"}
	for _, move := range mainLine {
		err := game.PushMove(move, nil)
		if err != nil {
			fmt.Printf("Error making move %s: %v\n", move, err)
			return
		}
	}

	// Add a variation that will become the main line
	err := game.PushMove("Nc6", &chess.PushMoveOptions{
		ForceMainline: true,
	})
	if err != nil {
		fmt.Printf("Error adding mainline move: %v\n", err)
		return
	}

	// Print current position
	fmt.Printf("\nCurrent position:\n%v\n", game.Position())

	// Print all moves made (main line)
	fmt.Printf("\nMain line moves:\n")
	for _, move := range game.Moves() {
		fmt.Printf("%s ", move.String())
	}
	fmt.Println()

	// Get all positions in the game
	positions := game.Positions()
	fmt.Printf("\nNumber of positions in game: %d\n", len(positions))

	// Check if we're at the end of the game
	if game.IsAtEnd() {
		fmt.Println("\nAt end of game")
	}

	// Go back one move if possible
	if game.GoBack() {
		fmt.Println("\nWent back one move")
		fmt.Printf("Current position:\n%v\n", game.Position())
	}
}
