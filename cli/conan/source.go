package main

import "github.com/urfave/cli/v2"

var SourceCMD = cli.Command{
	Name:     "source",
	Category: CmdCatePkgDev,
	Usage:    "Calls your local conanfile.py 'source()' method.",
}
