package router

import (
	"Crawler/conf"
	"Crawler/models"
	"Crawler/util/log"
	"github.com/flamego/flamego"
	"net/http"
	"strings"
)

type IplistApi struct {
	Evil bool          `json:"evil"`
	Data models.IpList `json:"data"`
}

type DomainListApi struct {
	Evil bool              `json:"evil"`
	Data models.DomainList `json:"data"`
}

// 检查是否为恶意domain
func CheckDomain(fCtx flamego.Context, r flamego.Render) {
	var domainListApi DomainListApi
	domain := strings.TrimSpace(fCtx.Param("domain"))

	// 在缓冲区寻找是否存在
	v, _ := models.CACHE_DOMAINS.Get(domain)

	data, ok := v.(models.DomainList)
	// 是否是调试态
	if conf.DEBUG {
		log.Log.Infof("domain:%v, data:%v", domain, data)
	}

	if ok {
		domainListApi.Evil = true
		domainListApi.Data = data
	}

	r.JSON(http.StatusOK, domainListApi)
}

// 检查是否为恶意ip
func CheckIp(fCtx flamego.Context, r flamego.Render) {
	var ipListApi IplistApi
	ip := strings.TrimSpace(fCtx.Param("ip"))

	// 在缓冲区寻找是否存在
	v, _ := models.CACHE_IPS.Get(ip)

	data, ok := v.(models.IpList)
	// 是否是调试态
	if conf.DEBUG {
		log.Log.Infof("ip:%v, data:%v", ip, data)
	}

	if ok {
		ipListApi.Evil = true
		ipListApi.Data = data
	}

	r.JSON(http.StatusOK, ipListApi)
}
