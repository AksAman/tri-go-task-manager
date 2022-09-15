/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AksAman/tri/models"
	"github.com/AksAman/tri/services"
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
	items, err := services.ReadItems()
	if err != nil {
		logger.Errorf("Error while reading todos: %v \n", err)
		fmt.Println("No TODOs, use add to create new")
		return
	}

	if len(items) == 0 {
		fmt.Println("No TODOs, use add to create new")
		return
	}

	services.ShowTridos(items, func(item models.Item) bool {
		return showAllOpt || item.Done == showDoneOpt
	})
}

func init() {

	listCmd.Flags().BoolVarP(&showDoneOpt, "done", "d", false, "shows only completed tasks")
	listCmd.Flags().BoolVarP(&showAllOpt, "all", "a", false, "shows all tasks")

	rootCmd.AddCommand(listCmd)
}
