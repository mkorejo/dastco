package dastco

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Display values from the configuration file.
var viewConfigCmd = &cobra.Command{
	Use:   "view-config",
	Short: "Display values from the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Key", "Value"})
		t.AppendRows([]table.Row{
			{"WebInspect API URL", viper.GetString("url")},
			{"WebInspect Username", viper.GetString("username")},
		})
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(viewConfigCmd)
}
