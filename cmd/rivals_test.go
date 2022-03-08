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

func TestWhenConfigHasNoRivalsWarningIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)

	reader.EXPECT().IsSet("rivals").Return(false)

	renderer.EXPECT().PrintHeader("Rivals").Times(1)
	renderer.EXPECT().PrintNoRivalsWarning().Times(1)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWhenBootstrapDataRetrievalErrorsItsHandledCorrectly_Rivals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, errors.New(expectedError))

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenCurrentGameweekRetrievalErrorsItsHandledCorrectly_Rivals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(-1, errors.New(expectedError))

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenGameweekLiveScoresRetrievalErrorsItsHandledCorrectly_Rivals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(-1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, errors.New(expectedError))

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenRivalReturnedIsNotIntErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "rivals must all be numeric ids"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	reader.
		EXPECT().
		GetStringSlice("rivals").
		Return([]string{"paddy"})

	renderer.EXPECT().PrintHeader("Rivals").Times(1)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenRivalGameweekPointsRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, errors.New(expectedError))

	reader.
		EXPECT().
		GetStringSlice("rivals").
		Return([]string{"123456"})

	renderer.EXPECT().PrintHeader("Rivals").Times(1)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenRivalManagerDetailsRetrievalErrorsItsHandledCorrectly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)
	expectedError := "Whoops"

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, nil)

	fpl.EXPECT().
		GetManagerDetails(gomock.Any()).
		Return(responses.ManagerDetails{}, errors.New(expectedError))

	reader.
		EXPECT().
		GetStringSlice("rivals").
		Return([]string{"123456"})

	renderer.EXPECT().PrintHeader("Rivals").Times(1)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err.Error() != expectedError {
		t.Fatalf("expected: %v; got %v", expectedError, err)
	}
}

func TestWhenRivalsReturnedSuccessfullyNoErrorReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, nil)

	fpl.EXPECT().
		GetManagerDetails(gomock.Any()).
		Return(responses.ManagerDetails{}, nil)

	reader.
		EXPECT().
		GetStringSlice("rivals").
		Return([]string{"123456"})

	renderer.EXPECT().PrintHeader("Rivals").Times(1)
	renderer.EXPECT().PrintRivalPoints(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWhenMultipleRivalsReturnedSuccessfullyNoErrorReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := helpers.NewMockConfigReader(ctrl)
	fpl := api.NewMockFplAPI(ctrl)
	renderer := ui.NewMockRenderer(ctrl)
	gameweekParser := helpers.NewMockGameweekParser(ctrl)

	reader.EXPECT().IsSet("rivals").Return(true)

	fpl.EXPECT().
		GetBootstrapData().
		Return(responses.BootstrapData{}, nil)

	gameweekParser.EXPECT().
		GetCurrentGameweek(gomock.Any(), gomock.Any()).
		Return(1, nil)

	fpl.EXPECT().
		GetGameweekLiveScores(gomock.Any()).
		Return(responses.GameweekLiveScores{}, nil)

	fpl.EXPECT().
		GetGameweekPoints(gomock.Any(), gomock.Any()).
		Return(responses.GameweekPoints{}, nil).Times(2)

	fpl.EXPECT().
		GetManagerDetails(gomock.Any()).
		Return(responses.ManagerDetails{}, nil).Times(2)

	reader.
		EXPECT().
		GetStringSlice("rivals").
		Return([]string{"123456", "789"})

	renderer.EXPECT().PrintHeader("Rivals").Times(1)
	renderer.EXPECT().PrintRivalPoints(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(2)

	cmd := BuildRivalsCommand(fpl, reader, gameweekParser, renderer)

	err := cmd.RunE(cmd, make([]string, 0))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
