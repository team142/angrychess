package model

import "testing"

func TestFindSpotEmptyGame(t *testing.T) {
	game := Game{Boards: 2}
	expectedFound := true
	found, spot := game.FindSpot()
	if !found {
		t.Errorf("Expected empty game to have found spot %v, got %v", expectedFound, found)
	}
	expectedSpot := 1
	if spot != expectedSpot {
		t.Errorf("Expected empty game to have TEAM spot %v, got %v", expectedSpot, spot)
	}

}

func TestFindSpotGameWithPlayers(t *testing.T) {
	game := Game{Boards: 2, Players: make(map[int]*Player)}
	expectedFound := true
	for i := 1; i <= 4; i++ {
		found, spot := game.FindSpot()
		if found != expectedFound {
			t.Errorf("Expected game to have found spot %v, got %v", expectedFound, found)
		}
		expectedSpot := i
		if spot != expectedSpot {
			t.Errorf("Expected game to have TEAM spot %v, got %v", expectedSpot, spot)
		}
		game.JoinGame(&Player{})
	}
}

func TestFindSpotGameFullPlayers(t *testing.T) {
	game := Game{Boards: 2, Players: make(map[int]*Player)}
	for i := 1; i <= 4; i++ {
		game.JoinGame(&Player{})
	}

	found, _ := game.findSpot()
	expectedFound := false
	if found != expectedFound {
		t.Errorf("Expected game to have found spot %v, got %v", expectedFound, found)
	}

}
