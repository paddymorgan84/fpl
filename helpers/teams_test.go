package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/paddymorgan84/fpl/responses"
)

func TestGetTeamID(t *testing.T) {
	tables := []struct {
		teamID            string
		configReaderValue int
		expectedResult    int
		expectedError     bool
	}{
		{"90210", 0, 90210, false},
		{"", 1984, 1984, false},
		{"90210", 1984, 90210, false},
		{"charlie", 0, 0, true},
	}

	for _, table := range tables {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockConfigReader(ctrl)

		reader.EXPECT().GetInt("team-id").Return(table.configReaderValue).AnyTimes()

		parser := new(FplTeamsParser)

		actualResult, err := parser.GetTeamID(table.teamID, reader)

		if err != nil && !table.expectedError {
			t.Fatalf("unexpected error: %v", err)
		}

		if actualResult != table.expectedResult {
			t.Fatalf("expected: %v; got %v", table.expectedResult, actualResult)
		}
	}
}

func TestGetTeam(t *testing.T) {
	tables := []struct {
		teamID       int
		expectedName string
	}{
		{1, "Arsenal"},
		{10, "Leeds"},
		{20, "Wolves"},
		{999999, ""},
	}

	jsonFile, err := os.Open("../sample-data/bootstrap-data.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var bootstrapData responses.BootstrapData
	json.Unmarshal(byteValue, &bootstrapData)
	parser := new(FplTeamsParser)

	for _, table := range tables {
		var actualResult = parser.GetTeam(table.teamID, bootstrapData)

		if actualResult != table.expectedName {
			t.Errorf("expected: %s; got %s", table.expectedName, actualResult)
		}
	}
}

func TestCalculateMonetaryValue(t *testing.T) {
	tables := []struct {
		value         int
		expectedValue float32
	}{
		{1000, 100.0},
		{500, 50.0},
		{1025, 102.5},
		{0, 0.0},
	}

	parser := new(FplTeamsParser)

	for _, table := range tables {
		var actualResult = parser.CalculateMonetaryValue(table.value)

		if actualResult != table.expectedValue {
			t.Errorf("expected: %v; got %v", table.expectedValue, actualResult)
		}
	}
}

func TestCalculateRankComparison(t *testing.T) {
	tables := []struct {
		newRank       int
		oldRank       int
		expectedValue string
	}{
		{5, 10, "🟢"},
		{10, 5, "🔴"},
		{10, 10, "⚪"},
		{0, 0, "⚪"},
	}

	for _, table := range tables {
		var league = responses.ClassicLeague{
			ID:            1,
			Name:          "Corp Street League",
			EntryRank:     table.newRank,
			EntryLastRank: table.oldRank,
		}

		parser := new(FplTeamsParser)

		var actualResult = parser.CalculateRankComparison(league)

		if actualResult != table.expectedValue {
			t.Errorf("expected: %v; got %v", table.expectedValue, actualResult)
		}
	}
}
