package helpers

import (
	"time"

	"github.com/paddymorgan84/fpl/responses"
)

// GetCurrentGameweek uses the bootstrap data to determine which gameweek is currently active in the game. If all are finished, defaults to 38
func GetCurrentGameweek(bootstrap responses.BootstrapResponse) int {
	for _, gameweek := range bootstrap.Gameweeks {
		if !gameweek.Finished {
			if gameweek.DeadlineTime.After(time.Now()) {
				return gameweek.ID - 1
			}

			return gameweek.ID
		}
	}

	return 38
}
