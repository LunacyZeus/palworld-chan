package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api"
	"palworld-chan/internal/service/api/config"
)

func init() {
	rootCmd.AddCommand(apiCMD)

	apiCMD.Flags().String("port", ":3001", "api port example: port :3001")
	apiCMD.Flags().String("authUser", "pal", "web auth username")
	apiCMD.Flags().String("authPassWd", "123", "web auth password")
}

var apiCMD = &cobra.Command{
	Use:   "api",
	Short: "Launch Server API",
	Long:  "Launch Server API",
	Run: func(cmd *cobra.Command, args []string) {
		//读取参数
		port := viper.GetString("port")
		authUser := viper.GetString("authUser")
		authPassWd := viper.GetString("authPassWd")
		LaunchServerAPI(port, authUser, authPassWd)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("port", cmd.Flags().Lookup("port"))
		viper.BindPFlag("authUser", cmd.Flags().Lookup("authUser"))
		viper.BindPFlag("authPassWd", cmd.Flags().Lookup("authPassWd"))

		viper.AutomaticEnv()
	},
}

func LaunchServerAPI(port, authUser, authPassWd string) {
	config.Auth().Set(authUser, authPassWd) //密码设置
	fmt.Printf("API Auth User:%s\n", authUser)
	fmt.Printf("API Auth PassWord:%s\n", authPassWd)
	fmt.Printf("PalChan Version:%s\n", consts.VERSION)

	api.Init(port)
	return
}
