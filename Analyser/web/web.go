package web

import (
	"FindGhost/Analyser/conf"
	"FindGhost/Analyser/web/routers"
	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"github.com/urfave/cli/v2"
)

var (
	HTTP_HOST string
	HTTP_PORT int
)

func init() {
	cfg := conf.Cfg
	HTTP_HOST = cfg.Section("").Key("HTTP_HOST").MustString("127.0.0.1")
	HTTP_PORT = cfg.Section("").Key("HTTP_PORT").MustInt(7777)
}

func RunWeb(cCtx *cli.Context) (err error) {
	f := flamego.Classic()
	// 使用template中间件渲染html页面
	f.Use(template.Templater())

	// 设置静态资源
	f.Use(flamego.Static(
		flamego.StaticOptions{
			Directory: "web/public",
			Index:     "ip.html",
		},
	))

	// 主页
	f.Get("/", routers.IpList)

	// 显示所有恶意ip
	f.Get("/ip/", routers.IpList)

	// ipInfo的审计路由
	f.Post("/api/ip/", routers.ProcessIpInfo)

	// dnsInfo的审计路由
	f.Post("/api/dns/", routers.ProcessDnsInfo)

	// httpInfo的审计路由
	f.Post("/api/http/", routers.ProcessHttpInfo)

	f.Run(HTTP_HOST, HTTP_PORT)
	return err
}
