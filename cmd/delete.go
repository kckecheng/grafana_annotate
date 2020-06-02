package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete annotations",
	Long: `
Delete annotations by their IDs
`,
	Run: func(cmd *cobra.Command, args []string) {
		g := login(cmd)

		ids, err := cmd.Flags().GetUintSlice("ids")
		bailout(err, "Annotation IDs are invalid", errExit)

		for _, id := range ids {
			err := g.DeleteAnnotation(id)
			if err != nil {
				fmt.Printf("Fail to delete annotation %d: %s", id, err)
			}
		}
	},
}

func init() {
	annotationCmd.AddCommand(deleteCmd)
	flags := deleteCmd.Flags()
	flags.UintSlice("ids", []uint{}, `Annotation IDs as --ids "id1, id2, ..."`)
	deleteCmd.MarkFlagRequired("ids")
}
