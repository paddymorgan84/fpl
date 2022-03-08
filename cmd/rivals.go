package cmd

import (
	"errors"
	"strconv"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// BuildRivalsCommand returns the rivals cobra command
func BuildRivalsCommand(c api.FplAPI, config helpers.ConfigReader, gameweekParser helpers.GameweekParser, renderer ui.Renderer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rivals",
		Short: "Show the points for all of your rivals (specified in config) for a specified gameweek",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRivals(c, config, gameweekParser, renderer)
		},
	}

	return cmd
}

func getRivals(c api.FplAPI, config helpers.ConfigReader, gameweekParser helpers.GameweekParser, renderer ui.Renderer) error {
	if !config.IsSet("rivals") {
		renderer.PrintHeader("Rivals")
		renderer.PrintNoRivalsWarning()
		return nil
	}

	bootstrap, err := c.GetBootstrapData()

	if err != nil {
		return err
	}

	gameweek, err := gameweekParser.GetCurrentGameweek(bootstrap.Gameweeks, config)

	if err != nil {
		return err
	}

	live, err := c.GetGameweekLiveScores(gameweek)

	if err != nil {
		return err
	}

	renderer.PrintHeader("Rivals")
	for _, rival := range config.GetStringSlice("rivals") {
		teamID, err := strconv.Atoi(rival)

		if err != nil {
			return errors.New("rivals must all be numeric ids")
		}

		points, err := c.GetGameweekPoints(teamID, gameweek)

		if err != nil {
			return err
		}

		detailsResponse, err := c.GetManagerDetails(teamID)

		if err != nil {
			return err
		}

		renderer.PrintRivalPoints(bootstrap, live, points, detailsResponse)
	}

	return err
}
