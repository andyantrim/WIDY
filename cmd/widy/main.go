package main

import (
	"fmt"
	"os"

	"github.com/andyantrim/WIDY/funcs"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	funcs.BuildCLI(app)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}