package main

import (
	"fmt"
	"os"

	"github.com/enr/clui"
	"github.com/urfave/cli/v2"

	"github.com/enr/qac-cli/lib/core"
)

var (
	ui              *clui.Clui
	versionTemplate = `%s
Revision: %s
Build date: %s
`
	appVersion        = fmt.Sprintf(versionTemplate, core.Version, core.GitCommit, core.BuildTime)
	ignoreMissingDirs bool
)

func main() {
	app := cli.NewApp()
	app.Name = "qac"
	app.Version = appVersion
	app.Description = "E2E tests for command line tools"
	app.Usage = "qac qac.yaml"
	app.Flags = []cli.Flag{
		// TBD
		// &cli.BoolFlag{Name: "debug", Aliases: []string{"D"}, Usage: "operates in debug mode: lot of output"},
		// &cli.BoolFlag{Name: "quiet", Aliases: []string{"q"}, Usage: "operates in quiet mode"},
	}
	app.EnableBashCompletion = true

	app.Action = doCheck

	app.Before = func(c *cli.Context) error {
		if ui != nil {
			return nil
		}
		verbosityLevel := clui.VerbosityLevelMedium
		if c.Bool("debug") {
			verbosityLevel = clui.VerbosityLevelHigh
		}
		if c.Bool("quiet") {
			verbosityLevel = clui.VerbosityLevelLow
		}
		var err error
		ui, err = clui.NewClui(func(ui *clui.Clui) {
			ui.VerbosityLevel = verbosityLevel
		})
		if err != nil {
			return err
		}
		return err
	}

	app.Commands = commands

	// returns an error
	app.Run(os.Args)
}
