package helpers

import (
	"sort"
	"strconv"
	"time"

	"github.com/paddymorgan84/fpl/responses"
)

// GameweekParser helps me abstract away the gaameweek related work so I can better unit test it
type GameweekParser interface {
	GetCurrentGameweek(gameweeks []responses.Gameweek, config ConfigReader) (int, error)
}

// FplGameweekParser handles any gameweek related parsing that needs to be done
type FplGameweekParser struct {
}

// GetCurrentGameweek uses the bootstrap data to determine which gameweek is currently active in the game. If all are finished, defaults to 38
func (g FplGameweekParser) GetCurrentGameweek(gameweeks []responses.Gameweek, config ConfigReader) (int, error) {
	gameweekParameter := config.GetString("gameweek")

	sort.SliceStable(gameweeks, func(i, j int) bool {
		return gameweeks[i].ID < gameweeks[j].ID
	})

	if gameweekParameter == "" {
		for _, gameweek := range gameweeks {
			if !gameweek.Finished {
				if gameweek.DeadlineTime.After(time.Now()) {
					return gameweek.ID - 1, nil
				}
				return gameweek.ID, nil
			}
		}
	} else {
		gameweek, err := strconv.Atoi(gameweekParameter)
		return gameweek, err
	}

	return gameweeks[len(gameweeks)-1].ID, nil
}

//go:generate mockgen -source=gameweeks.go -package=helpers -destination=mock_gameweek_parser.go
