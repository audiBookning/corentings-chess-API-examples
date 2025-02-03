package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/corentings/chess/v2"
)

func main() {
	fmt.Println("=== PGN Handling Examples ===")

	// Example 1: Reading a game with variations
	fmt.Println("\n1. Reading Game with Variations")
	variationsPath := filepath.Join("..", "..", "fixtures", "pgns", "variations.pgn")
	variationsData, err := os.ReadFile(variationsPath)
	if err != nil {
		log.Printf("Error reading variations PGN: %v\n", err)
	} else {
		pgnGame, err := chess.PGN(strings.NewReader(string(variationsData)))
		if err != nil {
			log.Printf("Error parsing variations PGN: %v\n", err)
		} else {
			game := chess.NewGame(pgnGame)
			fmt.Printf("Game with variations loaded. Moves:\n")
			for _, move := range game.Moves() {
				fmt.Printf("%s ", move.String())
			}
			fmt.Printf("\n\nFinal Position:\n%v\n", game.Position())
		}
	}

	// Example 2: Reading multiple games from a single file
	fmt.Println("\n2. Reading Multiple Games")
	multiPath := filepath.Join("..", "..", "fixtures", "pgns", "multi_game.pgn")
	multiData, err := os.ReadFile(multiPath)
	if err != nil {
		log.Printf("Error reading multi-game PGN: %v\n", err)
	} else {
		games := strings.Split(string(multiData), "\n\n\n")
		fmt.Printf("Found %d games in file\n", len(games))
		
		// Parse first game as example
		if len(games) > 0 {
			pgnGame, err := chess.PGN(strings.NewReader(games[0]))
			if err != nil {
				log.Printf("Error parsing first game: %v\n", err)
			} else {
				game := chess.NewGame(pgnGame)
				fmt.Printf("First game details:\n")
				fmt.Printf("Event: %s\n", game.GetTagPair("Event"))
				fmt.Printf("White: %s\n", game.GetTagPair("White"))
				fmt.Printf("Black: %s\n", game.GetTagPair("Black"))
				fmt.Printf("Result: %s\n", game.GetTagPair("Result"))
			}
		}
	}

	// Example 3: Reading a complete game with rich metadata
	fmt.Println("\n3. Reading Complete Game with Metadata")
	completePath := filepath.Join("..", "..", "fixtures", "pgns", "complete_game.pgn")
	completeData, err := os.ReadFile(completePath)
	if err != nil {
		log.Printf("Error reading complete game PGN: %v\n", err)
	} else {
		pgnGame, err := chess.PGN(strings.NewReader(string(completeData)))
		if err != nil {
			log.Printf("Error parsing complete game: %v\n", err)
		} else {
			game := chess.NewGame(pgnGame)
			fmt.Println("Game Information:")
			fmt.Printf("Event: %s\n", game.GetTagPair("Event"))
			fmt.Printf("White (%s): %s\n", game.GetTagPair("WhiteElo"), game.GetTagPair("White"))
			fmt.Printf("Black (%s): %s\n", game.GetTagPair("BlackElo"), game.GetTagPair("Black"))
			fmt.Printf("Opening: %s\n", game.GetTagPair("Opening"))
			fmt.Printf("Time Control: %s\n", game.GetTagPair("TimeControl"))
			fmt.Printf("Result: %s\n", game.GetTagPair("Result"))
		}
	}

	// Example 4: Creating and exporting PGN
	fmt.Println("\n4. Creating and Exporting PGN")
	game := chess.NewGame()
	
	// Add some moves
	moves := []string{"e4", "e5", "Nf3", "Nc6"}
	for _, move := range moves {
		game.PushMove(move, nil)
	}

	// Add game metadata
	game.AddTagPair("Event", "Example Game")
	game.AddTagPair("Site", "Local Analysis")
	game.AddTagPair("Date", time.Now().Format("2006.01.02"))
	game.AddTagPair("White", "Player 1")
	game.AddTagPair("Black", "Player 2")
	game.AddTagPair("Result", "*")

	// Export to PGN string
	fmt.Printf("\nExported PGN:\n%s\n", game.String())
}
