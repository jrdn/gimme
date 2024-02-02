package main

import (
	"fmt"
	"os"

	"github.com/jrdn/goutil/cli"

	"github.com/jrdn/gimme/cmd"
)

func main() {
	app := cli.NewCLI("gimme")
	cmd.SetupGimmeCLI(app)

	if err := app.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
