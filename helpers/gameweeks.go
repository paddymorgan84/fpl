package helpers

import (
	"log"
	"strconv"
	"time"

	"github.com/paddymorgan84/fpl/responses"
	"github.com/spf13/viper"
)

const lastGameweek int = 38

// GetCurrentGameweek uses the bootstrap data to determine which gameweek is currently active in the game. If all are finished, defaults to 38
func GetCurrentGameweek(bootstrap responses.BootstrapData) int {
	gameweekParameter := viper.GetString("gameweek")

	if gameweekParameter == "" {
		for _, gameweek := range bootstrap.Gameweeks {
			if !gameweek.Finished {
				if gameweek.DeadlineTime.After(time.Now()) {
					return gameweek.ID - 1
				}

				return gameweek.ID
			}
		}
	} else {
		gameweek, err := strconv.Atoi(gameweekParameter)
		if err != nil {
			log.Fatal(err)
		}

		return gameweek
	}

	return lastGameweek
}
