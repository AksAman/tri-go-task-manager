/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/AksAman/tri/cmd"
	"github.com/AksAman/tri/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	utils.InitializeLogger("test.log")
	logger = utils.Logger
}

func main() {
	cmd.Execute()
}
