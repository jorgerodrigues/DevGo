package cmd

import (
	"github.com/jorgerodrigues/devgo/pkg/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID and copy it to the clipboard",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		uuid.GenerateUUID()
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
