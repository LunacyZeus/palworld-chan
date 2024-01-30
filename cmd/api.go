package cmd

import (
	"github.com/spf13/cobra"
	"palworld-chan/internal/service/api"
)

func init() {
	rootCmd.AddCommand(apiCMD)
}

var apiCMD = &cobra.Command{
	Use:   "api",
	Short: "Launch Server API",
	Long:  "Launch Server API",
	Run: func(cmd *cobra.Command, args []string) {
		LaunchServerAPI()
	},
}

func LaunchServerAPI() {
	port := ":3000"
	api.Init(port)
	return
}
