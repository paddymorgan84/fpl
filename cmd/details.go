package cmd

import (
	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// DetailsArgs are the arguments you can pass to the history command
type DetailsArgs struct {
	TeamID string
}

var detailsArgs DetailsArgs

// BuildDetailsCommand returns the details cobra command
func BuildDetailsCommand(c api.FplAPI, config helpers.ConfigReader, teamParser helpers.TeamsParser, renderer ui.Renderer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "details",
		Short: "Returns details of manager for current season, e.g. league standings, cash in the bank, overall points etc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDetails(c, config, teamParser, renderer)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&detailsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}

func getDetails(c api.FplAPI, config helpers.ConfigReader, teamParser helpers.TeamsParser, renderer ui.Renderer) error {
	teamID, err := teamParser.GetTeamID(detailsArgs.TeamID, config)

	if err != nil {
		return err
	}

	detailsResponse, err := c.GetManagerDetails(teamID)

	if err != nil {
		return err
	}

	renderer.PrintHeader("Manager Details")
	renderer.PrintManagerDetails(detailsResponse)

	renderer.PrintHeader("Classic Leagues")
	renderer.PrintClassicLeagues(detailsResponse, teamParser)

	renderer.PrintHeader("Global Leagues")
	renderer.PrintGlobalLeagues(detailsResponse, teamParser)

	renderer.PrintHeader("Transfers & Finance")
	renderer.PrintTransfersAndFinance(detailsResponse, teamParser)

	return err
}
