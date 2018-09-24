package config

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var enforcer *casbin.Enforcer

func initCasbin() *casbin.Enforcer {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	database := viper.GetString("db.database")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?", user, password, host, port, database)
	a := gormadapter.NewAdapter("mysql", url, true) // Your driver and data source.
	e := casbin.NewEnforcer("rbac_model.conf", a)
	return e
}
func GetEnforcer() *casbin.Enforcer {
	if enforcer == nil {
		enforcer = initCasbin()
	}
	return enforcer
}
