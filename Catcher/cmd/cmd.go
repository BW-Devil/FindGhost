package cmd

import (
	"FindGhost/Catcher/catcher"
	"github.com/urfave/cli/v2"
)

func stringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

var StartUp = &cli.Command{
	Name:        "start",
	Usage:       "startup findghost catcher",
	Description: "startup findghost catcher",
	Action:      catcher.Start,
	Flags: []cli.Flag{
		boolFlag("debug", "debug mode"),
		stringFlag("filter", "", "set filters"),
		intFlag("length", 1024, "set snapshot length"),
	},
}
