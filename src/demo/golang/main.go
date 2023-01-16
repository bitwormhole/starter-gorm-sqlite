package main

import (
	"github.com/bitwormhole/starter"

	starter_gorm_sqlite "github.com/bitwormhole/starter-gorm-sqlite"
)

func main() {
	i := starter.InitApp()
	i.UseMain(starter_gorm_sqlite.DemoModule())
	i.Run()
}
