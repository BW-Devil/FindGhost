package audit

import (
	"FindGhost/Analyser/conf"
	"FindGhost/Analyser/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ApiUrl string
)

func init() {
	ApiUrl = conf.Cfg.Section("AUDIT_SERVER").Key("API_URL").MustString("")
}

// 审计ips
func AuditIpInfo(sensorIp string, ipInfo models.IpInfo) {
	// 获取ipInfo中的所有ip
	ips := make([]string, 0)
	ips = append(ips, ipInfo.SrcIp, ipInfo.DestIp)

	// 逐个审计ip
	for _, ip := range ips {
		if ip == sensorIp {
			continue
		}

		auditIpUrl := fmt.Sprintf("%v%v%v", ApiUrl, "/api/ip/", ip)
		resp, err := http.Get(auditIpUrl)
		var auditIpRes models.IpListApi
		if err == nil {
			respBody, err := io.ReadAll(resp.Body)
			if err == nil {
				err = json.Unmarshal(respBody, &auditIpRes)
				isEvil := auditIpRes.Evil
				if isEvil {
					evilIpInfo := models.NewEvilIpInfo(sensorIp, ipInfo, auditIpRes)
					evilIpInfo.Insert()
				}
			}
		}
	}
}
