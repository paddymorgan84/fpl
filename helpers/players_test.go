package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/paddymorgan84/fpl/responses"
)

func TestDetermineCaptainFlag(t *testing.T) {
	tables := []struct {
		isCaptain      bool
		isViceCaptain  bool
		expectedResult string
	}{
		{true, true, "(C)"},
		{true, false, "(C)"},
		{false, true, "(VC)"},
		{false, false, ""},
	}

	for _, table := range tables {
		var pick = responses.Pick{
			Element:       1,
			Position:      2,
			Multiplier:    3,
			IsCaptain:     table.isCaptain,
			IsViceCaptain: table.isViceCaptain,
		}

		var actualResult = DetermineCaptainFlag(pick)

		if actualResult != table.expectedResult {
			t.Fatalf("expected: %s; got %s", table.expectedResult, actualResult)
		}
	}
}

func TestGetPlayerName(t *testing.T) {
	tables := []struct {
		playerID     int
		expectedName string
	}{
		{1, "Özil"},
		{2, "Sokratis"},
		{3, "David Luiz"},
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

	for _, table := range tables {

		var pick = responses.Pick{
			Element:       table.playerID,
			Position:      2,
			Multiplier:    3,
			IsCaptain:     false,
			IsViceCaptain: false,
		}

		var actualResult = GetPlayerName(pick, bootstrapData)

		if actualResult != table.expectedName {
			t.Errorf("expected: %s; got %s", table.expectedName, actualResult)
		}
	}
}

func TestGetPoints(t *testing.T) {
	tables := []struct {
		playerID       int
		multiplier     int
		expectedPoints int
	}{
		{20, 0, 8},
		{20, 1, 8},
		{20, 2, 16},
		{20, 3, 24},
		{999999, 3, 0},
	}

	jsonFile, err := os.Open("../sample-data/gameweek-live-scores.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var liveScores responses.GameweekLiveScores
	json.Unmarshal(byteValue, &liveScores)

	for _, table := range tables {

		var pick = responses.Pick{
			Element:       table.playerID,
			Position:      2,
			Multiplier:    table.multiplier,
			IsCaptain:     false,
			IsViceCaptain: false,
		}

		var actualPoints = GetPoints(pick, liveScores)

		if actualPoints != table.expectedPoints {
			t.Errorf("expected: %v; got %v", table.expectedPoints, actualPoints)
		}

	}
}
