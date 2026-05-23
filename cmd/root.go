package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "onebookmark",
	Short: "onebookmark is a cli tool for Bookmark archaeology",
	Long:  "onebookmark is a cli tool for Bookmark archaeology — visualise your saved URLs and find the ones you forgot about.",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing 1bookmark '%s'\n", err)
		os.Exit(1)
	}
}
