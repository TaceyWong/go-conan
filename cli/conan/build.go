package main

import "github.com/urfave/cli/v2"

var BuildCMD = cli.Command{
	Name:     "build",
	Category: CmdCatePkgDev,
	Usage:    "Calls your local conanfile.py 'build()' method.",
}
