package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grafana_annotate",
	Short: "Annotate Grafana panels",
	Long: `
Annotate Grafana panels with customized text and tags`,
}

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringP("address", "a", "localhost", "Grafana server address")
	flags.Int64P("port", "l", 3000, "Grafana server listening port")
	flags.StringP("user", "u", "admin", "Grafana user name")
	flags.StringP("password", "p", "admin", "Grafana password")
}

// Execute entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
