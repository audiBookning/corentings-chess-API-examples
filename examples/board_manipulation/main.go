package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Board Manipulation Example ===")
	game := chess.NewGame()

	// Make some moves to get an interesting position
	moves := []string{"e4", "e5", "Nf3", "Nc6"}
	for _, move := range moves {
		game.PushMove(move, nil)
	}

	// Get the current board
	board := game.Position().Board()
	fmt.Printf("Original Board:\n%v\n", board)

	// Rotate board
	rotatedBoard := board.Rotate()
	fmt.Printf("\nRotated Board (90 degrees clockwise):\n%v\n", rotatedBoard)

	// Flip board vertically
	flippedBoardUD := board.Flip(chess.UpDown)
	fmt.Printf("\nFlipped Board (Up-Down):\n%v\n", flippedBoardUD)

	// Flip board horizontally
	flippedBoardLR := board.Flip(chess.LeftRight)
	fmt.Printf("\nFlipped Board (Left-Right):\n%v\n", flippedBoardLR)

	// Transpose board
	transposedBoard := board.Transpose()
	fmt.Printf("\nTransposed Board:\n%v\n", transposedBoard)
}
