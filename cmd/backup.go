package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"palworld-chan/internal/service/backup"
	"palworld-chan/pkg/logger"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "local backup",
	Long:  "local backup",
	RunE: func(command *cobra.Command, args []string) error {
		setups := []func() error{}
		for _, setup := range setups {
			if err := setup(); err != nil {
				panic(err)
			}
		}

		//读取参数
		source := viper.GetString("source")
		dest := viper.GetString("dest")
		backupCount := viper.GetInt("backupCount")

		err := backup.LocalBackUp(source, dest, backupCount)
		if err != nil {
			logger.Fatal("%v", err)
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("source", cmd.Flags().Lookup("source"))
		viper.BindPFlag("dest", cmd.Flags().Lookup("dest"))
		viper.BindPFlag("backupCount", cmd.Flags().Lookup("backupCount"))

		viper.AutomaticEnv()
	},

	//serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().String("source", "", "save file path")
	backupCmd.Flags().String("dest", "", "backup file path")
	backupCmd.Flags().Int("backupCount", 200, "max backup file count default is 200")
}
