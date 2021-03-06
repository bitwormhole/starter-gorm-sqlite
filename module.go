package sqlite

import (
	"embed"

	startergorm "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter-gorm-sqlite/src/main/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-gorm-sqlite"
	myVersion  = "v0.0.3"
	myRevision = 3
)

// DriverModule 导出本模块
func DriverModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.Dependency(startergorm.Module())
	mb.OnMount(etc.ExportConfig)
	mb.Resources(myResources())

	return mb.Create()
}

//go:embed src/main/resources
var theResFS embed.FS

func myResources() collection.Resources {
	return collection.LoadEmbedResources(&theResFS, "src/main/resources")
}
