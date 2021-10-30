package main

import "github.com/urfave/cli/v2"

var DownloadCMD = cli.Command{
	Name:     "download",
	Category: CmdCateMisc,
	Usage:    "Downloads recipe and binaries to the local cache, without using settings.",
}
