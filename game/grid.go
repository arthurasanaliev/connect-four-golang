package game

import "fmt"

type Grid struct {
	grid [6][7]string
}

func NewGrid() *Grid {
	return &Grid{
		grid: [6][7]string{
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
		},
	}
}

func (g *Grid) Display() {
	for _, row := range g.grid {
		fmt.Println(row)
	}
}

func (g *Grid) DropDisc(col int, disc string) bool {
	if col < 0 || col > 6 || g.grid[0][col] != " " {
		return false
	}
	row := 0
	for row < 6 && g.grid[row][col] == " " {
		row++
	}
	row--
	g.grid[row][col] = disc
	return true
}

func (g *Grid) CheckWin() bool {
	di := [8]int{1, 1, 1, 0, 0, -1, -1, -1}
	dj := [8]int{1, 0, -1, 1, -1, 1, 0, -1}

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if g.grid[i][j] == " " {
				continue
			}
			for k := 0; k < 8; k++ {
				I, J := i, j
				cnt := 0
				for cnt < 4 {
					cnt++
					I += di[k]
					J += dj[k]
					if !CheckBoundaries(I, J) || g.grid[I][J] != g.grid[i][j] {
						break
					}
				}
				if cnt == 4 {
					return true
				}
			}
		}
	}
	return false
}

func CheckBoundaries(r, c int) bool {
	return r >= 0 && r < 6 && c >= 0 && c < 7
}

func (g *Grid) IsFull() bool {
	for j := 0; j < 7; j++ {
		if g.grid[0][j] == " " {
			return false
		}
	}
	return true
}
