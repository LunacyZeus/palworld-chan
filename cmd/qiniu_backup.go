package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"palworld-chan/internal/service/backup"
	"palworld-chan/pkg/logger"
)

var qiNiuBackupCmd = &cobra.Command{
	Use: "7backup",
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
		accessKey := viper.GetString("accessKey")
		secretKey := viper.GetString("secretKey")
		bucket := viper.GetString("bucket")

		err := backup.QiNiuBackUp(source, dest, accessKey, secretKey, bucket)
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
		viper.BindPFlag("accessKey", cmd.Flags().Lookup("ak"))
		viper.BindPFlag("secretKey", cmd.Flags().Lookup("sk"))
		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))

		viper.AutomaticEnv()
	},

	//serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

}

func init() {
	rootCmd.AddCommand(qiNiuBackupCmd)

	qiNiuBackupCmd.Flags().String("source", "", "save file path")
	qiNiuBackupCmd.Flags().String("dest", "", "backup file path")

	qiNiuBackupCmd.Flags().String("ak", "", "qiniu accessToken")
	qiNiuBackupCmd.Flags().String("sk", "", "qiniu secretToken")
	qiNiuBackupCmd.Flags().String("bucket", "", "qiniu bucket")
}
