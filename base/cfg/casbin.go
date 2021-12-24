package cfg

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

// Initialize the casbin model from a string.
var casbinConfText = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`

var CasbinSE *casbin.SyncedEnforcer

func InitCasbin() {
	// Initialize a gorm adapter and use it in a Casbin enforcer
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&timeout=%s",
		MysqlConfObj.MysqlUsername,
		MysqlConfObj.MysqlPassword,
		MysqlConfObj.MysqlHost,
		MysqlConfObj.MysqlPort,
		MysqlConfObj.MysqlDatabase,
		MysqlConfObj.MysqlTimeout,
	)
	adapterObj, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		Mlog.Panic("failed to new casbin gorm adapter, error message: ", err.Error())
		return
	}

	modelObj, err := model.NewModelFromString(casbinConfText)
	if err != nil {
		Mlog.Panic("failed to new casbin model from string, error message: ", err.Error())
		return
	}

	CasbinSE, err = casbin.NewSyncedEnforcer(modelObj, adapterObj)
	if err != nil {
		Mlog.Panic("failed to new casbin synced enforcer, error message: ", err.Error())
		return
	}

	CasbinSE.EnableAutoSave(true)

	// 增加对url的正则支持，格式如: a URL path or a : pattern like /alice_data/:resource
	// 同时也需要更改model文件内容的[matchers]里面, 比如 ...r.dom == p.dom && keyMatch2(r.obj, p.obj)...
	// CasbinSE.GetRoleManager().(*defaultrolemanager.RoleManager).AddDomainMatchingFunc("KeyMatch2", util.KeyMatch2)
}
