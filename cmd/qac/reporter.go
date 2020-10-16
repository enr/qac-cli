package main

import (
	"github.com/enr/clui"
	"github.com/enr/qac"
)

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Printf("non abbastanza args %q \n", os.Args)
// 		os.Exit(1)
// 	}
// 	verbosity := func(ui *clui.Clui) {
// 		ui.VerbosityLevel = clui.VerbosityLevelHigh
// 	}
// 	ui, _ := clui.NewClui(verbosity)
// 	qf := os.Args[1]
// 	launcher := qac.NewLauncher()
// 	report := launcher.ExecuteFile(qf)
// 	reporter := newUIReporter(ui)
// 	reporter.Publish(report)
// }

func newUIReporter(ui *clui.Clui) qac.Reporter {
	return &uiReporter{
		ui: ui,
	}
}

type uiReporter struct {
	ui *clui.Clui
}

func (r *uiReporter) Publish(report *qac.TestExecutionReport) error {
	for _, block := range report.Blocks() {
		r.ui.Titlef("Phase <%s>", block.Phase())
		for _, entry := range block.Entries() {
			// r.ui.Lifecyclef(" - %s", entry.Description())
			switch entry.Kind() {
			case qac.ErrorType:
				r.ui.Errorf("  | - KO %s", entry.Description())
				for i, err := range entry.Errors() {
					r.ui.Lifecyclef("    %d %s", (i + 1), err.Error())
				}
				break
			case qac.InfoType:
				r.ui.Lifecyclef(`  | INFO %s`, entry.Description())
				break
			case qac.SuccessType:
				r.ui.Successf("  | - OK %s", entry.Description())
				break
			default:
				r.ui.Warnf("unexpected kind %v", entry.Kind())
			}
		}
	}
	return nil
}
