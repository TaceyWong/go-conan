package main

import "github.com/urfave/cli/v2"

var InfoCMD = cli.Command{
	Name:     "info",
	Category: CmdCateConsumer,
	Usage:    "Gets information about the dependency graph of a recipe.",
}
