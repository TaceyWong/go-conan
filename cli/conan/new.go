package main

import "github.com/urfave/cli/v2"

var NewCMD = cli.Command{
	Name:     "new",
	Category: CmdCateCreator,
	Usage: `Creates a new package recipe template with a 'conanfile.py' and optionally,
	'test_package' testing files.`,
	UsageText: `[-h] [-t] [-i] [-c] [-s] [-b] [-m TEMPLATE] [-cis] [-cilg] [-cilc]
	            [-cio] [-ciw] [-ciglg] [-ciglc] [-ciccg] [-ciccc] [-cicco] [-gi]
	            [-ciu CI_UPLOAD_URL] [-d DEFINE]
	            name

Creates a new package recipe template with a 'conanfile.py' and optionally,
'test_package' testing files.

positional arguments:
name        Package name, e.g.: "poco/1.9.4" or complete reference for CI
            scripts: "poco/1.9.4@user/channel"
`,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "test", Aliases: []string{"t"},
			Usage: "Create test_package skeleton to test package"},
		&cli.BoolFlag{Name: "header", Aliases: []string{"i"},
			Usage: "Create a headers only package template"},
		&cli.BoolFlag{Name: "pure-c", Aliases: []string{"c"},
			Usage: `Create a C language package only package, deleting"self.settings.compiler.libcxx" setting in the configure method`},
		&cli.BoolFlag{Name: "sources", Aliases: []string{"s"},
			Usage: `Create a package with embedded sources in "src" folder, using
"exports_sources" instead of retrieving external code with
the "source()" method`},
		&cli.BoolFlag{Name: "bare", Aliases: []string{"b"},
			Usage: ` Create the minimum package recipe, without build() method.
Useful in combination with "export-pkg" command`},
		&cli.StringFlag{Name: "template", Aliases: []string{"m"},
			Usage: "Use the given `template` to generate a conan project"},
		&cli.BoolFlag{Name: "ci-shared", Aliases: []string{"cis"},
			Usage: `Package will have a "shared" option to be used in CI`},
		&cli.BoolFlag{Name: "ci-travis-gcc", Aliases: []string{"cilg"},
			Usage: "Generate travis-ci files for linux gcc"},
		&cli.BoolFlag{Name: "ci-travis-clang", Aliases: []string{"cilc"},
			Usage: "Generate travis-ci files for linux clang"},
		&cli.BoolFlag{Name: "ci-travis-osx", Aliases: []string{"cio"},
			Usage: " Generate travis-ci files for OSX apple-clang"},
		&cli.BoolFlag{Name: "ci-appveyor-win", Aliases: []string{"ciw"},
			Usage: "Generate appveyor files for Appveyor Visual Studio"},
		&cli.BoolFlag{Name: "ci-gitlab-gcc", Aliases: []string{"ciglg"},
			Usage: "Generate GitLab files for linux gcc"},
		&cli.BoolFlag{Name: "ci-gitlab-clang", Aliases: []string{"ciglc"},
			Usage: "Generate GitLab files for linux clang"},
		&cli.BoolFlag{Name: "ci-circleci-gcc", Aliases: []string{"ciccg"},
			Usage: "Generate CircleCI files for linux gcc"},
		&cli.BoolFlag{Name: "ci-circleci-clang", Aliases: []string{"ciccc"},
			Usage: "Generate CircleCI files for linux clang"},
		&cli.BoolFlag{Name: "ci-circleci-osx", Aliases: []string{"cicco"},
			Usage: "Generate CircleCI files for OSX apple-clang"},
		&cli.BoolFlag{Name: "gitignore", Aliases: []string{"gi"},
			Usage: "Generate a .gitignore with the known patterns to excluded"},
		&cli.StringFlag{Name: "ci-upload-url", Aliases: []string{"ciu"},
			Usage: "Define `URL` of the repository to upload"},
		&cli.StringFlag{Name: "define", Aliases: []string{"d"},
			Usage: "custome `define`"},
	},
}
