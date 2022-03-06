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

func TestBuildHistoryCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	cmd := BuildHistoryCommand(fpl, reader, teamParser, renderer)

	var expectedShort = "Returns history for a managers current and past seasons"
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "history"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}

func TestWhenTeamIdRetrievalErrorsItsHandledCorrectly(t *testing.T) {
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

	cmd := BuildHistoryCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenManagerHistoryRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	expectedError := "Whoops"

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(0, nil)

	fpl.EXPECT().
		GetManagerHistory(0).
		Return(responses.ManagerHistory{}, errors.New(expectedError))

	cmd := BuildHistoryCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestHistoryReturnedSuccessfullyNoErrorReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	teamsParser := helpers.NewMockTeamsParser(ctrl)
	renderer := ui.NewMockRenderer(ctrl)

	teamsParser.EXPECT().
		GetTeamID(gomock.Any(), gomock.Any()).
		Return(0, nil)

	fpl.EXPECT().
		GetManagerHistory(0).
		Return(responses.ManagerHistory{}, nil)

	renderer.EXPECT().PrintHeader("This season").Times(1)
	renderer.EXPECT().PrintSeasonDetails(gomock.Any(), teamsParser).Times(1)
	renderer.EXPECT().PrintHeader("Chips").Times(1)
	renderer.EXPECT().PrintChipDetails(gomock.Any()).Times(1)
	renderer.EXPECT().PrintHeader("Previous Seasons").Times(1)
	renderer.EXPECT().PrintPreviousSeasonDetails(gomock.Any()).Times(1)

	cmd := BuildHistoryCommand(fpl, reader, teamsParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
