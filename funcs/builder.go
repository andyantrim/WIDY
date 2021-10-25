package funcs

import "github.com/urfave/cli/v2"

func BuildCLI(app *cli.App) {
	info(app)
	authorFlag(app)
	daysFlag(app)
	fileFlag(app)
	app.Action = runner
}
