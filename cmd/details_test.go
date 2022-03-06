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

func TestBuildDetailsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildDetailsCommand(fpl, reader, teamsParser, renderer)

	var expectedShort = "Returns details of manager for current season, e.g. league standings, cash in the bank, overall points etc"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "details"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}

func TestWhenTeamsParserReturnsErrorItsHandledAsExpected(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"
	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(-1, errors.New(expectedError))

	cmd := BuildDetailsCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenFetchingManagerDetailsReturnsErrorItsHandledAsExpected(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetManagerDetails(gomock.Any()).
		Return(responses.ManagerDetails{}, errors.New(expectedError))

	cmd := BuildDetailsCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenTeamIdAndManagerDetailsSuccessfulNoErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetManagerDetails(gomock.Any()).
		Return(responses.ManagerDetails{}, nil)

	renderer.EXPECT().PrintHeader("Manager Details").Times(1)
	renderer.EXPECT().PrintManagerDetails(gomock.Any()).Times(1)
	renderer.EXPECT().PrintHeader("Classic Leagues").Times(1)
	renderer.EXPECT().PrintClassicLeagues(gomock.Any(), teamsParser).Times(1)
	renderer.EXPECT().PrintHeader("Global Leagues").Times(1)
	renderer.EXPECT().PrintGlobalLeagues(gomock.Any(), teamsParser).Times(1)
	renderer.EXPECT().PrintHeader("Transfers & Finance").Times(1)
	renderer.EXPECT().PrintTransfersAndFinance(gomock.Any(), teamsParser).Times(1)

	cmd := BuildDetailsCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
