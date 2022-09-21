package main

import (
	"FindGhost/Catcher/cmd"
	"github.com/sirupsen/logrus"
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
	app.Description = "FindGhost Catcher"
	app.Usage = "FindGhost Catcher"
	app.Version = "1.0.0"
	app.Authors = []*cli.Author{{"BWFish", "weunknowing@gmail.com"}}
	app.Commands = []*cli.Command{
		cmd.Catch,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
