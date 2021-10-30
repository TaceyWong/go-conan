package main

import "github.com/urfave/cli/v2"

var CopyCMD = cli.Command{
	Name:     "copy",
	Category: CmdCateMisc,
	Usage:    "Copies conan recipes and packages to another user/channel.",
}
