package cmd

import "github.com/spf13/cobra"

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import bookmarks from chrome",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
