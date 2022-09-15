/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a given todo",
	Run:   editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	logger.Infoln("edit called")
}

func init() {
	rootCmd.AddCommand(editCmd)

}
