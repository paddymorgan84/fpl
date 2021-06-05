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
func BuildDetailsCommand(c api.FplAPI, config helpers.ConfigReader) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "details",
		Short: "Returns details of manager for current season, e.g. league standings, cash in the bank, overall points etc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDetails(c, config)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&detailsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}

func getDetails(c api.FplAPI, config helpers.ConfigReader) error {
	teamID, err := helpers.GetTeamID(detailsArgs.TeamID, config)

	if err != nil {
		return err
	}

	detailsResponse, err := c.GetManagerDetails(teamID)

	if err != nil {
		return err
	}

	ui.PrintHeader("Manager Details")
	ui.PrintManagerDetails(detailsResponse)

	ui.PrintHeader("Classic Leagues")
	ui.PrintClassicLeagues(detailsResponse)

	ui.PrintHeader("Global Leagues")
	ui.PrintGlobalLeagues(detailsResponse)

	ui.PrintHeader("Transfers & Finance")
	ui.PrintTransfersAndFinance(detailsResponse)

	return err
}
