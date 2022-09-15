/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/AksAman/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Marks todo as done",
	Aliases: []string{"do"},
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		log.Fatalln("No id provided")
	}

	idToMark, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln("Provided argument is not a valid id, it should be of type:int")
	}

	if idToMark < 0 {
		log.Fatalln("Provided argument is not a valid id, it should be a positive integer")
	}

	existingItems, err := todo.ReadItems(getDataFilePath())

	if idToMark > len(existingItems) {
		log.Fatalln("Provided ID exceeds maximum possible todo's id")
	}

	position := searchItemsForPosition(idToMark, existingItems)

	if position < 0 {
		log.Fatalln("No such position found")
	}

	item := existingItems[position]
	if item.Done {
		log.Println("[" + item.Label() + " : " + item.Text + "] is already complete")
		return
	}
	log.Println("Marking [" + item.Label() + " : " + item.Text + "] as Done.")

	existingItems[position].Done = true

	todo.SaveItems(getDataFilePath(), existingItems)
	ListRun(cmd, args)
}

func searchItemsForPosition(keyPosition int, items []todo.Item) int {
	for _, item := range items {
		if item.Position == keyPosition {
			return item.Position
		}
	}
	return -1
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
