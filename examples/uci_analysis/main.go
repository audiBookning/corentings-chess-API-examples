package main

import (
	"fmt"
	"log"
	"time"

	"github.com/corentings/chess/v2"
	"github.com/corentings/chess/v2/uci"
)

func main() {
	fmt.Println("=== UCI Engine Analysis Example ===")
	
	// Create a new game
	game := chess.NewGame()

	// Make some moves
	moves := []string{"e4", "e5", "Nf3", "Nc6", "Bb5"}
	for _, move := range moves {
		game.PushMove(move, nil)
	}

	fmt.Printf("Position to analyze:\n%v\n", game.Position())

	// Initialize UCI engine (assumes Stockfish is in PATH)
	enginePath := "stockfish"
	engine, err := uci.New(enginePath)
	if err != nil {
		fmt.Printf("Note: Stockfish engine not available (%v). Skipping analysis.\n", err)
		return
	}
	defer engine.Close()

	// Configure analysis
	analysis := uci.CmdGo{
		MoveTime: time.Second, // Analyze for 1 second
	}

	// Run analysis
	err = engine.Run(
		uci.CmdUCI,
		uci.CmdIsReady,
		uci.CmdPosition{Position: game.Position()},
		analysis,
	)
	if err != nil {
		log.Printf("Engine analysis error: %v\n", err)
		return
	}
}
