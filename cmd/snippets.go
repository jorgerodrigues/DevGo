package cmd

import (
	"github.com/jorgerodrigues/devgo/pkg/snippets"
	"github.com/spf13/cobra"
)

var reactComponentSnippetsCmd = &cobra.Command{
	Use:   "fc",
	Short: "Generates a snippet of a react functional component and copies it to the clipboard",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		snippets.GenerateReactComponent()
	},
}

func init() {
	rootCmd.AddCommand(reactComponentSnippetsCmd)
}
