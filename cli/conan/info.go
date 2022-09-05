package main

import "github.com/urfave/cli/v2"

var InfoCMD = cli.Command{
	Name:     "info",
	Category: CmdCateConsumer,
	Usage:    "Gets information about the dependency graph of a recipe.",
	UsageText: `[-h] [--paths] [-bo BUILD_ORDER] [-g GRAPH] [-if INSTALL_FOLDER]
               [-j [JSON]] [-n ONLY] [--package-filter [PACKAGE_FILTER]]
               [-db [DRY_BUILD]] [-b [BUILD]] [-r REMOTE] [-u] [-l LOCKFILE]
               [--lockfile-out LOCKFILE_OUT] [-e ENV_HOST] [-e:b ENV_BUILD]
               [-e:h ENV_HOST] [-o OPTIONS_HOST] [-o:b OPTIONS_BUILD]
               [-o:h OPTIONS_HOST] [-pr PROFILE_HOST] [-pr:b PROFILE_BUILD]
               [-pr:h PROFILE_HOST] [-s SETTINGS_HOST] [-s:b SETTINGS_BUILD]
               [-s:h SETTINGS_HOST] [-c CONF_HOST] [-c:b CONF_BUILD]
               [-c:h CONF_HOST]
               path_or_reference

It can be used with a recipe or a reference for any existing package in
your local cache.
			   
positional arguments:
 path_or_reference     Path to a folder containing a recipe (conanfile.py or
					   conanfile.txt) or to a recipe file. e.g.,
					   ./my_project/conanfile.txt. It could also be a reference
			  `,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "paths", Usage: "Show package paths in local cache"},
	},
}
