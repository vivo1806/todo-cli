/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vivo1806/mycli/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark item as done",
	Aliases: []string{"do"},
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label")
	}
	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked Done")
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesnt match any item")
	}
}
func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
