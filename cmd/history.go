package cmd

import (
	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// HistoryArgs are the arguments you can pass to the history command
type HistoryArgs struct {
	TeamID string
}

var historyArgs HistoryArgs

// BuildHistoryCommand returns the history cobra command
func BuildHistoryCommand(c api.FplAPI, config helpers.ConfigReader) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "history",
		Short: "Returns history for a managers current and past seasons",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getHistory(c, config)
		},
	}

	var flags = cmd.Flags()

	flags.StringVarP(&historyArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}

func getHistory(c api.FplAPI, config helpers.ConfigReader) error {
	teamID, err := helpers.GetTeamID(historyArgs.TeamID, config)

	if err != nil {
		return err
	}

	historyResponse, err := c.GetManagerHistory(teamID)

	if err != nil {
		return err
	}

	ui.PrintHeader("This season")
	ui.PrintSeasonDetails(historyResponse)

	ui.PrintHeader("Chips")
	ui.PrintChipDetails(historyResponse)

	ui.PrintHeader("Previous Seasons")
	ui.PrintPreviousSeasonDetails(historyResponse)

	return err
}
