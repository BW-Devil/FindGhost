package routers

import (
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"net/http"
)

// 列出所有恶意ip
func IpList(fCtx flamego.Context, t template.Template, tData template.Data) {
	tData["name"] = "test"
	tData["age"] = 18
	t.HTML(http.StatusOK, "ip")
}

// 列出所有恶意dns
func DnsList(fCtx flamego.Context, t template.Template, tData template.Data) {

}

// 列出所有恶意http
func HttpList(fCtx flamego.Context, t template.Template, tData template.Data) {

}
