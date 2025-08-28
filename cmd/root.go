/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "To-Do List",
	Short: "This app will help you manage your task priority wise",
}
var datafile string
var cfgFile string

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
	viper.AddConfigPath("/home")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tri")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /home/.tri.yaml)")
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("unable to set home directory please use --datafile to set it")
	}
	rootCmd.PersistentFlags().StringVar(&datafile, "datafile",
		home+string(os.PathSeparator)+".tridos.json", "data file to store in todo")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}
