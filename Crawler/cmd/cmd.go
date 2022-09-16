package cmd

import (
	"Crawler/feeds"
	"github.com/urfave/cli/v2"
)

// 启动web的指令
var StartUpWeb = &cli.Command{
	Name:        "web",
	Usage:       "start up web program",
	Description: "start up web program",
	Action:      feeds.StartUpWeb,
}

// 启动gui的指令
var StartUpGui = &cli.Command{
	Name:        "gui",
	Usage:       "start up gui program",
	Description: "start up gui program",
	Action:      feeds.StartUpGui,
}

// 导出文件到本地
var SaveFile = &cli.Command{
	Name:        "dump",
	Usage:       "fetch evil ips and domains to file",
	Description: "fetch evil ips and domains to file",
	Action:      feeds.Dump,
}
