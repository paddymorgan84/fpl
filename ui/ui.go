package ui

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/juju/ansiterm"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/responses"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// PrintHeader will print standardised header, calculating the appropriate length from the titles characters
func PrintHeader(title string) {

	var b bytes.Buffer

	for i := 0; i <= len(title); i++ {
		b.WriteString("-")
	}

	fmt.Printf("\n%s\n", b.String())
	fmt.Printf("%s\n", title)
	fmt.Printf("%s\n\n", b.String())
}

// PrintTeamPoints prints the points the team has
func PrintTeamPoints(bootstrap responses.BootstrapResponse, live responses.LiveResponse, points responses.PointsResponse) {
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, pick := range points.Picks {
		captain := helpers.DetermineCaptainFlag(pick)
		name := helpers.GetPlayerName(pick, bootstrap)
		playerPoints := helpers.GetPoints(pick, live)

		fmt.Fprintf(tr, "%s %s\t%d\n", name, captain, playerPoints)
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTotal points: %d\n", points.EntryHistory.Points)
}

// PrintManagerDetails prints the summary details for a manager
func PrintManagerDetails(details responses.DetailsResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	_, err := p.Fprintf(tr, "%s\t%s\n", "Team Name: ", details.Name)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Overall Points: ", details.SummaryOverallPoints)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Overall Rank: ", details.SummaryOverallRank)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Gameweek Points: ", details.SummaryEventPoints)

	if err != nil {
		log.Fatal(err)
	}

	err = tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintClassicLeagues prints all classic leagues and the current rank for each
func PrintClassicLeagues(details responses.DetailsResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, league := range details.Leagues.Classic {
		if league.LeagueType == "x" {
			_, err := p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, helpers.CalculateRankComparison(league))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintGlobalLeagues prints all global leagues and the current rank for each
func PrintGlobalLeagues(detailsResponse responses.DetailsResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, league := range detailsResponse.Leagues.Classic {
		if league.LeagueType == "s" {
			_, err := p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, helpers.CalculateRankComparison(league))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintTransfersAndFinance prints details for a teamsd transfers, value and money in the bank
func PrintTransfersAndFinance(detailsResponse responses.DetailsResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	_, err := p.Fprintf(tr, "%s\t%d\n", "Total transfers: ", detailsResponse.LastDeadlineTotalTransfers)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t£%.1f\n", "Squad value: ", helpers.CalculateMonetaryValue(detailsResponse.LastDeadlineValue))

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t£%.1f\n", "In the bank: ", helpers.CalculateMonetaryValue(detailsResponse.LastDeadlineBank))

	if err != nil {
		log.Fatal(err)
	}

	err = tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintGameweekFixtures prints the fixtures for the specified gameweek
func PrintGameweekFixtures(bootstrap responses.BootstrapResponse, fixtures responses.FixturesResponse, gameweek int) {
	for _, fixture := range fixtures {
		if fixture.Event == gameweek {
			var homeTeam = helpers.GetTeam(fixture.HomeTeam, bootstrap)
			var awayTeam = helpers.GetTeam(fixture.AwayTeam, bootstrap)
			fmt.Printf("%s vs %s\n", homeTeam, awayTeam)
		}
	}
}

// PrintSeasonDetails prints a teams current season details for each gameweek
func PrintSeasonDetails(history responses.HistoryResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	fmt.Fprintf(tr, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"Gameweek",
		"Points",
		"Bench Points",
		"Gameweek Rank",
		"Transfers made",
		"Transfers cost",
		"Overall Points",
		"Overall Rank",
		"Team Value")

	for _, gameweek := range history.Current {
		_, err := p.Fprintf(tr, "%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%.1f\n",
			gameweek.Event,
			gameweek.Points,
			gameweek.PointsOnBench,
			gameweek.Rank,
			gameweek.EventTransfers,
			gameweek.EventTransfersCost,
			gameweek.TotalPoints,
			gameweek.OverallRank,
			helpers.CalculateMonetaryValue(gameweek.Value))

		if err != nil {
			log.Fatal(err)
		}
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintChipDetails prints details around what chips have been used
func PrintChipDetails(history responses.HistoryResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	fmt.Fprintf(tr, "%s\t%s\t%s\n", "Date", "Name", "Gamweeek used")

	for _, chip := range history.Chips {
		_, err := p.Fprintf(tr, "%s\t%s\t%d\n", chip.Time, chip.Name, chip.Event)

		if err != nil {
			log.Fatal(err)
		}
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

// PrintPreviousSeasonDetails prints a teams record from past seasons
func PrintPreviousSeasonDetails(history responses.HistoryResponse) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	fmt.Fprintf(tr, "%s\t%s\t%s\n", "Season", "Points", "Rank")

	for _, season := range history.Past {
		_, err := p.Fprintf(tr, "%s\t%d\t%d\n", season.SeasonName, season.TotalPoints, season.Rank)

		if err != nil {
			log.Fatal(err)
		}
	}

	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}
