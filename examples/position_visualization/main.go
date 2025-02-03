package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/corentings/chess/v2"
	"github.com/corentings/chess/v2/image"
)

func main() {
	fmt.Println("=== Position Visualization Example ===")
	game := chess.NewGame()

	// Make some moves to get an interesting position
	moves := []string{"e4", "e5", "Nf3", "Nc6", "Bb5"}
	for _, move := range moves {
		game.PushMove(move, nil)
	}

	// Create SVG with marked squares
	var buf bytes.Buffer
	err := image.SVG(&buf, game.Position().Board(), 
		image.MarkSquares(&color.RGBA{255, 255, 0, 255}, chess.E4, chess.E5), // Yellow for central pawns
	)
	if err != nil {
		log.Fatal(err)
	}

	// Save SVG to file
	svgPath := filepath.Join(".", "ruy_lopez.svg")
	f, err := os.Create(svgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(buf.String())
	fmt.Printf("Position visualization saved to: %s\n", svgPath)
	fmt.Printf("\nCurrent position:\n%v\n", game.Position())
}
