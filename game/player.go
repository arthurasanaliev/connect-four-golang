package game

import "fmt"

type Player struct {
	name string
	disc string
}

func NewPlayer(name, disc string) *Player {
	return &Player{
		name: name,
		disc: disc,
	}
}

func (p *Player) GetMove() int {
	var col int
	fmt.Printf("%s, choose a column you want to drop your disc in: ", p.name)
	fmt.Scanf("%d", col)
	return col
}
