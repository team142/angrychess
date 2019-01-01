package model

type Board struct {
	Tiles [][]*Tile
}

type Tile struct {
	X, Y  int
	Color bool
}

func CreateBoard() *Board {
	board := &Board{}
	for x := 1; x <= 8; x++ {
		for y := 1; y <= 8; y++ {
			board.Tiles[x][y] = &Tile{X: x, Y: y, Color: getColor(x, y)}
		}
	}
	return board
}

func getColor(x, y int) bool {
	return (x+y)%2 == 0
}
