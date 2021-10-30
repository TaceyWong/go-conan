package main

import "github.com/urfave/cli/v2"

var SearchCMD = cli.Command{
	Name:     "search",
	Category: CmdCateConsumer,
	Usage: `Searches package recipes and binaries in the local cache or a remote. Unless a
	remote is specified only the local cache is searched.`,
}
