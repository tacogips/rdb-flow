package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tacogips/rdb-flow/conf"
)

var rootCmd = &cobra.Command{Use: "rdb-flow"}
var (
	configDir  string
	configName string
	config     conf.RDBFlowConfig
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(refreshCmd, upgradeCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err) //TODO()
		os.Exit(1)
	}
}

func initConfig() {
	if configDir != "" {
		viper.AddConfigPath(configDir)
	} else {
		viper.AddConfigPath(".")
	}

	if configName == "" {
		configName = "rdb-flow-conf" // default conf name
	}

	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("not read config file %+v\n", filepath.Join(configDir, configName))
		os.Exit(1)
	} else {

		viper.Unmarshal(&config)

	}

}
