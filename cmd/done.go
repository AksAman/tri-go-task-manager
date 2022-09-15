/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/AksAman/tri/services"
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
		logger.Fatalln("No id provided")
	}

	idToMark, err := strconv.Atoi(args[0])

	if err != nil {
		logger.Fatalln("Provided argument is not a valid id, it should be of type:int")
		return
	}

	if idToMark < 0 {
		logger.Fatalln("Provided argument is not a valid id, it should be a positive integer")
		return
	}

	err = services.MarkItemDoneByID(idToMark)
	if err != nil {
		logger.Fatalf("error while marking item done by id(%d): %v", idToMark, err)
		return
	}

	ListRun(cmd, args)
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
