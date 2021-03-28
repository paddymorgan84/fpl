package cmd

import (
	"fmt"
	"log"
	"strconv"

	fpl "github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
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

		gameweekParameter := viper.GetString("gameweek")

		if gameweekParameter == "" {
			gameweek = helpers.GetCurrentGameweek(bootstrap)
		} else {
			gameweek, err = strconv.Atoi(gameweekParameter)
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
				var homeTeam = helpers.GetTeam(fixture.HomeTeam, bootstrap)
				var awayTeam = helpers.GetTeam(fixture.AwayTeam, bootstrap)
				fmt.Printf("%s vs %s\n", homeTeam, awayTeam)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)
}
