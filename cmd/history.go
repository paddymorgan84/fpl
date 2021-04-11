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

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Returns history for a managers current and past seasons",
	Run: func(cmd *cobra.Command, args []string) {
		teamID := helpers.GetTeamID(historyArgs.TeamID)
		historyResponse := api.GetHistory(teamID)

		ui.PrintHeader("This season")
		ui.PrintSeasonDetails(historyResponse)

		ui.PrintHeader("Chips")
		ui.PrintChipDetails(historyResponse)

		ui.PrintHeader("Previous Seasons")
		ui.PrintPreviousSeasonDetails(historyResponse)

	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().StringVarP(&historyArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")
}
