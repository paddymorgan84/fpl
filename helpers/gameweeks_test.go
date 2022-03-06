package helpers

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/paddymorgan84/fpl/responses"
)

func TestGetCurrentGameweek(t *testing.T) {
	tables := []struct {
		gameweeks         []responses.Gameweek
		configReaderValue string
		expectedResult    int
		expectedError     bool
	}{
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "1", 1, false},
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "2", 2, false},
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "3", 3, false},
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "", 2, false},
		{[]responses.Gameweek{
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "", 3, false},
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, 1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "", 2, false},
		{[]responses.Gameweek{
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -1),
				ID:           3,
			},
			{
				Finished:     true,
				DeadlineTime: time.Now().AddDate(0, 0, -3),
				ID:           1,
			},
			{
				Finished:     false,
				DeadlineTime: time.Now().AddDate(0, 0, -2),
				ID:           2,
			}}, "Charlie", 0, true},
	}

	for _, table := range tables {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockConfigReader(ctrl)

		reader.EXPECT().GetString("gameweek").Return(table.configReaderValue).AnyTimes()

		parser := new(FplGameweekParser)

		actualResult, err := parser.GetCurrentGameweek(table.gameweeks, reader)

		if err != nil && !table.expectedError {
			t.Fatalf("unexpected error: %v", err)
		}

		if actualResult != table.expectedResult {
			t.Fatalf("expected: %v; got %v", table.expectedResult, actualResult)
		}
	}
}
