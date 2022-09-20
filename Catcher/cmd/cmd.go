package cmd

import (
	"FindGhost/Catcher/catch"
	"github.com/urfave/cli/v2"
)

var Catch = &cli.Command{
	Name:    "catch",
	Aliases: []string{"c"},
	Usage:   "catch packet off the wire",
	Action:  catch.Start,
}
