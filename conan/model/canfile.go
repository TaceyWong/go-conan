package model

type CanFile struct {
	Name    string //
	Version string // can be "1.1" or whatever
	URL     string // The URL where this File is located, as github, to collaborate in package
	// The license of the PACKAGE, just a shortcut, does not replace or
	// change the actual license of the source code
	License     string
	Author      string // Main maintainer/responsible for the package, any format
	Description string //
	Topics      []string
	HomePage    string
	BuildPolicy string
	ShortPath   bool
	//  Apply environment variables from requires deps_env_info and profiles
	ApplyENV      bool
	Exports       []string
	ExportsSource []string
	Generators    []string
	RevisionModev string

	// Vars to control the build steps (build(), package())
	ShouldConfigure bool
	ShouldBuild     bool
	ShouldInstall   bool
	ShouldTest      bool
}
