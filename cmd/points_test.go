package cmd

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
)

func TestBuildPointsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	var expectedShort = "Get the points for a specified gameweek (defaults to latest active gameweek)"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "points"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}
