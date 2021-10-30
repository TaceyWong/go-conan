package main

import "github.com/urfave/cli/v2"

var ImportsCMD = cli.Command{
	Name:     "imports",
	Category: CmdCateMisc,
	Usage:    "Calls your local conanfile.py or conanfile.txt 'imports' method.",
}
