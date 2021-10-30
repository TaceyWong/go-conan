package main

import "github.com/urfave/cli/v2"

var CreateCMD = cli.Command{
	Name:     "create",
	Category: CmdCateCreator,
	Usage:    "Builds a binary package for a recipe (conanfile.py).",
}
