/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AksAman/tri/models"
	"github.com/AksAman/tri/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Run:   AddRun,
}

func AddRun(cmd *cobra.Command, args []string) {

	logger.Debugln("Add called with args:")

	items := []models.Item{}

	for _, arg := range args {
		newItem := models.Item{
			Text: arg,
		}
		newItem.SetPriority(priority)
		items = append(items, newItem)
	}

	logger.Debugf("Items: %#v\n", items)

	err := services.SaveItems(items)
	if err != nil {
		logger.Errorf("Error while saving todos: %#v \n", err)
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority for TODO")
}
