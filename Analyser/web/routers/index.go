package routers

import (
	"FindGhost/Analyser/models"
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"net/http"
)

// 列出所有恶意ip
func IpList(fCtx flamego.Context, t template.Template, tData template.Data) {
	evilIpInfos := models.ListEvilIps()
	tData["evilIps"] = evilIpInfos
	t.HTML(http.StatusOK, "ip")
}

// 列出所有恶意dns
func DnsList(fCtx flamego.Context, t template.Template, tData template.Data) {
	evilDnsInfos := models.ListEvilDns()
	tData["evilDns"] = evilDnsInfos
	t.HTML(http.StatusOK, "dns")
}

// 列出所有恶意http
func HttpList(fCtx flamego.Context, t template.Template, tData template.Data) {
	evilHttpInfos := models.ListEvilHttp()
	tData["evilHttp"] = evilHttpInfos
	t.HTML(http.StatusOK, "http")
}
