// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

import (
	driver0x2cfc72 "github.com/bitwormhole/starter-gorm-sqlite/driver"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-driver0x2cfc72.SQLiteDriver
	cominfobuilder.Next()
	cominfobuilder.ID("com0-driver0x2cfc72.SQLiteDriver").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSQLiteDriver{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSQLiteDriver : the factory of component: com0-driver0x2cfc72.SQLiteDriver
type comFactory4pComSQLiteDriver struct {
	mPrototype *driver0x2cfc72.SQLiteDriver
}

func (inst *comFactory4pComSQLiteDriver) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComSQLiteDriver) newObject() *driver0x2cfc72.SQLiteDriver {
	return &driver0x2cfc72.SQLiteDriver{}
}

func (inst *comFactory4pComSQLiteDriver) castObject(instance application.ComponentInstance) *driver0x2cfc72.SQLiteDriver {
	return instance.Get().(*driver0x2cfc72.SQLiteDriver)
}

func (inst *comFactory4pComSQLiteDriver) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComSQLiteDriver) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComSQLiteDriver) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComSQLiteDriver) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComSQLiteDriver) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComSQLiteDriver) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}
