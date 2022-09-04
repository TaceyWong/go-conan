package main

import (
	"log"
	"os"

	cs "github.com/mitchellh/colorstring"
	"github.com/urfave/cli/v2"
)

const (
	CmdCateConsumer = "Consumer"
	CmdCateCreator  = "Creator"
	CmdCatePkgDev   = "Package development"
	CmdCateMisc     = "Misc"
)

/**
 * Entry point of `can`.
 */

func init() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate
	// cli.MarkdownDocTemplate = MarkdownDocTemplate
}

func main() {
	app := &cli.App{
		Name:    "gocona",
		Usage:   cs.Color("C/C++ package manager (a clone of [red][underline]conan[reset])."),
		Version: Version,
		Authors: []*cli.Author{{Name: "Tacey Wong",
			Email: "xinyong.w@gmail.com"}},
		Copyright: "©2021 Tacey Wong,All Rights Reserved Under MIT License ",
	}
	app.Commands = []*cli.Command{
		// Consumer Commands
		&InstallCMD, &ConfigCMD, &GetCMD, &InfoCMD, &SearchCMD,
		// Creator Commands
		&NewCMD, &CreateCMD, &UploadCMD, &ExportCMD, &ExportPkgCMD, &TestCMD,
		// Package Devlopment Commands
		&SourceCMD, &BuildCMD, &PackageCMD, &EditableCMD, &WorkspaceCMD,
		// Misc Commands
		&ProfileCMD, &RemoteCMD, &UserCMD, &ImportsCMD, &CopyCMD, &RemoveCMD,
		&AliasCMD, &DownloadCMD, &InspectCMD, &LockCMD, &FrogarianCMD,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
