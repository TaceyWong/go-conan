package main

import "github.com/urfave/cli/v2"

var LockCMD = cli.Command{
	Name:     "lock",
	Category: CmdCateMisc,
	Usage:    "Generates and manipulates lock files.",
}
