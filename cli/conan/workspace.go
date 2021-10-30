package main

import "github.com/urfave/cli/v2"

var WorkspaceCMD = cli.Command{
	Name:     "workspace",
	Category: CmdCatePkgDev,
	Usage: `Manages a workspace (a set of packages consumed from the user workspace that
	belongs to the same project).`,
}
