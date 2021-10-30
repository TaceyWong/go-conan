package main

import "github.com/urfave/cli/v2"

var GetCMD = cli.Command{
	Name:     "get",
	Category: CmdCateConsumer,
	Usage:    "Gets a file or list a directory of a given reference or package.",
}
