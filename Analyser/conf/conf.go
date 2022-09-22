package conf

import (
	"FindGhost/Analyser/util"
	"gopkg.in/ini.v1"
)

var (
	Cfg       *ini.File
	DebugMode bool
	SecureKey string
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		util.Log.Panicln(err)
	}

	DebugMode = Cfg.Section("").Key("DEBUG_MODE").MustBool(false)
	SecureKey = Cfg.Section("").Key("SECURE_KEY").MustString("")
}
