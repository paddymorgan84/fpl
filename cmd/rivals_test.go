package cmd

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
)

func TestBuildRivalsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	var expectedShort = "Show the points for all of your rivals (specified in config) for a specified gameweek"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "rivals"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}
