package main

import (
	"FindGhost/Catcher/cmd"
	"FindGhost/Catcher/util"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "FindGhost Catcher"
	app.Authors = []*cli.Author{{"BWFish", "weunknowing@gmail.com"}}
	app.Usage = "FindGhost Catcher"
	app.Commands = []*cli.Command{
		cmd.StartUp,
	}
	app.Flags = append(app.Flags, cmd.StartUp.Flags...)

	if err := app.Run(os.Args); err != nil {
		util.Log.Fatal(err)
	}
}
