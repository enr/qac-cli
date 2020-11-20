package e2e

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/enr/go-files/files"
	"github.com/enr/qac"
)

// go test -v -timeout 30s github.com/enr/qac-cli/e2e -run ^TestExecutions$
func TestExecutions(t *testing.T) {
	executions(`./basic.yaml`, t)
	// executions(`./process.yaml`, t)
	os := runtime.GOOS
	osTestFile := fmt.Sprintf(`./%s.yaml`, os)
	if files.Exists(osTestFile) {
		executions(osTestFile, t)
	}
}

func executions(f string, t *testing.T) {
	launcher := qac.NewLauncher()
	report := launcher.ExecuteFile(f)
	reporter := qac.NewTestLogsReporter(t)
	reporter.Publish(report)

	for ei, err := range report.AllErrors() {
		t.Errorf(`%d %s error %v`, ei, f, err)
	}
}
