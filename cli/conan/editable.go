package main

import "github.com/urfave/cli/v2"

var EditableCMD = cli.Command{
	Name:     "editable",
	Category: CmdCatePkgDev,
	Usage: `Manages editable packages (packages that reside in the user workspace, but are
	consumed as if they were in the cache).`,
}
