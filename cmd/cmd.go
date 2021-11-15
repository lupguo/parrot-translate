package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	cfgViper *viper.Viper
	rootCmd  = &cobra.Command{
		Use:   "parrot-translate",
		Short: "Parrot translate is a translate platform",
		// Long: "parrot-translate is a ",
		Run: func(cmd *cobra.Command, args []string) {
			s, _ := json.Marshal(viper.AllSettings())
			log.Printf("cfgFile info: %s", s)
			cfgViper = viper.GetViper()
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&cfgFile, "cfgFile", "c", "app.yaml", "app cfgFile file (default is app.yaml)")
	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("viper read config file got err: %v", err)
		}
	})
}

// Execute 执行任务
func Execute() error {
	return rootCmd.Execute()
}

func GetConfigViper() *viper.Viper {
	return cfgViper
}
