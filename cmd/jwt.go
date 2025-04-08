package cmd

import (
	"github.com/jorgerodrigues/devgo/pkg/jwt"
	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Parses and decodes JWT tokens",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if err := jwt.DecodeTokenFromArgs(args); err != nil {
			cmd.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}
