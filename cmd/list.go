/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vivo1806/mycli/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing the todos`,
	Run:   listRun,
}
var (
	doneOpt bool
)

func listRun(cmd *cobra.Command, args []string) {

	file := viper.GetString("datafile")
	if file == "" {
		file = datafile
	}

	items, err := todo.ReadItems(file)
	if err != nil {
		fmt.Println(err)
	}
	if len(items) == 0 {
		fmt.Println("No todos at the moment")
	}
	sort.Sort(todo.ByPri(items)) //:NOTE: todo.interface Len(), Swap(), Less() are needed by Sort interface
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	if doneOpt {
		for _, i := range items {
			if i.Done {
				fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t"+i.PrettyCreated())
			}

		}
	} else {
		for _, i := range items {

			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t"+i.PrettyCreated())

		}
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "show 'Done' todos")

}
