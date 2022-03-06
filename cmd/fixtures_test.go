package cmd

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/responses"
	"github.com/paddymorgan84/fpl/ui"
)

func TestBuildFixturesCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	var expectedShort = "Get the fixtures for a specific gameweek"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "fixtures"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}

func TestWhenBootstrapDataRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, errors.New(expectedError))

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenCurrentGameweekRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)
	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), reader).
		Return(-1, errors.New(expectedError))

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenGameweekFixtureRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)
	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), reader).
		Return(0, nil)
	fpl.EXPECT().
		GetGameweekFixtures(gomock.Any()).
		Return(responses.GameweekFixtures{}, errors.New(expectedError))

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestFixturesReturnedSuccessfullyNoErrorReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)
	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), reader).
		Return(0, nil)
	fpl.EXPECT().
		GetGameweekFixtures(gomock.Any()).
		Return(responses.GameweekFixtures{}, nil)

	renderer.EXPECT().PrintHeader("Gameweek 0 fixtures").Times(1)
	renderer.EXPECT().PrintGameweekFixtures(gomock.Any(), gomock.Any(), teamsParser, gomock.Any()).Times(1)

	cmd := BuildFixturesCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
