package cmd

import (
	"encoding/json"
	"log"

	"github.com/lupguo/parrot-translate/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 软件版控制
const version = `0.2.0`

// 应用配置文件读取
var cfgFile string

// 应用系统配置
var cfgViper *viper.Viper

// 命令行提示配置
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "parrot-translate",
	Short:   "Parrot translate is a translate platform",
	// Long: "parrot-translate is a ",
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := json.Marshal(viper.AllSettings())
		log.Printf("app config: %s", s)
		// 启动服务
		server.Start(viper.GetViper())
	},
}

func init() {
	rootCmd.Flags().StringVarP(&cfgFile, "config_file", "c", "app.yaml", "app cfgFile file")

	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("viper read config file got err: %v", err)
		}
	})
}

// Execute 执行任务
func Execute() error {
	err := rootCmd.Execute()
	return err
}
