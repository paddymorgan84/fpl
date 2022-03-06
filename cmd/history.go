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
func BuildHistoryCommand(c api.FplAPI, config helpers.ConfigReader, teamParser helpers.TeamsParser, renderer ui.Renderer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "history",
		Short: "Returns history for a managers current and past seasons",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getHistory(c, config, teamParser, renderer)
		},
	}

	var flags = cmd.Flags()

	flags.StringVarP(&historyArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}

func getHistory(c api.FplAPI, config helpers.ConfigReader, teamParser helpers.TeamsParser, renderer ui.Renderer) error {
	teamID, err := teamParser.GetTeamID(historyArgs.TeamID, config)

	if err != nil {
		return err
	}

	historyResponse, err := c.GetManagerHistory(teamID)

	if err != nil {
		return err
	}

	renderer.PrintHeader("This season")
	renderer.PrintSeasonDetails(historyResponse, teamParser)

	renderer.PrintHeader("Chips")
	renderer.PrintChipDetails(historyResponse)

	renderer.PrintHeader("Previous Seasons")
	renderer.PrintPreviousSeasonDetails(historyResponse)

	return err
}
