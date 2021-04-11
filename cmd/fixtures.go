package cmd

import (
	"fmt"

	fpl "github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Get the fixtures for a specific gameweek",
	Run: func(cmd *cobra.Command, args []string) {
		var bootstrap = fpl.GetBootstrapData()
		gameweek := helpers.GetCurrentGameweek(bootstrap)
		var fixtures = fpl.GetFixtures()

		ui.PrintHeader(fmt.Sprintf("Gameweek %d fixtures", gameweek))
		ui.PrintGameweekFixtures(bootstrap, fixtures, gameweek)

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)
}
