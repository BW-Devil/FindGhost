package routers

import (
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"net/http"
)

// 列出所有恶意ip
func IpList(fCtx flamego.Context, t template.Template, tData template.Data) {
	tData["name"] = "test"
	t.HTML(http.StatusOK, "ip")
}
