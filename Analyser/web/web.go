package web

import (
	"FindGhost/Analyser/conf"
	"github.com/flamego/flamego"
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

	f.Get("/")

	f.Run(HTTP_HOST, HTTP_PORT)
	return err
}
