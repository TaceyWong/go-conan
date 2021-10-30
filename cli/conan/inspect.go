package main

import "github.com/urfave/cli/v2"

var InspectCMD = cli.Command{
	Name:     "inspect",
	Category: CmdCateMisc,
	Usage: `Displays conanfile attributes, like name, version, and options. Works locally,
	in local cache and remote.`,
}
