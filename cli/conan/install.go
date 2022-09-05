package main

import "github.com/urfave/cli/v2"

var InstallCMD = cli.Command{
	Name:     "install",
	Category: CmdCateConsumer,
	Usage:    "Installs the requirements specified in a recipe (conanfile.py or conanfile.txt).",
	UsageText: `[-h] [-g GENERATOR] [-if INSTALL_FOLDER] [-of OUTPUT_FOLDER]
                  [-m [MANIFESTS]] [-mi [MANIFESTS_INTERACTIVE]] [-v [VERIFY]]
                  [--no-imports] [--build-require] [-j JSON] [-b [BUILD]]
                  [-r REMOTE] [-u] [-l LOCKFILE] [--lockfile-out LOCKFILE_OUT]
                  [-e ENV_HOST] [-e:b ENV_BUILD] [-e:h ENV_HOST] [-o OPTIONS_HOST]
                  [-o:b OPTIONS_BUILD] [-o:h OPTIONS_HOST] [-pr PROFILE_HOST]
                  [-pr:b PROFILE_BUILD] [-pr:h PROFILE_HOST] [-s SETTINGS_HOST]
                  [-s:b SETTINGS_BUILD] [-s:h SETTINGS_HOST] [-c CONF_HOST]
                  [-c:b CONF_BUILD] [-c:h CONF_HOST]
                  [--lockfile-node-id LOCKFILE_NODE_ID]
                  [--require-override REQUIRE_OVERRIDE]
                  path_or_reference [reference]
	
	
positional arguments:
path_or_reference     Path to a folder containing a recipe (conanfile.py or
					  conanfile.txt) or to a recipe file. e.g.,
					  ./my_project/conanfile.txt. It could also be a reference
reference             Reference for the conanfile path of the first argument:
					  user/channel, version@user/channel or
					  pkg/version@user/channel(if name or version declared in
					  conanfile.py, they should match)

`,
	Description: `Installs the requirements specified in a recipe (conanfile.py or conanfile.txt).

It can also be used to install a concrete package specifying a
reference. If any requirement is not found in the local cache, it will
retrieve the recipe from a remote, looking for it sequentially in the
configured remotes. When the recipes have been downloaded it will try
to download a binary package matching the specified settings, only from
the remote from which the recipe was retrieved. If no binary package is
found, it can be built from sources using the '--build' option. When
the package is installed, Conan will write the files for the specified
generators.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "generator", Aliases: []string{"g"},
			Usage: "`Generator`s to use",
		},
		&cli.PathFlag{
			Name: "install-folder", Aliases: []string{"if"},
			Usage: "Use this directory as the `directory` where to put the generatorfiles. e.g.,conaninfo/conanbuildinfo.txt",
		},
		&cli.PathFlag{
			Name: "output-folder", Aliases: []string{"of"},
			Usage: "The root `output folder` for generated and build files",
		},
		&cli.PathFlag{
			Name: "manifests", Aliases: []string{"m"},
			Usage: "Install dependencies `manifests` in folder for later verify.",
			Value: ".conan_manifests",
		},
		&cli.PathFlag{
			Name: "manifests-interactive", Aliases: []string{"mi"},
			Usage: "Install dependencies `manifests` in folder for later verify, asking user for confirmation.",
			Value: ".conan_manifests",
		},
	},
}

func InstallCMDAction(c *cli.Context) (err error) {
	return
}
