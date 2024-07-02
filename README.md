# Connect Four in Go

This is a simple terminal-based Connect Four game implemented in the Go programming language. The game allows two players to take turns dropping discs into a grid, with the objective of forming a line of four discs in a row either horizontally, vertically, or diagonally.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Game Rules](#game-rules)
- [Code Structure](#code-structure)
- [Contributing](#contributing)
- [License](#license)

## Features
- Two-player gameplay
- Terminal-based interface
- Detects wins and draws
- Simple and clear code structure

## Installation
1. Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).
2. Clone this repository:
    ```sh
    git clone https://github.com/yourusername/connect-four-golang.git
    ```
3. Navigate to the project directory:
    ```sh
    cd connect-four-golang
    ```
4. Build the project:
    ```sh
    go build
    ```

## Usage
1. Run the executable:
    ```sh
    ./connect-four-golang
    ```
2. Follow the on-screen instructions to play the game.

## Game Rules
1. Players take turns to drop one disc into any of the seven columns.
2. The disc falls to the lowest available space within the column.
3. The first player to form a horizontal, vertical, or diagonal line of four discs wins.
4. If the grid is completely filled without any player forming a line of four discs, the game ends in a draw.

## Code Structure
- `main.go`: Entry point of the application.
- `game/`: Contains the game logic.
  - `game.go`: Manages the game flow.
  - `grid.go`: Manages the game grid.
  - `player.go`: Manages player actions.