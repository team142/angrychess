package model

func GetTileColor(x, y int) bool {
	return (x+y)%2 == 0
}
