package cmd

import (
	"fmt"

	"github.com/dorianneto/1bookmark/internal/extractor"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import bookmarks from brave",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ex extractor.Extractor

		browser := args[0]
		switch browser {
		case "brave":
			fmt.Println("Importing from Brave...")
			ex = extractor.NewBraveExtractor()
		default:
			fmt.Printf("Browser %s not supported\n", browser)
			return
		}

		bookmarks, err := ex.Extract()
		if err != nil {
			fmt.Printf("Error extracting bookmarks: %v\n", err)
			return
		}

		fmt.Printf("Imported %d bookmarks\n", len(bookmarks))
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
