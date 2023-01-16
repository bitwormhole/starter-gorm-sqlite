package cfgsqlitedemo

import "github.com/bitwormhole/starter/application"

// ExportConfigForSQLiteDemo ...
func ExportConfigForSQLiteDemo(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
