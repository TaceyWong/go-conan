package main

import "github.com/urfave/cli/v2"

var ProfileCMD = cli.Command{
	Name:     "profile",
	Category: CmdCateMisc,
	Usage:    "Lists profiles in the '.conan/profiles' folder, or shows profile details.",
}
