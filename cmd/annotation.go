package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var annotationCmd = &cobra.Command{
	Use:   "annotation",
	Short: "List annotations",
	Long: `
List annotations`,
	Run: func(cmd *cobra.Command, args []string) {
		g := login(cmd)

		did, err := cmd.Flags().GetUint("dashboard")
		bailout(err, "Fail to parse option --dashboard", errExit)
		if did == 0 {
			fmt.Println("Dashboard ID is not specified with --dashboard")
			os.Exit(1)
		}

		_, err = g.GetDashboard(did)
		if err != nil {
			bailout(errors.New(""), "Fail to find dashboard with specified ID", errExit)
		}

		var tzero time.Time
		annotations, err := g.GetAnnotattions([]uint{did}, nil, tzero, tzero, nil)
		bailout(err, "Fail to list annotations for specified dashbaord", errExit)

		if len(annotations) > 0 {
			fmt.Println("Annotations(Format: AnnotationID PanelID From End Tags Text):")
		}
		for _, annotation := range annotations {
			fmt.Printf(
				"%4d %3d %v %v %+v %s\n",
				annotation.ID,
				annotation.PanelID,
				time.Unix(annotation.Time/1000, 0),
				time.Unix(annotation.TimeEnd/1000, 0),
				annotation.Tags,
				annotation.Text,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(annotationCmd)
	flags := annotationCmd.Flags()
	flags.Uint("dashboard", 0, "Dashboard ID")
	annotationCmd.MarkFlagRequired("dashboard")
}
