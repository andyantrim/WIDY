package funcs

import "github.com/urfave/cli/v2"

func authorFlag(app *cli.App) {
	flag := &cli.StringFlag{
		Name:     "author",
		Usage:    "the author to search commits for",
		Required: true,
	}
	app.Flags = append(app.Flags, flag)
}

func daysFlag(app *cli.App) {
	flag := &cli.IntFlag{
		Name:  "days",
		Usage: "Specify how many days ago to start search from",
	}
	app.Flags = append(app.Flags, flag)
}

func fileFlag(app *cli.App) {
	flag := &cli.StringFlag{
		Name:     "dir",
		Usage:    "The base directory that contains the git repos to search",
		Required: true,
	}
	app.Flags = append(app.Flags, flag)
}
