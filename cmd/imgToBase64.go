package cmd

import (
	"github.com/jorgerodrigues/devgo/pkg/imgBase64"
	"github.com/spf13/cobra"
)

var imgToBase64 = &cobra.Command{
	Use:   "imgToBase64",
	Short: "Converts an image to base64",
	Long:  `Converts an image to base64`,
	Run: func(cmd *cobra.Command, args []string) {
		imgBase64.ImgToBase64(args[0])
	},
}

func init() {
	rootCmd.AddCommand(imgToBase64)
}
