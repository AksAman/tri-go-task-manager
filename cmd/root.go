/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AksAman/tri/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile string
	logger     *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "A Short TODO cli app",
	Long: `
	Tri is a Todo cli library built using Golang and Cobra,
	that helps to CRUD, search todo's
	using your terminal. Created using Go, based on workshop/tutorial
	by spf13 (https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/)
`,
}

func init() {
	utils.InitializeLogger("tri.log")
	logger = utils.Logger
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file override (Default is $HOME/.tri.yml),")

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")

	viper.AutomaticEnv()

	viper.SetEnvPrefix("tri")

	// try reading config file

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file stored at %v\n", viper.ConfigFileUsed())
	}

}
