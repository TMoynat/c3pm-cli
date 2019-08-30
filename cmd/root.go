package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: os.Args[0],
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
    err := viper.WriteConfig()
    if err != nil {
    	os.Exit(0)
	}
	//fmt.Println("debug")
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.


	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configFile = "c3pm.yml"
	viper.SetConfigType("yml")
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("c3pm")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
