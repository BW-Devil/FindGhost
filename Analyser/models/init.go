package models

import (
	"FindGhost/Analyser/conf"
	"github.com/upper/db/v4"
)

var (
	DbConfig   DbCONF
	DbSettings db.ConnectionURL
)

type DbCONF struct {
	DbType string
	DbHost string
	DbPort int64
	DbUser string
	DbPass string
	DbName string
}

func init() {
	cfg := conf.Cfg
	sec := cfg.Section("database")
	DbConfig.DbType = sec.Key("DB_TYPE").MustString("mongodb")
	DbConfig.DbHost = sec.Key("DB_HOST").MustString("127.0.0.1")
	DbConfig.DbPort = sec.Key("DB_PORT").MustInt64(27017)
	DbConfig.DbUser = sec.Key("DB_USER").MustString("user")
	DbConfig.DbPass = sec.Key("DB_PASS").MustString("password")
	DbConfig.DbName = sec.Key("DB_NAME").MustString("proxy_honeypot")

}
