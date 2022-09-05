package main

import "github.com/urfave/cli/v2"

var LockCMD = cli.Command{
	Name:     "lock",
	Category: CmdCateMisc,
	Usage:    "Generates and manipulates lock files.",
	UsageText: `{update,build-order,clean-modified,install,create,bundle} ...
	
positional arguments:
	{update,build-order,clean-modified,install,create,bundle}`,
	Subcommands: []*cli.Command{
		{
			Name: "update",
			Usage: `Complete missing information in the first lockfile with information
			defined in the second lockfile. Both lockfiles must represent the same
			graph, and have the same topology with the same identifiers, i.e. the
			second lockfile must be an evolution based on the first one`,
		},
		{
			Name: "build-order",
			Usage: "Returns build-order",
		},
		{
			Name: "clean-modified",
			Usage: "Clean modified flags",
		},
		{
			Name: "install",
			Usage: "Install a lockfile",
		},
		{
			Name: "create",
			Usage: "Create a lockfile from a conanfile or a reference",
		},
		{
			Name: "bundle",
			Usage: "Manages lockfile bundles",
		},
	},
}
