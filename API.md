# API Documentation: corentings/chess/v2

This document provides a comprehensive overview of the API changes and features in corentings/chess/v2 compared to notnil/chess.

## Table of Contents
1. [Basic Game Operations](#basic-game-operations)
   - [Game Creation](#game-creation)
   - [Move Making](#move-making)
   - [Game Navigation](#game-navigation)
2. [Board Features](#board-features)
   - [Board Manipulation](#board-manipulation)
   - [Board Serialization](#board-serialization)
3. [Game State and Management](#game-state-and-management)
   - [Game State Access](#game-state-access)
   - [Game State Management](#game-state-management)
   - [Tag Pair Management](#tag-pair-management)
4. [Moves and Validation](#moves-and-validation)
   - [Move Types](#move-types)
   - [Move Validation](#move-validation)
   - [Notation Support](#notation-support)
5. [Chess Components](#chess-components)
   - [Piece Handling](#piece-handling)
   - [Square Handling](#square-handling)
   - [Position Handling](#position-handling)
   - [Position Hashing](#position-hashing)
6. [Advanced Features](#advanced-features)
   - [Variation Support](#variation-support)
   - [Comments Support](#comments-support)
7. [Game Outcomes](#game-outcomes)
   - [Outcome Detection](#outcome-detection)
   - [Draw Conditions](#draw-conditions)
8. [Format Support](#format-support)
   - [PGN Support](#pgn-support)
   - [PGN Lexer and Parser](#pgn-lexer-and-parser)
   - [FEN Support](#fen-support)
9. [Extended Features](#extended-features)
   - [UCI Engine Support](#uci-engine-support)
   - [Opening Book Support](#opening-book-support)
   - [Image Generation](#image-generation)

## Basic Game Operations

### Game Creation

```go
// Standard game
game := chess.NewGame()

// From FEN
game := chess.NewGame(chess.FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))

// From PGN
pgn := strings.NewReader("1. e4 e5 2. Nf3 Nc6")
game := chess.NewGame(chess.PGN(pgn))
```

### Move Making

```go
game := chess.NewGame()

// Basic move
err := game.PushMove("e4", &chess.PushMoveOptions{
    ForceMainline: true,  // Add as mainline move
})

// Special moves
game.PushMove("O-O", nil)     // Kingside castle
game.PushMove("O-O-O", nil)   // Queenside castle
game.PushMove("exd6", nil)    // En passant capture
game.PushMove("e8=Q", nil)    // Pawn promotion

// Valid moves
moves := game.ValidMoves()  // Get all legal moves
if len(moves) == 0 {
    // No valid moves - checkmate or stalemate
}
```

### Game Navigation

```go
game := chess.NewGame()

// Basic navigation
game.GoBack()           // Move back one move
game.GoForward()        // Move forward one move
game.IsAtStart()        // Check if at start
game.IsAtEnd()          // Check if at end
game.NavigateToMainLine() // Return to main line
```

## Board Features

### Board Manipulation

```go
game := chess.NewGame()
board := game.Position().Board()

// Rotation and Flipping
board = board.Rotate()            // Rotate 90 degrees clockwise
board = board.Flip(chess.UpDown)  // Flip vertically
board = board.Flip(chess.LeftRight) // Flip horizontally
board = board.Transpose()         // Transpose board
```

### Board Serialization

```go
// Text (FEN)
fen, err := board.MarshalText()    // Convert to FEN string
err = board.UnmarshalText(fen)     // Load from FEN string

// Binary
data, err := board.MarshalBinary()    // Convert to binary
err = board.UnmarshalBinary(data)     // Load from binary

// String representation
str := board.String()  // Get FEN string representation
```

## Game State and Management

### Game State Access

```go
game := chess.NewGame()
pos := game.Position()         // Get current position
curPos := game.CurrentPosition() // Get position at current move
mainLine := game.Moves()        // Get main line moves
variations := game.Variations(move) // Get alternative moves

// Position analysis
validMoves := pos.ValidMoves()  // Get valid moves from position
inCheck := pos.InCheck()        // Check if side to move is in check
hash := pos.Hash()              // Get unique hash of position
```

### Game State Management

```go
game := chess.NewGame()

// Get all positions
positions := game.Positions()  // Get all positions in game history

// State copying
newGame := game.Copy()       // Create a deep copy of the game
cloned := game.Clone()      // Create a shallow copy
// Both copy methods preserve:
// - All moves and variations
// - Current position
// - Game outcome and method
// - Comments
// - Tag pairs

// Position comparison
hash := pos.Hash()          // Get unique hash of position
```

### Tag Pair Management

```go
game := chess.NewGame()
game.AddTagPair("Event", "World Championship")  // Add/update tag
event := game.GetTagPair("Event")              // Get tag value
removed := game.RemoveTagPair("Event")         // Remove tag

// Tag pairs are preserved when copying game state
clone := game.Clone()
clone.GetTagPair("Event")  // Returns "World Championship"
```

## Moves and Validation

### Move Types

```go
// Move Types
move := &chess.Move{
    // Basic moves
    s1: chess.E2, s2: chess.E4,  // Regular move
    
    // Special moves
    tags: chess.KingSideCastle,   // Kingside castling
    tags: chess.QueenSideCastle,  // Queenside castling
    tags: chess.EnPassant,        // En passant capture
    tags: chess.Capture,          // Piece capture
    tags: chess.Check,            // Move gives check
    
    // Pawn promotion
    promo: chess.Queen,  // Promote to queen
}
```

### Move Validation

```go
game := chess.NewGame()
moves := game.ValidMoves()  // Get all valid moves

// The package validates:
// - Piece movement rules
// - Turn order
// - Check and checkmate
// - Castling rights
// - En passant opportunities
// - Pawn moves (including promotion)
// - Pin and discovered check situations

// Position Updates
// Moves automatically update:
// - Board position
// - Castling rights
// - En passant squares
// - Move counters
// - Game status

// Move validation includes:
// - Legal piece movements for each piece type
// - Detecting and preventing illegal moves
// - Handling special moves (castling, en passant, promotion)
// - Checking for check and checkmate conditions
// - Validating move notation (UCI, SAN, LAN)
// - Ensuring proper turn order
```

### Notation Support

```go
// 1. UCI Notation (Universal Chess Interface)
notation := chess.UCINotation{}

// Encode move to UCI
move := &chess.Move{s1: chess.E2, s2: chess.E4}
uci := notation.Encode(pos, move)    // Returns: "e2e4"

// Decode UCI move
move, err := notation.Decode(pos, "e2e4")

// UCI with promotion
move = &chess.Move{s1: chess.A7, s2: chess.A8, promo: chess.Queen}
uci = notation.Encode(pos, move)     // Returns: "a7a8q"

// 2. Algebraic Notation (SAN)
notation := chess.AlgebraicNotation{}

// Encode move to SAN
san := notation.Encode(pos, move)    // Returns: "e4"

// Special moves in SAN
"O-O"    // Kingside castle
"O-O-O"  // Queenside castle
"exd5"   // Pawn capture
"Nxe4"   // Piece capture
"e8=Q"   // Pawn promotion
"Nf3+"   // Check
"e4#"    // Checkmate

// 3. Long Algebraic Notation (LAN)
notation := chess.LongAlgebraicNotation{}

// Encode move to LAN
lan := notation.Encode(pos, move)    // Returns: "e2-e4"
```

## Chess Components

### Piece Handling

```go
// 1. Piece Types
piece := chess.King    // King piece
piece := chess.Queen   // Queen piece
piece := chess.Rook    // Rook piece
piece := chess.Bishop  // Bishop piece
piece := chess.Knight  // Knight piece
piece := chess.Pawn    // Pawn piece

// 2. Piece Colors
color := chess.White  // White piece
color := chess.Black  // Black piece

// 3. Piece Operations
str := piece.String()   // Get string representation ("k", "q", "r", etc.)
bytes := piece.Bytes()  // Get byte representation

// 4. Piece Placement
board.GetPiece(square)  // Get piece at square
board.SetPiece(square, piece)  // Place piece on square
board.RemovePiece(square)  // Remove piece from square

// 5. Piece Movement Rules
// The package handles all legal piece movements:
// - Kings: One square in any direction, castling
// - Queens: Any number of squares diagonally, horizontally, or vertically
// - Rooks: Any number of squares horizontally or vertically
// - Bishops: Any number of squares diagonally
// - Knights: L-shaped movement (2 squares in one direction, 1 square perpendicular)
// - Pawns: Forward movement, diagonal captures, en passant, promotion

// 6. Movement Validation
// Each piece type has specific rules:
// - Kings: Cannot move into check, special castling rules
// - Queens: Must have clear path to destination
// - Rooks: Must have clear path, affects castling rights
// - Bishops: Must have clear diagonal path
// - Knights: Can jump over other pieces
// - Pawns: Complex rules for first move, captures, promotion, en passant
```

### Square Handling

```go
// 1. Square Constants
square := chess.A1  // Bottom left square
square := chess.H8  // Top right square

// 2. Create Square from File and Rank
file := chess.FileA  // Files A through H
rank := chess.Rank1  // Ranks 1 through 8
square := chess.NewSquare(file, rank)

// 3. Square Information
file := square.File()  // Get file (A-H)
rank := square.Rank()  // Get rank (1-8)
str := square.String() // Get algebraic notation

// 4. Square Lists
squares := []Square{
    chess.A1, chess.B1, chess.C1,  // First rank
    chess.A2, chess.B2, chess.C2,  // Second rank
}
```

### Position Handling

```go
// 1. Create Position
pos := chess.StartingPosition()  // Standard starting position
pos := chess.NewPosition()       // Empty position

// 2. Position Information
turn := pos.Turn()              // Get current turn
moves := pos.ValidMoves()       // Get valid moves
clock := pos.HalfMoveClock()    // Get halfmove clock
fullMove := pos.FullMoveNumber() // Get fullmove number

// 3. Position Analysis
isCheck := pos.InCheck()        // Check if in check
isMate := pos.IsCheckmate()     // Check if checkmate
isDraw := pos.IsDraw()          // Check if draw
isStalemate := pos.IsStalemate() // Check if stalemate

// 4. Position Updates
newPos := pos.Update(move)      // Make a move
// Updates include:
// - Piece positions
// - Turn
// - Castling rights
// - En passant square
// - Move clocks

// 5. Position Serialization
data, err := pos.MarshalBinary()   // To binary
fen := pos.String()                // To FEN
```

### Position Hashing

```go
// 1. Create Hasher
hasher := chess.NewZobristHasher()

// 2. Hash Positions
fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
hash, err := hasher.HashPosition(fen)  // Returns "463b96181691fc9c"

// 3. Hash Components
// The hasher considers:
// - Piece positions
// - Side to move
// - Castling rights
// - En passant squares

// 4. Known Position Hashes
// Starting position: "463b96181691fc9c"
// After 1.e4: "823c9b50fd114196"
// After 1.e4 e5: "0844931a6ef4b9a0"
```

## Advanced Features

### Variation Support

#### API Reference
```go
// Core Variation API
game := chess.NewGame()

// Managing variations
parentMove := game.GetRootMove()                // Get root move
game.AddVariation(parentMove, newMove)          // Add variation to root
game.AddVariation(existingMove, alternateMove)  // Add variation to any move

// Variation navigation
variations := game.Variations(move)  // Get all alternative moves
game.NavigateToMainLine()           // Return to main line

// Move tree access
rootMove := game.GetRootMove()  // Get root of move tree
mainLine := game.Moves()        // Get main line moves

// Move information
move.Parent()   // Parent move
move.Children() // Child moves (variations)
move.Position() // Position after move
move.Comments() // Move comments
```

#### Practical Examples
```go
// Creating and Managing Variations
game := chess.NewGame()

// 1. Playing the main line
game.PushMove("e4", nil)      // 1. e4
game.PushMove("e5", nil)      // 1...e5
game.PushMove("Nf3", nil)     // 2. Nf3

// 2. Adding a variation to Black's first move
game.GoBack()                 // Go back to after 1. e4
game.PushMove("c5", nil)      // Add Sicilian Defense as variation

// 3. Promoting a variation to main line
game.GoBack()                 // Go back to after 1. e4
game.PushMove("c5", &chess.PushMoveOptions{
    ForceMainline: true,      // This makes c5 the main line
})                           // e5 becomes a variation

// 4. Working with multiple variations
game.PushMove("Nf3", nil)     // 2. Nf3 after 1...c5
game.PushMove("d6", nil)      // 2...d6

// Add another variation for Black's second move
game.GoBack()                 // Go back to after 2. Nf3
game.PushMove("Nc6", nil)     // 2...Nc6 as variation

// 5. Using the Move Tree API
rootMove := game.GetRootMove()
firstVariation := &chess.Move{
    // ... move details
}
game.AddVariation(rootMove, firstVariation)  // Add variation to root

// Get variations at current move
currentMove := game.CurrentMove()
variations := game.Variations(currentMove)

// Access move information
for _, variation := range variations {
    parent := variation.Parent()            // Parent move
    children := variation.Children()       // Child moves/variations
    position := variation.Position()       // Position after move
    comments := variation.Comments()       // Move comments
}

// 6. Practical Example: Building an Opening Repertoire
game := chess.NewGame()

// Main line: 1. e4 e5
game.PushMove("e4", nil)
game.PushMove("e5", nil)

// Add Sicilian Defense variation
game.GoBack()                 // Back to after 1. e4
game.PushMove("c5", nil)      // Add 1...c5 as variation

// Add French Defense variation using AddVariation
game.GoBack()                 // Back to after 1. e4
frenchDefense := &chess.Move{
    // ... move details for 1...e6
}
currentMove := game.CurrentMove()
game.AddVariation(currentMove, frenchDefense)

// Later decide to make Sicilian your main line
game.GoBack()                 // Back to after 1. e4
game.PushMove("c5", &chess.PushMoveOptions{
    ForceMainline: true,      // Promote Sicilian to main line
})

// Navigate and explore variations
game.NavigateToMainLine()      // Return to main line
mainLine := game.Moves()       // Get all main line moves
```

This section demonstrates:
- Complete API reference for variation handling
- Creating and adding variations using both PushMove and AddVariation
- Promoting variations to the main line
- Navigating between variations
- Accessing and manipulating the move tree
- Working with move properties (parent, children, position, comments)
- Building an opening repertoire with variations
- Working with multiple variations at the same position
- Common variation management patterns

### Comments Support

```go
game := chess.NewGame()

// Managing comments
comments := game.Comments()         // Get all comments
move.comments = "Good move!"        // Add comment to move

// Comments in variations
game.AddVariation(parentMove, newMove)
newMove.comments = "Interesting alternative"
```

## Game Outcomes

### Outcome Detection

```go
game := chess.NewGame()

// Game status
outcome := game.Outcome()  // Get result (WhiteWon, BlackWon, Draw, NoOutcome)
method := game.Method()    // Get end method (Checkmate, Stalemate, etc.)

// Ending the game
game.Resign(chess.White)   // White resigns
game.Resign(chess.Black)   // Black resigns

// Automatic detection
if game.Method() == Checkmate {
    // Checkmate
} else if game.Method() == Stalemate {
    // Stalemate
} else if game.Method() == InsufficientMaterial {
    // Insufficient material
} else if game.Method() == Resignation {
    // Resignation
}
```

### Draw Conditions

```go
game := chess.NewGame()

// Check available draw methods
draws := game.EligibleDraws()  // Get list of valid draw types
// Possible draws:
// - DrawOffer (always available)
// - ThreefoldRepetition (if position repeated 3 times)
// - FiftyMoveRule (if 50 moves without pawn move/capture)

// Draw types
game.Draw(ThreefoldRepetition)  // Draw by threefold repetition
game.Draw(FiftyMoveRule)        // Draw by fifty move rule
game.Draw(DrawOffer)            // Draw by mutual agreement

// Automatic draws
if game.Method() == FivefoldRepetition {
    // Automatic draw after five repetitions
}
if game.Method() == SeventyFiveMoveRule {
    // Automatic draw after 75 moves without pawn move/capture
}
if game.Method() == InsufficientMaterial {
    // Automatic draw by insufficient material
    // Examples of insufficient material:
    // - King vs King
    // - King + Knight vs King
    // - King + Bishop vs King
    // - King + Bishop vs King + Bishop (same color)
}
```

## Format Support

### PGN Support

```go
// Create game from PGN
pgn := `[Event "Example"]
[Site "Internet"]
[Date "2023.12.06"]
[Round "1"]
[White "Player1"]
[Black "Player2"]
[Result "1-0"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.} 4. Ba4 Nf6 1-0`

reader := strings.NewReader(pgn)
game := chess.NewGame(chess.PGN(reader))

// Read PGN components
event := game.GetTagPair("Event")    // "Example"
site := game.GetTagPair("Site")      // "Internet"

// Access moves and variations
moves := game.Moves()            // Get main line moves
variations := game.Variations()  // Get variations

// Parse large PGN files
scanner := chess.NewScanner(reader)
for {
    game, err := scanner.ScanGame()
    if err == io.EOF {
        break
    }
    // Process each game
}

// Advanced PGN features
// - Support for recursive variations
// - Comment preservation
// - NAG support
// - Move numbers
// - Game termination markers
// - Special symbols (!, ?, !!, ??, !?, ?!)
```

### PGN Lexer and Parser

```go
// 1. Lexer Features
lexer := chess.NewLexer(pgnText)

// Supports all PGN tokens:
// - Tag pairs: [Event "Example"]
// - Move numbers: 1.
// - Pieces: N (Knight), B (Bishop)
// - Squares: e4, f6
// - Special moves: O-O (castling)
// - Captures: exd5, Nxe4
// - Check/Mate: e4+, e4#
// - Comments: {This is a comment}
// - Variations: (1. e4 e5)
// - NAGs: $1, $2
// - Results: 1-0, 0-1, 1/2-1/2

// 2. Move Disambiguation
Nbd7  // Knight from b-file to d7
N4d7  // Knight from rank 4 to d7
Nbd7+ // Knight from b-file to d7, giving check

// 3. Special Notations
O-O    // Kingside castling
O-O-O  // Queenside castling
e.p.   // En passant capture
+      // Check
#      // Checkmate
!      // Good move
?      // Poor move
!!     // Brilliant move
??     // Blunder
!?     // Interesting move
?!     // Dubious move

// 4. Comments and Annotations
{Positional advantage}  // Regular comment
$1                     // NAG: Good move
$2                     // NAG: Poor move
$3                     // NAG: Very good move

// 5. Variations
1. e4 e5 (1... c5 2. Nf3) 2. Nf3  // Alternative lines
```

### FEN Support

```go
// Create game from FEN
fenStr := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
game := chess.NewGame(chess.FEN(fenStr))

// Get FEN from position
pos := game.Position()
fen := pos.String()  // Returns FEN string

// Common FEN positions
// Starting position
"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// After 1. e4
"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"

// FEN components
// A FEN string contains 6 fields:
// 1. Piece placement
// 2. Active color (w/b)
// 3. Castling availability (KQkq)
// 4. En passant target square
// 5. Halfmove clock
// 6. Fullmove number
```

## Extended Features

### UCI Engine Support

```go
import "github.com/corentings/chess/v2/uci"

// Initialize engine
enginePath := "path/to/stockfish"
engine, err := uci.New(enginePath)
if err != nil {
    log.Fatal(err)
}
defer engine.Close()

// Engine options
logger := log.New(os.Stdout, "", 0)
engine, err := uci.New(enginePath, 
    uci.Debug,                // Enable debug mode
    uci.Logger(logger),       // Set logger
)

// Set engine options
setOption := uci.CmdSetOption{
    Name: "UCI_Elo",         // Option name
    Value: "1500",           // Option value
}

// Game commands
newGame := uci.CmdUCINewGame
position := uci.CmdPosition{
    Position: chess.StartingPosition(),
}

// Configure analysis
analysis := uci.CmdGo{
    MoveTime: time.Second,    // Time to think
    // Depth: 20,            // Search depth
    // Nodes: 1000000,       // Nodes to search
    // Infinite: true,       // Search until stopped
}

// Run commands
err = engine.Run(
    uci.CmdUCI,              // Initialize UCI mode
    uci.CmdIsReady,          // Check if engine is ready
    setOption,               // Set engine options
    newGame,                 // Start new game
    position,                // Set position
    analysis,                // Start analysis
)
```

### Opening Book Support

```go
import "github.com/corentings/chess/v2/opening"

// Create opening book
book := opening.NewBookECO()  // Create ECO opening book

// Find specific opening
game := chess.NewGame()
game.PushMove("e4", nil)
game.PushMove("e6", nil)

opening := book.Find(game.Moves())  // Find exact opening
fmt.Println(opening.Title())        // Prints: "French Defense"

// Find possible openings
game = chess.NewGame()
game.PushMove("e4", nil)
game.PushMove("d5", nil)

// Get all possible variations
possibilities := book.Possible(game.Moves())
for _, o := range possibilities {
    fmt.Println(o.Title())  // Opening name
}
```

### Image Generation

```go
import (
    "github.com/corentings/chess/v2"
    "github.com/corentings/chess/v2/image"
    "image/color"
)

// Create SVG from position
game := chess.NewGame()
pos := game.Position()
var buf bytes.Buffer

// Generate basic SVG
err := image.SVG(&buf, pos.Board())

// Customize visualization
mark := image.MarkSquares(
    color.RGBA{255, 255, 0, 1},  // Yellow highlighting
    chess.D2, chess.D4,          // Squares to highlight
)

// Set perspective
perspective := image.Perspective(chess.Black)  // View from black's side

// Generate customized SVG
err = image.SVG(&buf, pos.Board(), 
    mark,         // Square highlighting
    perspective,  // Board perspective
)

// Save or use SVG
// Write to file
f, _ := os.Create("position.svg")
defer f.Close()
io.Copy(f, bytes.NewBufferString(buf.String()))

// Or use SVG string directly
svgString := buf.String()

// Use cases:
// - Generate position diagrams
// - Create game visualizations
// - Highlight moves and squares
// - Show board from different perspectives

// The image generation support allows you to:
// 1. Generate SVG representations of chess positions
// 2. Highlight specific squares with custom colors
// 3. View the board from either white's or black's perspective
// 4. Create visual chess diagrams for documentation or analysis

// Note: The image package generates clean, scalable SVG images 
// suitable for web or print use.
