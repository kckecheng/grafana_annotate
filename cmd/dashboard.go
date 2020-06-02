package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Query dashboards",
	Long: `
Query dashboards based on the title. If title is not specified, all dashboars will be retured.`,
	Run: func(cmd *cobra.Command, args []string) {
		g := login(cmd)

		title, err := cmd.Flags().GetString("title")
		bailout(err, "Fail to parse option --title", errExit)
		title = strings.Trim(title, " ")

		boards, err := g.GetAllDashboards()
		if err != nil {
			bailout(errors.New(""), "Fail to query dashboars", errExit)
		}

		if len(boards) > 0 {
			fmt.Println("Dashboards(Format: ID Title):")
		}
		for _, board := range boards {
			if title != "" {
				if strings.Contains(board.Title, title) {
					fmt.Printf("%3d %s\n", board.ID, board.Title)
				}
			} else {
				fmt.Printf("%3d %s\n", board.ID, board.Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
	flags := dashboardCmd.Flags()
	flags.String("title", "", "Dashboard title pattern, the whole title or a substring of the title")
}
