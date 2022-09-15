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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Run:   AddRun,
}

func AddRun(cmd *cobra.Command, args []string) {

	fmt.Println("Add called with args:")

	items, err := todo.ReadItems(getDataFilePath())

	if err != nil {
		log.Printf("Error while reading todos: %v \n", err)
	}

	lastPosition := -1
	if len(items) > 0 {
		lastPosition = items[len(items)-1].Position
	}

	for _, arg := range args {
		newItem := todo.Item{
			Text:     arg,
			Position: lastPosition + 1,
		}
		newItem.SetPriority(priority)
		items = append(items, newItem)
		lastPosition++
	}

	fmt.Printf("Items: %#v\n", items)

	err = todo.SaveItems(getDataFilePath(), items)
	if err != nil {
		log.Printf("Error while saving todos: %#v \n", err)
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority for TODO")
}
