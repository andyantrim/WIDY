package funcs

import "github.com/urfave/cli/v2"

func info(app *cli.App) {
	app.Name = "WIDY a tool to summarize your last day's work"
	app.Usage = "Use to recursively search git repos for contributions"
	app.Version = "0.0.1"

	// Author
	me := cli.Author{
		Name: "Andy Blair",
	}

	app.Authors = append(app.Authors, &me)
}
