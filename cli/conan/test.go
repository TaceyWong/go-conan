package main

import "github.com/urfave/cli/v2"

var TestCMD = cli.Command{
	Name:     "test",
	Category: CmdCateCreator,
	Usage:    "Tests a package consuming it from a conanfile.py with a test() method.",
}
