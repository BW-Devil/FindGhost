package conf

import (
	"Crawler/util/log"
	"gopkg.in/ini.v1"
)

var (
	Cfg    *ini.File
	DEBUG  bool
	SECRET string
	err    error
)

func init() {
	source := "conf/app.ini"
	//Cfg = new(ini.File)
	Cfg, err = ini.Load(source)
	if err != nil {
		log.Log.Logger.Panicln(err)
	}

	DEBUG = Cfg.Section("").Key("DEBUG").MustBool(true)
	SECRET = Cfg.Section("").Key("SECRET").MustString("SECRET_KEY")
}
