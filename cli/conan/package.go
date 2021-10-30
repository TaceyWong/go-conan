package main

import "github.com/urfave/cli/v2"

var PackageCMD = cli.Command{
	Name:     "package",
	Category: CmdCatePkgDev,
	Usage:    "Calls your local conanfile.py 'package()' method.",
}
