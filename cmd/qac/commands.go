package main

import (
	"fmt"

	"github.com/enr/qac"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	&checkCommand,
}

var checkCommand = cli.Command{
	Name:        "check",
	Aliases:     []string{"p"},
	Usage:       "",
	Description: `Check`,
	Action:      doCheck,
}

func doCheck(c *cli.Context) error {
	if !c.Args().Present() {
		return exitErrorf(1, "Missing file")

	}
	qf := c.Args().Get(0)
	launcher := qac.NewLauncher()
	report := launcher.ExecuteFile(qf)
	reporter := newUIReporter(ui)
	err := reporter.Publish(report)
	if err != nil {
		return exitErrorf(1, "Error check: %v", err)
	}
	if len(report.AllErrors()) > 0 {
		return exitErrorf(1, "The plan failed")
	}
	return nil
}

func exitErrorf(exitCode int, template string, args ...interface{}) error {
	ui.Errorf(`Something gone wrong:`)
	return cli.NewExitError(fmt.Sprintf(template, args...), exitCode)
}
