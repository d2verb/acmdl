package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "acmdl"
	app.Usage = "ACM Digital Library CLI tool"
	app.Version = "0.0.1"

	wordsFlag := cli.StringFlag{
		Name:  "words, w",
		Usage: "words",
	}

	yearFlag := cli.IntFlag{
		Name:  "year, y",
		Value: -1,
		Usage: "minimum publication year",
	}

	startFlag := cli.IntFlag{
		Name:  "start, s",
		Value: 1,
		Usage: "minimum numbering",
	}

	app.Flags = []cli.Flag{
		wordsFlag,
		yearFlag,
		startFlag,
	}

	app.Action = ExecCLI
	app.Run(os.Args)
}
