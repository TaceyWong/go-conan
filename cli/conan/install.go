package main

import "github.com/urfave/cli/v2"

var InstallCMD = cli.Command{
	Name:     "install",
	Category: CmdCateConsumer,
	Usage:    "Installs the requirements specified in a recipe (conanfile.py or conanfile.txt).",
	Description: `Installs the requirements specified in a recipe (conanfile.py or conanfile.txt).

It can also be used to install a concrete package specifying a
reference. If any requirement is not found in the local cache, it will
retrieve the recipe from a remote, looking for it sequentially in the
configured remotes. When the recipes have been downloaded it will try
to download a binary package matching the specified settings, only from
the remote from which the recipe was retrieved. If no binary package is
found, it can be built from sources using the '--build' option. When
the package is installed, Conan will write the files for the specified
generators.

positional arguments:
	path_or_reference     Path to a folder containing a recipe (conanfile.py or conanfile.txt) or to a recipe file. e.g.,
						./my_project/conanfile.txt. It could also be a reference
	reference             Reference for the conanfile path of the first argument: user/channel, version@user/channel or
						pkg/version@user/channel(if name or version declared in conanfile.py, they should match)
	`,
	Flags: []cli.Flag{},
}

func InstallCMDAction(c *cli.Context) (err error) {
	return
}
