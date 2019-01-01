package model

import "testing"

func TestPlayer_SetTeamAndColorNormal(t *testing.T) {
	p := Player{}
	p.SetTeamAndColor(1, 2)

	expectedColor := false
	if p.Color != expectedColor {
		t.Errorf("Expected color %v and got %v", expectedColor, p.Color)
	}

	expectedTeam := 1
	if p.Team != expectedTeam {
		t.Errorf("Expected team %v and got %v", expectedTeam, p.Team)
	}

}

func TestPlayer_SetTeamAndColorAbsurd(t *testing.T) {
	p := Player{}
	p.SetTeamAndColor(7, 4)

	expectedColor := false
	if p.Color != expectedColor {
		t.Errorf("Expected color %v and got %v", expectedColor, p.Color)
	}

	expectedTeam := 2
	if p.Team != expectedTeam {
		t.Errorf("Expected team %v and got %v", expectedTeam, p.Team)
	}

}
