package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var panelCmd = &cobra.Command{
	Use:   "panel",
	Short: "List panels",
	Long: `
List all panels of the specified dashboard`,
	Run: func(cmd *cobra.Command, args []string) {
		g := login(cmd)

		did, err := cmd.Flags().GetUint("dashboard")
		bailout(err, "Fail to parse option --dashboard", errExit)
		if did == 0 {
			fmt.Println("Dashboard ID is not specified with --dashboard")
			os.Exit(1)
		}

		board, err := g.GetDashboard(did)
		if err != nil {
			bailout(errors.New(""), "Fail to find dashboard with specified ID", errExit)
		}

		panels, err := g.GetAllPanels(board.UID)
		bailout(err, "Fail to list panels on specified dashboard", errExit)

		if len(panels) > 0 {
			fmt.Println("Panels(Format: PanelID Title):")
		}
		for _, panel := range panels {
			fmt.Printf("%3d %s\n", panel.ID, panel.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(panelCmd)
	flags := panelCmd.Flags()
	flags.Uint("dashboard", 0, "Dashboard ID")
	panelCmd.MarkFlagRequired("dashboard")
}
