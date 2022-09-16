package web

import (
	"Crawler/conf"
	"Crawler/web/router"
	"github.com/flamego/flamego"
	"github.com/urfave/cli/v2"
)

var (
	HTTP_HOST string
	HTTP_PORT int
)

func init() {
	HTTP_HOST = conf.Cfg.Section("").Key("HTTP_HOST").MustString("127.0.0.1")
	HTTP_PORT = conf.Cfg.Section("").Key("HTTP_PORT").MustInt(8888)

}

// 运行web
func RunWeb(cCtx *cli.Context) {
	f := flamego.Classic()

	// 设置渲染器
	f.Use(flamego.Renderer(
		flamego.RenderOptions{
			JSONIndent: " ",
		},
	))

	// 设置静态资源
	f.Use(flamego.Static(
		flamego.StaticOptions{
			Directory: "web/public",
			Index:     "index.html",
		},
	))

	// 主页
	f.Get("/")

	// 检查ip是否恶意
	f.Get("/api/ip/{ip}", router.CheckIp)

	// 检查域名是否恶意
	f.Get("/api/domain/{domain}", router.CheckDomain)

	f.Run(HTTP_HOST, HTTP_PORT)
}
