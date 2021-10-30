package main

import "github.com/urfave/cli/v2"

var UserCMD = cli.Command{
	Name:     "user",
	Category: CmdCateMisc,
	Usage:    "Authenticates against a remote with user/pass, caching the auth token.",
}
