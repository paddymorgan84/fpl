package cmd

import (
	"fmt"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// BuildFixturesCommand returns the fixtures cobra command
func BuildFixturesCommand(c *api.FplAPI) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "fixtures",
		Short: "Get the fixtures for a specific gameweek",
		Run: func(cmd *cobra.Command, args []string) {
			getFixtures(*c)
		},
	}

	return cmd
}

func getFixtures(c api.FplAPI) {
	var bootstrap = c.GetBootstrapData()
	gameweek := helpers.GetCurrentGameweek(bootstrap)
	var fixtures = c.GetFixtures()

	ui.PrintHeader(fmt.Sprintf("Gameweek %d fixtures", gameweek))
	ui.PrintGameweekFixtures(bootstrap, fixtures, gameweek)
}
