package main

import (
	"fmt"
	"log"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== Board Serialization Examples ===")

	// Create a game and make some moves
	game := chess.NewGame()
	moves := []string{"e4", "e5", "Nf3", "Nc6"}
	for _, move := range moves {
		if err := game.PushMove(move, nil); err != nil {
			log.Fatal(err)
		}
	}

	// Get the current board
	board := game.Position().Board()
	fmt.Printf("Current Board:\n%v\n", board)

	// Text (FEN) serialization
	fen, err := board.MarshalText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nBoard as FEN: %s\n", string(fen))

	// Binary serialization
	data, err := board.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nBoard as binary (length): %d bytes\n", len(data))

	// Create a new board from FEN
	newBoard := &chess.Board{}
	err = newBoard.UnmarshalText(fen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nBoard recreated from FEN:\n%v\n", newBoard)

	// String representation
	fmt.Printf("\nBoard as string:\n%s\n", board.String())
}
