package cmd

import (
	"fmt"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// BuildFixturesCommand returns the fixtures cobra command
func BuildFixturesCommand(c api.FplAPI, config helpers.ConfigReader) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "fixtures",
		Short: "Get the fixtures for a specific gameweek",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getFixtures(c, config)
		},
	}

	return cmd
}

func getFixtures(c api.FplAPI, config helpers.ConfigReader) error {
	var bootstrap, err = c.GetBootstrapData()

	if err != nil {
		return err
	}

	gameweek, err := helpers.GetCurrentGameweek(bootstrap.Gameweeks, config)

	if err != nil {
		return err
	}

	fixtures, err := c.GetGameweekFixtures(gameweek)

	if err != nil {
		return err
	}

	ui.PrintHeader(fmt.Sprintf("Gameweek %d fixtures", gameweek))
	ui.PrintGameweekFixtures(bootstrap, fixtures, gameweek)

	return err
}
