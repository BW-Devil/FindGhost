package main

import (
	"FindGhost/Analyser/cmd"
	"FindGhost/Analyser/models"
	"FindGhost/Analyser/util"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "FindGhost Analyser"
	app.Usage = "FindGhost Analyser"
	app.Description = "analyse evil ip or domain from catcher"
	app.Version = "1.0.0"
	app.Authors = []*cli.Author{{"BWFish", "weunknowing@gmail.com"}}
	app.Commands = []*cli.Command{
		cmd.StartUp,
	}

	if err := app.Run(os.Args); err != nil {
		util.Log.Fatal(err)
	}
	defer models.DisconnectDB()
}
