package cmd

import "github.com/jrdn/goutil/cli"

func SetupGimmeCLI(app *cli.CLI) {
	app.SetRoot(Root())
	app.Add("some", Some())
	app.Add("list", List())
	app.Add("edit", Edit())
	app.Add("dir", Dir())
}
