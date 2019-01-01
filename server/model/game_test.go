package model

import "testing"

func TestFindSpotEmptyGame(t *testing.T) {

	game := Game{Boards: MaxSupportedBoards}
	expectedFound := true
	found, spot := game.findSpot()
	if !found {
		t.Errorf("Expected empty game to have found spot %v, got %v", expectedFound, found)
	} else {
		t.Logf("Empty game returns spot found correctly")
	}

	expectedSpot := 1
	if spot != expectedSpot {
		t.Errorf("Expected empty game to have TEAM spot %v, got %v", expectedSpot, spot)
	} else {
		t.Logf("Empty game returns spot 1 correctly")
	}

}
