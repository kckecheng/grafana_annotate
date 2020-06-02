package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an annotation",
	Long: `
Create an annotation for specified panels or all panels on the same dashboard

- to create an annotation for the current time: ignore --from and --to
- to create an annotation for a previous time : specify --from and ignore --to
- to cerate an annotation for a time window   : specify both --from and --to
`,
	Run: func(cmd *cobra.Command, args []string) {
		g := login(cmd)

		flags := cmd.Flags()

		did, err := flags.GetUint("dashboard")
		bailout(err, "Fail to parse option --dashboard", errExit)
		if did == 0 {
			fmt.Println("Dashboard ID is not specified with --dashboard")
			os.Exit(1)
		}

		board, err := g.GetDashboard(did)
		if err != nil {
			bailout(errors.New(""), "Fail to find dashboard with specified ID", errExit)
		}

		pids, err := flags.GetUintSlice("panels")
		bailout(err, "Panel IDs are invalid", errExit)

		if len(pids) == 0 {
			panels, err := g.GetAllPanels(board.UID)
			if err != nil {
				bailout(errors.New(""), "Could not list panels on specified dashboard", errExit)
			}
			for _, panel := range panels {
				pids = append(pids, panel.ID)
			}
		}

		tags, err := flags.GetStringSlice("tags")
		bailout(err, "Tags are invalid", errExit)

		text, err := flags.GetString("text")
		bailout(err, "Annotation words are invalid", errExit)

		froms, err := flags.GetString("from")
		bailout(err, "From time is invalid", errExit)

		tos, err := flags.GetString("to")
		bailout(err, "End time is invalid", errExit)

		var from, to int64
		if froms != "" {
			t, err := time.Parse(time.RFC3339, froms)
			bailout(err, "From time is not formated as expected", errExit)
			from = t.Unix() * 1000
		}

		if tos != "" {
			t, err := time.Parse(time.RFC3339, tos)
			bailout(err, "End time is not formated as expected", errExit)
			to = t.Unix() * 1000
		}

		for _, pid := range pids {
			_, err = g.CreateAnnotation(did, pid, from, to, tags, text)
			if err != nil {
				fmt.Printf("Fail to create annotation for Panel %d\n", pid)
			}
		}

	},
}

func init() {
	annotationCmd.AddCommand(createCmd)
	flags := createCmd.Flags()
	flags.Uint("dashboard", 0, "Dashboard ID")
	flags.UintSlice("panels", []uint{}, `Panel IDs as --panels "id1, id2, ...", all panels will be used if not specified`)
	flags.String("from", "", "Annotation start time in RFC3339 format as 2020-06-02T21:51:00+08:00, current time will be used if not specified")
	flags.String("to", "", "Annotation end time, current time will be used if not specified")
	flags.StringSlice("tags", []string{}, `Annotation tags as --tags "tag1, tag2, ..."`)
	flags.String("text", "", "Annotation words")
	createCmd.MarkFlagRequired("dashboard")
	createCmd.MarkFlagRequired("text")
}
