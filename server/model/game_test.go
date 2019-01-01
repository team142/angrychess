package model

import "testing"

func TestFindSpot(t *testing.T) {

	game := Game{}
	expectedFound := true
	found, team, color := game.findSpot()
	if !found {
		t.Errorf("Expected empty game to have found spot %v, got %v", expectedFound, found)
	}

	expectedTeam := 1
	if team != expectedTeam {
		t.Errorf("Expected empty game to have TEAM spot %v, got %v", expectedTeam, team)
	}

	expectedColor := false
	if color != expectedColor {
		t.Errorf("Expected empty game to have COLOR spot %v, got %v", expectedColor, color)
	}

}
