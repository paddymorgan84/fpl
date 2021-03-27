package cmd

import (
	"fmt"
	"log"
	"strconv"

	fpl "github.com/paddymorgan84/fpl/api"
	responses "github.com/paddymorgan84/fpl/responses"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Get the fixtures for a specific gameweek",
	Run: func(cmd *cobra.Command, args []string) {
		var bootstrap = fpl.GetBootstrapData()
		var gameweek = 0
		var err error

		gameweekParamater := viper.GetString("gameweek")

		if gameweekParamater == "" {
			gameweek = getCurrentGameweek(bootstrap)
		} else {
			gameweek, err = strconv.Atoi(gameweekParamater)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err != nil {
			log.Fatal(err)
		}

		var fixtures = fpl.GetFixtures()

		fmt.Printf("Gameweek %d fixtures\n\n", gameweek)

		for _, fixture := range fixtures {
			if fixture.Event == gameweek {
				var homeTeam = getTeam(fixture.HomeTeam, bootstrap)
				var awayTeam = getTeam(fixture.AwayTeam, bootstrap)
				fmt.Printf("%s vs %s\n", homeTeam, awayTeam)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)
}

func getTeam(teamID int, bootstrap responses.BootstrapResponse) string {
	for _, team := range bootstrap.Teams {
		if team.ID == teamID {
			return team.Name
		}
	}

	return ""
}
