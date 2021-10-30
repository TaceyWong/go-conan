package main

import "github.com/urfave/cli/v2"

var AliasCMD = cli.Command{
	Name:     "alias",
	Category: CmdCateMisc,
	Usage:    "Creates and exports an 'alias package recipe'.",
}
