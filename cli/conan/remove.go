package main

import "github.com/urfave/cli/v2"

var RemoveCMD = cli.Command{
	Name:     "remove",
	Category: CmdCateMisc,
	Usage:    "Removes packages or binaries matching pattern from local cache or remote.",
}
