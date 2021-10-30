package main

import "github.com/urfave/cli/v2"

var ExportCMD = cli.Command{
	Name:     "export",
	Category: CmdCateCreator,
	Usage:    "Copies the recipe (conanfile.py & associated files) to your local cache.",
}
