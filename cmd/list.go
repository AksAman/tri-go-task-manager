/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/AksAman/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all Todos",
	Aliases: []string{"ls"},
	Run:     ListRun,
}

// flags
var (
	showDoneOpt bool
	showAllOpt  bool
)

func ListRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(getDataFilePath())
	if err != nil {
		log.Printf("Error while reading todos: %v \n", err)
	}

	if len(items) == 0 {
		fmt.Println("No TODOs, use add to create new")
		return
	}

	todo.ShowTridos(items, func(item todo.Item) bool {
		return showAllOpt || item.Done == showDoneOpt
	})
}

func init() {

	listCmd.Flags().BoolVarP(&showDoneOpt, "done", "d", false, "shows only completed tasks")
	listCmd.Flags().BoolVarP(&showAllOpt, "all", "a", false, "shows all tasks")

	rootCmd.AddCommand(listCmd)
}
