package main

import "github.com/urfave/cli/v2"

var RemoteCMD = cli.Command{
	Name:     "remote",
	Category: CmdCateMisc,
	Usage:    "Manages the remote list and the package recipes associated with a remote.",
}
