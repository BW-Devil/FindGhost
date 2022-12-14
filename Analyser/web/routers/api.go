package routers

import (
	"FindGhost/Analyser/audit"
	"FindGhost/Analyser/conf"
	"FindGhost/Analyser/models"
	"FindGhost/Analyser/util"
	"encoding/json"
	"github.com/flamego/flamego"
)

// 分析ipInfo
func ProcessIpInfo(fCtx flamego.Context) {
	_ = fCtx.Request().ParseForm()
	timeStamp := fCtx.Request().Form.Get("timeStamp")
	secureKey := fCtx.Request().Form.Get("secureKey")
	data := fCtx.Request().Form.Get("data")
	sensorIp := fCtx.Request().RemoteAddr

	// 验证签名
	if secureKey == util.MakeSign(timeStamp, conf.SecureKey) {
		var ipInfo models.IpInfo
		err := json.Unmarshal([]byte(data), &ipInfo)
		if err == nil {
			go func(sensorIp string, ipInfo models.IpInfo) {
				audit.AuditIpInfo(sensorIp, ipInfo)
			}(sensorIp, ipInfo)
		}
	}
}

// 分析dnsInfo
func ProcessDnsInfo(fCtx flamego.Context) {
	_ = fCtx.Request().ParseForm()
	timeStamp := fCtx.Request().Form.Get("timeStamp")
	secureKey := fCtx.Request().Form.Get("secureKey")
	data := fCtx.Request().Form.Get("data")
	sensorIp := fCtx.Request().RemoteAddr

	// 验证签名
	if secureKey == util.MakeSign(timeStamp, conf.SecureKey) {
		var dnsInfo models.DnsInfo
		err := json.Unmarshal([]byte(data), &dnsInfo)
		if err == nil {
			go func(sensorIp string, dnsInfo models.DnsInfo) {
				audit.AuditDnsInfo(sensorIp, dnsInfo)
			}(sensorIp, dnsInfo)
		}
	}
}

// 分析httpInfo
func ProcessHttpInfo(fCtx flamego.Context) {
	_ = fCtx.Request().ParseForm()
	timeStamp := fCtx.Request().Form.Get("timeStamp")
	secureKey := fCtx.Request().Form.Get("secureKey")
	data := fCtx.Request().Form.Get("data")
	sensorIp := fCtx.Request().RemoteAddr

	// 验证签名
	if secureKey == util.MakeSign(timeStamp, conf.SecureKey) {
		var httpInfo models.HttpInfo
		err := json.Unmarshal([]byte(data), &httpInfo)
		if err == nil {
			go func(sensorIp string, httpInfo models.HttpInfo) {
				audit.AuditHttpInfo(sensorIp, httpInfo)
			}(sensorIp, httpInfo)
		}
	}
}
