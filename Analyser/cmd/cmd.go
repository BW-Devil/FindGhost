package cmd

import (
	"FindGhost/Analyser/web"
	"github.com/urfave/cli/v2"
)

var StartUp = &cli.Command{
	Name:        "start",
	Usage:       "start up analyse program",
	Description: "start up analyse program",
	Action:      web.RunWeb,
}
