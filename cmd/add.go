/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vivo1806/mycli/todo"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new TO-DO",
	Long:  `add will create a new to-do item into the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	//magical things happen here

	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(Priority)
		item.CreatedAt = time.Now()
		items = append(items, item)
	}
	e := todo.SaveItems(viper.GetString("datafile"), items)
	if e != nil {
		fmt.Errorf("%v", e)
	}

	fmt.Println("todo added!!!.make sure to do it")
}

var Priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
	addCmd.Flags().IntVarP(&Priority, "priority", "p", 2, "Priority:1,2,3") //:NOTE p for short-term flag
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
