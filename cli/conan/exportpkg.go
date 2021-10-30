package main

import "github.com/urfave/cli/v2"

var ExportPkgCMD = cli.Command{
	Name:     "export-pkg",
	Category: CmdCateCreator,
	Usage:    "xports a recipe, then creates a package from local source and build folders.",
}
