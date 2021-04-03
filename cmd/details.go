package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/juju/ansiterm"
	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// DetailsArgs are the arguments you can pass to the history command
type DetailsArgs struct {
	TeamID string
}

var detailsArgs DetailsArgs

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Returns details of manager for current season, e.g. league standings, cash in the bank, overall points etc",
	Run: func(cmd *cobra.Command, args []string) {
		teamID := 0
		var err error

		if detailsArgs.TeamID == "" {
			teamID = viper.GetInt("team-id")
		} else {
			teamID, err = strconv.Atoi(detailsArgs.TeamID)
			if err != nil {
				log.Fatal(err)
			}
		}

		p := message.NewPrinter(language.English)
		detailsResponse := api.GetDetails(teamID)

		tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

		fmt.Println("\n------------")
		fmt.Println("Manager Details")
		fmt.Println("------------")

		p.Fprintf(tr, "%s\t%s\n", "Team Name: ", detailsResponse.Name)
		p.Fprintf(tr, "%s\t%d\n", "Overall Points: ", detailsResponse.SummaryOverallPoints)
		p.Fprintf(tr, "%s\t%d\n", "Overall Rank: ", detailsResponse.SummaryOverallRank)
		p.Fprintf(tr, "%s\t%d\n", "Gameweek Points: ", detailsResponse.SummaryEventPoints)

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n------------")
		fmt.Println("Classic Leagues")
		fmt.Println("------------")

		for _, league := range detailsResponse.Leagues.Classic {
			if league.LeagueType == "x" {
				p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, calculateRankComparison(league))
			}
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n------------")
		fmt.Println("Global Leagues")
		fmt.Println("------------")

		for _, league := range detailsResponse.Leagues.Classic {
			if league.LeagueType == "s" {
				p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, calculateRankComparison(league))
			}
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n------------")
		fmt.Println("Transfers & Finance")
		fmt.Println("------------")

		p.Fprintf(tr, "%s\t%d\n", "Total transfers: ", detailsResponse.LastDeadlineTotalTransfers)
		p.Fprintf(tr, "%s\t£%.1f\n", "Squad value: ", helpers.CalculateMonetaryValue(detailsResponse.LastDeadlineValue))
		p.Fprintf(tr, "%s\t£%.1f\n", "In the bank: ", helpers.CalculateMonetaryValue(detailsResponse.LastDeadlineBank))

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(detailsCmd)
	detailsCmd.Flags().StringVarP(&detailsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")
}
