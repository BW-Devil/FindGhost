package main

import (
	"Crawler/cmd"
	"Crawler/util/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crawler"
	app.Usage = "FindGhost Crawler"
	app.Description = "FindGhost Crawler"
	app.Version = "1.0.0"
	app.Authors = []*cli.Author{{"BWFish", "weunknowing@gmail.com"}}

	app.Commands = []*cli.Command{
		cmd.StartUpWeb,
		cmd.StartUpGui,
		cmd.SaveFile,
	}

	if err := app.Run(os.Args); err != nil {
		log.Log.Fatal(err)
	}
}
