/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/AksAman/tri/models"
	"github.com/AksAman/tri/services"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches all todos using a keyword",
	Run:   searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("No search keyword passed, exiting!")
	}

	items, err := services.ReadItems()
	if err != nil {
		log.Printf("Error while reading todos: %v \n", err)
	}

	if len(items) == 0 {
		fmt.Println("No TODOs, use add to create new")
		return
	}

	keyword := args[0]

	services.ShowTridos(items, func(item models.Item) bool {
		textToSearch := strings.ToLower(item.Text)
		query := strings.ToLower(keyword)
		return strings.Contains(textToSearch, query)
	})
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
