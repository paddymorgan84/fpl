package cmd

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
)

func TestBuildFixturesCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, renderer)

	var expectedShort = "Get the fixtures for a specific gameweek"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "fixtures"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}
