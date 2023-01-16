package cfgsqlite

import "github.com/bitwormhole/starter/application"

// ExportConfigForSQLite ...
func ExportConfigForSQLite(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
