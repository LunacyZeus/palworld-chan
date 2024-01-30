package cmd

import (
	"github.com/spf13/cobra"
	"palworld-chan/internal/service/wails"
	"palworld-chan/pkg/utility/embeds"
)

func init() {
	rootCmd.AddCommand(uiCMD)
}

var uiCMD = &cobra.Command{
	Use:   "ui",
	Short: "Launch UI",
	Long:  "Launch UI",
	Run: func(cmd *cobra.Command, args []string) {
		LaunchUI()
	},
}

func LaunchUI() {
	wails.Init(embeds.EmbedDir)
	return
}
