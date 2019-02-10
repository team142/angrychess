package model

import (
	"fmt"
	"testing"
)

func TestTile_EqualSameTile(t *testing.T) {
	tile := Tile{X: 1, Y: 2}
	expected := true
	equal := tile.Equal(&tile)

	if equal != expected {
		t.Error("Tile should be equal to itself")
	}
}

func TestTile_EqualDifferentTile(t *testing.T) {
	tile := Tile{X: 1, Y: 1}
	dTile := Tile{X: 2, Y: 2}
	expected := false
	equal := tile.Equal(&dTile)

	if equal != expected {
		t.Error("Different tiles should not be equal")
	}
}

func TestTile_GetTilesUntil(t *testing.T) {
	start := &Tile{X: 1, Y: 1}
	end := &Tile{X: 8, Y: 8}
	expected := 7

	c := start.GetTilesUntil(end)
	count := 0
	for item := range c {
		fmt.Println(item)
		count++
	}
	if count != expected {
		t.Error(fmt.Sprintf("Got this many %v exptected %v", count, expected))
	}

}
