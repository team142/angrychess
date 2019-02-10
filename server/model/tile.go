package model

import "github.com/team142/angrychess/util"

type Tile struct {
	X, Y int
}

func (t *Tile) Equal(other *Tile) bool {
	return t.X == other.X && t.Y == other.Y
}

func (t *Tile) GetTilesUntil(end *Tile) chan Tile {
	c := make(chan Tile)
	go func() {
		xd := util.GetDirection(t.X, end.X)
		yd := util.GetDirection(t.Y, end.Y)
		current := Tile{X: t.X, Y: t.Y}
		for {
			current.X += xd
			current.Y += yd
			c <- current
			if current.Equal(end) {
				close(c)
				return
			}
		}
	}()
	return c
}
