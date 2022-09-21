package conf

import (
	"FindGhost/Catcher/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	Cfg        *ini.File
	DeviceName string
	DebugMode  bool
	FilterRule string
	err        error
	Ips        []string
)

func init() {
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	DeviceName = Cfg.Section("").Key("DEVICE_NAME").MustString("eth0")
	DebugMode = Cfg.Section("").Key("DEBUG_MODE").MustBool(false)
	FilterRule = Cfg.Section("").Key("FILTER_RULE").MustString("tcp or (udp and dst port 53)")

	// 获取网络设备ip
	Ips, _ = util.GetIps(DeviceName)
	logrus.Infof("deviceName: %v, ips: %v", DeviceName, Ips)
}
