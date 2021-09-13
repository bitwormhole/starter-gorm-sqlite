package golang

import (
	"testing"

	root "github.com/bitwormhole/starter-gorm-sqlite"
	"github.com/bitwormhole/starter/tests"
)

func TestSQLServerDriver(t *testing.T) {

	i := tests.Starter(t)
	i.Use(root.ExportTestModule())
	rt := i.RunTest()
	rt.Exit()
}
