# Chess Package Examples

Just a temp and unmaintained repo.
This directory contains examples demonstrating various features of the chess package and info about the API of https://github.com/CorentinGS/chess.

## Documentation

For more detailed information about the chess package API, please refer to:
- The [API Reference](API.md) with detailed method descriptions, code examples, and feature comparisons
- The package documentation on pkg.go.dev
- The example source code for detailed implementation details


## Example Categories

### Game Fundamentals
- `basic_moves/`: Basic chess moves and rules
  - Making legal moves
  - Capturing pieces
  - Basic game flow

- `game_navigation/`: Navigating through a chess game
  - Moving forward/backward through moves
  - Checking game position
  - Navigating to specific positions

- `game_outcomes/`: Game termination and results
  - Checkmate detection
  - Stalemate scenarios
  - Draw conditions
  - Game resignation

### Board and Position
- `board_manipulation/`: Working with the chess board
  - Board setup and modification
  - Piece placement
  - Board visualization

- `position_visualization/`: Visualizing chess positions
  - Displaying board state
  - FEN string handling
  - Position analysis

- `board_serialization/`: Saving and loading board states
  - FEN string conversion
  - Position serialization
  - State persistence

### Advanced Features
- `variations/`: Managing chess variations
  - Adding alternative lines
  - Managing multiple variations
  - Navigating variation trees

- `pgn_handling/`: PGN format support
  - Reading PGN files
  - Writing PGN notation
  - Game metadata handling

- `opening_book/`: Opening book functionality
  - Opening detection
  - Move suggestions
  - Opening statistics

- `uci_analysis/`: UCI engine integration
  - Engine communication
  - Position analysis
  - Move evaluation

### Chess Components
- `chess_components/`: Core chess components
  - Square handling
  - Position management
  - Move validation

- `comments/`: Game commentary
  - Adding move comments
  - Managing annotations
  - NAG support

- `notation_support/`: Chess notation systems
  - SAN (Standard Algebraic Notation)
  - LAN (Long Algebraic Notation)
  - UCI notation

## Running the Examples

Each example can be run individually from its directory using:

```bash
go run main.go
```

Or you can run all examples at once using the provided scripts:

### Windows
```powershell
.\run_examples.ps1
```

### Linux/macOS
```bash
# First make the script executable
chmod +x run_examples.sh
./run_examples.sh
```

The scripts will:
- Run all examples in sequence
- Create timestamped logs for each example
- Store logs in the `logs` directory
- Show execution status for each example

Make sure you have the chess package installed:

```bash
go get github.com/CorentinGS/chess/v2
```

## Example Output

Each example will print its output to the console, demonstrating specific functionality:

- Game positions and moves
- Board visualization
- Valid moves and captures
- Game outcomes and analysis
- Variations and annotations
- Opening information (where applicable)



## Notes

- All examples use documented API methods and follow best practices
- Examples are designed to be self-contained and easy to understand
- Each example focuses on a specific aspect of the chess package
- Code comments explain key concepts and functionality
