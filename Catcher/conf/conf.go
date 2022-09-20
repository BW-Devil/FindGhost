package conf

import "gopkg.in/ini.v1"

var (
	Cfg        *ini.File
	DeviceName string
	DebugMode  bool
	FilterRule string
	err        error
)

func init() {
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	DeviceName = Cfg.Section("").Key("DEVICE_NAME").MustString("eth0")
	DebugMode = Cfg.Section("").Key("DEBUG_MODE").MustBool(false)
	FilterRule = Cfg.Section("").Key("FILTER_RULE").MustString("tcp or (udp and dst port 53)")
}
