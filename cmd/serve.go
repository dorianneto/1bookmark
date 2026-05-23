package cmd

import "github.com/spf13/cobra"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web dashboard",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
