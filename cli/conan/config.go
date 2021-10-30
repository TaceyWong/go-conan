package main

import "github.com/urfave/cli/v2"

var ConfigCMD = cli.Command{
	Name:     "config",
	Category: CmdCateConsumer,
	Usage:    "Manages Conan configuration",
}
