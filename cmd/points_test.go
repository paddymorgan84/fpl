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

func TestWhenTeamIDRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(-1, errors.New(expectedError))

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenBootstrapDataRetrievalErrorsItsHandledCorrectly_Points(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, errors.New(expectedError))

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenCurrentGameweekRetrievalErrorsItsHandledCorrectly_Points(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(-1, errors.New(expectedError))

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenGameweekPointsRetrievalErrorsItsHandledCorrectly_Points(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, errors.New(expectedError))

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenGameweekLiveScoresRetrievalErrorsItsHandledCorrectly_Points(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, errors.New(expectedError))

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestPointsReturnedSuccessfullyNoErrorReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	renderer.EXPECT().PrintHeader("Gameweek 1 points").Times(1)
	renderer.EXPECT().PrintTeamPoints(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	cmd := BuildPointsCommand(fpl, reader, teamsParser, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
