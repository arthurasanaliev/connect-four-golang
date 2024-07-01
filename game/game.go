package game

import "fmt"

type Game struct {
	grid          *Grid
	player1       *Player
	player2       *Player
	currentPlayer *Player
}

func NewGame() *Game {
	grid := NewGrid()
	player1 := NewPlayer("Player 1", "X")
	player2 := NewPlayer("Player 2", "O")

	return &Game{
		grid:          grid,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (g *Game) Start() {
	for {
		g.grid.Display()
		col := g.currentPlayer.GetMove()
		if !g.grid.DropDisc(col, g.currentPlayer.disc) {
			fmt.Println("Invalid move, try again.")
			continue
		}
		if g.grid.CheckWin() { // can be optimized
			g.grid.Display()
			fmt.Printf("%s wins!\n", g.currentPlayer.name)
			return
		}
		if g.grid.IsFull() {
			g.grid.Display()
			fmt.Println("Draw!")
			return
		}
		g.switchPlayer()
	}
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}
