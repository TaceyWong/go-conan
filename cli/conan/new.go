package main

import "github.com/urfave/cli/v2"

var NewCMD = cli.Command{
	Name:     "new",
	Category: CmdCateCreator,
	Usage: `Creates a new package recipe template with a 'conanfile.py' and optionally,
	'test_package' testing files.`,
}
