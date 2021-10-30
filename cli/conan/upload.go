package main

import "github.com/urfave/cli/v2"

var UploadCMD = cli.Command{
	Name:     "upload",
	Category: CmdCateCreator,
	Usage:    "Uploads a recipe and binary packages to a remote.",
}
